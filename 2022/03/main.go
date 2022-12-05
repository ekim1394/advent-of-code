package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func catch(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func load_data(path string) []string {
	var input []string
	file, err := os.Open(path)
	catch(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func main() {
	data := load_data("input.txt")
	var priorities int
	for i := 0; i < len(data); i += 3 {
		compartment_one := data[i]
		compartment_two := data[i+1]
		compartment_three := data[i+2]
		fmt.Println(compartment_one, " ", compartment_two, " ", compartment_three)
		for _, char := range compartment_one {
			if strings.ContainsRune(compartment_two, char) && strings.ContainsRune(compartment_three, char) {
				fmt.Println(strconv.QuoteRune(char), char)
				if char >= 97 {
					priorities += int(char) - 96
				} else {
					priorities += int(char) - 38
				}
				break
			}
		}
	}
	fmt.Println(priorities)
}
