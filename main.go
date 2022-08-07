package main

import (
	"StudentScore/database"
	"fmt"
)

func main() {
	db := database.NewDB()
	studentRepo := database.NewStudentRepo(db)
	err := studentRepo.DeleteStudent("272763")
	if err != nil {
		fmt.Println("xoa k thanh cong")
		return
	}
	fmt.Println("xoa student thanh cong")
}
