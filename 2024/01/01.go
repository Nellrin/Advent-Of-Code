package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
	"math"
)

func main() {
    input, _ := os.ReadFile("2024.12.01.txt")

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
    nums := make([][]float64, len(lines))

    fmt.Println(len(lines))

    for i, line := range lines {
        line = strings.ReplaceAll(line, "   ", " ")
        parts := strings.Fields(line)

		nums[i] = make([]float64, 2)
        nums[i][0], _ = strconv.ParseFloat(parts[0], 64)
        nums[i][1], _ = strconv.ParseFloat(parts[1], 64)

    }

	for i := 0; i < 1000; i++{
		for j := 0; j < 1000; j++{
			
			if(nums[i][0] > nums[j][0]){
				x := nums[i][0]
				nums[i][0] = nums[j][0]
				nums[j][0] = x
			}

			if(nums[i][1] > nums[j][1]){
				x := nums[i][1]
				nums[i][1] = nums[j][1]
				nums[j][1] = x
			}
		}	
	}
	
	var total float64
	total = 0

	for i := 0; i < 1000; i++{
		total += math.Abs(nums[i][0] - nums[i][1])
	}


	fmt.Println(uint64(total))
}