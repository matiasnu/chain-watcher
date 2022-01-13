/*
   Copyright 2021 TEAM-A

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matiasnu/chain-watcher/config"
	"github.com/matiasnu/chain-watcher/controller"
)

// Route estructura basica para encapsular los elementos basicos de Route
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

var ApiUuid string

// App es una estructura para abstraer el Router y usarlo desde los Tests
type App struct {
	Router *mux.Router
	Server *http.Server
}

// Routes Array
type Routes []Route

// Run ejecuta el servicio
func (a *App) Run() {
	server := http.ListenAndServe(config.ConfMap.APIRestServerHost+":"+config.ConfMap.APIRestServerPort, a.Router)
	log.Fatal(server)
}

// InitializeAndServerRestAPI generates a new Router
func (a *App) InitializeRestAPI() {

	fmt.Println("Inicializando APP API")

	a.Router = mux.NewRouter().StrictSlash(true)
	a.Router.Use(commonMiddleware)
	a.Router.Use(loggingMiddleware)
	//a.Router.PathPrefix("/api")
	for _, route := range routes {
		a.Router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
}

var routes = Routes{
	Route{"AddContractWatch", "PATCH", "/watch", controller.AddContractWatch},
}
