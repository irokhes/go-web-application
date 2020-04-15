package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
	shopController shop
)

func Startup(template map[string]*template.Template) {
	homeController.homeTemplate = template["home.html"]
	shopController.shopTemplate = template["shop.html"]
	shopController.categoryTemplate = template["shop_details.html"]
	homeController.registerRoutes()
	shopController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
