package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var TLOC int
var functoLOCmap = make(map[string]int)
var funcLOCs []int
var funcnum int

type user struct {
	userinput_obj string
}

//PROXY BEGIN printMap
func printMap(mp map[string]int) {
	var maxLenKey int
	for k, _ := range mp {
		if len(k) > maxLenKey {
			maxLenKey = len(k)
		}
	}
	fmt.Println("    Function:	LOC")
	fmt.Println("__________________")
	for k, v := range mp {
		fmt.Printf("%*s:	%d\n", maxLenKey, k, v)
	}
}
//PROXY END

//PROXY BEGIN refresh
func refresh() {
	funcnum = 0
	functoLOCmap = make(map[string]int)
	funcLOCs = nil
	TLOC = 0
}
//PROXY END

//PROXY BEGIN readFile
func readFile(filename string) {
	if filename == "quit" || filename == "q" {
		fmt.Println("Exiting")
		os.Exit(0)
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("ERROR: %s", err)
	}
	fmt.Printf("\nData: %s\n", data)
}

//PROXY END

//PROXY BEGIN findFuncName
func findFuncName(linestr string) {
	// Iterative method:
	// for i := 0; (i + 14) < len(linestr); i++ {
	// 	funcname = funcname + string([]rune(linestr)[14+i]) 
	// 	// fmt.Println(funcname)
	// }

	// Slice method
	funcname := linestr[14:]
	//fmt.Println("func: ",funcname)
	functoLOCmap[funcname] = funcLOCs[funcnum]
}
//PROXY END

//PROXY BEGIN this_Line_Contains
func this_Line_Contains(linestr string, strcheck string) bool {
	if len(linestr)<len(strcheck){
		return false
	}
	// Iterative method
	// proxycheck := ""
	// for i := 0; i < len(strcheck); i++ {
	// 	proxycheck = proxycheck + string([]rune(linestr)[i])
	// }

	// Slice method
	proxycheck := linestr[0:len(strcheck)]
	// fmt.Println("check:",proxycheck)
	if proxycheck == strcheck {
		return true
	}
	return false
}
//PROXY END

//PROXY BEGIN linesInFile
func linesInFile(fileName string) (int, int, int, error) {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	loc := 0
	bloc := 0
	cloc := 0
	proxyB := loc
	tempLine := ""
	funcnum = 0
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		line = string(line)
		line = strings.Trim(line, " \t")
		sc := scanner
		if this_Line_Contains(line, "//") {
			cloc++
		}
		if this_Line_Contains(line, "//PROXY BEGIN") { // I coded this initially while keeping in mind that no function name will be provided after // PROXY BEGIN.
			proxyB = loc // As a result I extracted the name from the function declaration below it .
			// sc.Scan() 							 // we can use a similar strstr() function in <string.h> and rfind in C++
			tempLine = string(sc.Text()) // To use the method above, just uncomment the line 118
		}
		if this_Line_Contains(line, "//PROXY END") { // we can use the strstr() function in string.h and rfind in C++
			proxyE := loc
			funcLOCs = append(funcLOCs, (proxyE - proxyB - 1))
			// fmt.Println(tempLine)
			findFuncName(tempLine)
			funcnum++
		}
		if len(line) == 0 {
			bloc++
			continue
		}
		loc++
	}
	printMap(functoLOCmap)
	return loc, bloc, cloc, scanner.Err()
}

//PROXY END

//PROXY BEGIN main
func main() {
	var userinput string
	for true {
		fmt.Println("Type your input (filenames only) (type 'quit' to quit giving input)")
		fmt.Scanf("%s", &userinput)
		filename := userinput
		readFile(filename)
		TLOC, bloc, cloc, err := linesInFile(filename)
		if err != nil {
			fmt.Printf("ERROR: %s", err)
		}
		fmt.Println("LOC: ", TLOC, ", BLOC: ", bloc, ", CLOC: At least ", cloc)
		refresh()
	}
}
//PROXY END