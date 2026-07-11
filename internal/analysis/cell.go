package analysis

import (
	"math"

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

func AddLightning(grid map[models.Cell]int, ch <-chan *models.Lightning, lat0, lon0, cellSize float64) {
	for lightning := range ch {
		dx, dy := latLonToXY(lightning.Lat, lightning.Lon, lat0, lon0)
		i, j := toCell(dx, dy, cellSize)

		cell := models.Cell{I: i, J: j}
		grid[cell]++
	}
}
