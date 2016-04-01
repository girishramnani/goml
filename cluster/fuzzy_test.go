package cluster

import "testing"

func TestFuzzyClusterNew(t *testing.T) {
	fuzzyCl := NewFuzzyCluster(2)

	// if fuzzyCl.SetDataPoints([][]int{{1, 2}, {2, 3}}) != 2 {
	// 	t.Error("the dimensions dont match")
	//
	// }

	for _, val := range fuzzyCl.clusters {
		if val.x == 0 || val.y == 0 {
			t.Error("the random numbers were not correctly generated")
		}
	}
	// fmt.Println("Passed the Test Fuzzy Cluster New")
}

func TestClusterStore(t *testing.T) {
	fuzzCtl := NewFuzzyCluster(2)
	fuzzCtl.clusterStore[0] = []Point{
		{1, 2},
		{2, 5},
	}
	if len(fuzzCtl.clusterStore[0]) != 2 {
		t.Error("not inserting")
	}
}

func TestClusterProbabilityMatrix_ErrorThrow(t *testing.T) {
	fuzzCtl := NewFuzzyCluster(2)

	_, err := fuzzCtl.computeProbabilityMatrix()

	if err == nil {
		t.Fatal("Error not thrown")
	}

}

func TestComputeProbabilityMatrix_NoError(t *testing.T) {

	points := []Point{
		{1, 0},
		{2, 0},
		{9, 0},
		{10, 0},
	}

	FuzzyCtl := NewFuzzyCluster(2)
	FuzzyCtl.SetDataPoints(points)

	// For consistancy
	FuzzyCtl.clusters = []Point{
		{3, 0},
		{11, 0},
	}

	answer, _ := FuzzyCtl.computeProbabilityMatrix()
	if answer[0][0] != 1.0 && answer[1][0] != 1 {
		t.Fatal("The answer is wrong, Actual output ", answer[0][0], answer[0][1])
	}

}

func Test_Easy(t *testing.T) {

	points := []Point{
		{1, 1},
		{2, 1},
		{9, 1},
		{1000, 100},
		{1001, 103},
	}

	FuzzyCtl := NewFuzzyCluster(2)
	FuzzyCtl.SetDataPoints(points)

	// For consistancy
	FuzzyCtl.fuzzFactor = 1.25
	FuzzyCtl.clusters = []Point{
		{10, 1},
		{1005, 80},
	}

	FuzzyCtl.GenerateClusters()
	for i := 0; i < 10; i++ {
		FuzzyCtl.recomputeClusters()
	}
	clusters := FuzzyCtl.clusters
	// fmt.Println(clusters)
	if clusters[0].x != 1.668205210179437 && clusters[0].y == 1 && clusters[1].x != 1000.5002523738128 && clusters[1].y != 1 {
		t.Error("Cluster algorithm not working correctly")
	}

}

func TestIntegration(t *testing.T) {
	points := []Point{
		{12.0, 3504},
		{11.5, 3693},
		{11.0, 3436},
		{12.0, 3433},
		{10.5, 3449},
		{10.0, 4341},
		{9.0, 4354},
		{8.5, 4312},
		{10.0, 4425},
		{8.5, 3850},
		{10.0, 3563},
		{8.0, 3609},
		{9.5, 3761},
		{10.0, 3086},
		{15.0, 2372},
		{15.5, 2833},
		{15.5, 2774},
		{16.0, 2587},
	}
	if points[0].x != 12.0 {
		t.Error("precession error")
	}
	// fmt.Println(len(points))
	fuzzyCtl := NewFuzzyCluster(2)
	fuzzyCtl.SetDataPoints(points)
	fuzzyCtl.clusters = []Point{
		{9.9, 3000},
		{15, 3500},
	}
	fuzzyCtl.fuzzFactor = 1.25

	fuzzyCtl.GenerateClusters()

	// fmt.Println(cluster)
	// for i := 0; i < 5; i++ {
	// 	cluster, _ := fuzzyCtl.recomputeClusters()
	// 	fmt.Println(cluster[1])
	// }

}

func TestPoint(t *testing.T) {

	p := Point{2, 5}
	p2 := Point{2, 6}

	if p2.Distance(p) != 1 {
		t.Error("Linear distance not working")
	}

	p = Point{3, 0}
	p2 = Point{0, 4}

	if p2.Distance(p) != 5 {
		t.Error("Unrelated fields distance not working")
	}

}
