package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Приложение, которое:

// Принимает операцию (AVG - среднее, SUM - сумму, MED - медиану)
// Принимает неограниченное число чисел через запятую (2, 10, 9)
// Разбивает строку чисел по запятым и затем делает расчёт в зависимости от операции выводя результат

func main() {
	operation, numbers := getUserInput()
	result := calculateStat(operation, numbers)
	fmt.Println("Answer", result)
}

func getUserInput() (string, []int) {
	operation, err := scanOpeation()
	if err != nil {
		fmt.Println("Error:", err)
	}
	numbers, err := scanNumbers()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return operation, numbers
}

func scanOpeation() (string, error) {
	var operation string
	fmt.Println("Enter operation (available ops: AVG, SUM or MED)")
	fmt.Scan(&operation)
	switch operation {
	case "AVG", "MED", "SUM":
		return operation, nil
	default:
		return "", fmt.Errorf("Not in available values")
	}
}

func scanNumbers() ([]int, error) {
	var numbers string
	fmt.Println("Enter any integer numbers separated by a comma")
	fmt.Scan(&numbers)

	numbersStringArray := strings.Split(numbers, ",")
	result := make([]int, 0, len(numbersStringArray))
	for _, part := range numbersStringArray {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("Wrong number: %s", part)
		}
		result = append(result, num)
	}
	return result, nil
}

func calculateStat(stat string, numbers []int) float64 {
	switch stat {
	case "AVG":
		return calcAverage(numbers)
	case "SUM":
		return calcSum(numbers)
	case "MED":
		return calcMedian(numbers)
	default:
		return 0.0
	}
}

func calcSum(numbers []int) float64 {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return float64(total)
}

func calcAverage(numbers []int) float64 {
	return float64(calcSum(numbers)) / float64(len(numbers))
}

func calcMedian(numbers []int) float64 {
	sorted := make([]int, len(numbers))
	copy(sorted, numbers)
	sort.Ints(sorted)

	n := len(sorted)
	if n%2 == 1 {
		return float64(sorted[n/2])
	}
	return float64(sorted[n/2-1]+sorted[n/2]) / 2
}
