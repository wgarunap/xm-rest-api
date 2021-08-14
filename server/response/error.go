package response

import (
	"encoding/json"
	"fmt"
	"github.com/tryfix/log"
	"net/http"
	"strings"
)

// Error is the error message response object which will be sent
// to the client on event of any error
type Error struct {
	// Error code to be send to the client
	// 		ex:- http.StatusOk
	Code int `json:"-"`

	// Mgs is the descriptive error message sent to the client
	Mgs string `json:"message"`

	// AppErrorCode is the application specific error code which is
	// documented against all the possible errors
	AppErrorCode int `json:"app_error_code"`

	// Actual error occured
	Error error `json:"-"`
}

// Encoder encode the error response which sends to the client
func ErrEncoder(w http.ResponseWriter, e Error) {
	w.Header().Set(ContentTypeApplicationJson.Key,
		strings.Join(ContentTypeApplicationJson.Value, ","))

	w.WriteHeader(e.Code)

	err := json.NewEncoder(w).Encode(e)
	if err != nil {
		log.Error(`error encoding httperror`, err)
	}
	if e.Error != nil {
		fmt.Println(e.Error)
	}

}
