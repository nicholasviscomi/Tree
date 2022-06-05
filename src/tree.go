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
	numFiles := FullLength(dir, pwd)
	if numFiles > 100 {
		fmt.Printf("There are %s%s%d elements%s. Would you like to continue? [y/n]\n", Red, Bold, numFiles, Reset)
		_,err = fmt.Scan(&cont)

		if err != nil || cont != 'y' {
			fmt.Println("Exiting...")
			return
		}
	}
	
	PrintDir(dir, 0, pwd)
}

var count = 0
func FullLength(dir []os.DirEntry, pwd string) int {
	for _,entry := range dir {
		if entry.Name() == ".DS_Store" { continue }
		if entry.Name() == ".git" { continue }

		if entry.IsDir() {
			//Exit point for recursion
			path := pwd + "/" + entry.Name()
			subdir,err := os.ReadDir(path)
			if err != nil {
				fmt.Println("length: error reading sub directory!")
			}

			FullLength(subdir, path)
			continue
		}
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
func PrintDir(dir []os.DirEntry, level int, pwd string) {
	for _,entry := range dir {
		for i := 0; i < level; i++ {
			fmt.Print("  ")
		}
		if entry.Name() == ".DS_Store" { continue }
		if entry.Name() == ".git" { continue }

		if entry.IsDir() {
			//Exit point for recursion
			fmt.Printf("%s%s——> %s\n%s", Bold, Green, entry.Name(), Reset)

			path := pwd + "/" + entry.Name()
			subdir,err := os.ReadDir(path)
			if err != nil {
				fmt.Println("printdir: error reading sub directory!")
			}

			PrintDir(subdir, level + 1, path)
			continue
		}


		fmt.Printf("——> %s\n", entry.Name())
	}
}