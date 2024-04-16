package greater_average

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(reader *bufio.Reader) {
	input, _ := reader.ReadString('\n') // Read the input until the newline
	input = strings.TrimSpace(input)
	var a, b, c int64
	_, err := fmt.Sscan(input, &a, &b, &c)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	avg := (float64(a) + float64(b)) / 2.0
	if avg > float64(c) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func Solution() {
	// your code goes here
	reader := bufio.NewReader(os.Stdin)
	var n int64
	input, _ := reader.ReadString('\n') // Read the input until the newline
	input = strings.TrimSpace(input)
	_, err := fmt.Sscan(input, &n) // Parse input into variables
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	for i := int64(0); i < n; i++ {
		solve(reader)
	}
}
