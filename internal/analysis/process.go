package analysis

import (
	"log"

	"mod.com/m/internal/config"
	"mod.com/m/internal/models"
)

func ProcessLightningData(ch <-chan *models.Lightning, chPost chan<- *models.Lightning) {
	for lightning := range ch {
		if IsWithinRadius(config.CityCoords, lightning) {
			log.Println("Lightning detected within radius:", lightning)
			chPost <- lightning
		} else {
			log.Println("Lightning detected outside radius:", lightning)
		}
	}
}
