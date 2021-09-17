package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/soguazu/clean_arch/services"
)

type carDetailsController struct{}

var (
	carDetailsService services.CarDetailsService
)

type CarDetailsController interface {
	GetCarDetails(rw http.ResponseWriter, r *http.Request)
}

func NewCarDetailsControllers(service services.CarDetailsService) CarDetailsController {
	carDetailsService = service
	return &carDetailsController{}
}

func (*carDetailsController) GetCarDetails(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	result := carDetailsService.GetDetails()
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(result)
}
