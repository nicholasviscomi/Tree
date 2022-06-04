package main

import (
	"fmt"
	"os"
)

var pwd string
func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed Getting Current Working Directory!")
		return
	}

	dir, err := os.ReadDir(pwd)
	if err != nil {
		fmt.Println("Error reading current directory!")
		return
	}

	var cont  rune
	numFiles := FullLength(dir, "")
	if numFiles > 100 {
		fmt.Printf("There are %s%s%d elements%s. Would you like to continue? [y/n]\n", Red, Bold, numFiles, Reset)
		_,err = fmt.Scan(&cont)

		if err != nil || cont != 'y' {
			fmt.Println("Exiting...")
			return
		}
	}
	
	PrintDir(dir, 0, "", pwd)
}

var count = 0
func FullLength(dir []os.DirEntry, parent string) int {
	path := pwd
	if parent != "" {
		path = path + "/" + parent
	}

	for _,entry := range dir {
		if entry.IsDir() {
			count++
			//Exit point for recursion
			subdir,_ := os.ReadDir(path)
			FullLength(subdir, entry.Name())
			continue
		}

		if entry.Name() == ".DS_Store" { continue }
		count++
	}
	return count
}

/*
dir = the directory being listed out
err = the return value from reading a directory
level = how many subdirectories have been explored. 
	--> Used to format the output with tabs
*/
func PrintDir(dir []os.DirEntry, level int, parent, pwd string) {
	path := pwd
	if parent != "" {
		path = pwd + "/" + parent
	}

	for _,entry := range dir {
		for i := 0; i < level; i++ {
			fmt.Print("  ")
		}

		if entry.IsDir() {
			//Exit point for recursion
			subdir,_ := os.ReadDir(path)

			fmt.Printf("%s%s——> %s\n%s", Bold, Green, entry.Name(), Reset)
			PrintDir(subdir, level + 1, entry.Name(), pwd)
			continue
		}

		if entry.Name() == ".DS_Store" { continue }

		
		fmt.Printf("——> %s\n", entry.Name())
	}
}