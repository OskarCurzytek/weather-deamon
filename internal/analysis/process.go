package analysis

import (
	"mod.com/m/internal/config"
	"mod.com/m/internal/models"
)

func ProcessLightningData(ch <-chan *models.Lightning, chPost chan<- *models.Lightning) {
	for lightning := range ch {
		if IsWithinRadius(config.CityCoords, lightning) {
			chPost <- lightning
		}
	}
}
