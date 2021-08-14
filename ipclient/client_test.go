package ipclient

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIpCountryClient(t *testing.T) {
	c := NewIpCountryClient()
	country, err := c.GetCountry("31.153.207.255")
	require.Equal(t, "CY", country)
	require.Equal(t, nil, err)
}
