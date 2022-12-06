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

type Buffer []string

func (b *Buffer) queue(elem string, size int) {
	if len(*b) == size {
		(*b).deque(size)
	}
	*b = append(*b, elem)
}

func (b *Buffer) deque(size int) {
	if !(*b).IsEmpty() {
		*b = (*b)[1:size]
	}
}

func (b *Buffer) IsEmpty() bool {
	return len(*b) == 0
}

func (b *Buffer) answer(size int) bool {
	if len(*b) != size {
		return false
	}
	char_map := map[string]bool{}
	for _, char := range *b {
		if _, ok := char_map[char]; ok {
			char_map[char] = true
		} else {
			char_map[char] = false
		}
	}
	for _, v := range char_map {
		if v {
			return false
		}
	}
	return true
}

func main() {
	data := load_data("input.txt")[0]
	var buf1 Buffer
	var buf2 Buffer
	i := 1
	for _, v := range data {
		buf1.queue(strconv.QuoteRune(v), 4)
		if buf1.answer(4) {
			break
		}
		i++
	}
	fmt.Println("==== PART 1 ====")
	fmt.Println(i)

	i = 1
	for _, v := range data {
		buf2.queue(strconv.QuoteRune(v), 14)
		if buf2.answer(14) {
			break
		}
		i++
	}
	fmt.Println("==== PART 2 ====")
	fmt.Println(i)

}
