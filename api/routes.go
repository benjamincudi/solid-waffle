package api

import (
	"context"
	"fmt"
	"reflect"
	"solid-waffle/api/handlers"
	"solid-waffle/waffledb"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source routes.go -destination mock_routes_test.go -package api
type methodRegistrar interface {
	gin.IRoutes
}

type getHandler interface{ HandleGet(c *gin.Context) }
type postHandler interface{ HandlePost(c *gin.Context) }
type putHandler interface{ HandlePut(c *gin.Context) }
type deleteHandler interface{ HandleDelete(c *gin.Context) }

func mapRoutesToHandlers(r methodRegistrar, routeMap gin.H) {
	for path, controller := range routeMap {
		matchedAnyMethod := false
		if getter, ok := controller.(getHandler); ok {
			matchedAnyMethod = true
			r.GET(path, getter.HandleGet)
		}
		if poster, ok := controller.(postHandler); ok {
			matchedAnyMethod = true
			r.POST(path, poster.HandlePost)
		}
		if putter, ok := controller.(putHandler); ok {
			matchedAnyMethod = true
			r.PUT(path, putter.HandlePut)
		}
		if deleter, ok := controller.(deleteHandler); ok {
			matchedAnyMethod = true
			r.DELETE(path, deleter.HandleDelete)
		}
		if !matchedAnyMethod {
			panic(fmt.Sprintf("handler %s for %s did not match any mapped method", reflect.TypeOf(controller), path))
		}
	}
}

func AttachRouteHandlers(r methodRegistrar) {
	ds := waffledb.MustGetDS(context.Background())
	routeMap := gin.H{
		"/api/sellers":  handlers.SellersController{DS: ds},
		"/api/products": handlers.ProductsController{DS: ds},
	}
	mapRoutesToHandlers(r, routeMap)
}
