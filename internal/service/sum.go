package util

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/zhas-off/test-task/pkg/models"
)

func Sum(fileName string, blockSize int, goroutinesNum int) (int, error) {
	// Читаем данные из файла в срез []Object
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var objects []models.Object

	for {
		if err := decoder.Decode(&objects); err == io.EOF {
			break
		} else if err != nil {
			return 0, err
		}
	}

	// Вычисляем параллельно суммы чисел
	sum := 0
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	waitChan := make(chan struct{}, goroutinesNum)

	for i := 0; i < len(objects); i += blockSize {
		end := i + blockSize
		if end > len(objects) {
			end = len(objects)
		}
		block := objects[i:end]

		wg.Add(1)
		waitChan <- struct{}{} // Занимаем место в канале, чтобы использовать заданное количество горутин
		go func() {
			defer func() {
				wg.Done()
				<-waitChan
			}() // Освобождаем место в канале, когда горутина завершается

			blockSum := 0
			for _, obj := range block {
				blockSum += obj.A + obj.B
			}

			mutex.Lock()
			sum += blockSum
			mutex.Unlock()
		}()
	}

	wg.Wait()

	return sum, nil
}
