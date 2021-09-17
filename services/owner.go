package services

import (
	"fmt"
	"net/http"
)

type OwnerService interface {
	FetchData()
}

const (
	ownerServiceUrl = "https://myfakeapi.com/api/users/1"
)

type fetchOwnerDataService struct{}

func NewOwnerService() OwnerService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}

	fmt.Println("Fetching owner data")

	resp, _ := client.Get(ownerServiceUrl)

	ownerDataChannel <- resp
}
