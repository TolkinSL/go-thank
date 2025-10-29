package user

import "errors"

type Name struct {
	First string
	Last  string
}

type User struct {
	Name      Name
	birthYear int
}

func (u *User) ChangeYear(num int) error {
	// if num < 1950 {
	// 	return errors.New("Smal Year")
	// }

		if isGoodYear(num) {
		return errors.New("Smal Year")
	}

	u.birthYear = num
	return nil
}

func isGoodYear(year int) bool {
	return year < 1950
}