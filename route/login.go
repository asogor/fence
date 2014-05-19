package route

import (
	"net/http"
	"github.com/asogor/fence/util"
)

type login struct {
	config util.Config
}

type LoginContext struct {
	Title string
	FbAppId string
	RedirectURL string
}

const TEMPLATE util.TemplateType = "login"

func NewLogin(c util.Config) http.Handler {
	return &login{config:c}
}

func (l *login) ServeHTTP(w http.ResponseWriter, r *http.Request){
	l.config.Template(TEMPLATE).Execute(w,LoginContext{Title:"NonSense",FbAppId:"313331798817136",RedirectURL:"http://dev.bluefung.us:8000/fblogin"})
}

