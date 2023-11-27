package rnsgo

import (
	"github.com/go-resty/resty/v2"
	"strings"
)

type RNS struct {
	client *resty.Client
}

type Name string

func (n Name) String() string {
	return string(n)
}

func (n Name) Valid() bool {
	return strings.HasSuffix(string(n), RNSLabelSuffix) && len(strings.TrimSuffix(string(n), RNSLabelSuffix)) >= 3
}

type Address string

func (a Address) String() string {
	return string(a)
}

func (a Address) Normalize() Address {
	return Address(strings.Replace(string(a), RoninAddressPrefix, EthAddressPrefix, 1))
}

func (a Address) Valid() bool {
	return strings.HasPrefix(string(a), EthAddressPrefix) && len(string(a)) == 42
}

type BatchRequest interface {
	Route() string
}

type BatchAddressRequest struct {
	Addresses []Address `json:"addresses"`
}

func (b *BatchAddressRequest) Route() string {
	return "/batch/lookup"
}

type BatchNameRequest struct {
	Names []Name `json:"names"`
}

func (b *BatchNameRequest) Route() string {
	return "/batch/resolve"
}

type BatchAddressResponse map[Name]*Address

type BatchNameResponse map[Address]*Name
