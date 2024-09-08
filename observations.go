package clusters

import (
	"fmt"
	"math"
)

// Coordinates is a struct with a slice of float64 and an additional string field
type Coordinates struct {
	Values []float64
	Label  string
}

// Observation is a data point (float64 between 0.0 and 1.0) in n dimensions
type Observation interface {
	Coordinates() Coordinates
	Distance(point Coordinates) float64
}

// Observations is a slice of observations
type Observations []Observation

// Coordinates implements the Observation interface for a plain set of float64
// coordinates
func (c Coordinates) Coordinates() Coordinates {
	return c
}

// Distance returns the euclidean distance between two coordinates
func (c Coordinates) Distance(p2 Coordinates) float64 {
	var r float64
	for i, v := range c.Values {
		r += math.Pow(v-p2.Values[i], 2)
	}
	return r
}

// Center returns the center coordinates of a set of Observations
func (c Observations) Center() (Coordinates, error) {
	var l = len(c)
	if l == 0 {
		return Coordinates{}, fmt.Errorf("there is no mean for an empty set of points")
	}

	cc := make([]float64, len(c[0].Coordinates().Values))
	for _, point := range c {
		for j, v := range point.Coordinates().Values {
			cc[j] += v
		}
	}

	var mean Coordinates
	mean.Values = make([]float64, len(cc))
	for i, v := range cc {
		mean.Values[i] = v / float64(l)
	}
	return mean, nil
}

// AverageDistance returns the average distance between o and all observations
func AverageDistance(o Observation, observations Observations) float64 {
	var d float64
	var l int

	for _, observation := range observations {
		dist := o.Distance(observation.Coordinates())
		if dist == 0 {
			continue
		}

		l++
		d += dist
	}

	if l == 0 {
		return 0
	}
	return d / float64(l)
}
