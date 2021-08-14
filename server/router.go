package server

import (
	"github.com/gorilla/mux"
	"github.com/wgarunap/xm-rest-api/config"
	"github.com/wgarunap/xm-rest-api/domain"
	"github.com/wgarunap/xm-rest-api/ipclient"
	"github.com/wgarunap/xm-rest-api/repositories/mysql"
	"github.com/wgarunap/xm-rest-api/server/company"
	"github.com/wgarunap/xm-rest-api/server/middlewares"
	"github.com/wgarunap/xm-rest-api/stream/producer"
	"github.com/wgarunap/xm-rest-api/validator"
	"net/http"
)

type Router struct {
	cfg *config.Conf
}

func (ro *Router) Route() http.Handler {
	r := mux.NewRouter()

	database := mysql.NewRamDBInstance(ro.cfg)
	validateOb := validator.NewValidator()
	prodcuer := producer.NewProducer(ro.cfg)
	ipcli := ipclient.NewIpCountryClient()

	r.Handle("/company", middlewares.IpCheck{
		IpClient: ipcli,
		Next: company.Create{
			Validator: validateOb,
			Producer:  prodcuer,
			DB:        database,
		},
	}).Methods(http.MethodPut)

	r.Handle("/company", middlewares.IpCheck{
		IpClient: ipcli,
		Next: company.Update{
			Validator: validateOb,
			Producer:  prodcuer,
			DB:        database,
		},
	}).Methods(http.MethodPost)

	r.Handle("/company", middlewares.IpCheck{
		IpClient: ipcli,
		Next: company.Delete{
			Producer: prodcuer,
			DB:       database,
		}}).Methods(http.MethodDelete)

	r.Handle("/company", company.Read{
		DB: database,
	}).Methods(http.MethodGet)

	return r
}

func NewRouter(conf *config.Conf) domain.Router {
	return &Router{conf}
}
