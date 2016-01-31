# gomolog
MongoDB backed Golang net/http logger

## How to use?

```
$ go get github.com/toyoshim/gomolog
```

*web.go*
```go
package main

import (
  "net/http"
  "os"
  
  "github.com/toyoshim/gomolog"
)

func main() {
  log := gomolog.Open(os.Getenv("MONGOLAB_URI"), "log")
  defer log.Close()
  
  http.Handle("/", http.FileServer(http.Dir("static")))
  err := http.ListenAndServe(":"+os.Getenv("PORT"), log.Logger())
  if err != nil {
    panic(err)
  }
}
```

## Log format example
```json
{
    "_id": {
        "$oid": "56aa1e536c00944eaa237329"
    },
    "format": 1,
    "date": "2016-01-28T13:57:39Z",
    "referrer": "",
    "request": {
        "method": "GET",
        "host": "crlogo.herokuapp.com",
        "url": "/mozapp/chrome_logo.html",
        "protocol": "HTTP/1.1",
        "acceptLanguage": "ja,en-US;q=0.8,en;q=0.6"
    },
    "response": {
        "status": 200,
        "contentLength": 15194,
        "responseTime": 4.267561000000001
    },
    "remote": {
        "addr": "10.71.199.148:39047",
        "user": "-",
        "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.106 Safari/537.36"
    }
}
```

## For other languages

- Express / Node.js: [momolog](https://github.com/toyoshim/momolog)
- Rack / Ruby: [ramolog](https://github.com/toyoshim/ramolog)
