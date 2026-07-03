package decoder

import (
	"bytes"
	"regexp"
	"strconv"

	"mod.com/m/internal/models"
)

type Decoder struct {
	substitutions map[string]string
}

func NewDecoder() *Decoder {
	return &Decoder{
		substitutions: map[string]string{
			"\xc4\x88": "0", "\xc4\x89": "1", "\xc4\x8a": "2", "\xc4\x8b": "3",
			"\xc4\x8c": "4", "\xc4\x8d": "5", "\xc4\x8e": "6", "\xc4\x8f": "7",
			"\xc4\x90": "8", "\xc4\x91": "9", "\xc4\x86": ": ",
			"\xc4\x87": ".", "\xc4\xb8": `"`, "\xc4\xb9": `"`, "\xc4\xba": `"`},
	}
}

func (d *Decoder) Decode(data []byte) *models.Lightning {
	result := data
	for k, v := range d.substitutions {
		result = bytes.ReplaceAll(result, []byte(k), []byte(v))
	}

	text := string(result)

	timeRe := regexp.MustCompile(`time[":]+(\d+)`)
	latRe := regexp.MustCompile(`lat[:\s]*([0-9\.-]+)`)
	lonRe := regexp.MustCompile(`lon[:\s]*([0-9\.-]+)`)

	timeMatch := timeRe.FindStringSubmatch(text)
	latMatch := latRe.FindStringSubmatch(text)
	lonMatch := lonRe.FindStringSubmatch(text)

	if timeMatch == nil || latMatch == nil || lonMatch == nil {
		return nil
	}

	t, _ := strconv.ParseInt(timeMatch[1], 10, 64)
	lat, _ := strconv.ParseFloat(latMatch[1], 64)
	lon, _ := strconv.ParseFloat(lonMatch[1], 64)

	return &models.Lightning{
		Time: t,
		Lat:  lat,
		Lon:  lon,
	}

}
