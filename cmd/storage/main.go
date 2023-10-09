package main

import (
	"fmt"
	"log"

	"github.com/Gear19/STORAGE/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(file)
}
