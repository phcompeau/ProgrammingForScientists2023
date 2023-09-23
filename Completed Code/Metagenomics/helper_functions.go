package main

//InitializeSquareMatrix takes an integer n and returns an nxn slice of
//floats initialized to 0.0.
func InitializeSquareMatrix(n int) [][]float64 {
	mtx := make([][]float64, n)

	for i := range mtx {
		mtx[i] = make([]float64, n)
	}
	return mtx
}

// FrequencyMap forms the frequency map of a collection of input patterns.
// Input: one collection of strings patterns
// Output: a frequency map of strings to their # of counts in patterns
func FrequencyMap(patterns []string) map[string]int {
	freqMap := make(map[string]int)
	for _, val := range patterns {
		freqMap[val]++
	}
	return freqMap
}

//Average takes two floats and returns their average.
func Average(x, y float64) float64 {
	return (x + y) / 2.0
}

//SampleTotal takes a frequency map as input.
//It returns the sum of the counts for each string in a sample.
func SampleTotal(freqMap map[string]int) int {
	total := 0
	for _, val := range freqMap {
		if val > 0 {
			total += val
		}
	}
	return total
}

//SumOfMinima takes two frequency maps as input.
//It identifies the keys that are shared by the two frequency maps
//and returns the sum of the corresponding values.
func SumOfMinima(map1 map[string]int, map2 map[string]int) int {
	c := 0

	for key := range map1 {
		_, exists := map2[key]
		if exists {
			c += Min2(map1[key], map2[key])
		}
	}

	return c
}

//Min2 takes two integers and returns their minimum.
func Min2(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
