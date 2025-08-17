//Запуск Stepik заданий
// $ echo "1 1 4 5" | go run ./dev_01_05
// 5

package main

import (
	"fmt"
	"math"
)

func main() {
	// объявите переменные x1, y1, x2, y2 типа float64
	var x1, y1, x2, y2 float64 

	// считываем числа из os.Stdin
	// гарантируется, что значения корректные
	// не меняйте этот блок
	fmt.Scan(&x1, &y1, &x2, &y2)

	// рассчитайте d по формуле эвклидова расстояния
	// используйте math.Pow(x, 2) для возведения в квардрат
	// используйте math.Sqrt(x) для извлечения корня
    sum := math.Pow((x1 - x2), 2) + math.Pow((y1 - y2), 2)
    d := math.Sqrt(sum)

	// выводим результат в os.Stdout
	fmt.Println(d)
}