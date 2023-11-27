package rnsgo

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

const (
	// RNSRestApiUrl is the base url for the RNS REST API
	RNSRestApiUrl = "https://rns.rest"
	// RNSRestApiHeaderName is the header name for the RNS REST API
	RNSRestApiHeaderName = "X-LIB"
	// RNSRestApiHeaderNameValue is the header value for the RNS REST API
	RNSRestApiHeaderNameValue = "rns.go lib"
	// RNSLabelSuffix is the suffix for RNS labels
	RNSLabelSuffix = ".ron"
	// RoninAddressPrefix is the prefix for Ronin addresses
	RoninAddressPrefix = "ronin:"
	// EthAddressPrefix is the prefix for Ethereum addresses
	EthAddressPrefix = "0x"

	// ErrInvalidName Error Invalid Name
	ErrInvalidName = "invalid_name"
	// ErrInvalidAddress Error Invalid Address
	ErrInvalidAddress = "invalid_address"
	// ErrInvalidBatchRequest Error Invalid Batch Request
	ErrInvalidBatchRequest = "invalid_batch_request"
)

func NewClient(options ...string) RNS {

	host := RNSRestApiUrl

	if len(options) > 0 {
		host = options[0]
	}

	return RNS{
		client: resty.
			New().
			SetBaseURL(host).
			SetHeader(RNSRestApiHeaderName, RNSRestApiHeaderNameValue),
	}
}

func (r *RNS) GetAddress(name Name) (Address, error) {

	if !name.Valid() {
		return "", errors.New(ErrInvalidName)
	}

	var response struct {
		Address Address `json:"address"`
	}

	_, err := r.client.R().
		SetResult(&response).
		SetPathParams(map[string]string{
			"name": name.String(),
		}).
		Get("/resolve/{name}")

	if err != nil {
		return "", err
	}

	return response.Address, nil

}

func (r *RNS) GetName(address Address) (Name, error) {

	if !address.Valid() {
		return "", errors.New(ErrInvalidAddress)
	}

	var response struct {
		Name Name `json:"name"`
	}

	_, err := r.client.R().
		SetResult(&response).
		SetPathParams(map[string]string{
			"address": address.String(),
		}).
		Get("/lookup/{address}")

	if err != nil {
		return "", err
	}

	return response.Name, nil
}

func (r *RNS) GetNameBatch(addresses []Address) (BatchNameResponse, error) {
	var response BatchNameResponse

	for _, address := range addresses {
		if !address.Valid() {
			return nil, errors.New(ErrInvalidBatchRequest)
		}

	}

	_, err := r.client.R().
		SetResult(&response).
		SetBody(BatchAddressRequest{Addresses: addresses}).
		Post("/batch/lookup")

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *RNS) GetAddressBatch(names []Name) (BatchAddressResponse, error) {

	for _, name := range names {
		if !name.Valid() {
			return nil, errors.New(ErrInvalidBatchRequest)
		}
	}

	var response BatchAddressResponse

	_, err := r.client.R().
		SetResult(&response).
		SetBody(BatchNameRequest{Names: names}).
		Post("/batch/resolve")

	if err != nil {
		return nil, err
	}

	return response, nil
}
