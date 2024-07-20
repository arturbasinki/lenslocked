package controllers

import (
	"net/http"

	"github.com/arturbasinki/lenslocked/views"
)

type Static struct {
	Template views.Template
}

func (statis Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statis.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
