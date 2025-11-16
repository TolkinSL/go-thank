package main

import (
	"errors"
	"fmt"
)

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

// Бизнес Логика
type SimpleLogic struct {
	l Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log(fmt.Sprintf("\nin method SayHello for ID - %s\n", userID))
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return fmt.Sprintf("Hello - %s\n", name), nil
}

func main() {
	my_dataStruct := NewSimpleDataStore()
	fmt.Printf("%+v\n", my_dataStruct)
	
	name, ok := my_dataStruct.UserNameForID("7")
	fmt.Printf("%#v %#v\n", name, ok)

	my_logger := LoggerAdapter(LogOutput)
	my_logger.Log("Hello")

	fmt.Println("-------------")

	my_simpleStruct := SimpleLogic{
		l: my_logger,
		ds: my_dataStruct,
	}

	my_simpleStruct.l.Log("Test")

	name, ok = my_simpleStruct.ds.UserNameForID("1")
	fmt.Printf("%#v %#v\n", name, ok)

	str, err := my_simpleStruct.SayHello("2")
	fmt.Printf("%#v %#v\n", str, err)
	fmt.Println(str)

	str, err = my_simpleStruct.SayHello("7")
	fmt.Printf("%#v %#v\n", str, err)
	fmt.Println(str)
}
