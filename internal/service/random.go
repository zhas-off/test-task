package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/zhas-off/test-task/pkg/models"
)

func Generate(fileName string, number int) error {

	// Генерируем случайные числа
	rand.Seed(time.Now().UnixNano())

	// Создаем срез объектов Object
	data := make([]models.Object, number)

	// Заполняем срез случайными числами
	for i := 0; i < len(data); i++ {
		data[i] = models.Object{
			A: rand.Intn(100),
			B: rand.Intn(100) - 50,
		}
	}

	// Создаем JSON файл
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Кодируем срез в JSON и записываем в него
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	fmt.Println("Готово!")

	return nil
}
