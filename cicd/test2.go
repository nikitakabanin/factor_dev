package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func runTest(testName, inputFile, outputFile string) bool {
	fmt.Printf("Запуск теста:", testName)
	cmd := exec.Command("../usr/bin/factorial")

	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Ошибка открытия input:", err)
		return false
	}
	defer inFile.Close()
	cmd.Stdin = inFile

	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Ошибка создания output:", err)
		return false
	}
	defer outFile.Close()
	cmd.Stdout = outFile

	err = cmd.Run()
	if err != nil {
		fmt.Println("Ошибка выполнения программы:", err)
		return false
	}
	fmt.Println("Программа выполнилась успешно")
	return true
}

func checkOutput(outputFile, expectedSubstring string) bool {
	data, err := ioutil.ReadFile(outputFile)
	if err != nil {
		fmt.Println("Не удалось открыть файл вывода:", err)
		return false
	}
	content := string(data)
	if len(content) == 0 {
		fmt.Println("Файл вывода пуст")
		return false
	}
	if strings.Contains(content, expectedSubstring) {
		fmt.Printf("Содержит ожидаемое:\n", expectedSubstring)
		return true
	}
	fmt.Printf("Не содержит полный вывод:\n", expectedSubstring, content)
	return false
}

func main() {
	fmt.Println("тест 2")

	inputData := "0\n"
	err := ioutil.WriteFile("test2_input.txt", []byte(inputData), 0644)
	if err != nil {
		fmt.Println("Ошибка создания:", err)
		return
	}

	passed := false
	if runTest("Факториал 0", "test2_input.txt", "test2_output.txt") {
		if checkOutput("test2_output.txt", "0! = 1") {
			passed = true
		}
	}

	fmt.Println("тест 2")
	if passed {
		fmt.Println("тест пройден")
	} else {
		fmt.Println("тест не пройден")
	}
}