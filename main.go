package main

import (
	"StudentScore/student"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a1 student.StudentA
	a1.AddStudentA()
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("Nhap ID sinh vien muon tim\n")
	id, _ := consoleReader.ReadString('\n')
	id = strings.Replace(id, "\n", "", -1)
	IDInput, _ := strconv.ParseInt(id, 10, 64)

	student.FindStudent(IDInput)
}
