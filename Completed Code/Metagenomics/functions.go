package main

import "sort"

// BetaDiversityMatrix takes a map of frequency maps along with a distance metric
// ("Bray-Curtis" or "Jaccard") as input.
// It returns a slice of strings corresponding to the sorted names of the keys
// in the map, along with a matrix of distances whose (i,j)-th
// element is the distance between the i-th and j-th samples using the input metric.
func BetaDiversityMatrix(allMaps map[string](map[string]int), distMetric string) ([]string, [][]float64) {
	//first, grab all the sample IDs, put them in a list, and sort them
	numSamples := len(allMaps)

	sampleNames := make([]string, 0)

	//i := 0
	//get the sample IDs by ranging through allMaps
	for sampleName := range allMaps {
		sampleNames = append(sampleNames, sampleName)
		//sampleNames[i] = sampleName
		//i++
	}

	//sort the sample IDs
	sort.Strings(sampleNames)

	//build the beta diversity matrix
	mtx := InitializeSquareMatrix(numSamples)

	//diagonal is good to go because these entries should be 0.0

	//range through the matrix and set mtx[i][j] equal to beta diversity distance between sampleNames[i] and sampleNames[j]

	for i := range mtx {
		for j := i + 1; j < len(mtx[i]); j++ {
			//grab the two maps we need
			freqMap1 := allMaps[sampleNames[i]]
			freqMap2 := allMaps[sampleNames[j]]
			var d float64 // will store distance
			//which metric we using?
			if distMetric == "Bray-Curtis" {
				d = BrayCurtisDistance(freqMap1, freqMap2)
			} else if distMetric == "Jaccard" {
				d = JaccardDistance(freqMap1, freqMap2)
			} else {
				panic("Error: invalid distance metric given.")
			}
			mtx[i][j] = d
			//also enter the corresponding symmetric value
			mtx[j][i] = d
		}
	}
	return sampleNames, mtx
}

// SimpsonsMap takes an array of frequency maps as input. It returns a
// map mapping each sample name to its Simpson's index.
func SimpsonsMap(allMaps map[string](map[string]int)) map[string]float64 {
	s := make(map[string]float64)

	for sampleName, freqMap := range allMaps {
		s[sampleName] = SimpsonsIndex(freqMap)
	}

	return s
}

// RichnessMap takes a map of frequency maps as input.  It returns a map
// whose values are the richness of each sample.
func RichnessMap(allMaps map[string](map[string]int)) map[string]int {
	r := make(map[string]int)

	//range over sample IDs and set associated richness value
	for sampleID, freqMap := range allMaps {
		r[sampleID] = Richness(freqMap)
	}

	return r
}
