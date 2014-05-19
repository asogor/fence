package util

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

type TemplateType string

type config struct {
	log     *log.Logger
	dataDir string
}

type Config interface {
	Template(t TemplateType) (template *template.Template)
	Log() (log *log.Logger)
}

func (c *config) Log() (log *log.Logger) {
	return c.log
}

func (c *config) Template(ttype TemplateType) (t *template.Template) {
	c.log.Println("Load Template", t)
	t, err := template.ParseFiles(c.dataDir + "/view/" + string(ttype) + ".html")
	if err != nil {
		c.log.Fatal("Template loading failed", err)
	}
	return t
}

func NewConfig(data string) (c Config) {
	secret_config := data + "/config/secret"

	l := log.New(os.Stdout, "FENCE ", log.LstdFlags)

	l.Println("Start loading:", secret_config)
	secretdata, err := ioutil.ReadFile(secret_config)
	if err != nil {
		l.Fatalln("Failed to load secret", secret_config)
	}

	var secret appSecret
	err = json.Unmarshal(secretdata, &secret)
	if err != nil {
		panic(err)
	}
	l.Println("Secret loaded", secret)

	return &config{log: l, dataDir: data}
}

type appSecret struct {
	FbAppId string
	FbRedirectUrl string
}
