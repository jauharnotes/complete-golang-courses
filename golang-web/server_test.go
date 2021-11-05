package golang_web

import (
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


func TestServer(t *testing.T) {
	//var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer,"Test Server")
	//}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"Home")
	})
	mux.HandleFunc("/server1", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"Server1")
	})
	mux.HandleFunc("/server2", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"Server2")
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
