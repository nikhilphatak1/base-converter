package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

func main() {
	/*
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("start number: ")
	start, _ := reader.ReadString('\n')
	start = strings.TrimSuffix(num, "\n")
	fmt.Print("start base: ")
	start, _ := reader.ReadString("\n")
	start = strings.TrimSuffix(num, "\n")
	fmt.Print("end base")
	*/
	n := getVal("start number: ")
	start := getVal("start base: ")
	end := getVal("end base: ")
	fmt.Println(n)
	fmt.Println(start)
	fmt.Println(end)

	

}

func getVal(input string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(input)
	val, _ := reader.ReadString('\n')
	val = strings.TrimSuffix(val, "\n")
	
	i, err := strconv.Atoi(val)
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
	}
	
	return i
}