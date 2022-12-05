package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func catch(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file, err := os.Open("input.txt")
	catch(err)
	defer file.Close()

	var calorie_cnt int = 0
	var top_3 []int
	top_3 = []int{0, 0, 0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			// Compare with top 3 calories
			for i, cal := range top_3 {
				if calorie_cnt > cal {
					top_3 = append(top_3[:i+1], top_3[i:]...)
					top_3[i] = calorie_cnt
					// Only have 3 elements
					top_3 = top_3[:3]
					break
				}
			}

			// Reset
			calorie_cnt = 0
			continue
		}
		calories, err := strconv.Atoi(scanner.Text())
		catch(err)
		calorie_cnt += calories
	}

	fmt.Println(top_3)
	var total int = 0
	for _, cal := range top_3 {
		total += cal
	}
	fmt.Println(total)
}
