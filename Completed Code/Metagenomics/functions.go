package main

import "sort"

//BrayCurtisDistance takes two frequency maps and returns the Bray-Curtis
//distance between them.
func BrayCurtisDistance(map1 map[string]int, map2 map[string]int) float64 {
	c := SumOfMinima(map1, map2)
	s1 := SampleTotal(map1)
	s2 := SampleTotal(map2)

	//don't allow the situation in which we have zero richness.
	if s1 == 0 || s2 == 0 {
		panic("Error: sample given to BrayCurtisDistance() has no positive values.")
	}

	av := Average(float64(s1), float64(s2))
	return 1 - (float64(c) / av)
}

//JaccardDistance takes two frequency maps and returns the Jaccard
//distance between them.
func JaccardDistance(map1 map[string]int, map2 map[string]int) float64 {
	c := SumOfMinima(map1, map2)
	t := SumOfMaxima(map1, map2)
	j := 1 - (float64(c) / float64(t))
	return j
}

//SumOfMaxima takes two frequency maps as input.
//It identifies the keys that are shared by the two frequency maps
//and returns the sum of the corresponding values. (a.k.a. "union")
//SumOfMaxima takes two frequency maps as input.
//It identifies the keys that are shared by the two frequency maps
//and returns the sum of the corresponding values. (a.k.a. "union")
func SumOfMaxima(map1 map[string]int, map2 map[string]int) int {
	c := 0

	for key := range map2 {
		_, exists := map1[key]
		if exists {
			c += Max2(map1[key], map2[key])
		} else {
			c += map2[key]
		}
	}
	for key := range map1 {
		_, exists := map2[key]
		if !exists {
			c += map1[key]
		}
	}

	// panic if c is equal to zero since we don't want 0/0
	if c == 0 {
		panic("Error: no species common to two maps given to SumOfMaxima")
	}

	return c
}

//Max2 takes two integers and returns their maximum.
func Max2(n1, n2 int) int {
	if n1 < n2 {
		return n2
	}
	return n1
}

//SimpsonsIndex takes a frequency map and returns a decimal corresponding to Simpson's index:
//the sum of (n/N)^2 where n is the number of individuals found for a given string/species
//and N is the total number of individuals. The sum is over all species present.
func SimpsonsIndex(sample map[string]int) float64 {
	total := SampleTotal(sample)
	simpson := 0.0

	if total == 0 {
		panic("Error: Empty frequency map given to SimpsonsIndex()!")
	}

	for _, val := range sample {
		x := float64(val) / float64(total)
		simpson += x * x
	}
	return simpson
}

//BetaDiversityMatrix takes a map of frequency maps along with a distance metric
//("Bray-Curtis" or "Jaccard") as input.
//It returns a slice of strings corresponding to the sorted names of the keys
//in the map, along with a matrix of distances whose (i,j)-th
//element is the distance between the i-th and j-th samples using the input metric.
func BetaDiversityMatrix(allMaps map[string](map[string]int), distMetric string) ([]string, [][]float64) {

	//first grab all strings
	numSamples := len(allMaps)
	sampleNames := make([]string, 0)
	for name := range allMaps {
		sampleNames = append(sampleNames, name)
	}

	// now sort sample names to make matrix ordered
	sort.Strings(sampleNames)

	// now form the distance matrix

	mtx := InitializeSquareMatrix(numSamples)

	for i := 0; i < numSamples; i++ {
		for j := i; j < numSamples; j++ {
			if distMetric == "Bray-Curtis" {
				d := BrayCurtisDistance(allMaps[sampleNames[i]], allMaps[sampleNames[j]])
				mtx[i][j] = d
				mtx[j][i] = d
			} else if distMetric == "Jaccard" {
				d := JaccardDistance(allMaps[sampleNames[i]], allMaps[sampleNames[j]])
				mtx[i][j] = d
				mtx[j][i] = d
			} else {
				panic("Error: Invalid distance metric name given to BetaDiversityMatrix().")
			}
		}
	}
	return sampleNames, mtx
}

//SimpsonsMap takes an array of frequency maps as input. It returns a
//map mapping each sample name to its Simpson's index.
func SimpsonsMap(allMaps map[string](map[string]int)) map[string]float64 {

	s := make(map[string]float64)

	for sampleName, freqMap := range allMaps {
		s[sampleName] = SimpsonsIndex(freqMap)
	}
	return s
}

//RichnessMap takes a map of frequency maps as input.  It returns a map
//whose values are the richness of each sample.
func RichnessMap(allMaps map[string](map[string]int)) map[string]int {

	r := make(map[string]int)

	for sampleName, freqMap := range allMaps {
		r[sampleName] = Richness(freqMap)
	}

	return r
}

//Richness takes a frequency map. It returns the richness of the frequency map
//(i.e., the number of keys in the map corresponding to nonzero values.)
func Richness(sample map[string]int) int {
	c := 0

	for _, val := range sample {
		if val < 0 {
			panic("Error: negative value in frequency map given to Richness()")
		}
		c++
	}

	return c
}
