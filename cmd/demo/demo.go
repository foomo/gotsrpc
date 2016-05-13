package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/foomo/gotsrpc/demo"
)

type Demo struct {
	proxy *demo.ServiceGoTSRPCProxy
}

func (d *Demo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/demo.js":
		serveFile("demo.js", w)
	case "/":
		serveFile("index.html", w)
	default:
		switch true {
		case strings.HasPrefix(r.URL.Path, "/service"):
			d.proxy.ServeHTTP(w, r)
			return
		}
	}
}

func serveFile(name string, w http.ResponseWriter) {
	index, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	w.Write(index)

}

func main() {
	d := &Demo{
		proxy: demo.NewServiceGoTSRPCProxy(&demo.Service{}, "/service"),
	}
	fmt.Println(http.ListenAndServe(":8080", d))
}
