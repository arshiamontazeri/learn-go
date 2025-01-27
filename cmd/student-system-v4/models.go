package main

type Student struct {
	ID   int
	Name string
	Age  int
}

type Grade struct {
	ID         int
	StudentID  int
	LessonName string
	Score      float64
}

type StudentGradeView struct {
	ID          int
	StudentName string
	StudentAge  int
	LessonName  string
	Score       float64
}

type StudentAverageView struct {
	StudentName string
	Average     float64
	Grades      []Grade
}

type StudentSearch struct {
	M map[Student][]Grade
}

type StudentWithGrades struct {
	Student
	Grades []Grade
}

type TemplateData struct {
	Students []StudentWithGrades
}
type GradesSearch struct {
	G map[Grade][]Student
}
type GradeWithStudents struct {
	Grade    Grade
	Students []Student
}
type GradeTemplateDate struct {
	Grades []GradeWithStudents
}
