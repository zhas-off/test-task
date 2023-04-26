package main

import (
	"flag"
	"fmt"

	service "github.com/zhas-off/test-task/internal/service"
)

func main() {
	// Закомментировал кусок кода который создает случайные
	// числа в массиве на количество который вы укажете для проверки кода

	// service.Generate("pkg/data/data.json", 1000000)

	// Определяем флаг с именем "goroutines" и с типом int
	goroutinesNum := flag.Int("goroutines", 2, "number of goroutines")

	// Парсим аргумент командной строки
	flag.Parse()

	// Создаем количество блоков на которую хотим поделить
	blockSize := 100
	sum, _ := service.Sum("pkg/data/data.json", blockSize, *goroutinesNum)

	// Выводим результат
	fmt.Printf("Итоговая сумма: %d\n", sum)
}
