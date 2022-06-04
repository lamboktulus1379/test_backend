package main

import (
	"fmt"
	db "test_backend_2/DB"
	repository "test_backend_2/Repository"
	service "test_backend_2/Service"
)

func main() {
	DB, _ := db.NewRepositories()
	areaRepository := repository.NewAreaRepository(DB)
	s := service.NewService(areaRepository)
	fmt.Println(s.Insert())
}
