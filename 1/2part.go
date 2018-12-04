package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)

func main() {

	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

    var line string
    var num int

    set := make(map[int]bool)

    for {

        fileHandle.Seek(0, 0)
        fileScanner := bufio.NewScanner(fileHandle)

        for fileScanner.Scan() {
            line = fileScanner.Text()
            if line[0] == '-' {
                i, _ := strconv.Atoi(line[1:len(line)])
                num = num - i
            } else {
                i, _ := strconv.Atoi(line[0:len(line)])
                num = num + i
            }

            if _, ok := set[num]; ok {
                fmt.Println("-------")
                fmt.Println(num)
                os.Exit(0)
            }

            set[num] = true
        }
    }

}
