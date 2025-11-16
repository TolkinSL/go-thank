package main

import (
	"fmt"
	"os"
	"path/filepath"
	// "path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Provide programm")
		return
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)

		// Проверка существования пути
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					fmt.Printf("%#v\n", fileInfo)
					// return
				}
			}
		}

		// if err != nil {
		// Будет возвращать -Err:0x2 — код ошибки ОС. то есть "No such file or directory" — файл или папка не существует.
		// fmt.Printf("%#v\n", err)
		// }
	}
}
