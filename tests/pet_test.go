package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ROOT_URL                 = "https://petstore.swagger.io"
	FIND_PET_BY_STATUS_ROUTE = "/v2/pet/findByStatus?status="
	PET_PER_ID               = "/v2/pet/"
)

var client = &http.Client{}

func TestFindPetByStatusAvailableReturnStatusOk(t *testing.T) {
	s := "available"
	req, _ := http.NewRequest("GET", ROOT_URL+FIND_PET_BY_STATUS_ROUTE+s, nil)
	res, _ := client.Do(req)
	assert.Equal(t, res.StatusCode, 200)
}
func TestFindByStatusAvailableBodyIsNotEmpty(t *testing.T) {
	s := "available"
	req, _ := http.NewRequest("GET", ROOT_URL+FIND_PET_BY_STATUS_ROUTE+s, nil)
	res, _ := client.Do(req)
	assert.NotEmpty(t, res.Body)
}

func TestFindPetByStatusPendingeReturnStatusOk(t *testing.T) {
	s := "pending"
	req, _ := http.NewRequest("GET", ROOT_URL+FIND_PET_BY_STATUS_ROUTE+s, nil)
	res, _ := client.Do(req)
	assert.Equal(t, res.StatusCode, 200)
}
func TestFindByStatusPendingBodyIsNotEmpty(t *testing.T) {
	s := "pending"
	req, _ := http.NewRequest("GET", ROOT_URL+FIND_PET_BY_STATUS_ROUTE+s, nil)
	res, _ := client.Do(req)
	assert.NotEmpty(t, res.Body)
}

func TestFindPetByStatusSoldeReturnStatusOk(t *testing.T) {
	s := "sold"
	req, _ := http.NewRequest("GET", ROOT_URL+FIND_PET_BY_STATUS_ROUTE+s, nil)
	res, _ := client.Do(req)
	assert.Equal(t, res.StatusCode, 200)
}
func TestFindByStatusSoldBodyIsNotEmpty(t *testing.T) {
	s := "sold"
	req, _ := http.NewRequest("GET", ROOT_URL+FIND_PET_BY_STATUS_ROUTE+s, nil)
	res, _ := client.Do(req)
	assert.NotEmpty(t, res.Body)
}

func TestFindPetByIdReturnPetNotFound(t *testing.T) {
	s := "0"
	req, _ := http.NewRequest("GET", ROOT_URL+PET_PER_ID+s, nil)
	res, _ := client.Do(req)
	assert.Equal(t, res.StatusCode, 404)
}

func TestAddNewPetToStoreReturnsStatusOk(t *testing.T) {
	body := &Pet{
		Id:     12345678,
		Name:   "Fiddo",
		Status: "available",
	}
	payloadBuffer := new(bytes.Buffer)
	json.NewEncoder(payloadBuffer).Encode(body)
	req, _ := http.NewRequest("POST", ROOT_URL+PET_PER_ID, payloadBuffer)
	req.Header.Add("Content-Type", "application/json")
	res, _ := client.Do(req)
	assert.Equal(t, res.StatusCode, 200)
}
func TestAddNewPetToStoreReturnsPayload(t *testing.T) {
	body := &Pet{
		Id:     12345678,
		Name:   "Fiddo",
		Status: "available",
	}
	payloadBuffer := new(bytes.Buffer)
	json.NewEncoder(payloadBuffer).Encode(body)
	req, _ := http.NewRequest("POST", ROOT_URL+PET_PER_ID, payloadBuffer)
	req.Header.Add("Content-Type", "application/json")
	res, _ := client.Do(req)
	assert.NotEmpty(t, res.Body)
}
