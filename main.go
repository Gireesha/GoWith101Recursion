package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	//Statements to run the test
	//fmt.Println(reverseString_001("Hello World!")) // "Hello World!" ==> "!dlroW olleH"
	//fmt.Println(printNaturalNumbersDesc_002(12))   // 12 ==> "12 11 10 9 8 7 6 5 4 3 2 1"
	//fmt.Println(printNaturalNumbersAsc_003(12))    // 12 ==> "1 2 3 4 5 6 7 8 9 10 11 12"

	//arr := [10]float64{67.8, 29.23, 35.22, 784.19, 5.6, 34.48, 39.55, 67.33, 689.33, 6.37}
	//fmt.Println(printSmallestNumberInAnArray_004(arr[:], 10)) // [] ==> 5.6

	//printOddNEvenNumbers_005(1, 3) // 1 is Odd number, 2 is Even number, 3 is Odd number

	arr := [10]int{67, 29, 345, 7849, 562, 13448, 395, 633, 6893, 37}
	printOddNEvenNumbersInAnArray_005(arr[:], 10)
}

func reverseString_001(raw string) (reversed string) {
	//reverseString_001("123") = "3" + reverseString_001("12")
	//reverseString_001("12") = "2" + reverseString_001("1")
	//reverseString_001("1") = "1"  Base condition

	if len([]rune(raw)) == 0 { //Base condition
		return raw
	}
	return raw[len(raw)-1:] + reverseString_001(raw[:len(raw)-1])
}

func printNaturalNumbersDesc_002(upperLimit int) (series string) {
	//printNaturalNumbersAsc_002(3) = "3" + printNaturalNumbersAsc_002(2)
	//printNaturalNumbersAsc_002(2) = "2" + printNaturalNumbersAsc_002(1)
	//printNaturalNumbersAsc_002(1) = "1" Base condition

	if upperLimit == 1 {
		return "1"
	}
	return strconv.Itoa(upperLimit) + " " + printNaturalNumbersDesc_002(upperLimit-1)
}

func printNaturalNumbersAsc_003(upperLimit int) (series string) {
	//printNaturalNumbersAsc_003(3) = "3" + printNaturalNumbersAsc_003(2)
	//printNaturalNumbersAsc_003(2) = "2" + printNaturalNumbersAsc_003(1)
	//printNaturalNumbersAsc_003(1) = "1" Base condition

	if upperLimit == 1 {
		return "1"
	}
	return printNaturalNumbersAsc_003(upperLimit-1) + " " + strconv.Itoa(upperLimit)
}

func printSmallestNumberInAnArray_004(arrInput []float64, size int) float64 {
	//printSmallestNumberInAnArray_004([45,72,33,49], 4) = "49" + printSmallestNumberInAnArray_004([45,72,33,49],3)
	//printSmallestNumberInAnArray_004([45,72,33,49], 3) = "33" + printSmallestNumberInAnArray_004([45,72,33,49],2)
	//printSmallestNumberInAnArray_004([45,72,33,49], 2) = "72" + printSmallestNumberInAnArray_004([45,72,33,49],1)
	//printSmallestNumberInAnArray_004([45,72,33,49], 1) = "49" Base condition

	if size == 1 {
		return arrInput[0]
	}
	//fmt.Println("arrInput[size-1] = ", arrInput[size-1], "size = ", size)
	return math.Min(arrInput[size-1], printSmallestNumberInAnArray_004(arrInput, size-1))
}

func printOddNEvenNumbers_005(startNumber int, endNumber int) {
	//Print the odd and even # between the number range(startNumber and endNumber)
	//printOddNEvenNumbers_005(1,4) = printOddNEvenNumbers_005(2,4)
	//printOddNEvenNumbers_005(2,4) = printOddNEvenNumbers_005(3,4)
	//printOddNEvenNumbers_005(4, 4) = print; Base condition

	if startNumber == endNumber {
		if startNumber%2 == 0 {
			fmt.Println(startNumber, "is Even number")
		} else {
			fmt.Println(startNumber, "is Odd number")
		}
		return
	}

	if startNumber%2 == 0 {
		fmt.Println(startNumber, "is Even number")
	} else {
		fmt.Println(startNumber, "is Odd number")
	}

	printOddNEvenNumbers_005(startNumber+1, endNumber)
}

func printOddNEvenNumbersInAnArray_005(inputArray []int, size int) {
	//printOddNEvenNumbersInAnArray_005([],4) = printOddNEvenNumbersInAnArray_005([],3)
	//printOddNEvenNumbersInAnArray_005([],3) = printOddNEvenNumbersInAnArray_005([],2)
	//printOddNEvenNumbersInAnArray_005([],2) = printOddNEvenNumbersInAnArray_005([],1)
	//printOddNEvenNumbersInAnArray_005(1, 1) = just print and exit; Base condition

	if size == 1 {
		if inputArray[size-1]%2 == 0 {
			fmt.Println(inputArray[size-1], " is Even number")
		} else {
			fmt.Println(inputArray[size-1], " is Odd number")
		}
		return
	}

	if inputArray[size-1]%2 == 0 {
		fmt.Println(inputArray[size-1], " is Even number")
	} else {
		fmt.Println(inputArray[size-1], " is Odd number")
	}

	printOddNEvenNumbersInAnArray_005(inputArray, size-1)
}
