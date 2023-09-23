package main

import (
	"fmt"
)

func main() {
	fmt.Println("Metagenomics!")

	// step 1: reading input from a single file.

	filename := "Data/Fall_Allegheny_1.txt"
	freqMap := ReadFrequencyMapFromFile(filename)
	fmt.Println("File read successfully! We have", len(freqMap), "total patterns.")

	// we may as well do something with our file. For example, let's print its Simpson's Index.
	fmt.Println("Simpson's Index:", SimpsonsIndex(freqMap))

	//step 2: reading input from a directory

	path := "Data/"
	allMaps := ReadSamplesFromDirectory(path)
	fmt.Println(len(allMaps))

	// step 3: processing the data that we have received.

	// now all of our maps have been processed and we can start working with them.

	// for example, let's compute the richness and evenness of each sample.

	richness := RichnessMap(allMaps)
	for key, val := range richness {
		fmt.Println(key, val)
	}

	simpson := SimpsonsMap(allMaps)
	for key, val := range simpson {
		fmt.Println(key, val)
	}

	// now let's look at beta diversity.

	distMetric := "Bray-Curtis"
	sampleNames, mtx := BetaDiversityMatrix(allMaps, distMetric)
	for i := range sampleNames {
		fmt.Println(sampleNames[i], mtx[i])
	}

	// this is all well and good but we cannot really analyze anything from this printing.

	// It would be better to print to a file.  Hence, we will need to learn writing to a file.

	simpsonFile := "Matrices/SimpsonMatrix.csv"
	WriteSimpsonsMapToFile(simpson, simpsonFile)

	outFilename := "Matrices/BetaDiversityMatrix.csv"
	WriteBetaDiversityMatrixToFile(mtx, sampleNames, outFilename)

	fmt.Println("Success! Now we are ready to do something cool with our data.")
}
