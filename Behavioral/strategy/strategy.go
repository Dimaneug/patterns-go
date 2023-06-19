package main

import "fmt"

type SortAlgorithm interface {
	Sort(arr []int)
}

type BubbleSort struct{}

func (bs *BubbleSort) Sort(arr []int) {
	fmt.Println("Bubble sort is completed")
}

type QuickSort struct{}

func (qs *QuickSort) Sort(arr []int) {
	fmt.Println("Quick sort is completed")
}

type Sorter struct {
	sorter SortAlgorithm
}

func (s *Sorter) SortArr(arr []int) {
	s.sorter.Sort(arr)
}

func main() {
	arr := []int{3, 1, 2}

	sorter := Sorter{&BubbleSort{}}
	sorter.SortArr(arr)

	newSorter := Sorter{&QuickSort{}}
	newSorter.SortArr(arr)
}
