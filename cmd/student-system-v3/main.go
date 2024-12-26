package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Age  int    `json:"age"`
}

type Grade struct {
	StudentId  int    `json:"studentId"`
	LessonName string `json:"lessonName"`
	Score      int    `json:"score"`
}

// 1-----------------------------------------------------------------------------------------------------------------------------------
func main() {
	fmt.Print("-------------------------------------------------------------")
	fmt.Println()
	first := "\t1- Add student :\n"
	second := "\t2- Add grade :\n"
	third := "\t3- Add student :\n"
	fourth := "\t4- Add grade to student :\n"
	fifth := "\t5- See student average score :\n"
	sixth := "\t6- Search student :\n"
	for {
		fmt.Println(first, second, third, fourth, fifth, sixth)
		var choose int
		fmt.Scan(&choose)
		fmt.Print("---------------------------------------------------------\n")
		switch choose {
		case 1:
			fmt.Println("Add students :")
			StudentFile()
		case 2:
			fmt.Println("Add grades :")
			GradeFile()
		case 3:
			fmt.Println("See a list student :")
			SeeAListStudent()
		case 4:
			fmt.Println("Add grade to student :")
			AddGeradeToStudent()
		case 5:
			fmt.Println("See student average score :")
			SeeStudentAverageScore()
		case 6:
			fmt.Println("Search student :")
			Search()
		default:
			fmt.Println("\t<*******> Nothing was found about it <*******>")
			return

		}
		fmt.Print("---------------------------------------------------------\n")
	}
}

// 2-----------------------------------------------------------------------------------------------------------------------------------
func StudentFile() {
	StudentBytes, err := os.ReadFile("./cmd/student-system-v3/students.json")
	if err != nil {
		panic(err)
	}
	students := []Student{}

	err = json.Unmarshal(StudentBytes, &students)
	if err != nil {
		panic(err)
	}
	s := Student{}

	fmt.Println("students :")
	_, err = fmt.Scan(&s.Name, &s.Age)
	if err != nil {
		panic(err)
	}
	s.Id = len(students) + 1

	students = append(students, s)

	StudentBytes, err = json.Marshal(students)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./cmd/student-system-v3/students.json", StudentBytes, 0644)
	if err != nil {
		panic(err)
	}
}

// 3----------------------------------------------------------------------------------------------------------------------------------
func GradeFile() {
	GradeBytes, err := os.ReadFile("./cmd/student-system-v3/grades.json")
	if err != nil {
		panic(err)
	}

	grades := []Grade{}

	err = json.Unmarshal(GradeBytes, &grades)
	if err != nil {
		panic(err)
	}

	g := Grade{}

	fmt.Println("grades :")
	_, err = fmt.Scan(&g.LessonName, &g.Score)
	if err != nil {
		panic(err)
	}

	g.StudentId = len(grades) + 1

	grades = append(grades, g)

	GradeBytes, err = json.Marshal(grades)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./cmd/student-system-v3/grades.json", GradeBytes, 0644)
	if err != nil {
		panic(err)
	}
}

// 4---------------------------------------------------------------------------------------------------------------------------------------------------------
func SeeAListStudent() {
	liststudent, err := os.ReadFile("./cmd/student-system-v3/students.json")
	if err != nil {
		panic(err)
	}
	list := []Student{}

	err = json.Unmarshal(liststudent, &list)
	if err != nil {
		panic(err)
	}
	fmt.Print(list, "\n")
}

// 5------------------------------------------------------------------------------------------------------------------------------------------------------
func AddGeradeToStudent() {
	GradeBytes, err := os.ReadFile("./cmd/student-system-v3/grades.json")
	if err != nil {
		panic(err)
	}

	grades := []Grade{}

	err = json.Unmarshal(GradeBytes, &grades)
	if err != nil {
		panic(err)
	}

	g := Grade{}

	fmt.Println("grades :")
	_, err = fmt.Scan(&g.LessonName, &g.Score, &g.StudentId)
	if err != nil {
		panic(err)
	}

	grades = append(grades, g)

	GradeBytes, err = json.Marshal(grades)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./cmd/student-system-v3/grades.json", GradeBytes, 0644)
	if err != nil {
		panic(err)
	}
}

// 6---------------------------------------------------------------------------------------------------------------------------------
func SeeStudentAverageScore() {
	var id int
	fmt.Print("ID: ")
	fmt.Scan(&id)
	// grade
	GradeBytes, err := os.ReadFile("./cmd/student-system-v3/grades.json")
	if err != nil {
		panic(err)
	}

	grades := []Grade{}

	err = json.Unmarshal(GradeBytes, &grades)
	if err != nil {
		panic(err)
	}
	var studentGrades []Grade
	for _, grade := range grades {
		if grade.StudentId == id {
			studentGrades = append(studentGrades, grade)
		}
	}
	if len(studentGrades) == 0 {
		fmt.Println("------------------------------------------------------------")
		fmt.Println("No grades available for this student.")
	}
	// average
	totalScore := 0.0
	for _, grade := range studentGrades {
		totalScore += float64(grade.Score)
	}
	averageScore := totalScore / float64(len(studentGrades))
	// student
	liststudent, err := os.ReadFile("./cmd/student-system-v3/students.json")
	if err != nil {
		panic(err)
	}
	students := []Student{}

	err = json.Unmarshal(liststudent, &students)
	if err != nil {
		panic(err)
	}
	// print
	for _, student := range students {
		if student.Id == id {
			fmt.Println("------------------------------------------------------------")
			fmt.Println("\t Student name is :", student.Name)
			fmt.Println("\t Student ID is :", student.Id)
			fmt.Println("\t list grade :", grades)
			fmt.Println("\t Student average score is :", averageScore)
		}
	}
}

// 7----------------------------------------------------------------------------------------------------------------------------------
func Search() {
	var name string
	fmt.Print("name :")
	fmt.Scan(&name)
	// student
	liststudent, err := os.ReadFile("./cmd/student-system-v3/students.json")
	if err != nil {
		panic(err)
	}
	students := []Student{}

	err = json.Unmarshal(liststudent, &students)
	if err != nil {
		panic(err)
	}
	// grade
	GradeBytes, err := os.ReadFile("./cmd/student-system-v3/grades.json")
	if err != nil {
		panic(err)
	}

	grades := []Grade{}

	err = json.Unmarshal(GradeBytes, &grades)
	if err != nil {
		panic(err)
	}
	// search
	ok := false
	for _, student := range students {
		if strings.Contains(student.Name, name) {
			ok = true
			fmt.Println("\t students name :", student.Name)
			fmt.Println("\t students id :", student.Id)

		}
	}
	var scanId int
	fmt.Print("\t Choose student id :")
	fmt.Scan(&scanId)
	s := 0
	if ok {
		s = scanId
	}
	for _, student := range students {
		if student.Id == s {
			fmt.Println("\t students name :", student.Name)
			fmt.Println("\t students id :", student.Id)
			fmt.Println("\t students age :", student.Age)
			for _, grade := range grades {
				if student.Id == grade.StudentId {
					fmt.Println("\t Lesson name :", grade.LessonName)
					fmt.Println("\t student id :", grade.StudentId)
					fmt.Println("\t score :", grade.Score)
				}
			}
		}
	}

	if !ok {
		fmt.Println("------------------------------------------------------------")
		fmt.Println("\t<*****> No such name found <*****>")
		return
	}
}
