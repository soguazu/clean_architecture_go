package services

import (
	"fmt"
	"net/http"
)

type CarService interface {
	FetchData()
}

const (
	carServiceUrl = "https://myfakeapi.com/api/cars/1"
)

type fetchCarDataService struct{}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	fmt.Println("Fetching car data")

	resp, _ := client.Get(carServiceUrl)

	carDataChannel <- resp
}
