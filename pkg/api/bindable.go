package api

import "github.com/gin-gonic/gin"

// Bindable allows to bind routes or endpoints to other router.
type Bindable interface {
	// Bind binds object to router.
	Bind(router gin.IRouter)
}
