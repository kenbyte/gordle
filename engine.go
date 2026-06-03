package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func check_error (e error) {
	if (e != nil){
		panic(e)
	}
}

func get_words(i string, w string){

	getcwd,err := os.Getwd()
	check_error(err)
	input_file, err := os.Open(i)
	check_error(err)
	defer input_file.Close()

	output_file:= filepath.Join(getcwd,w)
	output,err := os.Create(output_file)
	check_error(err)
	
	scanner := bufio.NewScanner(input_file)

	for scanner.Scan() {
		line := scanner.Text()
		if (len(line) == 5 ) {
			_,err := output.WriteString(line + "\n") 
			check_error(err)
		}
		
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}


