package main

import (
	"mod.com/m/internal/analysis"
	"mod.com/m/internal/models"
	"mod.com/m/internal/scheduler"
)

func main() {
	ch := make(chan *models.Lightning, 1000)
	go scheduler.Loop(ch)
	go analysis.ProcessLightningData(ch)

	select {}
}
