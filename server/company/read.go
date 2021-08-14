package company

import (
	"github.com/wgarunap/xm-rest-api/domain/repository"
	"github.com/wgarunap/xm-rest-api/server/response"
	"net/http"
)

type Read struct {
	DB repository.Company
}

func (c Read) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryies := r.URL.Query()

	var filters []repository.Filter
	for field, values := range queryies {
		for _, v := range values {
			filters = append(filters, repository.Filter{FieldName: field, Value: v})
		}
	}

	companies, err := c.DB.Get(filters...)
	if err != nil {
		response.ErrEncoder(w, response.Error{
			Code:         http.StatusInternalServerError,
			Mgs:          "unable to retrive data",
			AppErrorCode: 4002,
			Error:        err,
		})
		return
	}

	response.Encoder(w, response.Response{
		Body:   companies,
		Code:   http.StatusOK,
		Header: response.WithHeader(response.ContentTypeApplicationJson),
	})
}
