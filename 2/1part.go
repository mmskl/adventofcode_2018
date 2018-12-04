package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()


    var line string
    // var num int

        fileHandle.Seek(0, 0)
        fileScanner := bufio.NewScanner(fileHandle)

        masterset := make(map[int]int)

        for fileScanner.Scan() {

            line = fileScanner.Text()
            set := make(map[string]int)
            for i := 0; i < len(line); i++ {
                ch := string(rune(line[i]))

                if _, ok := set[ch]; ok {
                    set[ch]++
                } else {
                    set[ch] = 1
                }
            }
            verifyset := make(map[int]bool)

            for _, v := range set {
                if v <= 1 {
                    continue
                }


                if _, ok := verifyset[v]; ok {
                   continue
                } else {
                    verifyset[v] = true
                }

                if _, ok := masterset[v]; ok {
                   masterset[v]++
                } else {
                    masterset[v] = 1
                }
            }

        }

        var result int
        result = 1

        fmt.Println(masterset)

        for k, v := range masterset {
            if k == 3 || k == 2 {
                result *= v
            }
        }

        fmt.Println(result)


}
