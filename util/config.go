package util

import (
	"log"
	"html/template"
	"os"
)

type TemplateType string 

type config struct{
	log *log.Logger
	dataDir string
}

type Config interface {
	Template(t TemplateType) (template *template.Template)
	Log()(log *log.Logger)
}

func (c *config) Log() (log *log.Logger) {
	return c.log
}

func (c *config) Template(ttype TemplateType) (t *template.Template) {
	log.Println("Load Template", t)
	t, err := template.ParseFiles(c.dataDir + "/view/" + string(ttype) + ".html")
	if(err != nil) {
		log.Fatal("Template loading failed", err)
	}
	return t
}

func NewConfig(data string)(c Config){
	l := log.New(os.Stdout,"FENCE",log.LstdFlags)
	return &config{log:l,dataDir:data}
}
	