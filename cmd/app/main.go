package main

import (
	"mod.com/m/internal/analysis"
	"mod.com/m/internal/config"
	"mod.com/m/internal/models"
	"mod.com/m/internal/scheduler"
)

func main() {
	grid := make(map[models.Cell]int)
	lightningsPre := make(chan *models.Lightning, 1000)
	lightningsPost := make(chan *models.Lightning, 1000)
	go scheduler.GetData(lightningsPre)
	go analysis.ProcessLightningData(lightningsPre, lightningsPost)
	go analysis.AddLightning(grid, lightningsPost, config.CityCoords.Lat, config.CityCoords.Lon, 3.0)

	select {}
}
