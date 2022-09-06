package router

import (
	"github.com/gorilla/mux"
	"github.com/martinustrinc/BE_Test/Shopping/Controller"
	"net/http"
)

func (server *Server) ApiController() {
	handler := mux.NewRouter()

	//handler.HandleFunc(setPath("/users/signup"), ).Methods("POST", "OPTIONS")
	//handler.HandleFunc(setPath("/users/login"), ).Methods("POST", "OPTIONS")
	//handler.HandleFunc(setPath("/product/insert"), Controller.AddProduct()).Methods("POST", "OPTIONS")
	handler.HandleFunc(setPath("/products"), Controller.GetProduct()).Methods("GET", "OPTIONS")
	handler.HandleFunc(setPath("/product/search"), Controller.GetProductByID()).Methods("GET", "OPTIONS")


	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
	//handler.Use(Middleware)
	//fmt.Println(http.ListenAndServe(":"+strconv.Itoa(port), handler))
}

func setPath(path string) string {
	prefixPath := "app"
	return prefixPath + path
}
