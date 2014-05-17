package route

import (
	"net/http"
	"github.com/asogor/fence/util"
)

type login struct {
	config util.Config
}

const TEMPLATE util.TemplateType = "login"

func NewLogin(c util.Config) http.Handler {
	return &login{config:c}
}

func (l *login) ServeHTTP(w http.ResponseWriter, r *http.Request){
	l.config.Template(TEMPLATE).Execute(w,nil)
}
