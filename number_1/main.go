package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan jumlah string: ")
	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("Input jumlah string tidak valid")
		return
	}

	stringsList := make([]string, n)
	fmt.Println("Masukkan string satu per satu:")
	for i := 0; i < n; i++ {
		str, _ := reader.ReadString('\n')
		stringsList[i] = strings.TrimSpace(str)
	}

	result := ValidateStringInput(n, stringsList)
	fmt.Println("Output:", result)

}

func ValidateStringInput(n int, stringsList []string) string {
	checked := make(map[string][]int)
	maxOccurrences := 0
	maxString := ""

	for i := 0; i < n; i++ {
		lowerStr := strings.ToLower(stringsList[i])
		checked[lowerStr] = append(checked[lowerStr], i+1)

		if len(checked[lowerStr]) > maxOccurrences {
			maxOccurrences = len(checked[lowerStr])
			maxString = lowerStr
		}
	}

	if maxOccurrences > 1 {
		return fmt.Sprintf("%v", checked[maxString])
	}

	return "false"
}
