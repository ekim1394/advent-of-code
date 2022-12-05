package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func catch(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Rock Paper Scissor
// A B C
// X Y Z

// A > Z; A < Y; A == X
// B > X; B < Z; B == Y
// C > Y; C < X; C == Z
func main() {
	// var score int = 0
	var game string
	// var scores []int
	var elf string
	var me string
	var total_score int

	var scoring_map map[string]map[string]int
	scoring_map = map[string]map[string]int{
		"A": {"X": 3, "Y": 4, "Z": 8},
		"B": {"X": 1, "Y": 5, "Z": 9},
		"C": {"X": 2, "Y": 6, "Z": 7},
	}

	// var play_map map[string]int
	// play_map = map[string]int{
	// 	"X": 1,
	// 	"Y": 2,
	// 	"Z": 3,
	// }

	file, err := os.Open("input.txt")
	catch(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total_score = 0
	for scanner.Scan() {
		game = scanner.Text()
		elf = string(game[0])
		me = string(game[2])
		total_score += scoring_map[elf][me]

	}

	fmt.Println(total_score)
}
