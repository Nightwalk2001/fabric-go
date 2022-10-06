package schedules

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

var (
	c  *cron.Cron
	id cron.EntryID

	//t cron.EntryID
)

func Setup() {
	c = cron.New()

	id, _ = c.AddFunc("@daily", ResetId)
	//t, _ = c.AddFunc("@every 10s", IncrBy)

	c.Start()
}

func CleanUp() {
	c.Remove(id)
	//c.Remove(t)
	fmt.Println("定时任务清理完成")
}
