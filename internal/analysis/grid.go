package analysis

import (
	"math"
	"time"

	"mod.com/m/internal/config"
	"mod.com/m/internal/models"
)

func latLonToXY(lat, lon, lat0, lon0 float64) (float64, float64) {
	dy := (lat - lat0) * config.KmPerDegree
	dx := (lon - lon0) * config.KmPerDegree * math.Cos(lat0*math.Pi/180.0)
	return dx, dy
}

func toCell(dx, dy, cellSize float64) (int, int) {
	i := int(math.Floor(dx / cellSize))
	j := int(math.Floor(dy / cellSize))
	return i, j
}

func handleLightning(grid map[models.Cell]int, lightning *models.Lightning, lat0, lon0, cellSize float64) {
	dx, dy := latLonToXY(lightning.Lat, lightning.Lon, lat0, lon0)
	i, j := toCell(dx, dy, cellSize)
	cell := models.Cell{I: i, J: j}
	grid[cell]++
}

func createCopyOfGrid(grid map[models.Cell]int) map[models.Cell]int {
	newGrid := make(map[models.Cell]int, len(grid))
	for cell, count := range grid {
		newGrid[cell] = count
	}
	return newGrid
}

func GridSnapshot(grid map[models.Cell]int, ch <-chan *models.Lightning, lat0, lon0, cellSize float64, snapshot chan<- map[models.Cell]int) {
	snapshotTicker := time.NewTicker(5 * time.Minute)
	cleanTicker := time.NewTicker(10 * time.Minute)
	defer snapshotTicker.Stop()
	defer cleanTicker.Stop()
	for {
		select {
		case lightning, ok := <-ch:
			if !ok {
				return
			}
			handleLightning(grid, lightning, lat0, lon0, cellSize)
		case <-snapshotTicker.C:
			snapshot <- createCopyOfGrid(grid)
		case <-cleanTicker.C:
			for cell := range grid {
				delete(grid, cell)
			}
		}

	}
}
