package response

import (
	"encoding/json"
	"fmt"
	"github.com/tryfix/log"
	"net/http"
)

type header struct {
	Key   string
	Value []string
}

// NewHeader returns a pointer to header with given key and values which can be passed to
// http response
func NewHeader(key string, values ...string) *header {
	return &header{
		Key:   key,
		Value: values,
	}
}

func WithHeader(h ...*header) (out []*header) {
	for _, t := range h {
		out = append(out, t)
	}
	return out
}

var (
	ContentTypeApplicationJson = &header{
		Key:   "Content-Type",
		Value: []string{"application/json"},
	}

	ContentTypeText = &header{
		Key:   "Content-Type",
		Value: []string{"text/plain; charset=utf-8"},
	}
)

type Response struct {
	Body   interface{}
	Code   int
	Header []*header
}

// Encoder encode the response
func Encoder(w http.ResponseWriter, res Response) {
	if len(res.Header) > 0 {
		for _, v := range res.Header {
			for _, v2 := range v.Value {
				w.Header().Add(v.Key, v2)
			}
		}
	}

	switch res.Code {
	case 0:
	default:
		w.WriteHeader(res.Code)
	}

	if res.Body != nil {
		if w.Header().Get("Content-Type") == "application/json" {
			err := json.NewEncoder(w).Encode(res.Body)
			if err != nil {
				log.Error(`error encoding response`, err)
			}
			return
		}
		_, err := fmt.Fprintln(w, res.Body)
		if err != nil {
			log.Error(`error writing response`, err)
		}
		return

	}

}
