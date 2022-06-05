package main

import (
	"fmt"
	"math/rand"
	"os"
)

var pwd string
func main() {
	pwd := os.Args[1]
	// if err != nil {
	// 	fmt.Println("Failed Getting Current Working Directory!")
	// 	return
	// }

	dir, err := os.ReadDir(pwd)
	if err != nil {
		fmt.Println("Error reading current directory!")
		return
	}

	var cont  rune
	numFiles := FullLength(dir, pwd)
	if numFiles > 100 {
		fmt.Printf("There are %s%s%d elements%s. Exit now if you wish\n", 
						colors[Red], colors[Bold], numFiles, colors[Reset])
		_,err = fmt.Scanln(&cont)
	}
	
	PrintDir(dir, 0, pwd, 10)
}

/*
dir = the directory being listed out
err = the return value from reading a directory
level = how many subdirectories have been explored. 
	--> Used to format the output with tabs
*/
func PrintDir(dir []os.DirEntry, level int, pwd string, parentColor int) {
	for _,entry := range dir {
		for i := 0; i < level; i++ {
			fmt.Print("  ")
		}
		if entry.Name() == ".DS_Store" { continue }
		if entry.Name() == ".git" { continue }

		currColorInd := rand.Intn(7) + 2
		if entry.IsDir() {
			//Exit point for recursion
			fmt.Printf("%s%s└──%s\n%s", 
			colors[Bold], colors[currColorInd], entry.Name(), colors[Reset])

			path := pwd + "/" + entry.Name()
			subdir,err := os.ReadDir(path)
			if err != nil {
				fmt.Println("printdir: error reading sub directory!")
			}

			PrintDir(subdir, level + 1, path, Green)
			continue
		}


		fmt.Printf("├──%s\n", entry.Name())
	}
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