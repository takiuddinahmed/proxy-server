package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	url2 "net/url"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method)
		switch request.Method {
		case "GET":
			http.ServeFile(writer, request, "./src/templates/index.html")
		case "POST":
			url := request.FormValue("url")
			if len(url) > 0 {
				targetUrl, err := url2.Parse(url)
				if err != nil {
					log.Fatal(err)
				}
				reverseProxy := httputil.NewSingleHostReverseProxy(targetUrl)
				reverseProxy.ServeHTTP(writer, request)
			} else {
				fmt.Println("Invalid")
			}
		default:
			fmt.Println("NOT Implemented")
		}

	})

	http.ListenAndServe(":8080", nil)
}
