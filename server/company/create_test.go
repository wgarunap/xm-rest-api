package company

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/wgarunap/xm-rest-api/domain"
	"github.com/wgarunap/xm-rest-api/mocks"
	"github.com/wgarunap/xm-rest-api/server/response"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestCreate_Success(t *testing.T) {
	testData := CreateRequest{
		Name:    "ABC2",
		Code:    "123123",
		Country: "SL",
		Website: "www.sl.com",
		Phone:   2342234,
	}
	b, _ := json.Marshal(testData)

	ctrl := gomock.NewController(t)

	repo := mocks.NewMockCompany(ctrl)
	repo.EXPECT().
		Create(gomock.Any()).
		Do(func(arg1 interface{}) {
			_, ok := arg1.(domain.Company)
			if !ok {
				t.Fail()
			}
		}).
		Return(nil).Times(1)

	validator := mocks.NewMockValidator(ctrl)
	validator.EXPECT().
		Validate(gomock.Any()).Return(nil).Times(1)

	prod := mocks.NewMockProducer(ctrl)
	prod.EXPECT().
		Produce(gomock.Any(), gomock.Eq([]byte(strconv.Itoa(123123))), gomock.Eq(b)).
		Return(nil).Times(1)

	createApi := Create{
		Validator: validator,
		Producer:  prod,
		DB:        repo,
	}

	respWriter := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/company", bytes.NewBuffer([]byte(b)))

	createApi.ServeHTTP(respWriter, req)
	require.Equal(t, http.StatusCreated, respWriter.Code)

	resp := CreateResponse{}
	dd, err := ioutil.ReadAll(respWriter.Body)
	_ = json.Unmarshal(dd, &resp)
	require.Equal(t, nil, err)
	require.Equal(t, testData.Code, resp.Code)

}

func TestCreate_ValidationFail(t *testing.T) {
	testData := CreateRequest{
		Name:    "ABC2",
		Code:    "123123",
		Country: "SL",
		Website: "www.sl.com",
		Phone:   2342234,
	}
	b, _ := json.Marshal(testData)

	ctrl := gomock.NewController(t)

	repo := mocks.NewMockCompany(ctrl)

	validator := mocks.NewMockValidator(ctrl)
	validator.EXPECT().
		Validate(gomock.Any()).Return(errors.New(`validation failed`)).Times(1)

	prod := mocks.NewMockProducer(ctrl)

	createApi := Create{
		Validator: validator,
		Producer:  prod,
		DB:        repo,
	}

	respWriter := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPut, "/company", bytes.NewBuffer([]byte(b)))

	createApi.ServeHTTP(respWriter, req)
	require.Equal(t, http.StatusBadRequest, respWriter.Code)

	resp := response.Error{}
	dd, err := ioutil.ReadAll(respWriter.Body)
	_ = json.Unmarshal(dd, &resp)
	require.Equal(t, nil, err)
	require.Equal(t, `request body data validation error`, resp.Mgs)
}
