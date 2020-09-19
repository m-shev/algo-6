package main

import (
	"algo-6/sorting"
	"algo-6/utils"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var testArr = []int{7, 0, 6, 1, 3, 2, 8, 5, 4, 9}

func main() {
	utils.GenerateFile("test", 1_000_000, 1, math.MaxInt16)
	//t := utils.BytesToInt16(utils.Read("test"))
	//fmt.Println("test before sorting", utils.BytesToInt16(utils.Read("test")))

	//fmt.Println(sorting.MergeSort{}.Sort(x))
	//fmt.Println("quick sort", sorting.MergeSort{}.Sort(t))
	start := time.Now()
	//sorting.External{MaxMemoryUse: 4096, InternalSort: sorting.Shell{
	//	StepType: sorting.SedgewickSteps,
	//}}.Sort("test")
	sorting.External{MaxMemoryUse: 4096, InternalSort: sorting.Quick{}}.Sort("test")
	//sorting.Shell{StepType: sorting.SedgewickSteps}.Sort(t)
	stop := time.Since(start)
	//fmt.Println("test after sorting", utils.BytesToInt16(utils.Read("test")))
	fmt.Println("execution time", stop)
	//utils.SplitFile("test", "1", "2", 10)
	//fmt.Println(utils.BytesToInt16(utils.Read("2")))

	//sorting.MergeSort{}.Sort(testArr)
	//fmt.Println(testArr)
}

func randomString(l int) []int {
	bytes := make([]int, l)
	for i := 0; i < l; i++ {
		bytes[i] = randInt(65, 90)
	}
	return bytes
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
