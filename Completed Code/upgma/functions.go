package main

import "strconv"

// UPGMA takes a distance matrix and a collection of species names as input.
// It returns a Tree (an array of nodes) resulting from applying
// UPGMA to this dataset.
func UPGMA(mtx DistanceMatrix, speciesNames []string) Tree {
	AssertSquareMatrix(mtx)
	AssertSameNumberSpecies(mtx, speciesNames)

	t := InitializeTree(speciesNames)
	clusters := t.InitializeClusters() //clusters are []*Node

	numLeaves := len(speciesNames) //OR numLeaves := CountRows(mtx)

	for k := numLeaves; k < 2*numLeaves-1; k++ {
		//engine of UPGMA
		//find minimum element
		row, col, val := FindMinElement(mtx)
		//set age of current node
		t[k].Age = val / 2.0
		//then, set the two children of t[k]
		t[k].Child1 = clusters[row]
		t[k].Child2 = clusters[col]
		mtx = AddRowCol(mtx, clusters, row, col)
		mtx = DelRowCol(mtx, row, col)
		clusters = append(clusters, t[k])
		clusters = DeleteClusters(clusters, row, col)
	}

	return t
}

// DeleteClusters takes as input a slice of Node pointers (clusters) and two integers row and col.
// It deletes the elements from the slice having indices row and col and returns the updated slice.
// It assumes that row < col
func DeleteClusters(clusters []*Node, row, col int) []*Node {
	if row == col {
		panic("no")
	}
	clusters = append(clusters[:col], clusters[col+1:]...)
	clusters = append(clusters[:row], clusters[row+1:]...)
	return clusters
}

// DeleteRowCol takes as input a distance matrix and row and column indices.
// It deletes the row and column with both of these indices.
// It assumes that row < col
func DeleteRowCol(mtx DistanceMatrix, row, col int) DistanceMatrix {
	// first, delete rows
	mtx = append(mtx[:col], mtx[col+1:]...)
	mtx = append(mtx[:row], mtx[row+1:]...)
	//next, delete the columns
	for i := range mtx {
		mtx[i] = append(mtx[i][:col], mtx[i][col+1:]...)
		mtx[i] = append(mtx[i][:row], mtx[i][row+1:]...)
	}
	return mtx
}

// InitializeClusters is a Tree method.
// It returns a slice of Node pointers pointing at the leaves of the Tree.
func (t Tree) InitializeClusters() []*Node {
	//len(t) = 2*numLeaves - 1
	//(len(t)+1)/2 = numLeaves
	numLeaves := (len(t) + 1) / 2
	clusters := make([]*Node, numLeaves)
	for i := range clusters {
		clusters[i] = t[i]
	}

	return clusters
}

// InitializeTree takes a slice of strings speciesNames as input of length n.
// It returns a tree having 2n-1 total nodes. The first n will represent leaf nodes and correspond to the species in speciesNames, so that node i has label equal to speciesNames[i].
func InitializeTree(speciesNames []string) Tree {

	numLeaves := len(speciesNames)
	t := make([]*Node, 2*numLeaves-1)

	//point these default nil pointers at real nodes
	for i := range t {

		//create a node
		var vx Node
		//set its fields
		vx.Num = i
		if i < numLeaves {
			//leaf
			vx.Label = speciesNames[i]
		} else {
			//ancestor
			vx.Label = "Ancestor species: " + strconv.Itoa(i)
		}
		//age is zero by default
		//no need to set children (will be set for internal nodes later)
		//point t[i] at vx
		t[i] = &vx
	}

	return t
}

/*
UPGMA(D, speciesNames)
	t  InitializeTree(speciesNames)
	clusters  t.InitializeClusters() // clusters are []*Node
	numLeaves  CountRows(D) // = |speciesNames|
	for every integer k from numLeaves to 2*numLeaves–2
		row, col, val  FindMinElt(D)
		t[k].age  val/2
		t[k].child1  clusters[row]
		t[k].child2  clusters[col]
		D  AddRowCol(D, clusters, row, col)
		D  DelRowCol(D, row, col)
		clusters  append(clusters, t[k])
		clusters  DelClusters(clusters, row, col)
	return t

*/

func AssertSameNumberSpecies(mtx DistanceMatrix, speciesNames []string) {
	if len(mtx) != len(speciesNames) {
		panic("Differing number of species names given to number of rows in distance matrix.")
	}
}

func AssertSquareMatrix(mtx DistanceMatrix) {
	numRows := len(mtx)
	for i := range mtx {
		if len(mtx[i]) != numRows {
			panic("Non-square matrix given.")
		}
	}
}

func CountRows(mtx DistanceMatrix) int {
	return len(mtx)
}

// CountLeaves
// Input: a Node pointer v
// Output: the number of leaves in the tree whose root is v. If v is a leaf, return 1.
func CountLeaves(v *Node) int {
	//base case: both children are nil and we are at a leaf
	if v.Child1 == nil && v.Child2 == nil {
		return 1
	} else if v.Child1 == nil {
		return CountLeaves(v.Child2)
	} else if v.Child2 == nil {
		return CountLeaves(v.Child1)
	}
	return CountLeaves(v.Child1) + CountLeaves(v.Child2)
}
