package app

import (
	"encoding/json"
	"fmt"
	"github.com/babbageLabs/brinch/bin"
	"github.com/babbageLabs/brinch/bin/core/methods"
	"github.com/babbageLabs/brinch/bin/core/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Instance struct {
	Name    string
	Address string
	Routes  map[string]*types.Route
}

func (app *Instance) GetRoute(name string) (*types.Route, error) {
	r, ok := app.Routes[name]
	if ok {
		return r, nil
	}
	return nil, fmt.Errorf(fmt.Sprintf("Route %s not found", name))
}

func (app *Instance) Start() error {
	err := app.CollectRoutes()
	if err != nil {
		return err
	}
	server := &http.Server{
		Addr:         app.Address,
		Handler:      app.RegisterHandlers(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	bin.Logger.Info("Starting application ", app.Name, " on address ", app.Address)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (app *Instance) RegisterHandlers() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(bin.JSONLogMiddleware())

	bin.Logger.Info("Registering route handlers...")
	for s, route := range app.Routes {
		r := *route
		routePath := strings.ReplaceAll(s, ".", "/")
		bin.Logger.Debug("Registering route handler ", routePath)
		e.Handle(r.GetMethod(), routePath, func(c *gin.Context) {
			r.SetContext(c)
			res, err := methods.Call(r)
			if err != nil {
				var responseCode = http.StatusBadRequest
				if res != nil {
					responseCode = res.Meta.ResponseCode
				}
				c.PureJSON(responseCode, gin.H{
					"error": err.Error(),
				})
				return
			}
			var result interface{}
			err = json.Unmarshal(res.Data, &result)
			if err != nil {
				c.PureJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.PureJSON(r.GetResponseCode(), gin.H{
				"result": result,
			})
		})
	}
	return e
}

func (app *Instance) CollectRoutes() error {
	return nil
}
