package rnsgo

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	TestRNSName      Name    = "jihoz.ron"
	TestRoninAddress Address = "ronin:2d62c27ce2e9e66bb8a667ce1b60f7cb02fa9810"
	TestEthAddress   Address = "0x2d62c27ce2e9e66bb8a667ce1b60f7cb02fa9810"
)

// TestNewClient function tests the NewClient function with a real API call
func TestNewClient(t *testing.T) {
	client := NewClient()

	// You can add tests here to ensure that the client is correctly configured
	assert.NotNil(t, client, "Client should not be nil")
}

// TestGetAddress function tests the GetAddress method with a real API call
func TestGetAddress(t *testing.T) {
	client := NewClient()

	address, err := client.GetAddress(TestRNSName)

	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, address, "Address should not be empty")
	assert.Equal(t, strings.ToLower(TestEthAddress.String()), strings.ToLower(address.String()), "Address should be equal to the expected value")
}

// TestGetName function tests the GetName method with a real API call
func TestGetNameInvalidPrefix(t *testing.T) {
	client := NewClient()

	_, err := client.GetName(TestRoninAddress)

	assert.NotNil(t, err, "Error should not! be nil")
}

// TestGetName function tests the GetName method with a real API call
func TestAddress_Normalize(t *testing.T) {
	assert.Equal(t, TestEthAddress, TestRoninAddress.Normalize(), "Address should be equal to the expected value")
}

// TestGetName function tests the GetName method with a real API call
func TestAddress_Valid(t *testing.T) {
	assert.False(t, TestRoninAddress.Valid(), "Address should not be valid")
	assert.True(t, TestEthAddress.Valid(), "Address should be valid")
}

// TestGetName function tests the GetName method with a real API call
func TestGetName(t *testing.T) {
	client := NewClient()

	name, err := client.GetName(TestEthAddress)

	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, name, "Name should not be empty")
	assert.Equal(t, strings.ToLower(TestRNSName.String()), strings.ToLower(name.String()), "Name should be equal to the expected value")
}

func TestRNS_GetNameBatch(t *testing.T) {
	client := NewClient()

	names, err := client.GetNameBatch([]Address{TestEthAddress})

	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, names, "Names should not be empty")
	assert.Len(t, names, 1, "Names should have a length of 1")
	assert.Equal(t, strings.ToLower(TestRNSName.String()), strings.ToLower(names[TestEthAddress].String()), "Name should be equal to the expected value")
}

func TestRNS_GetAddressBatch(t *testing.T) {
	client := NewClient()

	addresses, err := client.GetAddressBatch([]Name{TestRNSName})

	assert.Nil(t, err, "Error should be nil")
	assert.NotEmpty(t, addresses, "Addresses should not be empty")
	assert.Len(t, addresses, 1, "Addresses should have a length of 1")
	assert.Equal(t, strings.ToLower(TestEthAddress.String()), strings.ToLower(addresses[TestRNSName].String()), "Address should be equal to the expected value")
}
