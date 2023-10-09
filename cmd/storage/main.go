package main

import (
	"fmt"

	"github.com/Gear19/STORAGE/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Print(st)
}
