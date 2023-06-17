package main

import (
	"net/http"

	"github.com/bagasjs/fuyubase/mux"
	"github.com/bagasjs/fuyubase/sqlex"
)

type AppConfig struct {
	Hostname    string
	Port        string
	Database    sqlex.Config
	ResourceDir string
}

type App struct {
	conf   AppConfig
	Router *mux.Router
}

func (self *App) Run() {
	self.OnSetup()
	self.OnProcess()
}

func (self *App) OnSetup() {
}

func (self *App) OnProcess() {
	http.ListenAndServe(":8000", self.Router)
}
