package main

import (
	"fmt"
	"strconv"
)

type xy struct {
	X float64
	Y float64
}

var b0 []float64
var b1 []float64
var entries []xy

func average(s []xy) (float64, float64) {
	averagex := 0.00
	averagey := 0.00
	temp := s
	for i := 0; i < len(entries); i++ {
		averagex = averagex + temp[i].X
	}
	for i := 0; i < len(entries); i++ {
		averagey = averagey + temp[i].Y
	}
	averagex = averagex/float64(len(entries))
	averagey = averagey/float64(len(entries))
	return averagex, averagey
}

func get_number(str string) (string, string) {
	flag := 0
	for i := 0; i < len(str); i++ {
		if str[i] ==  ','{
			flag += 1
			if flag > 1 {
				break
			}
			return str[0 : i], str[i+1 : (len(str))]
		}
	}
	return "", ""
}

func linear_Reg() {
	var beta0 float64
	var beta1 float64
	var tempa float64
	var tempb float64
	tempa = 0
	for i := 0; i < len(entries); i++ {
		tempa = tempa + entries[i].X*entries[i].Y
	}
	a,b := average(entries)
	tempa = tempa - float64(len(entries))*a*b
	tempb = 0
	for i := 0; i < len(entries); i++ {
		tempb = tempb + entries[i].X*entries[i].X
	}
	tempb = tempb - float64(len(entries))*a*a
	beta1 = tempa/tempb
	beta0 = b - beta1*a
	fmt.Println("Intercept: ", beta0, "\nSlope: ", beta1)
}

func main() {
	var userinput_XY xy
	var userinput string
	var tempX string
	var tempY string
	// var learning_Rate float64
	// var epochs int64
	// fmt.Println("Type your preferred learning rate:")
	// fmt.Scanf("%f", &learning_Rate)
	// fmt.Println("Type your preferred epochs:")
	// fmt.Scanf("%d", &epochs)
	for true {
		fmt.Println("Enter X and Y separated by a \",\" (type 'quit' to quit giving input):")
		fmt.Scanf("%s", &userinput)
		if userinput == "quit" {
			break
		}

		tempX, tempY = get_number(userinput)
		if tempX == ""{
			fmt.Println("Invalid format")
			continue
		}
		userinput_XY.X, _ = strconv.ParseFloat(tempX, 3)
		userinput_XY.Y, _ = strconv.ParseFloat(tempY, 3)
		entries = append(entries, userinput_XY)
		b0 = append(b0, 0.0)
		b1 = append(b1, 0.0)
	}
	//fmt.Println("b0: ", len(b0),"\tb1: ", len(b1),"\tentries: ", len(entries),)
	linear_Reg()
}