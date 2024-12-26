package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Id   int
	Name string
	Age  int
}

// ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
type Grade struct {
	Id         int
	StudentId  int
	LessonName string
	Score      int
}

// ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
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
		Name: "mahyer",
		Age:  14,
	},
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
var greds []Grade = []Grade{
	{
		StudentId:  1,
		LessonName: "math",
		Score:      18,
	},
	{
		StudentId:  2,
		LessonName: "Chemistry",
		Score:      18,
	},
	{
		StudentId:  3,
		LessonName: "physics",
		Score:      19,
	},
	{
		StudentId:  4,
		LessonName: "Geometry",
		Score:      17,
	},
	{
		StudentId:  5,
		LessonName: "english",
		Score:      20,
	},
	{
		StudentId:  6,
		LessonName: "discrete",
		Score:      19,
	},
	{
		StudentId:  7,
		LessonName: "Statistics",
		Score:      15,
	},
	{
		StudentId:  8,
		LessonName: "accountants",
		Score:      16,
	},
	{
		StudentId:  9,
		LessonName: "Geography",
		Score:      13,
	},
	{
		StudentId:  10,
		LessonName: "history",
		Score:      11,
	},
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func main() {
	fmt.Println("------------------------------------------------------------")
	fmt.Println()
	first := "\t 1- See a list student :\n"
	second := "\t 2- Add student :\n"
	third := "\t 3- Add grade :\n"
	fourth := "\t 4- Add grade to student :\n"
	fifth := "\t 5- See student average score :\n"
	sixth := "\t 6- search student :\n"
	Seventh := "\t 7- Delete Student :\n"
	for {
		fmt.Print(first, second, third, fourth, fifth, sixth, Seventh)
		var choose int
		fmt.Scan(&choose)
		fmt.Println("------------------------------------------------------------")
		switch choose {
		case 1:
			fmt.Println("See a list student :")
			SeeAListStudent()
		case 2:
			fmt.Println("Add student :")
			AddStudent()
		case 3:
			fmt.Println("Add grade :")
			AddGrade()
		case 4:
			fmt.Println("Add grade to student :")
			AddGradeToStudent()
		case 5:
			fmt.Println("See student average score :")
			SeeStudentAverageScore()
		case 6:
			fmt.Println("search student :")
			Search()
		case 7:
			fmt.Println("Delete Student :")
		default:
			fmt.Println("\t<*******> Nothing was found about it <*******>")
		}
		fmt.Println("------------------------------------------------------------")
		fmt.Println()
	}
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func SeeAListStudent() string {
	idlen := 3
	namelen := 6
	agelen := 5
	line := "-"

	for i := 0; i < idlen+namelen+agelen+10; i++ {
		fmt.Print(line)
	}
	fmt.Println()
	fmt.Print("|")
	fmt.Print(" ", "Id", " ")
	fmt.Print("|")
	fmt.Print(" ", "Name", " ")
	fmt.Print("|")
	fmt.Print(" ", "Age", " ")
	fmt.Print("|")
	fmt.Println()
	for i := 0; i < idlen+namelen+agelen+10; i++ {
		fmt.Print(line)
	}

	fmt.Println()
	var w2 string
	for _, s := range Students {
		fmt.Print("|")
		if s.Id < 10 {
			fmt.Print(" ", s.Id, "  ")
		} else {
			fmt.Print(" ", s.Id, " ")
		}
		fmt.Print("|")
		w2 = s.Name
		fmt.Print(Makespacestudent(Maxstudent(), w2), " ", w2, " ")
		fmt.Print("|")
		fmt.Print(" ", s.Age, "  ")
		fmt.Print("|")
		fmt.Println()

		for i := 0; i < idlen+namelen+agelen+10; i++ {
			fmt.Print(line)
		}

		fmt.Println()

	}
	return w2
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func AddStudent() {
	id := MakeStudentId()
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

// ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func AddGrade() {
	id := MakeGradeId()
	var lessonName string
	var score int
	fmt.Print("LessonName :")
	fmt.Scan(&lessonName)
	fmt.Print("Score :")
	fmt.Scan(&score)

	f := true
	for _, g := range greds {
		if g.StudentId == id {
			f = false
		}
	}
	if f {
		greds = append(greds, Grade{Id: id, LessonName: lessonName, Score: score})
	} else {
		fmt.Println("------------------------------------------------------------")
		fmt.Println("**********  error ----->This ID already existed  **********")
	}

}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func AddGradeToStudent() {
	var id int
	var lessonName string
	var score int
	fmt.Print("Id :")
	fmt.Scan(&id)
	fmt.Print("lessonName :")
	fmt.Scan(&lessonName)
	fmt.Print("score :")
	fmt.Scan(&score)

	greds = append(greds, Grade{StudentId: id, LessonName: lessonName, Score: score})
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func SeeStudentAverageScore() string {
	var id int
	fmt.Print("ID: ")
	fmt.Scan(&id)
	// grades
	var studentGrades []Grade
	for _, g := range greds {
		if g.StudentId == id {
			studentGrades = append(studentGrades, g)
		}
	}

	if len(studentGrades) == 0 {
		fmt.Println("------------------------------------------------------------")
		fmt.Println("No grades available for this student.")
	}
	//  average
	totalScore := 0.0
	for _, grade := range studentGrades {
		totalScore += float64(grade.Score)
	}
	averageScore := totalScore / float64(len(studentGrades))
	// Print
	for _, s := range Students {
		if s.Id == id {
			fmt.Println("------------------------------------------------------------")
			fmt.Println("\t Student name is :", s.Name)
			fmt.Println("\t Student ID is :", s.Id)
		}
	}
	// Table
	LessonNamelen := 11
	StudentIdlen := 10
	Scorelen := 6
	Idlen := 3

	line := "-"
	for i := 0; i < StudentIdlen+LessonNamelen+Scorelen+Idlen+6; i++ {
		fmt.Print(line)
	}

	fmt.Println()
	fmt.Print("|")
	fmt.Print(" ", "Student Id", " ")
	fmt.Print("|")
	fmt.Print(" ", "lesson Name", " ")
	fmt.Print("|")
	fmt.Print(" ", "Score", " ")
	fmt.Print("|")
	fmt.Println()

	for i := 0; i < StudentIdlen+LessonNamelen+Scorelen+Idlen+6; i++ {
		fmt.Print(line)
	}
	fmt.Println()
	var w string
	for _, g := range greds {
		fmt.Print("|")
		if g.StudentId < 10 {
			fmt.Print(" ", g.StudentId, "  ")
		} else {
			fmt.Print(" ", g.StudentId, " ")
		}
		fmt.Print("|")
		if g.Id < 10 {
			fmt.Print(" ", g.Id, "  ")
		} else {
			fmt.Print(" ", g.Id, " ")
		}
		fmt.Print("|")
		w := g.LessonName
		fmt.Print(MakeSpacegrade(maxgrade(), w), " ", w, " ")
		fmt.Print("|")
		fmt.Print(" ", g.Score, " ")
		fmt.Print("|")
		fmt.Println()
		for i := 0; i < StudentIdlen+LessonNamelen+Scorelen+Idlen+6; i++ {
			fmt.Print(line)
		}
		fmt.Println()
	}
	fmt.Println("\t Student average score is :", averageScore)
	return w
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func Search() (string, string) {
	var name string
	fmt.Print("name :")
	fmt.Scan(&name)

	idlen := 3
	namelen := 6
	agelen := 5

	LessonNamelen := 11
	StudentIdlen := 10
	Scorelen := 6
	Idlen := 3

	line := "-"
	for i := 0; i < idlen+namelen+agelen+27; i++ {
		fmt.Print(line)
	}
	fmt.Println()
	fmt.Print("|")
	fmt.Print(" ", "Id", " ")
	fmt.Print("|")
	fmt.Print(" ", "Name", " ")
	fmt.Print("|")
	fmt.Print(" ", "Age", " ")
	fmt.Print("|")
	fmt.Println()
	for i := 0; i < idlen+namelen+agelen+27; i++ {
		fmt.Print(line)
	}
	fmt.Println()
	ok := false
	var w string
	var w2 string
	for _, s := range Students {
		if strings.Contains(s.Name, name) {
			ok = true
			fmt.Print("|")
			if s.Id < 10 {
				fmt.Print(" ", s.Id, "  ")
			} else {
				fmt.Print(" ", s.Id, " ")
			}
			fmt.Print("|")
			w2 = s.Name
			fmt.Print(Makespacestudent(Maxstudent(), w2), " ", w2, " ")
			fmt.Print("|")
			fmt.Print(" ", s.Age, "  ")
			fmt.Print("|")
			fmt.Println()

			for i := 0; i < idlen+namelen+agelen+27; i++ {
				fmt.Print(line)

			}

			fmt.Println()
			fmt.Print("|")
			fmt.Print(" ", "Student Id", " ")
			fmt.Print("|")
			fmt.Print(" ", "Id", " ")
			fmt.Print("|")
			fmt.Print(" ", "lesson Name", " ")
			fmt.Print("|")
			fmt.Print(" ", "Score", " ")
			fmt.Print("|")
			fmt.Println()
			for i := 0; i < StudentIdlen+LessonNamelen+Scorelen+Idlen+11; i++ {
				fmt.Print(line)
			}
			fmt.Println()
			for _, g := range greds {
				if g.StudentId == s.Id {
					fmt.Print("|")
					if g.StudentId < 10 {
						fmt.Print(" ", g.StudentId, "  ")
					} else {
						fmt.Print(" ", g.StudentId, " ")
					}
					fmt.Print("|")
					if g.Id < 10 {
						fmt.Print(" ", g.Id, "  ")
					} else {
						fmt.Print(" ", g.Id, " ")
					}
					fmt.Print("|")
					w := g.LessonName
					fmt.Print(MakeSpacegrade(maxgrade(), w), " ", w, " ")
					fmt.Print("|")
					fmt.Print(" ", g.Score, " ")
					fmt.Print("|")
					fmt.Println()
					for i := 0; i < StudentIdlen+LessonNamelen+Scorelen+Idlen+11; i++ {
						fmt.Print(line)

					}
				}
			}

		}

	}
	if !ok {
		fmt.Println("------------------------------------------------------------")
		fmt.Println("\t<*****> No such name found <*****>")
	}
	return w2, w
}

// ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||

// func DeleteStudent() {
// 	var id int
// 	fmt.Print("id :")
// 	fmt.Scan(&id)

// 	for i, s := range Students {
// 		if s.Id == id {

// 		}
// 	}
// }

// ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func MakeStudentId() int {
	return len(Students) + 1
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func MakeGradeId() int {
	return len(greds) + 1
}

// ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func maxgrade() int {
	max := 0
	for _, gg := range greds {
		if max < len(gg.LessonName) {
			max = len(gg.LessonName)
		}

	}
	return max
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func MakeSpacegrade(max int, w string) string {

	space := ""

	numberofspace := max - len(w)
	for range numberofspace {
		space += " "
	}
	return space
}

// ||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func Maxstudent() int {
	max2 := 0
	for _, ss := range Students {
		if max2 < len(ss.Name) {
			max2 = len(ss.Name)
		}
	}
	return max2
}

// |||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
func Makespacestudent(max2 int, w2 string) string {
	space2 := ""
	numderofspace := max2 - len(w2)
	for range numderofspace {
		space2 += " "
	}
	return space2
}
