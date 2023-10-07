package controller

import (
	"golang-restapi-noframework/dto/request"
	"golang-restapi-noframework/dto/response"
	"golang-restapi-noframework/helper"
	"golang-restapi-noframework/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (controller *ProductController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	createProductRequest := request.CreateProductRequest{}
	helper.ReadRequestBody(requests, &createProductRequest)

	controller.ProductService.Create(requests.Context(), createProductRequest)
	apiResponse := response.ApiResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, apiResponse)
}

func (controller *ProductController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	updateProductRequest := request.UpdateProductRequest{}
	helper.ReadRequestBody(requests, &updateProductRequest)

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)
	updateProductRequest.Id = id

	controller.ProductService.Update(requests.Context(), updateProductRequest)
	apiResponse := response.ApiResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, apiResponse)

}

func (controller *ProductController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	controller.ProductService.Delete(requests.Context(), id)
	apiResponse := response.ApiResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, apiResponse)

}

func (controller *ProductController) GetAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.ProductService.GetAll(requests.Context())
	apiResponse := response.ApiResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, apiResponse)
}

func (controller *ProductController) GetById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	result := controller.ProductService.GetById(requests.Context(), id)
	apiResponse := response.ApiResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, apiResponse)

}
