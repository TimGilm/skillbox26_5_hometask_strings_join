/*
Что нужно сделать
Программа должна получать на вход имена двух файлов, необходимо  конкатенировать их
содержимое, используя strings.Join.
При получении одного файла на входе программа должна печатать его содержимое на экран.
При получении двух файлов на входе программа соединяет их и печатает содержимое обоих
файлов на экран. Если программа запущена командой go run firstFile.txt secondFile.txt
resultFile.txt, то она должна написать два соединённых файла в результирующий.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()                        //парсим аргументы командной строки
	commandLineArguments := flag.Args() //инициализируем слайс из этих аргументов
	if len(commandLineArguments) == 1 { //выполняем первое условие задачи
		fileName := commandLineArguments[0]
		fmt.Println(readFromFile(fileName)) //печатаем в StdOut содержимое из файла
	} else if len(commandLineArguments) == 2 { //переходим ко второму условию задачи
		var resultString []string
		var result = ""
		for _, fileName := range commandLineArguments {
			resultString = append(resultString, readFromFile(fileName)) //передаем в слайс считанное содержимое файлов
			result = strings.Join(resultString, "\n")                   //конкатенируем строки в слайсе
		}
		fmt.Println(result) //печатаем в StdOut результат конкатенации
		//переходим к последнему условию задачи
	} else if len(commandLineArguments) == 3 && commandLineArguments[0] == "firstFile.txt" && commandLineArguments[1] == "secondFile.txt" && commandLineArguments[2] == "resultFile.txt" {
		var resultString []string
		var result = ""
		for i := 0; i < 2; i++ {
			filename := commandLineArguments[i]
			resultString = append(resultString, readFromFile(filename)) //передаем в слайс считанное содержимое файлов
			result = strings.Join(resultString, "\n")                   //конкатенируем строки в слайсе
		}

		file, err := os.OpenFile("resultFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600) //открываем файл для записи
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(3)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("error closing file: err:", err)
				os.Exit(4)
			}
		}(file)

		_, err = file.WriteString(result) //записываем в файл из переменной
		if err != nil {
			os.Exit(5)
		}

	} else {
		fmt.Println("Вы ввели некорректные аргументы командной строки")
		os.Exit(6) //обработка ошибки ввода некорректных данных командной строки
	}
}

func readFromFile(fileName string) string {
	f, err := os.Open(fileName) //открываем файл по аргументу ком. строки
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1) //обрабатываем ошибку открытия файла, если неверно указано название файла
	}
	defer func(f *os.File) { //закрываем файл после выполнения функции
		err := f.Close()
		if err != nil {
			os.Exit(2) //обрабатываем ошибку закрытия файла
		} //обрабатываем ошибку закрытия файла
	}(f)

	buf := bufio.NewScanner(f) //читаем содержимое из файла
	buf.Scan()                 //сохраняем в буфер
	return buf.Text()
}
