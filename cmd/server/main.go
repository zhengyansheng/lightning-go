package main

import (
	"github.com/douyu/jupiter/pkg/xlog"
	"lightning-go/internal/app"
)

// @title lightning-go api
// @version 1.0
// @description lightning-go api
// @termsOfService https://github.com/zhengyansheng/lightning-go
// @license.name MIT
// @license.url https://github.com/zhengyansheng/lightning-go

func main() {

	eng := app.NewEngine()
	if err := eng.Run(); err != nil {
		xlog.Error("service could not run", xlog.FieldErr(err))
	}
}
