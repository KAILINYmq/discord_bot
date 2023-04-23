package Scheduled_work

import (
	"DiscordRolesBot/internal/async_work"
	"fmt"
	"github.com/robfig/cron/v3"
)

func ScheduledWork() {
	cornJob := cron.New(cron.WithSeconds())
	fmt.Println("Init ScheduledWork")
	cornJob.AddFunc("0 */5 * * * *", async_work.ScheduledJob)
	cornJob.Start()
}
