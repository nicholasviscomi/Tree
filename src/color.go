package main

const (
	Reset  = 0
	Red    = 1
	Green  = 2
	Yellow = 3
	Blue   = 4
	Purple = 5
	Cyan   = 6
	Gray   = 7
	White  = 8
	Bold   = 9
	Underline = 10
)

var colors = []string{
	"\033[0m",
	"\033[31m",
	"\033[32m",
	"\033[33m",
	"\033[34m",
	"\033[35m",
	"\033[36m",
	"\033[37m",
	"\033[97m",
	"\033[1m", 	
	"\u001b[4m",
}