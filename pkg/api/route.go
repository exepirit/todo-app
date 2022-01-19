package api

import "github.com/gin-gonic/gin"

// Route creates new group of Bindable objects under generic prefix.
func Route(prefix string, bindable ...Bindable) Bindable {
	return &route{
		prefix:    prefix,
		subRoutes: bindable,
	}
}

type route struct {
	prefix    string
	subRoutes []Bindable
}

func (r *route) Bind(router gin.IRouter) {
	router = router.Group(r.prefix)
	for _, subRoute := range r.subRoutes {
		subRoute.Bind(router)
	}
}
