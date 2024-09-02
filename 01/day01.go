package main

import (
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	a1 := []string{}
	b1 := [][]string{}
	buffer := ""
	largestSum := 0
	for index, i := range string(dat) {
		if i == 10 {
			a1 = append(a1, buffer)
			buffer = ""
			if i == 10 && dat[index-1] == 10 {
				b1 = append(b1, a1)
				a1 = []string{}
			}
		} else {
			fmt.Println(string(i))
			buffer += string(i)
		}
		if index == len(dat)-1 {
			a1 = append(a1, buffer)
			b1 = append(b1, a1)
		}
	}

	for _, i := range b1 {
		fmt.Println(i)
		sum := 0
		for _, str := range i {
			num, err := strconv.Atoi(str)
			if err != nil {
				continue
			}
			sum += num
		}
		if sum > largestSum {
			largestSum = sum
		}

	}
	fmt.Println("OneStar: Largest Sum:", largestSum)
}
