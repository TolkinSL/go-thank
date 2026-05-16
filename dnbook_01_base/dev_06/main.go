package main

import (
	"fmt"
	"net/http"
	"sync"
)

type App struct {
	mu *sync.Mutex
	count int
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mu.Lock()
	a.count++
	a.mu.Unlock()
	fmt.Fprintf(w, "Count: %d\n", a.count)
}

func main() {
	myApp := App{
		mu: &sync.Mutex{},
	}

	http.Handle("/", &myApp)
	http.ListenAndServe(":8080", nil)
}

