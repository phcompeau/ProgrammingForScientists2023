package main

import (
	// "fmt"
	"reflect"
	"testing"
)

//type Tree Tree
/********************************************
 FIND_MIN_ELEMENT
*********************************************/

type fme_TestingPair struct {
	matrixPath string
	minElt     float64
	rowLoc     int
	colLoc     int
}

var tests = []fme_TestingPair{
	{"Tests/matrix1.txt", 2.0, 0, 1},
	{"Tests/matrix5.txt", 2.0, 2, 3},
	{"Tests/matrix6.txt", 8.0, 2, 3},
	{"Tests/matrix7.txt", 1.0, 0, 6},
	{"Tests/matrix8.txt", 2.0, 2, 3},
	{"Tests/matrix9.txt", 2.0, 0, 7},
	{"Tests/matrix10.txt", 2.0, 0, 7}}

func TestFindMinElement(t *testing.T) {
	for _, test := range tests {

		mtx, _ := ReadMatrixFromFile(test.matrixPath)
		r, c, val := FindMinElement(mtx)

		if !(val == test.minElt && r == test.rowLoc && c == test.colLoc) {

			t.Error(
				"For:", mtx,
				"Expected Minimum Element:", test.minElt,
				"Got:", val)

		}

	}
}

/********************************************
 ADD_ROW_COL
*********************************************/

func createTree1() Tree {
	T := make([]*Node, 7)

	var NODE0 Node
	var NODE1 Node
	var NODE2 Node
	var NODE3 Node
	var NODE4 Node
	var NODE5 Node
	var NODE6 Node

	NODE4.Child1 = &NODE0
	NODE4.Child2 = &NODE3
	NODE5.Child1 = &NODE1
	NODE5.Child2 = &NODE2
	NODE6.Child1 = &NODE4
	NODE6.Child2 = &NODE6

	T[0] = &NODE0
	T[1] = &NODE1
	T[2] = &NODE2
	T[3] = &NODE3
	T[4] = &NODE4
	T[5] = &NODE5
	T[6] = &NODE6

	return T
}

func createTree2() Tree {
	T := make([]*Node, 7)

	var NODE0 Node
	var NODE1 Node
	var NODE2 Node
	var NODE3 Node
	var NODE4 Node
	var NODE5 Node
	var NODE6 Node

	NODE4.Child1 = &NODE2
	NODE4.Child2 = &NODE3
	NODE5.Child1 = &NODE0
	NODE5.Child2 = &NODE4
	NODE6.Child1 = &NODE1
	NODE6.Child2 = &NODE5

	T[0] = &NODE0
	T[1] = &NODE1
	T[2] = &NODE2
	T[3] = &NODE3
	T[4] = &NODE4
	T[5] = &NODE5
	T[6] = &NODE6

	return T
}

func createTree3() Tree {
	T := make([]*Node, 15)

	var NODE0 Node
	var NODE1 Node
	var NODE2 Node
	var NODE3 Node
	var NODE4 Node
	var NODE5 Node
	var NODE6 Node
	var NODE7 Node
	var NODE8 Node
	var NODE9 Node
	var NODE10 Node
	var NODE11 Node
	var NODE12 Node
	var NODE13 Node
	var NODE14 Node

	NODE8.Child1 = &NODE3
	NODE8.Child2 = &NODE4
	NODE9.Child1 = &NODE0
	NODE9.Child2 = &NODE1
	NODE10.Child1 = &NODE8
	NODE10.Child2 = &NODE5
	NODE11.Child1 = &NODE9
	NODE11.Child2 = &NODE2
	NODE12.Child1 = &NODE10
	NODE12.Child2 = &NODE6
	NODE13.Child1 = &NODE11
	NODE13.Child2 = &NODE12
	NODE14.Child1 = &NODE13
	NODE14.Child2 = &NODE7

	T[0] = &NODE0
	T[1] = &NODE1
	T[2] = &NODE2
	T[3] = &NODE3
	T[4] = &NODE4
	T[5] = &NODE5
	T[6] = &NODE6
	T[7] = &NODE7
	T[8] = &NODE8
	T[9] = &NODE9
	T[10] = &NODE10
	T[11] = &NODE11
	T[12] = &NODE12
	T[13] = &NODE13
	T[14] = &NODE14

	return T
}

var TREE1 = createTree1()
var TREE2 = createTree2()
var TREE3 = createTree3()

type arc_TestingPair struct {
	tree Tree
	mtx  DistanceMatrix
	cltr []*Node
	row  int
	col  int
	ans  DistanceMatrix
}

var arc_tests = []arc_TestingPair{
	{TREE2,
		[][]float64{[]float64{0, 16},
			[]float64{16, 0}},
		[]*Node{TREE2[1], TREE1[5]},
		0,
		1,
		[][]float64{[]float64{0, 16, 0},
			[]float64{16, 0, 0},
			[]float64{0, 0, 0}}},
	{TREE1,
		[][]float64{[]float64{0, 1, 3},
			[]float64{1, 0, 3},
			[]float64{3, 3, 0}},
		[]*Node{TREE1[0], TREE1[1], TREE1[4]},
		0,
		1,
		[][]float64{[]float64{0, 1, 3, 0},
			[]float64{1, 0, 3, 0},
			[]float64{3, 3, 0, 3},
			[]float64{0, 0, 3, 0}}},
	{TREE1,
		[][]float64{[]float64{0, 3, 3, 1},
			[]float64{3, 0, 1, 3},
			[]float64{3, 1, 0, 3},
			[]float64{1, 3, 3, 0}},
		[]*Node{TREE1[0], TREE1[1], TREE1[2], TREE1[3]},
		0,
		3,
		[][]float64{[]float64{0, 3, 3, 1, 0},
			[]float64{3, 0, 1, 3, 3},
			[]float64{3, 1, 0, 3, 3},
			[]float64{1, 3, 3, 0, 0},
			[]float64{0, 3, 3, 0, 0}}},
	{TREE2,
		[][]float64{[]float64{0, 20, 9, 11},
			[]float64{20, 0, 17, 11},
			[]float64{9, 17, 0, 8},
			[]float64{11, 11, 8, 0}},
		[]*Node{TREE2[0], TREE2[1], TREE2[2], TREE2[3]},
		2,
		3,
		[][]float64{[]float64{0, 20, 9, 11, 10},
			[]float64{20, 0, 17, 11, 14},
			[]float64{9, 17, 0, 8, 0},
			[]float64{11, 11, 8, 0, 0},
			[]float64{10, 14, 0, 0, 0}}},
	{TREE3,
		[][]float64{[]float64{0, 5, 8, 25, 27, 72, 26},
			[]float64{5, 0, 9, 22, 28, 69, 23},
			[]float64{8, 9, 0, 22, 25, 71, 19.5},
			[]float64{25, 22, 22, 0, 14, 54, 6.5},
			[]float64{27, 28, 25, 14, 0, 63, 11.5},
			[]float64{72, 69, 71, 54, 63, 0, 64.5},
			[]float64{26, 23, 19.5, 6.5, 11.5, 64.5, 0}},
		[]*Node{TREE3[0], TREE3[1], TREE3[2], TREE3[5], TREE3[6], TREE3[7], TREE3[8]},
		0,
		1,
		[][]float64{[]float64{0, 5, 8, 25, 27, 72, 26, 0},
			[]float64{5, 0, 9, 22, 28, 69, 23, 0},
			[]float64{8, 9, 0, 22, 25, 71, 19.5, 8.5},
			[]float64{25, 22, 22, 0, 14, 54, 6.5, 23.5},
			[]float64{27, 28, 25, 14, 0, 63, 11.5, 27.5},
			[]float64{72, 69, 71, 54, 63, 0, 64.5, 70.5},
			[]float64{26, 23, 19.5, 6.5, 11.5, 64.5, 0, 24.5},
			[]float64{0, 0, 8.5, 23.5, 27.5, 70.5, 24.5, 0}}}}

func TestAddRowCol(t *testing.T) {
	for _, test := range arc_tests {

		val := AddRowCol(test.mtx, test.cltr, test.row, test.col)
		if !reflect.DeepEqual(val, test.ans) {
			t.Error(
				"For Matrix:", test.mtx,
				"Expected Matrix:", test.ans,
				"Got:", val)
		}

	}
}

/********************************************
 DEL_ROW_COL
*********************************************/

type drc_TestingPair struct {
	matrixPath string
	r          int
	c          int
	matrixSol  [][]float64
}

var drc_tests = []drc_TestingPair{
	{"Tests/matrix11.txt", 1, 2,
		[][]float64{[]float64{0.0, 12.0},
			[]float64{12.0, 0.0}}},
	{"Tests/matrix12.txt", 1, 3,
		[][]float64{[]float64{0, 1},
			[]float64{1, 0}}},
	{"Tests/matrix13.txt", 2, 3,
		[][]float64{[]float64{0, 3, 5, 2, 6.8},
			[]float64{3, 0, 6, 15, 7},
			[]float64{5, 6, 0, 35, 5},
			[]float64{2, 15, 35, 0, 17},
			[]float64{6.7, 7, 5, 17, 0}}},
	{"Tests/matrix14.txt", 0, 3,
		[][]float64{[]float64{0, 43, 4, 6.1, 74},
			[]float64{43, 0, 25, 76, 4},
			[]float64{4, 25, 0, 1.4, 9.2},
			[]float64{6.1, 76, 1.4, 0, 23},
			[]float64{74, 4, 9.2, 23, 0}}},
	{"Tests/matrix15.txt", 2, 3,
		[][]float64{[]float64{0, 16, 67, 25, 18, 2, 6.5},
			[]float64{16, 0, 6, 57, 7, 44, 54},
			[]float64{67, 6, 0, 8, 26, 12, 9.4},
			[]float64{25, 57, 8, 0, 77.4, 4, 18},
			[]float64{18, 7, 26, 77.4, 0, 5, 12},
			[]float64{2, 44, 12, 4, 5, 0, 3},
			[]float64{6.5, 54, 9.4, 18, 12, 3, 0}}},
	{"Tests/matrix16.txt", 0, 4,
		[][]float64{[]float64{0, 2, 18, 4.15, 7, 13, 4},
			[]float64{2, 0, 24, 9, 4, 21, 3},
			[]float64{18, 24, 0, 13, 8.4, 32, 8},
			[]float64{4.15, 9, 13, 0, 17, 4, 18},
			[]float64{7, 4, 8.4, 17, 0, 64, 4},
			[]float64{13, 21, 32, 4, 64, 0, 2.3},
			[]float64{4, 3, 8, 18, 4, 2.3, 0}}}}

func TestDeleteRowCol(t *testing.T) {
	for _, test := range drc_tests {

		mtx, _ := ReadMatrixFromFile(test.matrixPath)
		val := DeleteRowCol(mtx, test.r, test.c)
		if !MtxEq(val, test.matrixSol) {
			t.Error(
				"For Matrix:", mtx,
				"Deleting Row, Col: ", test.r, ",", test.c,
				"Expected: ", test.matrixSol,
				"Got: ", val)
		}
	}
}

func MtxEq(mtx1, mtx2 [][]float64) bool {
	if len(mtx1) != len(mtx2) {
		return false
	}
	for i := range mtx1 {
		if len(mtx1[i]) != len(mtx2[i]) {
			return false
		}
		for j := range mtx1[i] {
			if mtx1[i][j] != mtx2[i][j] {
				return false
			}
		}
	}
	return true
}

/********************************************
 COUNT_LEAVES
*********************************************/

func createTree() Tree {
	T := make([]*Node, 9)

	var NODE0 Node
	var NODE1 Node
	var NODE2 Node
	var NODE3 Node
	var NODE4 Node
	var NODE5 Node
	var NODE6 Node
	var NODE7 Node
	var NODE8 Node

	NODE5.Child1 = &NODE0
	NODE5.Child2 = &NODE1
	NODE6.Child1 = &NODE2
	NODE6.Child2 = &NODE3
	NODE7.Child1 = &NODE6
	NODE7.Child2 = &NODE4
	NODE8.Child1 = &NODE5
	NODE8.Child2 = &NODE7

	T[0] = &NODE0
	T[1] = &NODE1
	T[2] = &NODE2
	T[3] = &NODE3
	T[4] = &NODE4
	T[5] = &NODE5
	T[6] = &NODE6
	T[7] = &NODE7
	T[8] = &NODE8

	return T
}

var TREE = createTree()

type cleaves_TestingPair struct {
	node int
	ans  int
}

var cleaves_tests = []cleaves_TestingPair{
	{1, 1},
	{4, 1},
	{5, 2},
	{7, 3},
	{8, 5}}

func TestCountLeaves(t *testing.T) {
	for _, test := range cleaves_tests {

		val := CountLeaves(TREE[test.node])
		if val != test.ans {
			t.Error(
				"For Node:", test.node,
				"Expected Number of Leaves:", test.ans,
				"Got:", val)
		}

	}
}

/********************************************
 UPGMA
*********************************************/

type upgma_TestingPair struct {
	matrixPath string
	inordSol   string
	postordSol string
}

var upgma_tests = []upgma_TestingPair{
	{"Tests/matrix1.txt", "((i,j));", "((j,i));"},
	{"Tests/matrix2.txt", "(((i,l),(j,k)));", "(((k,j),(l,i)));"},
	{"Tests/matrix3.txt", "((j,(i,(k,l))));", "((((l,k),i),j));"},
	{"Tests/matrix4.txt", "((7,((2,(0,1)),(6,(5,(3,4))))));",
		"((((((4,3),5),6),((1,0),2)),7));"},
	{"Tests/matrix5.txt", "(((k,l),(i,j)));", "(((j,i),(l,k)));"},
	{"Tests/matrix6.txt", "((j,(i,(k,l))));", "((((l,k),i),j));"},
	{"Tests/matrix7.txt", "((m,((k,l),(n,(j,(i,t))))));", "((((((t,i),j),n),(l,k)),m));"},
	{"Tests/matrix8.txt", "(((n,(j,m)),(i,(t,(k,l)))));", "(((((l,k),t),i),((m,j),n)));"},
	{"Tests/matrix9.txt", "((((i,v),(t,z)),((l,n),(m,(j,k)))));",
		"(((((k,j),m),(n,l)),((z,t),(v,i))));"},
	{"Tests/matrix10.txt", "(((i,v),((l,n),(m,(t,(z,(j,k)))))));",
		"(((((((k,j),z),t),m),(n,l)),(v,i)));"},
	{"Tests/matrix14.txt", "(((t,(k,l)),(i,(j,(n,m)))));",
		"(((((m,n),j),i),((l,k),t)));"},
	{"Tests/matrix15.txt", "(((n,m),((z,(i,v)),(l,(t,(j,k))))));",
		"((((((k,j),t),l),((v,i),z)),(m,n)));"},
	{"Tests/matrix16.txt", "(((l,n),((i,v),(m,(t,(z,(j,k)))))));",
		"(((((((k,j),z),t),m),(v,i)),(n,l)));"}}

func TestUPGMA(t *testing.T) {
	for _, test := range upgma_tests {

		mtx, labels := ReadMatrixFromFile(test.matrixPath)
		mtxCopy, _ := ReadMatrixFromFile(test.matrixPath)
		T := UPGMA(mtx, labels)
		strT := ToNewick(T)
		if strT != test.inordSol && strT != test.postordSol {
			t.Error(
				"For Matrix:", mtxCopy,
				"Expected:", test.inordSol,
				"OR", test.postordSol,
				"Got:", strT)
		}
	}
}
