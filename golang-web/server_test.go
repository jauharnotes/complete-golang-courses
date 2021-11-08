package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


func TestServer2(t *testing.T) {
	//var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer,"Test Server")
	//}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"Home")
	})
	mux.HandleFunc("/image", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"image")
	})
	mux.HandleFunc("/image/thumnail", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"thumnail")
	})
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServerRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}
	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}