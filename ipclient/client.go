package ipclient

import (
	"fmt"
	"github.com/wgarunap/xm-rest-api/domain"
	"io/ioutil"
	"net/http"
)

type client struct {
	client http.Client
}

func (c *client) GetCountry(ip string) (country string, err error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(`https://ipapi.co/%s/country`, ip), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.3")
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func NewIpCountryClient() domain.IpClient {
	return &client{
		client: http.Client{},
	}
}
