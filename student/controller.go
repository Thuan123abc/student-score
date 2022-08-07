package student

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ListStudentA []StudentA
var ListStudentB []StudentB
var ListStudentC []StudentC

func (s *Student) Input() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("Nhap so bao danh cua sinh vien\n")
	ID, _ := consoleReader.ReadString('\n')
	ID = strings.Replace(ID, "\n", "", -1)
	IDInt, _ := strconv.ParseInt(ID, 10, 64)
	s.ID = IDInt

	fmt.Println("Nhap ten cua sinh vien\n")
	nameInp, _ := consoleReader.ReadString('\n')
	nameInp = strings.Replace(nameInp, "\n", "", -1)
	s.Name = nameInp

	fmt.Println("Nhap dia chi cua sinh vien\n")
	addr, _ := consoleReader.ReadString('\n')
	addr = strings.Replace(addr, "\n", "", -1)
	s.Address = addr

	fmt.Println("Nhap so diem uu tien cua sinh vien\n")
	ScorePri, _ := consoleReader.ReadString('\n')
	ScorePri = strings.Replace(ScorePri, "\n", "", -1)
	ScorePriFloat, _ := strconv.ParseFloat(ScorePri, 64)
	s.ScorePrio = ScorePriFloat
}

func (a *StudentA) AddStudentA() {
	a.Input()
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("Nhap so diem mon toan\n")
	mathSco, _ := consoleReader.ReadString('\n')
	mathSco = strings.Replace(mathSco, "\n", "", -1)
	mathScoInt, _ := strconv.ParseFloat(mathSco, 64)
	a.MathScore = mathScoInt

	fmt.Println("Nhap so diem mon ly\n")
	physSco, _ := consoleReader.ReadString('\n')
	physSco = strings.Replace(physSco, "\n", "", -1)
	physScoInt, _ := strconv.ParseFloat(mathSco, 64)
	a.PhysicalScore = physScoInt

	fmt.Println("Nhap so diem mon hoa\n")
	chemSco, _ := consoleReader.ReadString('\n')
	chemSco = strings.Replace(chemSco, "\n", "", -1)
	chemScoInt, _ := strconv.ParseFloat(chemSco, 64)
	a.ChemistryScore = chemScoInt

	ListStudentA = append(ListStudentA, *a)
}

func (b *StudentB) AddStudentB() {
	b.Input()
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("Nhap so diem mon toan\n")
	mathSco, _ := consoleReader.ReadString('\n')
	mathSco = strings.Replace(mathSco, "\n", "", -1)
	mathScoInt, _ := strconv.ParseFloat(mathSco, 64)
	b.MathScore = mathScoInt

	fmt.Println("Nhap so diem mon sinh\n")
	bioSco, _ := consoleReader.ReadString('\n')
	bioSco = strings.Replace(bioSco, "\n", "", -1)
	bioScoInt, _ := strconv.ParseFloat(bioSco, 64)
	b.ChemistryScore = bioScoInt

	fmt.Println("Nhap so diem mon hoa\n")
	chemSco, _ := consoleReader.ReadString('\n')
	chemSco = strings.Replace(chemSco, "\n", "", -1)
	chemScoInt, _ := strconv.ParseFloat(chemSco, 64)
	b.ChemistryScore = chemScoInt

	ListStudentB = append(ListStudentB, *b)
}

func (c *StudentC) AddStudentC() {
	c.Input()
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Println("Nhap so diem mon van\n")
	liteSco, _ := consoleReader.ReadString('\n')
	liteSco = strings.Replace(liteSco, "\n", "", -1)
	liteScoInt, _ := strconv.ParseFloat(liteSco, 64)
	c.LiteScore = liteScoInt

	fmt.Println("Nhap so diem mon su\n")
	hisSco, _ := consoleReader.ReadString('\n')
	hisSco = strings.Replace(hisSco, "\n", "", -1)
	hisScoInt, _ := strconv.ParseFloat(hisSco, 64)
	c.HistoryScore = hisScoInt

	fmt.Println("Nhap so diem mon dia\n")
	geoSco, _ := consoleReader.ReadString('\n')
	geoSco = strings.Replace(geoSco, "\n", "", -1)
	geoScoInt, _ := strconv.ParseFloat(geoSco, 64)
	c.GeoScore = geoScoInt

	ListStudentC = append(ListStudentC, *c)
}
func FindStudent(IDInput int64) {
	for i := 0; i < len(ListStudentA)+len(ListStudentB)+len(ListStudentC); i++ {
		if a := ListStudentA[i]; IDInput == a.ID {
			fmt.Println("Sinh vien can tim la:\n", a)
			break
		} else if b := ListStudentB[i]; IDInput == b.ID {
			fmt.Println("Sinh vien can tim la:\n", b)
			break
		} else if c := ListStudentC[i]; IDInput == c.ID {
			fmt.Println("Sinh vien can tim la:\n", c)
			break
		} else {
			fmt.Println("Khong co sinh vien co ID nhu vay")
		}
	}
}
