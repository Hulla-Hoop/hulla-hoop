package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetMap := make(map[rune]int)

	for i, let := range alphabet {
		alphabetMap[let] = i + 1
	}

	var data string
	charachter := make(map[rune]int)

	fmt.Fscan(in, &data)
	dataI := strings.SplitN(data, ",", 4)
	dataII := strings.Join(dataI[0:3], "")
	dataIII := strings.Split(dataI[3], ",")
	daySl := strings.Split(dataIII[0], "")
	monthSl := strings.Split(dataIII[1], "")

	var counter int

	for _, d := range daySl {
		day, _ := strconv.Atoi(d)
		counter = counter + day
	}

	for _, m := range monthSl {
		month, _ := strconv.Atoi(m)
		counter = counter + month
	}

	for _, ch := range dataII {
		charachter[ch] = +1
	}

	first := alphabetMap[rune(dataII[0])]

	total := first*256 + counter*64 + len(charachter)
	totalHex := strconv.FormatInt(int64(total), 16)

	var fini []rune

	for i := 3; i > 0; i-- {
		fini = append(fini, rune(totalHex[4-i]))

	}

	fmt.Fprintln(out, string(fini))
}
