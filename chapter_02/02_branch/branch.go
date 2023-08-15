package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// if 语句1
func readFile1() {
	wd, _ := os.Getwd()
	fmt.Printf("%s\n", wd)
	filename := "./abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

// if 语句2
func readFile2() {
	filename := "./abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

// switch 语句
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score %d", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	readFile1()
	fmt.Println()
	readFile2()

	fmt.Println(grade(30))
	fmt.Println(grade(60))
	fmt.Println(grade(70))
	fmt.Println(grade(80))
	fmt.Println(grade(90))
	fmt.Println(grade(100))
	// fmt.Println(grade(-1))
	// fmt.Println(grade(101))
}
