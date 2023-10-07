package router

import (
	"fmt"
	"golang-restapi-noframework/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController *controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Welcome Home")
	})

	router.GET("/api/product", productController.GetAll)
	router.GET("/api/product/:productId", productController.GetById)
	router.POST("/api/product", productController.Create)
	router.PATCH("/api/product/:productId", productController.Update)
	router.DELETE("/api/product/:productId", productController.Delete)

	return router
}
