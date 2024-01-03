package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Основная логика программы. Допускается передать в качестве аргументов один или
// три файла.
// Если передано один или два файла, то программа при необходимости объединяет
// их содержимое и выводит на экран.
// Если передано три файла, то третий файл используется для записи содержимого
// первых двух.
// В случае отсутствия или большого числа аргументов программа уведомит об этом
// и не пойдет дальше.
func main() {
	if len(os.Args) <= 1 || len(os.Args) > 4 {
		fmt.Println("Ошибка! Аргументы не переданы или их слишком много!")
		fmt.Println("Пример: go run main.go firstFile.txt")
		fmt.Println("Еще пример: go run main.go firstFile.txt secondFile.txt")
		return
	} else if len(os.Args) == 4 {
		resultStr := fileOps()
		writingToFile(resultStr)
		fmt.Printf("Произведена запись содержимого %s и %s в итоговый файл %s \n", os.Args[1], os.Args[2], os.Args[3])
	} else {
		resultStr := fileOps()
		fmt.Println(resultStr)
	}
}

// Функция, которая отвечает за чтение и объединение содержимого из переданных файлов в аргументах.
// Стоит ограничение на количество аргументов if i > 1, так как по логике программы с содержимым
// может быть только один или два файла.
func fileOps() string {
	var tmpSlice []string

	for i, filename := range os.Args[1:] {
		if i > 1 {
			break
		}
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Ошибка при открытии файла:", err)
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			tmpSlice = append(tmpSlice, line+"\n")
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка при чтении из файла:", err)
		}

		file.Close()
	}

	finalString := strings.Join(tmpSlice, "")

	return finalString
}

// Функция предназначена для передачи готовой объединенной строки в файл.
// При необходимости создается новый файл, а при наличии файла и содержимого - оно обнуляется.
func writingToFile(s string) {
	file, err := os.OpenFile(os.Args[3], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(s)
	if err != nil {
		panic(err)
	}
}
