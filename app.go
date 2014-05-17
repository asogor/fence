package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/asogor/fence/route"
	"github.com/asogor/fence/util"
)

func main() {
	var data string
	flag.StringVar(&data, "data", "data", "location of the data dir")
	flag.Parse()
	
	secret_config := data + "/config/secret"
	fmt.Println("Start loading:", secret_config)
	secretdata, err := ioutil.ReadFile(secret_config)
	if (err != nil) {
		panic("Failed to load secret >>" + secret_config) 
	}
	
	var secret AppSecret 
	err = json.Unmarshal(secretdata,&secret)
	if (err != nil) {
		panic(err) 
	}
	fmt.Println("Secret Loaded: ",secret.FBClientId)
	http.ListenAndServe(":8000",registerHandlers(util.NewConfig(data)))
}

type AppSecret struct {
	FBClientId string
}

func registerHandlers(c util.Config)(mux *http.ServeMux){
	mux = http.NewServeMux()
	mux.Handle("/login",route.NewLogin(c))
	return mux
}
