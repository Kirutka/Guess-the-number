package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var level, diap string
	var a, attempt, minAttempt, number int
	fmt.Println("Выберите уровень сложность: ") // выбор сложности
	fmt.Println("1. Лёгкий (1 - 50 диапазон, 10 попыток)")
	fmt.Println("2. Средний (1 - 100 диапазон, 7 попыток)")
	fmt.Println("3. Сложный (1 - 200 диапазон, 5 попыток)")

	for {
		_, err := fmt.Scanln(&a)
		if err != nil || a < 1 || a > 3 {
			fmt.Println("Ошибка: введите число от 1 до 3")
			continue
		}
		break
	}

	rand.Seed(time.Now().UnixNano()) // генерация рандомного числа

	switch a { // проверка выбора уровня
	case 1:
		number = rand.Intn(50) + 1
		attempt = 10
		level = "лёгкий"
		diap = "1 - 50"
	case 2:
		number = rand.Intn(100) + 1
		attempt = 7
		level = "средний"
		diap = "1 - 100"
	case 3:
		number = rand.Intn(200) + 1
		attempt = 5
		level = "сложный"
		diap = "1 - 200"
	default:
		fmt.Println("Вы ввели неправильную цифру. Установлен средний уровень")
		number = rand.Intn(100) + 1
		attempt = 7
		level = "средний"
	}

	var num int
	minAttempt = 0

	fmt.Printf("Вы выбрали %s уровень, у вас есть %d попыток, диапазон чисел - (%s)", level, attempt, diap)
	fmt.Println()

	for attempt > minAttempt { // цикл игры

		fmt.Println("Напишите число: ")
		for {
			_, err := fmt.Scan(&num)
			if err != nil {
				fmt.Println("Ошибка: введите число.")
				continue
			}
			switch a {
			case 1:
				if num < 1 || num > 50 {
					fmt.Println("Ошибка: число должно быть в диапазоне от 1 до 50.")
					continue
				}
			case 2:
				if num < 1 || num > 100 {
					fmt.Println("Ошибка: число должно быть в диапазоне от 1 до 100.")
					continue
				}
			case 3:
				if num < 1 || num > 200 {
					fmt.Println("Ошибка: число должно быть в диапазоне от 1 до 200.")
					continue
				}
			}
			break
		}
		attempt--
		if num < number {
			fmt.Println("*Загаданное число больше*")
			fmt.Printf("Осталось попыток: %d", attempt)
			fmt.Println()
		} else if num > number {
			fmt.Println("*Загаданное число меньше*")
			fmt.Printf("Осталось попыток: %d", attempt)
			fmt.Println()
		} else {
			fmt.Println("Поздравляю, вы выиграли!")
			return
		}
	}
	fmt.Println()
	fmt.Println("К сожаление ваши попытки закончились(")
	fmt.Printf("Загаданное число - %d", number)
	bufio.NewReader(os.Stdin).ReadBytes(' ')
}
