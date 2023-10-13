package main

import (
	"fmt"
)

func main() {

	fmt.Println("Read in Hemoglobin alpha subunit 1.")

	mtx, speciesNames := ReadMatrixFromFile("Data/HBA1/hemoglobin_distance.mtx")

	fmt.Println("Starting UPGMA.")

	// generate UPGMA tree

	hemoglobinTree := UPGMA(mtx, speciesNames)

	fmt.Println("UPGMA tree built.")

	WriteNewickToFile(hemoglobinTree, "Output/HBA1", "hba1.tre")

}
