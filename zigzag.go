package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("input text :")
	input := bufio.NewReader(os.Stdin)

	var lines []string

	for {
		line, err := input.ReadString('\n')

		if err != nil {
			panic(err)
		}

		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		lines = append(lines, line)
	}

	fmt.Println(lines[2])
	fmt.Println("output :")
	for _, v := range lines {
		fmt.Println(v)
	}
	var arr []int
	stringInt := strings.Split(strings.TrimSpace(lines[2]), " ")
	for _, v := range stringInt {
		value, err := strconv.Atoi(string(v))

		if err != nil {
			panic(err)
		}

		arr = append(arr, value)
	}
	fmt.Println("output array:")
	fmt.Println(arr)
	fmt.Println("output string:")
	fmt.Println(string(arr))

}
