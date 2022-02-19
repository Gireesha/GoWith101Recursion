package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(reverseString_001("Hello World!")) // "Hello World!" ==> "!dlroW olleH"
	fmt.Println(printNaturalNumbersDesc_002(12))   // 12 ==> "12 11 10 9 8 7 6 5 4 3 2 1"
	fmt.Println(printNaturalNumbersAsc_003(12))    // 12 ==> "1 2 3 4 5 6 7 8 9 10 11 12"

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
