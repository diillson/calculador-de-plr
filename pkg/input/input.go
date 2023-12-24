package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LeFloat64(prompt string) float64 {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	valor, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	if err != nil {
		fmt.Println("Entrada inválida. Por favor, insira um número.")
		return LeFloat64(prompt)
	}
	return valor
}

func LeInt(prompt string) int {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	valor, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		fmt.Println("Entrada inválida. Por favor, insira um número inteiro.")
		return LeInt(prompt)
	}
	return valor
}
