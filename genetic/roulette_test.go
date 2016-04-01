package genetic

import (
	"math/rand"
	"testing"
	"time"
)

func testableRoulette() *RouletteWheel {

	rw := NewRouletteWheelSelector(8, 2)
	rw.sample[0] = "00000101"
	rw.sample[1] = "11111100"
	rw.totalStrength = getTotalStrength(rw.sample)

	return rw

}

func TestSampleGeneration(t *testing.T) {

	if len(generateSample(5)) != 5 {
		t.Error("Length doesnt match")
	}
}

func TestRouletteStrength(t *testing.T) {

	rw := testableRoulette()
	if rw.Strength(0) != 0.25 {
		t.Errorf("The strength isnt comming same. Actual  %f", rw.Strength(0))
	}

}

func TestRoulette(t *testing.T) {

	rw := testableRoulette()
	wheel := rw.GetDistribution()
	var new_chr chromosomes
	new_chr = make([]string, rw.numSamples)
	t.Log(wheel)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rw.numSamples; i++ {
		r := rand.Float64()
		new_chr[i] = rw.sample[rw.getNewIndex(r, wheel)]
	}

	rw.sample = new_chr
	t.Log(rw)

}
