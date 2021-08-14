package server

import (
	"fmt"
	"github.com/tryfix/log"
	"github.com/wgarunap/xm-rest-api/config"
	"github.com/wgarunap/xm-rest-api/domain"
	"net/http"
)

// Serve start the http server in background and return a closer func
// this server config.PORT in order to run
func Serve(c *config.Conf, r domain.Router) <-chan bool {

	srv := &http.Server{
		//ReadTimeout:  30 * time.Second,
		//WriteTimeout: 30 * time.Second,
		Addr:    fmt.Sprintf(":%d", c.ServicePort),
		Handler: r.Route(),
	}

	closed := make(chan bool, 1)
	go func() {
		log.Info(`server is starting on port:(` + srv.Addr + ")")
		err := srv.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				log.Error(err)
			}
		}
		closed <- true
	}()

	return closed

}
