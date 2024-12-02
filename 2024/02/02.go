package main

import (
    "fmt"
    "os"
    "strings"
	"strconv"
	"math"
)

func GetLines(filePath string) ([]string,int){
    input, _ := os.ReadFile(filePath)
	lines := strings.Split(string(input), "\n")
	return lines, len(lines)
}

func main() {

	x,y := GetLines("2024.12.02.txt")
	total := y

	for i := 0; i < y; i++{
		j := 1

		x[i] = strings.TrimSpace(x[i])
		zz := strings.Split(x[i], " ")
		z:=make([]int64,len(zz))

		for j := 0; j < len(zz); j++{
			z[j],_ = strconv.ParseInt(zz[j],10,64)
		}


		for ; j < len(zz); j++{
			diff := z[j-1] - z[j]

			condition1 := math.Abs(float64(diff)) > 3
			condition2 := z[j-1] - z[j] == 0
			condition3 := (diff < 0) != (z[0] - z[1] < 0)

	
			if condition1 || condition2 || condition3 {
				total --
				break
			}
		}
	}
	fmt.Println(total)
}