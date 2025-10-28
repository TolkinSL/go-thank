package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// User представляет пользователя
type User struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	Phone   string     `json:"phone"`
	Address Address    `json:"address"`
	Cart    []CartItem `json:"cart"`
}

// Address представляет адрес пользователя
type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
}

// CartItem представляет элемент в корзине
type CartItem struct {
	Product  Product `json:"product"`
	Quantity int     `json:"quantity"`
}

// Product представляет продукт в корзине
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       int     `json:"price"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
	Rating      float64 `json:"rating"`
	Reviews     int     `json:"reviews"`
}

func main() {
	// Читаем JSON из stdin (можно заменить на файл)
	data, err := os.ReadFile("./user.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return
	}

	// fmt.Printf("Имя: %s\nEmail: %s\nГород: %s\n", user.Name, user.Email, user.Address.City)
	// fmt.Println("Корзина:")
	// for _, item := range user.Cart {
	// 	fmt.Printf("- %s x%d (%.2f)\n", item.Product.Name, item.Quantity, item.Product.Rating)
	// }

	fmt.Printf("Покупатель %s. Телефон: %s. Адрес: г. %s, %s.\n", user.Name, user.Phone, user.Address.City, user.Address.Street)
	var strElectro string = "не является"
	var arrGoods []string = []string{}
	var sum int = 0

	for _, item := range user.Cart {
		// fmt.Println(item.Product.Category)
		if item.Product.Category == "Электроника" {
			strElectro = "является"
		}

		if item.Product.Price >= 10000 {
			arrGoods = append(arrGoods, item.Product.Name)
		}

		sum += item.Product.Price * item.Quantity
	}

	fmt.Printf("Пользователь %s покупателем электроники.\n", strElectro)

	var strGoods string = ""
	if len(arrGoods) == 0 {
		strGoods = fmt.Sprintf("Товары в корзине, где цена 10000 и более: %s.\n", "отсутствуют")
	} else {
		strGoods = fmt.Sprintf("Товары в корзине, где цена 10000 и более: %s.\n", strings.Join(arrGoods, ", "))
	}

	fmt.Print(strGoods)

	fmt.Printf("Общая сумма покупки: %d руб.\n", sum)
}
