package array

import (
	"ds/utility"
	"fmt"
	"math"
)

// function to reverse a given array
// this particular implementation uses recursive approach
func ReverseArray(arr [5]int) {
	utility.Swap(&arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// func to print the smallest and largest element of an array
func MinMaxElement(arr [5]int) {
	fmt.Println("test", arr)
	var minElement = arr[0]
	var maxElement = arr[0]
	for _, v := range arr {
		if v < minElement {
			minElement = v
		} else if v > maxElement {
			maxElement = v
		}
	}
	fmt.Println("Min element:", minElement)
	fmt.Println("Max element:", maxElement)
}

// function to print the kth smallest element in array
func KSmallestElement(arr [5]int) {
	var tempArr [5]int
	minElem := arr[0]
	for _, v := range arr {
		if v < minElem {
			minElem = v
		}
	}
	tempArr[0] = minElem
	i := 1
	for _, v := range arr {
		if v > minElem {
			tempArr[i] = v
			minElem = v
			i++
		}
	}
	fmt.Println(tempArr)
}

// function to sort and array of 1s, 2s, 0s also called the Dutch National flag problemr
func SortZOT(arr [5]int) {
	var low, mid, high = 0, 0, len(arr) - 1
	for mid <= high {
		if arr[mid] == 0 {
			temp := arr[mid]
			arr[mid] = arr[low]
			arr[low] = temp
			low++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else if arr[mid] == 2 {
			temp := arr[mid]
			arr[mid] = arr[high]
			arr[high] = temp
			high--
		}
	}
	fmt.Println(arr)
}

// function to Move all negative numbers to beginning and positive to end
func Move(arr [5]int) {
	low, high := 0, len(arr)-1
	for low < high {
		if arr[low] < 0 {
			low++
		} else if arr[low] > 0 {
			temp := arr[low]
			arr[low] = arr[high]
			arr[high] = temp
			high--
		}
	}
	fmt.Println(arr)
}

//Union and Intersection of sorted arrays
func UnionAndIntersection(arr1 [6]int, arr2 [4]int) {
	var tempArr [10]int
	var counter = 0
	//Union
	//INCOMPLETE
	for i := 0; i < len(tempArr); i++ {
		if i < 6 || arr1[i] == arr2[0] {
			tempArr[i] = arr2[counter]
			counter++
		}

	}
	fmt.Println(tempArr)
	//Intersection
	counter = 0
	var tempArr1 [4]int
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < len(arr1); j++ {
			if arr2[i] == arr1[j] {
				tempArr1[counter] = arr2[i]
				counter++
			}
		}
	}
	fmt.Println(tempArr1)
}

// roteates a given array in clockwise direction ONCE and in-place
func CycRotate(arr []int) {
	l := len(arr)
	arr = append(arr[len(arr)-1:], arr[0:l-1]...)
	fmt.Println(arr)
}

// Kadane's Algorithm OR largest sum contiguous sub-array
func KadanesAlgo(arr []int) {
	var maxTillNow, maxCurrent = math.MinInt, arr[0]
	//maxSum = max(a[i], a[i] + a[i-1])
	for i := 1; i < len(arr); i++ {
		maxCurrent = maxCurrent + arr[i]
		if maxCurrent > maxTillNow {
			maxTillNow = maxCurrent
		}
		if maxCurrent < 0 {
			maxCurrent = 0
		}
	}
	fmt.Println(maxTillNow)
}

// func MinHeight(arr []int) {

// }

func MinJumps(arr []int) {
	var jumpCount = 0
	var i = 1
	for i < len(arr) {
		fmt.Println(i)
		if arr[i] == 0 {
			jumpCount = -1
			break
		} else {
			fmt.Println("herer", arr[i], " ", jumpCount)
			jumpCount++
			i = i + arr[i]
			if arr[i] == arr[(len(arr)-1)] {
				break
			}
		}
	}
	fmt.Println("Min. Jumps = ", jumpCount)
}
