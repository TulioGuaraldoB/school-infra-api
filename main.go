package main

import (
	"github.com/TulioGuaraldoB/school-report/config/env"
	"github.com/TulioGuaraldoB/school-report/server"
)

func main() {
	env.GetEnvironmentVariables()

	server := server.NewServer()
	server.Run()
}
