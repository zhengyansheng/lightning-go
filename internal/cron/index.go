package cron

import (
	"github.com/douyu/jupiter/pkg/worker/xcron"
)

func InitShedule(cron *xcron.Cron) {
	//add cron task here

	// ex
	//if !conf.GetBool("kjcloud.disable_iso") {
	//	cron.Schedule(xcron.Every(time.Minute*5), xcron.FuncJob(ResetInstallStatus))
	//}

}
