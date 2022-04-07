package app

import (
	"github.com/robfig/cron/v3"
	"stock/crons"
)

func handleCron() {
	c := cron.New()
	c.AddFunc("30 18 * * *", crons.SetStockCurrentPrice)
	c.Start()
}
