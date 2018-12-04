package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
    fileScanner := bufio.NewScanner(fileHandle)

	fileHandle2, _ := os.Open("input.txt")
	defer fileHandle2.Close()
    fileScanner2 := bufio.NewScanner(fileHandle2)

    var line string
    var line2 string
    checker := 0


        for fileScanner.Scan() {

            line = fileScanner.Text()

            for fileScanner2.Scan() {

                line2 = fileScanner2.Text()
                if line == line2 {
                    continue
                }

                checker = 0

                for i := 0; i < len(line); i++ {
                    ch := string(rune(line[i]))
                    ch2 := string(rune(line2[i]))

                    if checker > 1 {
                        break
                    }

                    if ch != ch2 {
                        checker++
                    }

                }

                if checker <= 1 {
                    fmt.Println(line)
                    fmt.Println(line2)
                    os.Exit(0)
                }


            }

            fileHandle2.Seek(0, 0)
            fileScanner2 = bufio.NewScanner(fileHandle2)


        }


}
