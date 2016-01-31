package gomolog

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

type wrappedWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *wrappedWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *wrappedWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	w.length = len(b)
	return w.ResponseWriter.Write(b)
}

type Gomolog struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func Open(uri string, collection string) *Gomolog {
	fmt.Printf("Connecting MongoDB %s; collection = %s\n", uri, collection)
	log := Gomolog{nil, nil}
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}
	log.session = session
	log.collection = session.DB("").C(collection)
	return &log
}

func (log *Gomolog) Close() {
	log.session.Close()
}

type request struct {
	Method         string
	Host           string
	Url            string
	Protocol       string
	AcceptLanguage string `bson:"acceptLanguage"`
}

type response struct {
	Status        int
	ContentLength int     `bson:"contentLength"`
	ResponseTime  float64 `bson:"responseTime"`
}

type remote struct {
	Addr      string
	User      string
	UserAgent string `bson:"userAgent"`
}

type logFormat struct {
	Format   int
	Date     string
	Referrer string
	Request  request
	Response response
	Remote   remote
}

func (log *Gomolog) Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := wrappedWriter{w, 0, 0}
		start := time.Now()
		http.DefaultServeMux.ServeHTTP(&writer, r)
		duration := time.Now().Sub(start)
		forwarded := strings.Split(
			strings.Join(r.Header["X-Forwarded-For"], ","), ",")
		raddr := r.RemoteAddr
		if len(forwarded) != 0 {
			raddr = forwarded[len(forwarded) - 1]
		}
		log.collection.Insert(&logFormat{
			Format:   1,
			Date:     time.Now().Format(time.RFC3339),
			Referrer: strings.Join(r.Header["Referer"], ","),
			Request: request{
				Method:         r.Method,
				Host:           r.Host,
				Url:            r.URL.Path,
				Protocol:       r.Proto,
				AcceptLanguage: strings.Join(r.Header["Accept-Language"], ",")},
			Response: response{
				Status:        writer.status,
				ContentLength: writer.length,
				ResponseTime:  duration.Seconds() * 1000},
			Remote: remote{
				Addr:      raddr,
				User:      "-",
				UserAgent: r.UserAgent()}})
	})
}
