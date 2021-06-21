package main

import (
	"fmt"
	"os"
	"os/exec"
)

func PrettyPrintArrays(array [ROW_SIZE][COLUMN_SIZE]int) {
	for i := 0; i < ROW_SIZE; i++ {
		for j := 0; j < COLUMN_SIZE; j++ {
			fmt.Printf("%d\t", array[i][j])
		}
		fmt.Printf("\n")
 	}
}

func ClearConsole() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}