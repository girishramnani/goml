package genetic

import (
	"math/rand"
	"strings"
)

type chromosomes []string

var CHR = [2]byte{'0', '1'}

type RouletteWheel struct {
	length        int
	sample        chromosomes
	totalStrength int
	numSamples    int
}

func generateSample(length int) string {
	var temp []byte
	temp = make([]byte, length)

	for i := 0; i < length; i++ {
		temp[i] = CHR[rand.Intn(2)]
	}

	chr := string(temp[:])
	return chr

}

func (rw *RouletteWheel) getNewIndex(prob float64, wheel []float64) int {
	for t := 0; t < rw.numSamples; t++ {
		if prob < wheel[t] {
			return t
		}
	}
	return -1
}

func getTotalStrength(c chromosomes) int {
	d := 0

	for i := 0; i < len(c); i++ {
		d += strings.Count(c[i], "1")
	}
	return d
}

func NewRouletteWheelSelector(length int, numSamples int) *RouletteWheel {

	rw := &RouletteWheel{length: length,
		numSamples: numSamples,
	}
	rw.sample = make([]string, numSamples)

	for i := 0; i < numSamples; i++ {
		rw.sample[i] = generateSample(length)
	}

	rw.totalStrength = getTotalStrength(rw.sample)

	return rw
}

func (rw *RouletteWheel) Strength(index int) float64 {

	return float64(strings.Count(rw.sample[index], "1")) / float64(rw.totalStrength)

}

func (rw *RouletteWheel) GetDistribution() []float64 {
	temp := make([]float64, rw.numSamples)

	for i := 0; i < rw.numSamples; i++ {
		temp[i] = rw.Strength(i)

		if i != 0 {
			temp[i] += temp[i-1]
		}

	}
	return temp
}
