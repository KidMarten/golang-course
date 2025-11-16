package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// **Описание**: Создайте программу для анализа продаж магазина, которая создает слайс из массива транзакций и выводит информацию о выбранном периоде
	//
	// **Входные данные**: Программа получает три аргумента командной строки:
	// - start (int) - начальный индекс периода (включительно)
	// - end (int) - конечный индекс периода (исключительно)
	// - Массив транзакций встроен в код: [7]int{150, -50, 200, -30, 100, -25, 75}
	//
	// **Выходные данные**: Программа выводит:
	// - Количество транзакций в выбранном периоде
	// - Все транзакции выбранного периода через запятую
	// - Общую сумму транзакций периода
	//
	// **Ограничения**:
	// - 0 <= start < end <= 7
	// - Индексы должны быть валидными для массива
	//
	// **Примеры**:
	// Input: ./program 1 4
	// Output: Период: 3 транзакции
	// Транзакции: -50, 200, -30
	// Общая сумма: 120
	//
	// Input: ./program 0 2
	// Output: Период: 2 транзакции
	// Транзакции: 150, -50
	// Общая сумма: 100

	var transactions = [7]int{150, -50, 200, -30, 100, -25, 75}

	// Ваш код здесь
	if len(os.Args) != 3 {
		fmt.Println("enter two integeres: start and end period")
		return
	}
	start, _ := strconv.Atoi(os.Args[1])
	end, _ := strconv.Atoi(os.Args[2])
	if start < 0 {
		fmt.Println("Start must be more than 0")
		return
	}
	if end > len(transactions) {
		fmt.Println("End must be less than transactions count")
		return
	}
	if start > end {
		fmt.Println("End must be more than start")
		return
	}
	getResult(start, end, transactions)
}

func getResult(start int, end int, transactions [7]int) {
	result := transactions[start:end]
	fmt.Printf("Период: %d транзакции\n", len(result))

	strNumbers := make([]string, len(result))
	for i, num := range result {
		strNumbers[i] = strconv.Itoa(num)
	}
	resultStr := strings.Join(strNumbers, ", ")
	fmt.Printf("Транзакции: %s\n", resultStr)

	var sum int
	for _, num := range result {
		sum += num
	}
	fmt.Printf("Общая сумма: %d\n", sum)
}
