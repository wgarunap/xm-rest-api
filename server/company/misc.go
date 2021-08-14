package company

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func decodeBody(r *http.Request, req interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, req)
	if err != nil {
		return err
	}
	return nil
}

