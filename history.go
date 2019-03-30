package main

import (
	"log"
	"time"
)

type History struct {
	Movements []*Movement
}

type Movement struct {
	x         float64
	y         float64
	timestamp time.Time
}

func NewHistory() *History {
	return &History{}
}

func (h *History) PrintDebug() {
	log.Printf("History Items: %v", h)

	for _, item := range h.Movements {
		log.Printf("item: %#v", item)
	}
}

//@todo function to remove applied history change
