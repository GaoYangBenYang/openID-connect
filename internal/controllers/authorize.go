package controllers

import (
	"net/http"
	"text/template"
	"github.com/astaxie/beego"
)

//授权
type Authorize struct {
	beego.Controller
}

func (a *Authorize) Authorize() {

}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	tpl := `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Login</title>
		</head>
		<body>
			<form method="POST" action="/login">
				<input name="client"/>
				<button type="submit">Login</button>
			</form>
		</body>
	</html>`
	t, err := template.New("login").Parse(tpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	client := r.FormValue("client")
	http.Redirect(w, r, "/authorize/callback?id="+client, http.StatusFound)
}
