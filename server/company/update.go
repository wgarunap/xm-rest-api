package company

import (
	"encoding/json"
	"github.com/wgarunap/xm-rest-api/domain"
	"github.com/wgarunap/xm-rest-api/domain/repository"
	"github.com/wgarunap/xm-rest-api/server/response"
	"net/http"
)

type UpdateRequest struct {
	Name    string `json:"name" validate:"required"`
	Code    string `json:"code" validate:"required"`
	Country string `json:"country" validate:"required"`
	Website string `json:"website"`
	Phone   int    `json:"phone" validate:"required"`
}

type UpdateResponse struct {
	Code string `json:"code"`
}

type Update struct {
	Validator domain.Validator
	Producer  domain.Producer
	DB        repository.Company
}

func (c Update) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := UpdateRequest{}

	err := decodeBody(r, &req)
	if err != nil {
		response.ErrEncoder(w, response.Error{
			Code:         http.StatusBadRequest,
			Mgs:          "unable to decode the request body",
			AppErrorCode: 4000,
			Error:        err,
		})
		return
	}

	err = c.Validator.Validate(req)
	if err != nil {
		response.ErrEncoder(w, response.Error{
			Code:         http.StatusBadRequest,
			Mgs:          "request body data validation error",
			AppErrorCode: 4001,
			Error:        err,
		})
		return
	}

	data, err := json.Marshal(req)
	if err != nil {
		response.ErrEncoder(w, response.Error{
			Code:         http.StatusInternalServerError,
			Mgs:          "request body data validation error",
			AppErrorCode: 4002,
			Error:        err,
		})
		return
	}

	err = c.Producer.Produce(domain.TopicCompany, []byte(req.Code), data)
	if err != nil {
		response.ErrEncoder(w, response.Error{
			Code:         http.StatusInternalServerError,
			Mgs:          "request body data validation error",
			AppErrorCode: 4004,
			Error:        err,
		})
		return
	}

	err = c.DB.Update(domain.Company{
		Name:    req.Name,
		Code:    req.Code,
		Country: req.Country,
		Website: req.Website,
		Phone:   req.Phone,
	})
	if err != nil {
		response.ErrEncoder(w, response.Error{
			Code:         http.StatusInternalServerError,
			Mgs:          "error creating database entry",
			AppErrorCode: 4003,
			Error:        err,
		})
		return
	}

	response.Encoder(w, response.Response{
		Body:   UpdateResponse{req.Code},
		Code:   http.StatusCreated,
		Header: response.WithHeader(response.ContentTypeApplicationJson),
	})
}
