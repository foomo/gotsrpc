package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/foomo/gotsrpc/demo"
)

type Demo struct {
	proxy *demo.DemoGoTSRPCProxy
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
		proxy: demo.NewDefaultDemoGoTSRPCProxy(&demo.Demo{}, []string{}),
	}
	fmt.Println("staring a demo on http://127.0.0.1:8080 - open it and take a look at the console")
	fmt.Println(http.ListenAndServe(":8080", d))
}
