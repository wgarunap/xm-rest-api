package metrics

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/wgarunap/xm-rest-api/config"
	"net/http"
)

func MetricsRouter(cfg *config.Conf) {
	r := mux.NewRouter()
	r.Handle(`/metrics`, promhttp.Handler()).Methods(http.MethodGet)

	port := cfg.MetricsPort

	fmt.Println(fmt.Sprintf(`Http server started on port %+v`, port))

	err := http.ListenAndServe(fmt.Sprintf(`:%v`, port), r)
	if err != nil {
		fmt.Println(`error on metrics router,closing the endpoint : `, err)
	}
}
