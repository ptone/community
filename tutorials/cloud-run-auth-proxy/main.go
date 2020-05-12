package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var targetService string
var oauthClient string

func init() {
	value, exists := os.LookupEnv("TARGET_SERVICE")
	if !exists {
		panic("You need to set a TARGET_SERVICE env var for the service behind the proxy")
	}
	targetService = value
	value, exists = os.LookupEnv("OAUTH_CLIENT_ID")
	if !exists {
		panic("You need to set a OAUTH_CLIENT_ID env var for the Google login flow")
	}
	oauthClient = value
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	info := struct {
		OauthClient string
	}{
		oauthClient,
	}
	t, _ := template.ParseFiles("login.html")
	t.Execute(w, info)
}

func main() {
	service := targetService
	proxy, _ := getProxy(service)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/__hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})
	r.Get("/_login*", loginPage)
	r.Post("/_login", verify)
	r.With(authGuard).Mount("/", proxy)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)

}
