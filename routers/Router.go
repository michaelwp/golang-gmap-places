package routers

import (
	"github.com/gorilla/mux"
	"github.com/michaelwp/golang-gmap-places/routers/v1"
)

func Router(m *mux.Router) {
	api := m.PathPrefix("/api").Subrouter()
	routerVersion1(api)
}

/* api version 1 */
func routerVersion1(m *mux.Router) {
	ver1 := m.PathPrefix("/v1").Subrouter()
	v1.RouterProfile(ver1)
}
