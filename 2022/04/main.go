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

func set(start int, end int) map[int]bool {
	set := make(map[int]bool)
	for i := start; i <= end; i++ {
		set[i] = true
	}
	return set
}

func string_to_int(str string) int {
	num, err := strconv.Atoi(str)
	catch(err)
	return num
}

func set_diff(set_one map[int]bool, set_two map[int]bool) []int {
	var ret []int
	for k, _ := range set_one {
		if !set_two[k] { // If set_two does not contain k
			ret = append(ret, k)
		}
	}
	return ret
}

func set_intersection(s1 map[int]bool, s2 map[int]bool) []int {
	s_intersection := []int{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1 // better to iterate over a shorter set
	}
	for k, _ := range s1 {
		if s2[k] {
			s_intersection = append(s_intersection, k)
		}
	}
	return s_intersection
}

func main() {
	data := load_data("input.txt")
	// Write code here
	cnt := 0
	for _, line := range data {
		pair := strings.Split(line, ",")
		group_one := pair[0]
		group_two := pair[1]
		seats := strings.Split(group_one, "-")
		set_one := set(string_to_int(seats[0]), string_to_int(seats[1]))

		seats = strings.Split(group_two, "-")
		set_two := set(string_to_int(seats[0]), string_to_int(seats[1]))

		diff1 := set_intersection(set_one, set_two)
		fmt.Println(diff1)

		if len(diff1) > 0 {
			cnt += 1
		}

	}
	fmt.Println(cnt)
}
