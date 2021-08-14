package domain

//go:generate mockgen -destination=../mocks/router.go -package=mocks github.com/wgarunap/xm-rest-api/domain Router
//go:generate mockgen -destination=../mocks/producer.go -package=mocks github.com/wgarunap/xm-rest-api/domain Producer
//go:generate mockgen -destination=../mocks/validator.go -package=mocks github.com/wgarunap/xm-rest-api/domain Validator
//go:generate mockgen -destination=../mocks/ipclient.go -package=mocks github.com/wgarunap/xm-rest-api/domain IpClient

import "net/http"

type Router interface {
	Route() http.Handler
}

type Producer interface {
	Produce(topic string, key, value []byte) error
}

type Validator interface {
	Validate(_struct interface{}) error
}

type IpClient interface {
	GetCountry(ip string) (country string, err error)
}
