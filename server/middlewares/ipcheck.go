package middlewares

import (
	"github.com/wgarunap/xm-rest-api/domain"
	"github.com/wgarunap/xm-rest-api/server/response"
	"net/http"
	"strings"
)

type IpCheck struct {
	IpClient domain.IpClient
	Next     http.Handler
}

func (m IpCheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)

	country, err := m.IpClient.GetCountry(ip)
	if err != nil {
		response.ErrEncoder(w, response.Error{
			Code:         http.StatusInternalServerError,
			Mgs:          "adaptor error, unable to retrieve clients origin country",
			AppErrorCode: 4001,
			//Error:        err,
		})
		return
	}

	if country != domain.CyprusCountryCode {
		response.ErrEncoder(w, response.Error{
			Code:         http.StatusForbidden,
			Mgs:          "invalid request origin country",
			AppErrorCode: 4002,
			//Error:        err,
		})
		return
	}

	m.Next.ServeHTTP(w, r)
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		ips := strings.Split(forwarded, ",")
		clientIp := strings.TrimSpace(ips[0])
		return clientIp
	}

	ip := strings.Split(r.RemoteAddr, ":")
	return ip[0]
}
