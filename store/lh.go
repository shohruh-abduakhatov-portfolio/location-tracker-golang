package store

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type LocationHistory struct {
	ID           string      `json:"id"`
	Coordincates Coordinates `json:"coordinates"` // ex. [lng, lat]
	Datetime     time.Time   `json:"datetime"`
}

type Coordinates []float64

func NewPost(data string) (*LocationHistory, error) {
	loc, err := ParseData(data)
	if err != nil {
		return nil, err
	}
	return loc, nil
}

func ParseData(data string) (*LocationHistory, error) {
	splitData := strings.Split(data, string(byte(31))) // unit separator
	if len(splitData) != 2 {
		return nil, errIncorrectData
	}

	lng, err := strconv.ParseFloat(splitData[0], 64)
	if err != nil {
		log.Println(errLng)
		return nil, errLng
	}

	lat, err := strconv.ParseFloat(splitData[1], 64)
	if err != nil {
		log.Println(errLat)
		return nil, errLat
	}

	lh := &LocationHistory{
		Coordincates: Coordinates{lng, lat},
		Datetime:     time.Now(),
	}

	return lh, nil
}
