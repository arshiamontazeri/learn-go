package main

import (
	"fmt"
)

type Student struct {
	Id     int
	Name   string
	Age    int
	Grades []Grade
}
type Grade struct {
	lessonName string
	Score      int
}

var Students []Student = []Student{
	{
		Id:   1,
		Name: "Ali",
		Age:  18,
	},
	{
		Id:   2,
		Name: "Ahmad",
		Age:  17,
	},
	{
		Id:   3,
		Name: "mohammad",
		Age:  18,
	},
	{
		Id:   4,
		Name: "sara",
		Age:  18,
	},
	{
		Id:   5,
		Name: "negin",
		Age:  15,
	},
	{
		Id:   6,
		Name: "baran",
		Age:  16,
	},
	{
		Id:   7,
		Name: "mehdi",
		Age:  19,
	},
	{
		Id:   8,
		Name: "mahan",
		Age:  20,
	},
	{
		Id:   9,
		Name: "saeed",
		Age:  22,
	},
	{
		Id:   10,
		Name: "mahyar",
		Age:  14,
	},
}

func main() {
	fmt.Println("------------------------------------------------------------")
	fmt.Println()
	first := "\t 1-See a list student :\n"
	second := "\t 2- Add student :\n"
	third := "\t 3- Add grade to student :\n"
	fourth := "\t 4- See student ararega score :\n"
	for {
		fmt.Print(first, second, third, fourth)
		var choose int
		fmt.Scan(&choose)
		fmt.Println("------------------------------------------------------------")
		if choose == 1 {
			fmt.Println("See a list student :")
			SeeAListStudent()
		}
		if choose == 2 {
			fmt.Println("Add student :")
			AddStudent()
		}
		if choose == 3 {
			fmt.Println("Add grade to student :")
			AddGradeToStudent()
		}
		if choose == 4 {
			fmt.Println("See student ararega score :")
			SeeStudentAraregaScore()
		}
		fmt.Println("------------------------------------------------------------")
		fmt.Println()
	}
}
func SeeAListStudent() {
	for _, s := range Students {
		fmt.Print("Id :", s.Id, "\t", "Name :", s.Name, "\t", "Age :", s.Age, "\n")
	}
}

func AddStudent() {
	id := MakeId()
	var name string
	var age int
	fmt.Print("Name :")
	fmt.Scan(&name)
	fmt.Print("Age :")
	fmt.Scan(&age)
	ok := true
	for _, s := range Students {
		if s.Id == id {
			ok = false
		}
	}
	if ok {
		Students = append(Students, Student{Id: id, Name: name, Age: age})
	} else {
		fmt.Println("------------------------------------------------------------")
		fmt.Println("**********  error ----->This ID already existed  **********")
	}
}

func AddGradeToStudent() {
	var id int
	var lessonName string
	var score int
	fmt.Print("ID :")
	fmt.Scan(&id)
	fmt.Print("LessonName :")
	fmt.Scan(&lessonName)
	fmt.Print("Score :")
	fmt.Scan(&score)

	for i, g := range Students {
		if g.Id == id {
			Students[i].Grades = append(Students[i].Grades, Grade{lessonName: lessonName, Score: score})
		}
	}
}

func SeeStudentAraregaScore() {
	var id int
	fmt.Print("ID :")
	fmt.Scan(&id)
	for _, s := range Students {
		if s.Id == id {
			if len(s.Grades) == 0 {
				fmt.Println("No grades available for this student.")
				return
			}
			t := 0.0
			for _, grade := range s.Grades {
				t += float64(grade.Score)
			}
			averageScore := t / float64(len(s.Grades))
			fmt.Println(averageScore)
		}
	}
}

func MakeId() int {
	id := len(Students)
	id = id + 1
	return id
}
