package restapi

import (
	"log"
	"ockernuts/goroborally/handlers"
	"net/http"
	"strings"
)




func otherPathHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("method: ", r.Method, "url:", r.URL.Path, "accessed")
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/index.html", http.StatusFound)
			return
		}
		if r.URL.Path == "/index.html" {
			handlers.DefaultPageHandler(w, r)
			return
		}
		if strings.Index(r.URL.Path, "/api/swagger.json") == 0 {
			handlers.GetSwaggerJson(w, r)
			return
		}
		if strings.Index(r.URL.Path, "/static") == 0 {
			http.StripPrefix("/static", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
