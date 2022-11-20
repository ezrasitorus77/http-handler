package main

import (
	"fmt"
	"net/http"

	"github.com/ezrasitorus77/http-handler/config"
	"github.com/ezrasitorus77/http-handler/domain/delivery"
	log "github.com/ezrasitorus77/http-handler/internal/helper"
	"github.com/ezrasitorus77/http-handler/service"

	_ "github.com/ezrasitorus77/http-handler/config"
	"github.com/ezrasitorus77/http-handler/controller"
)

type a struct {
	b string
}

func main() {
	var (
		router        delivery.Router            = service.RouterService
		midService    delivery.MiddlewareService = service.MiddlewareService
		midController delivery.Handler
		server        http.Server
		e             error
	)

	router.POST("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("pre method", r.Header.Get("Access-Control-Request-Method"))
		fmt.Println("origin", r.Header.Get("Origin"))
		// if r.Header.Get("Access-Control-Request-Method") == "DELETE" {
		// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		// 	return
		// }
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		// w.Header().Set("Access-Control-Allow-Methods", "POST")
		fmt.Println("req method", r.Method)
		fmt.Println("")

		// tmpl, err := template.ParseFiles("test.html")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// b := a{
		// 	b: "b",
		// }
		// err = tmpl.Execute(w, b)
	})

	midController = controller.NewMiddleware(router, midService.NotFound)

	server = http.Server{
		Addr:    config.ServerAddress + ":" + config.ServerPort,
		Handler: midController,
	}

	e = server.ListenAndServe()
	if e != nil {
		log.ERROR(e.Error())
	}
}
