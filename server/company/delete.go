package company

import (
	"github.com/wgarunap/xm-rest-api/domain"
	"github.com/wgarunap/xm-rest-api/domain/repository"
	"github.com/wgarunap/xm-rest-api/server/response"
	"net/http"
)

type DeleteResponse struct {
	DeletedCode []string `json:"deleted_codes"`
	FailedCode  []Failed `json:"failed_codes"`
}

type Failed struct {
	Code string `json:"code"`
	Err  error  `json:"err"`
}

type Delete struct {
	//Validator domain.Validator
	Producer  domain.Producer
	DB        repository.Company
}

func (c Delete) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryies := r.URL.Query()

	var filters []repository.Filter
	for field, values := range queryies {
		for _, v := range values {
			filters = append(filters, repository.Filter{FieldName: field, Value: v})
		}
	}

	res := DeleteResponse{}


	err := c.DB.Delete(filters...)
	if err != nil {
		response.ErrEncoder(w, response.Error{
			Code:         http.StatusInternalServerError,
			Mgs:          "error deleting database entry",
			AppErrorCode: 4003,
			Error:        err,
		})
		return
	}
	//
	//err := c.Producer.Produce(domain.TopicCompany, []byte(code), nil)
	//if err != nil {
	//	res.FailedCode = append(res.FailedCode, Failed{
	//		Code: code,
	//		Err:  err,
	//	})
	//	continue
	//}
	//
	//res.DeletedCode = append(res.DeletedCode, code)

	response.Encoder(w, response.Response{
		Body:   res,
		Code:   http.StatusOK,
		Header: response.WithHeader(response.ContentTypeApplicationJson),
	})

}
