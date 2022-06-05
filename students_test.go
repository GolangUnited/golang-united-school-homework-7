package coverage

import (
	"os"
	"testing"
	"time"

)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLen(t *testing.T) {
	person01 := Person {"Jon", "Karter", time.Date(1982, time.February,15,0,0,0,0, time.UTC)}
	person02 := Person {"Billy", "Wiliams", time.Date(1988, time.September,01,0,0,0,0, time.UTC)}
	person03 := Person {"Ben", "Pupkin", time.Date(1982,time.June,29,0,0,0,0, time.UTC)}
	peopleList := People{person01, person02, person03}
	if peopleList.Len() != 3{
		//req.NotEqual(peopleList.Len(), 3)
		//t.Errorf("fail orders lenght")
	}
}

func TestLess(t *testing.T){
	person01 := Person {"Jon", "Karter", time.Date(1982, time.February,15,0,0,0,0, time.UTC)}
	person02 := Person {"Jon", "Pup", time.Date(1982,time.February,15,0,0,0,0, time.UTC)}
	person03 := Person {"Jeyn", "Pupis", time.Date(1982,time.February,15,0,0,0,0, time.UTC)}
	person04 := Person {"Billy", "Wiliams", time.Date(1982, time.June,29,0,0,0,0, time.UTC)}
	peopleList := People{person01, person02, person03, person04}
	
	peopleList.Less(0,1)
	peopleList.Less(0,2)
	peopleList.Less(0,3)

}

func TestSwap(t *testing.T){
	person01 := Person {"Jon", "Karter", time.Date(1982, time.February,15,0,0,0,0, time.UTC)}
	person02 := Person {"Jon", "Pup", time.Date(1982,time.February,15,0,0,0,0, time.UTC)}
	peopleList := People{person01, person02}

	peopleList.Swap(0,1)
	
}
//////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestNew(t *testing.T){
	str01 := "sequoia"
	str02 := "1\n2 3"
	str03 := "3"
	New(str01)
	New(str02)
	New(str03)
}

func TestRows(t *testing.T) {
	mx01 := Matrix{1,1,[]int{0}}
	mx01.Rows()
}

func TestCols(t *testing.T) {
	mx01 := Matrix{1,1,[]int{0}}
	mx01.Cols()
}

func TestSet(t *testing.T) {
	mx01 := Matrix{1,1,[]int{0}}
	mx01.Set(1,1,1)
	mx01.Set(0,0,0)
	}