package main

import (
	"io"
	"log"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")

	log.Println(req.FormValue("param1"))

	io.WriteString(res,
		`<doctype html>
<html>
	<head>
    	<title>Hello World!</title>
	</head>
	<body>
    	Hello World!
	</body>
</html>`)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
