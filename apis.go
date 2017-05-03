package webx

import (
	"strings"

	"github.com/ije/webx/user"
)

var xapis = APIService{}

type apiHandler struct {
	privileges user.Privileges
	handle     func(*Context, *XService)
}

type APIService map[string]map[string]apiHandler

func (s APIService) Get(endpoint string, handler func(*Context, *XService), privileges user.Privileges) {
	s.register("GET", endpoint, handler, privileges)
}

func (s APIService) Post(endpoint string, handler func(*Context, *XService), privileges user.Privileges) {
	s.register("POST", endpoint, handler, privileges)
}

func (s APIService) Put(endpoint string, handler func(*Context, *XService), privileges user.Privileges) {
	s.register("PUT", endpoint, handler, privileges)
}

func (s APIService) Delete(endpoint string, handler func(*Context, *XService), privileges user.Privileges) {
	s.register("DELETE", endpoint, handler, privileges)
}

func (s APIService) register(method string, endpoint string, handler func(*Context, *XService), privileges user.Privileges) {
	if s[method] == nil {
		s[method] = map[string]apiHandler{}
	}
	endpoint = strings.Trim(endpoint, "/")
	s[method][endpoint] = apiHandler{privileges, handler}
}

func RegisterAPIService(apis APIService) {
	for method, handlers := range apis {
		for endpoint, handler := range handlers {
			xapis.register(method, endpoint, handler.handle, handler.privileges)
		}
	}
}
