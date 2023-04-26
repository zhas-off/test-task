package main

import (
	"flag"
	"fmt"

	"github.com/zhas-off/test-task/pkg/util"
)

func main() {
	// Закомментировал кусок кода который создает случайные
	// числа в массиве на количество который вы укажете для проверки кода

	// err := util.Generate("pkg/data/data.json", 1000)
	// if err != nil {
	// 	panic(err)
	// }

	// Определяем флаг с именем "goroutines" и с типом int
	goroutinesNum := flag.Int("goroutines", 2, "number of goroutines")

	// Парсим аргумент командной строки
	flag.Parse()

	// Создаем количество блоков на которую хотим поделить
	blockSize := 100
	sum := util.Sum("pkg/data/data.json", blockSize, *goroutinesNum)

	// Выводим результат
	fmt.Printf("Total sum: %d\n", sum)
}
