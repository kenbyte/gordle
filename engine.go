package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func check_error(e error) {
	if e != nil {
		panic(e)
	}
}

func get_words(i string, w string) {

	getcwd, err := os.Getwd()
	check_error(err)
	input_file, err := os.Open(i)
	check_error(err)
	defer input_file.Close()

	output_file := filepath.Join(getcwd, w)
	output, err := os.Create(output_file)
	check_error(err)

	scanner := bufio.NewScanner(input_file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 5 {
			_, err := output.WriteString(line + "\n")
			check_error(err)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func compare(targeted []byte, tried []byte) {
	if len(targeted) != len(tried) {
		fmt.Println("The arrays are not of the same length.")
		return
	}

	used := make([]bool, len(targeted))
	result := make([]int, len(tried))

	for i := range tried {
		if tried[i] == targeted[i] {
			result[i] = 2
			used[i] = true
		}
	}

	for i := range tried {
		if result[i] == 2 {
			continue
		}
		for j := range targeted {
			if !used[j] && tried[i] == targeted[j] {
				result[i] = 1
				used[j] = true
				break
			}
		}
	}

	for i, ch := range tried {
		switch result[i] {
		case 2:
			fmt.Printf("\033[32m%c\033[0m ", ch)
		case 1:
			fmt.Printf("\033[33m%c\033[0m ", ch)
		default:
			fmt.Printf("\033[31m%c\033[0m ", ch)
		}
	}
	fmt.Println()
}
