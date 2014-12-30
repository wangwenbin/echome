//[echome] is a echoServer which help you get your public ip & port.
//Service is available at http://echome.org
//In this way you can build app server on machine that have no public ip.
//At this time [echome] is a test server. Many code need to be write.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var serverAddr string = ":80"
var addrApi = []addrType{"/ip", "/port", "/addr", "/ip/", "/port/", "/addr/"}

func main() {
	for _, pattern := range addrApi {
		http.Handle(string(pattern), pattern)
	}
	http.HandleFunc("/debug", debugHandleFunc)
	http.HandleFunc("/", sayHello)

	fmt.Printf("Try http://localhost%s\nPress Ctrl+C to exit!\n",
		serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type addrType string

func (t addrType) ServeHTTP(rw http.ResponseWriter, requ *http.Request) {
	var re string
	switch t {
	case "/ip", "/ip/":
		re = strings.Split(requ.RemoteAddr, ":")[0]
	case "/port", "/port/":
		re = strings.Split(requ.RemoteAddr, ":")[1]
	case "/addr", "/addr/":
		re = requ.RemoteAddr
	default:
		re = "404 this api not exist!"
	}
	fmt.Fprintln(rw, re)
}

func debugHandleFunc(w http.ResponseWriter, requ *http.Request) {
	fmt.Fprintf(w, "%+v", *requ)
}

func sayHello(w http.ResponseWriter, requ *http.Request) {
	fmt.Fprintln(w, "Hello,this is just a echoserver!")
	who := requ.RemoteAddr
	fmt.Fprintf(w, "you are: [%v].\n", who)
	url := requ.Host + requ.RequestURI
	fmt.Fprintf(w, "your Request url is: [%v].\n", url)
	fmt.Fprintln(w, "The available GET api is:")
	fmt.Fprintln(w, addrApi)
}
