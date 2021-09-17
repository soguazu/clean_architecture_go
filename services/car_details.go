package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/soguazu/clean_arch/entity"
)

type CarDetailsService interface {
	GetDetails() entity.CarDetails
}

var (
	carService       CarService   = NewCarService()
	ownerService     OwnerService = NewOwnerService()
	carDataChannel                = make(chan *http.Response)
	ownerDataChannel              = make(chan *http.Response)
)

type carDetailService struct{}

func NewCarDetailsService() CarDetailsService {
	return &carDetailService{}
}

func (*carDetailService) GetDetails() entity.CarDetails {
	go carService.FetchData()
	go ownerService.FetchData()

	car, _ := GetCarData()
	owner, _ := GetOwnerData()

	return entity.CarDetails{
		ID:        1,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}
}

func GetCarData() (entity.Car, error) {
	carResp := <-carDataChannel
	var car entity.Car
	err := json.NewDecoder(carResp.Body).Decode(&car)
	if err != nil {
		log.Println(err.Error())
		return car, err
	}
	return car, nil
}

func GetOwnerData() (entity.Owner, error) {
	ownerResp := <-ownerDataChannel
	var owner entity.Owner
	err := json.NewDecoder(ownerResp.Body).Decode(&owner)
	if err != nil {
		log.Println(err.Error())
		return owner, err
	}
	return owner, nil
}
