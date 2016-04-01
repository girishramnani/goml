package cluster

import (
	"errors"
	"math"
	"math/rand"
)

// FuzzyCluster is the fuzzy Clustering algorithm

const (
	MAX_X      = 100
	MAX_Y      = 100
	fuzzFactor = 1.01
)

// Point is the feature set
type Point struct {
	x float64
	y float64
}

// Less method to adhere to the Comparable
func (self Point) Less(other Point) bool {
	selfDistanceFromOrigin := self.Distance(Point{0.0, 0.0})
	otherDistanceFromOrigin := self.Distance(Point{0.0, 0.0})
	if selfDistanceFromOrigin > otherDistanceFromOrigin {
		return true
	}
	return false
}

// Distance provides euclidian distance between two points
func (self Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow((self.x-other.x), 2) + math.Pow((self.y-other.y), 2))
}

// FuzzyCluster is a clustering algorithm
type FuzzyCluster struct {
	numCluster   int
	params       map[string]int
	Points       []Point
	clusters     []Point
	clusterStore [][]Point
	dim          int
	oldClusters  []Point
	fuzzFactor   float64
}

// NewFuzzyCluster creates a new Fuzzy Clustering algorithm instance
func NewFuzzyCluster(numCluster int) *FuzzyCluster {
	return newFuzzyCluster(numCluster, fuzzFactor)
}

func newFuzzyCluster(numCluster int, fuzzier float64) *FuzzyCluster {

	newCluster := new(FuzzyCluster)
	newCluster.params = make(map[string]int, 0)
	newCluster.numCluster = numCluster
	newCluster.fuzzFactor = fuzzier
	newCluster.clusters = make([]Point, numCluster)
	newCluster.oldClusters = make([]Point, numCluster)
	for i := 0; i < numCluster; i++ {
		newCluster.clusters[i] = Point{rand.Float64() * MAX_X, rand.Float64() * MAX_Y}
	}
	newCluster.clusterStore = make([][]Point, numCluster)
	return newCluster
}

// SetDataPoints is used to Point data elements
func (self *FuzzyCluster) SetDataPoints(dataPoints []Point) {
	self.Points = dataPoints
	self.dim = 2

}

// GenerateClusters provides the final cluster
func (self *FuzzyCluster) GenerateClusters() ([]Point, error) {
	matrix, err := self.computeProbabilityMatrix()
	if err != nil {
		return nil, err
	}
	// fmt.Println(matrix)
	numPoints := len(self.Points)

	for i := 0; i < self.numCluster; i++ {
		self.clusters[i] = Point{0, 0}

		upper_x := 0.0
		for j := 0; j < numPoints; j++ {
			upper_x += math.Pow(matrix[i][j]*self.Points[j].x, self.fuzzFactor) * self.Points[j].x

		}
		upper_y := 0.0
		for j := 0; j < numPoints; j++ {
			upper_y += math.Pow(matrix[i][j]*self.Points[j].y, self.fuzzFactor) * self.Points[j].y
		}

		temp_x := 0.0
		for j := 0; j < numPoints; j++ {
			temp_x += math.Pow((matrix[i][j])*self.Points[j].x, self.fuzzFactor)
		}
		self.clusters[i].y = upper_x / temp_x

		temp_y := 0.0
		for j := 0; j < numPoints; j++ {

			temp_y += math.Pow((matrix[i][j])*self.Points[j].y, self.fuzzFactor)
		}

		self.clusters[i].y = upper_y / temp_y
	}

	return self.clusters, nil
}

func (self *FuzzyCluster) recomputeClusters() ([]Point, error) {
	copy(self.oldClusters, self.clusters)
	return self.GenerateClusters()
}

func (self *FuzzyCluster) computeProbabilityMatrix() ([][]float64, error) {

	if self.Points == nil {
		return nil, errors.New("Points not set")
	}

	var probMatrix [][]float64
	probMatrix = make([][]float64, self.numCluster)
	numPoints := len(self.Points)

	for i := 0; i < self.numCluster; i++ {
		probMatrix[i] = make([]float64, numPoints)
	}

	for i := 0; i < numPoints; i++ {
		totalDistance := 0.0
		for j := 0; j < self.numCluster; j++ {
			totalDistance += math.Pow(1/self.clusters[j].Distance(self.Points[i]), 1/(self.fuzzFactor-1))

		}

		for j := 0; j < self.numCluster; j++ {
			probMatrix[j][i] = math.Pow(1/self.clusters[j].Distance(self.Points[i]), 1/(self.fuzzFactor-1)) / totalDistance
		}
	}

	return probMatrix, nil
}
