package app

import (
	"fmt"
	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/server/xgin"
	"github.com/douyu/jupiter/pkg/util/xcolor"
	"github.com/douyu/jupiter/pkg/util/xgo"
	"github.com/douyu/jupiter/pkg/worker/xcron"
	"github.com/douyu/jupiter/pkg/xlog"
	"go-ops/internal/api/message"
	"go-ops/internal/db"
	"go-ops/internal/http"
)

type Engine struct {
	jupiter.Application
}

func NewEngine() *Engine {
	eng := &Engine{}
	eng.HideBanner = true
	if err := eng.Startup(
		xgo.ParallelWithError(
			eng.printBanner,
			eng.serveHTTP,
			eng.initApp,
		),
	); err != nil {
		xlog.Panic("startup engine", xlog.Any("err", err))
	}

	return eng
}

func (eng *Engine) printBanner() error {

	const banner = `
   ______ ____   ____   ____  _____
  / ____// __ \ / __ \ / __ \/ ___/
 / / __ / / / // / / // /_/ /\__ \ 
/ /_/ // /_/ // /_/ // ____/___/ / 
\____/ \____/ \____//_/    /____/  

 Welcome to LIGHTNING-OPS  API, starting application ...
`

	fmt.Println(xcolor.Green(banner))
	return nil
}
func (eng *Engine) initApp() error {

	db.Init()
	eng.initCron()
	eng.printBanner()
	go message.GobalSend.ConsumerMsg()
	//redis.InitRedis()
	return nil
}

func (eng *Engine) initCron() error {
	cron := xcron.StdConfig("chaos").Build()
	eng.Schedule(cron)
	return nil
}

//http服务
func (eng *Engine) serveHTTP() error {
	server := xgin.StdConfig("http").Build()
	http.InitRouters(server)
	return eng.Serve(server)
}
