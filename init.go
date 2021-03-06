package main

import (
	"log"

	ecomapp "github.com/codedv8/go-ecom-app"
	ecomstructsapi "github.com/codedv8/go-ecom-structs/API"
	"github.com/gin-gonic/gin"
)

// SysInit - Pre initialization of this object
func (api *API) SysInit(app *ecomapp.Application) {
	// Create apiRouter
	apiRouter := app.Router.Group("/api/v1")
	// Create basic auth object
	basicAuth := app.UseBasicAuth("API")
	// Add basic auth to qpiRouter
	apiRouter.Use(basicAuth.Use())
	// Add handlers
	apiRouter.Handle("GET", "/", func(c *gin.Context) {
		payload := &ecomstructsapi.Root{S: "Hej"}
		_, _, err := app.CallHook("API_CALL", payload)
		if err != nil {
			log.Printf("Returned with error: %+v\n", err)
			c.JSON(200, payload)
		} else {
			log.Printf("Returned without error: %+v\n", payload)
			c.JSON(200, payload)
		}
	})

	// Add hook
	app.ListenToHook("API_ADD_ROUTER_HANDLE", func(payload interface{}) (bool, error) {
		switch v := payload.(type) {
		case *ecomstructsapi.Root:
			log.Printf("API_CALL: %+v\n", v)
			v.S = "Bye bye"
		}
		return false, nil
	})
}

// Init - Initialization of this object
func (api *API) Init(app *ecomapp.Application) {
	// Add hook
	app.ListenToHook("API_CALL", func(payload interface{}) (bool, error) {
		switch v := payload.(type) {
		case *ecomstructsapi.Root:
			log.Printf("API_CALL: %+v\n", v)
			v.S = "Bye bye"
		}
		return true, nil
	})
}
