package main

import (
	"flag"
	"github.com/asogor/fence/route"
	"github.com/asogor/fence/util"
	"net/http"
)

func main() {
	var data string
	flag.StringVar(&data, "data", "data", "location of the data dir")
	flag.Parse()
	var config = util.NewConfig(data)

	http.ListenAndServe(":8000", registerHandlers(config))
}

func registerHandlers(c util.Config) (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.Handle("/login/start", route.NewLogin(c))
	return mux
}
