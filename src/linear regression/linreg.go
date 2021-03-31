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

func average(s []float64) float64 {
	average := 0.00
	entries := s
	for i := 0; i < len(entries); i++ {
		average = average + entries[i]
	}
	average = average / float64(len(entries))
	return average
}

func get_number(str string) (string, string) {
	flag := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ',' {
			flag += 1
			if flag > 1 {
				break
			}
			return str[0:i], str[i+1 : (len(str))]
		}
	}
	return "", ""
}

func linear_Reg(learning_rate float64, epochs int64) {
	var i int64
	var temperr float64
	var temperrX float64

	for i = 0; i < epochs; i++ {
		var error []float64
		var y []float64
		for i := 0; i < len(b0); i++ {
			y = append(y, (b0[i] + b1[i]*(float64(entries[i].X))))
			error = append(error, (y[i] - entries[i].Y))
		}
		temperr = 0
		temperrX = 0
		for j := 0; j < len(b0); j++ {
			for k := 0; k < len(entries); k++ {
				temperr = temperr + error[k]
				temperrX = temperrX + error[k]*(float64(entries[k].X))
			}
			temperr = temperr / float64(len(error))
			temperrX = temperrX / float64(len(error))

			b0[j] = b0[j] - learning_rate*2*temperr
			b1[j] = b1[j] - learning_rate*2*temperrX
		}
	}
	b0avg := average(b0)
	b1avg := average(b1)
	fmt.Println("Intercept: ", b0avg, "\nSlope: ", b1avg)

}

func main() {
	var userinput_XY xy
	var userinput string
	var tempX string
	var tempY string
	var learning_Rate float64
	var epochs int64
	fmt.Println("Type your preferred learning rate:")
	fmt.Scanf("%f", &learning_Rate)
	fmt.Println("Type your preferred epochs:")
	fmt.Scanf("%d", &epochs)
	for true {
		fmt.Println("Enter X and Y separated by a \",\" (type 'quit' to quit giving input):")
		fmt.Scanf("%s", &userinput)
		if userinput == "quit" {
			break
		}

		tempX, tempY = get_number(userinput)
		if tempX == "" {
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
	linear_Reg(learning_Rate, epochs)
}
