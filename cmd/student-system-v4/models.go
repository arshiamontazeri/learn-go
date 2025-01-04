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
