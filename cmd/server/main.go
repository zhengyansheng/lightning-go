package main

import (
	"github.com/douyu/jupiter/pkg/xlog"
	_ "go-ops/docs"
	"go-ops/internal/app"
)

// @title go-ops api
// @version 1.0
// @description go-ops api
// @termsOfService https://github.com/51Reboot/go-ops
// @license.name MIT
// @license.url https://github.com/51Reboot/go-ops

func main() {

	eng := app.NewEngine()
	if err := eng.Run(); err != nil {
		xlog.Error("service could not run", xlog.FieldErr(err))
	}
}
