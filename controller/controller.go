package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController         home
	shopController         shop
	standLocatorController standLocator
)

func Startup(template map[string]*template.Template) {
	homeController.homeTemplate = template["home.html"]
	homeController.loginTemplate = template["login.html"]
	standLocatorController.standLocatorTemplate = template["stand_locator.html"]
	shopController.shopTemplate = template["shop.html"]
	shopController.categoryTemplate = template["shop_details.html"]
	homeController.registerRoutes()
	shopController.registerRoutes()
	standLocatorController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
