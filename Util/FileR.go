package Util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLineByMap(filePth string) (data map[string]int, err error) {
	file, err := os.Open(filePth); if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data = make(map[string]int)

	countLines(file,data)

	return data, err
}

func countLines(f *os.File,counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err()
}

func ReadLineBySlice(filePth string) (lines []string, err error) {
	file, err := os.Open(filePth)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println("read lines:", len(lines))
	//for _, line := range lines {
	//	fmt.Println(line)
	//}
	return
}
