package main

import "fmt"

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(nums []int) PosPeaks {
	pp := PosPeaks{
		Pos:   []int{},
		Peaks: []int{},
	}

	for i, j := range nums {
		if !(i == 0 || i == len(nums)-1) {
			previousNum := nums[i-1]
			nextNum := nums[i+1]

			if j > previousNum && j > nextNum {
				pp.Pos = append(pp.Pos, i)
				pp.Peaks = append(pp.Peaks, j)
			}

			if j > previousNum && j == nextNum {
				for x, y := range nums {
					if x <= i {
						// don't start working until we're at the subset
					} else if !(x <= i) && j > y {
						pp.Pos = append(pp.Pos, i)
						pp.Peaks = append(pp.Peaks, j)
						break
					} else if j < y {
						break
					}
				}
			}
		}
	}
	return pp
}

func main() {
	fmt.Println(PickPeaks([]int{1, 2, 3, 4, 4, 3, 2, 1}))                      // {[3] [4]}
	fmt.Println(PickPeaks([]int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}))          // {[3 7] [6 3]}
	fmt.Println(PickPeaks([]int{1, 2, 2, 2, 1}))                               // {[1] [2]}
	fmt.Println(PickPeaks([]int{1, 2, 2, 2, 3}))                               // {[] []}
	fmt.Println(PickPeaks([]int{1, 2, 2, 2, 2}))                               // {[] []}
	fmt.Println(PickPeaks([]int{2, 1, 3, 1, 2, 2, 2, 2, 1}))                   // {[2 4] [3 2]}
	fmt.Println(PickPeaks([]int{2, 14, 13, -5, -5, 6, 4, 9, 0, 1, 1}))         // {[1 5 7] [14 6 9]}
	fmt.Println(PickPeaks([]int{13, 9, -2, -5, 8, 8, 14, -2, -3}))             // {[6] [14]}
	fmt.Println(PickPeaks([]int{6, -1, -2, 5, 13, 6, 9, -2, 3, 3, 5, -5}))     // {[4 6 10] [13 9 5]}
	fmt.Println(PickPeaks([]int{14, 10, 13, 0, 6, 7, 7, -1, 4, 4, 11, 7, -1})) //{[2 5 10] [13 7 11]}
}
