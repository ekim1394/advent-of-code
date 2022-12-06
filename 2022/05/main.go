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

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	if strings.TrimSpace(str) == "" {
		return
	}
	if str == "[" {
		return
	}
	if str == "]" {
		return
	}
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	if (*s).IsEmpty() {
		return ""
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element
}

func (s *Stack) Reverse() {
	if s.IsEmpty() {
		return
	}
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func part_one(instructions []string, crate_map map[int]*Stack, num_stacks int) {
	for _, line := range instructions {
		instruction := strings.Split(line, " ")
		move, err := strconv.Atoi(instruction[1])
		catch(err)
		from, err := strconv.Atoi(instruction[3])
		catch(err)
		to, err := strconv.Atoi(instruction[5])
		catch(err)
		for i := 0; i < move; i++ {
			crate_map[to].Push(crate_map[from].Pop())
		}
	}
}

func part_two(instructions []string, crate_map map[int]*Stack, num_stacks int) {
	for _, line := range instructions {
		instruction := strings.Split(line, " ")
		move, err := strconv.Atoi(instruction[1])
		catch(err)
		from, err := strconv.Atoi(instruction[3])
		catch(err)
		to, err := strconv.Atoi(instruction[5])
		catch(err)
		temp_stack := Stack{}
		for i := 0; i < move; i++ {
			temp_stack.Push(crate_map[from].Pop())
		}
		temp_stack.Reverse()
		for _, val := range temp_stack {
			crate_map[to].Push(val)
		}
	}
}

func setup(data []string) ([]string, map[int]*Stack, int) {
	instructions := []string{}
	crate_map := map[int]*Stack{}
	num_stacks := (len(data[0]) + 1) / 4
	// Setup data
	for i, line := range data {
		crates := strings.Split(line, "")
		if len(crates) == 0 || crates[1] == "1" {
			// This is the end of the stack
			for n := i + 2; n < len(data); n++ {
				instructions = append(instructions, data[n])
			}
			break
		}
		index := 1
		for m := 1; m <= num_stacks; m++ {
			if _, ok := crate_map[m]; !ok {
				crate_map[m] = &Stack{}
			}

			crate_map[m].Push(crates[index])
			index += 4
		}

	}
	for _, stack := range crate_map {
		stack.Reverse()
	}
	return instructions, crate_map, num_stacks
}

// Number of spaces
// 4N - 1
func main() {
	data := load_data("input.txt")
	// Write code here
	instructions, crate_map_1, num_stacks := setup(data)
	part_one(instructions, crate_map_1, num_stacks)
	answer1 := make([]string, num_stacks)
	for k, v := range crate_map_1 {
		answer1[k-1] = (*v)[len(*v)-1]
	}
	fmt.Println("==== Part 1 ====")
	fmt.Println(strings.Join(answer1, ""), "\n")

	instructions, crate_map_2, num_stacks := setup(data)
	part_two(instructions, crate_map_2, num_stacks)
	answer2 := make([]string, num_stacks)
	for k, v := range crate_map_2 {
		answer2[k-1] = (*v)[len(*v)-1]
	}
	fmt.Println("==== Part 2 ====")
	fmt.Println(strings.Join(answer2, ""), "\n")

}
