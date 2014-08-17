package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

var server *gin.Engine
var templates map[string]*template.Template

func init() {
	loadTemplates()
}

func main() {

	server = gin.Default()

	server.Static("/public/css/", "./public/css")
	server.Static("/public/js/", "./public/js/")
	server.Static("/public/fonts/", "./public/fonts/")
	server.Static("/public/img/", "./public/img/")

	server.GET("/", IndexRouter)
	server.GET("/about", AboutRoute)
	server.GET("/contact", ContactRoute)
	server.GET("/signin", SigninRoute)
	server.GET("/signup", SignupRoute)

	server.Run(":3000")
}

func IndexRouter(g *gin.Context) {
	server.SetHTMLTemplate(templates["index"])
	g.HTML(200, "_base.html", nil)
}

func AboutRoute(g *gin.Context) {
	server.SetHTMLTemplate(templates["about"])
	g.HTML(200, "_base.html", nil)
}

func ContactRoute(g *gin.Context) {
	server.SetHTMLTemplate(templates["contact"])
	g.HTML(200, "_base.html", nil)
}

func SigninRoute(g *gin.Context) {
	server.SetHTMLTemplate(templates["signin"])
	g.HTML(200, "_base.html", nil)
}

func SignupRoute(g *gin.Context) {
	server.SetHTMLTemplate(templates["signup"])
	g.HTML(200, "_base.html", nil)
}

func loadTemplates(){
	var baseTemplate = "templates/layout/_base.html"
	templates = make(map[string]*template.Template)

	templates["index"] = template.Must(template.ParseFiles(baseTemplate, "templates/home/index.html",))
	templates["about"] = template.Must(template.ParseFiles(baseTemplate, "templates/home/about.html",))
	templates["contact"] = template.Must(template.ParseFiles(baseTemplate, "templates/home/contact.html",))
	templates["signin"] = template.Must(template.ParseFiles(baseTemplate, "templates/account/signin.html",))
	templates["signup"] = template.Must(template.ParseFiles(baseTemplate, "templates/account/signup.html",))
}
