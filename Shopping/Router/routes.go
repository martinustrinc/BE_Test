package Router

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/martinustrinc/BE_Test/Shopping/Controller"
	"net/http"
	"strconv"
)

func ApiController(incomingRoute *Server) {
	handler := mux.NewRouter()

	//handler.HandleFunc(setPath("/users/signup"), ).Methods("POST", "OPTIONS")
	//handler.HandleFunc(setPath("/users/login"), ).Methods("POST", "OPTIONS")
	handler.HandleFunc(setPath("/product/addproduct"), Controller.AddProduct()).Methods("POST", "OPTIONS")
	handler.HandleFunc(setPath("/product/getproduct"), ).Methods("POST", "OPTIONS")
	handler.HandleFunc(setPath("/product/search"), ).Methods("POST", "OPTIONS")


	handler.Use(Middleware)
	fmt.Println(http.ListenAndServe(":"+strconv.Itoa(port), handler))
}

func setPath(path string) string {
	prefixPath := "app"
	return prefixPath + path
}
