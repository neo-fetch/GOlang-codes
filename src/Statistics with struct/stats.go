package main

import (
    "fmt"
    "os"
	"strconv"
	"math"
	"sort"
)
var entries []float64

type user struct {
	userinput_obj string
	userinputd float64
	inputlength_obj int
}

//PROXY BEGIN
func (r user) isDigit(ch rune) bool {
	if ch >= '0' && ch <= '9' {
		return true
	}
	return false
}
//PROXY END

//PROXY BEGIN
func (r user) exitApp() {
	fmt.Println("Exiting")
	os.Exit(0)
}
//PROXY END

//PROXY BEGIN
func (r user) verifyInputAllDigits() bool {
	str := r.userinput_obj
	length := r.inputlength_obj

	if length > 16 {
		fmt.Println("Input range exceeded!")
		return false
	}
	
	if str == "-." || str == ".-" {
		fmt.Println("what?")
		return false
	}

	if str == "quit"{
		r.exitApp();
	}

	decimalpoint := 0
	negativeflag := 1

	for i := 0; i < length; i++ {
		if str[i] == '.' {
			decimalpoint++
			if decimalpoint > 1 || i == 0 {
				fmt.Println("Additional decimal point entered./ Decimal point entered at the start.")
				return false
			} else{
				fmt.Println("Decimal space entered")
				continue
			}
		}
		if str[i] == '-' {
			negativeflag = negativeflag - 2
			fmt.Println("Negative space entered.")
			if negativeflag < -2 || i != 0 {
				fmt.Println("Additional negative symbol entered./ negative '-' not entered at the start.\n")
				return false	
			} else{
				continue
			}
		}

		if r.isDigit(rune(str[i])) { // If the current character is a digit
			fmt.Printf("%c is a valid digit\n", str[i])
			continue
		} else{
			fmt.Printf("%c is an invalid digit. Please check.\n", str[i])
			return false
		}
	}

	fmt.Println("All clear!\n Updated list:")
	r.userinputd, _ = strconv.ParseFloat(str, 32)
	entries = append(entries, r.userinputd)
	for i := 0; i < len(entries); i++ {
		fmt.Print(entries[i]," ")
	}
	fmt.Println("\n")
	return true
}
//PROXY END

//PROXY BEGIN
func (r user) average(s []float64) float64 {
	average := 0.00
	entries = s
	for i := 0; i < len(entries); i++ {
		average = average + entries[i]
	}
	average = average/float64(len(entries))
	return average
}
//PROXY END

//PROXY BEGIN
func (r user) standardDeviationSampled(s []float64) float64 {
	average := 0.00
	entries = s
	for i := 0; i < len(entries); i++ {
		average = average + entries[i]
	}
	average = average/float64(len(entries))
	sumSq := 0.00
	for i := 0; i < len(entries); i++ {
		sumSq = sumSq + (entries[i] - average)*(entries[i] - average) //sample standard deviation = squareroot(deviations/(size-1))
	}
	samstd := math.Sqrt(sumSq/float64(len(entries)-1))
	return samstd
}
//PROXY END

//PROXY BEGIN
func (r user) median(s []float64) float64 {
	entries = s
	n := len(entries)/2
	if n == 0 {
		fmt.Println("Nothing present")
		return -1;
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i] < entries[j]
	})
	if n%2==0 {
		return (entries[n] + entries[n-1])/2.0
	} else{
		return (entries[n])
	}
}
//PROXY END

//PROXY BEGIN
func main() {
	var userinput string
	for true {
		fmt.Println("Type your input (integers, decimals only) tip: You can press Enter ⏎ to automatically input previous input: \n(type 'quit' to quit giving input)")
		fmt.Scanf("%s", &userinput)
		fmt.Println("verifying input...")
		inputlen := len(userinput)
		r := user{userinput_obj: userinput , inputlength_obj: inputlen }
		if r.verifyInputAllDigits(){
			fmt.Println("Verified ✓. Proceeding...")
			fmt.Println("Average: ", r.average(entries))
			fmt.Println("Sample standard deviation: ", r.standardDeviationSampled(entries))
			fmt.Println("Median: ", r.median(entries))
		} else{
			fmt.Println("Invaild entry ✘. Re-enter input:")
		}

	}
	
}
//PROXY END