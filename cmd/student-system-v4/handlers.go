package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// --------------------------------------------------------------------------------------------
func HandleGetAllGrades(w http.ResponseWriter, r *http.Request) {

	var studentGradeViews []StudentGradeView
	for _, grade := range grades {
		for _, student := range students {
			if grade.StudentID == student.ID {
				studentGradeViews = append(studentGradeViews, StudentGradeView{
					StudentName: student.Name,
					StudentAge:  student.Age,
					LessonName:  grade.LessonName,
					Score:       grade.Score,
				})
			}
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/grades.html"))
	tmpl.Execute(w, studentGradeViews)
}

// --------------------------------------------------------------------------------------------
func HandleGetAllStudents(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/students.html"))
	tmpl.Execute(w, students)
}

// --------------------------------------------------------------------------------------------
func HandleAddStudentForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/add_student.html"))
	tmpl.Execute(w, nil)
}

// --------------------------------------------------------------------------------------------
func HandleAddStudent(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}

	if age > 100 {
		http.Error(w, "Invalid age must be less than 100", http.StatusBadRequest)
		return
	}

	// Add the new student
	newStudent := Student{
		ID:   len(students) + 1,
		Name: name,
		Age:  age,
	}
	students = append(students, newStudent)

	// Redirect back to the grades page
	http.Redirect(w, r, "/grades", http.StatusSeeOther)
}

// --------------------------------------------------------------------------------------------
func HandleAddGradeForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/add_grade.html"))
	tmpl.Execute(w, nil)
}

// --------------------------------------------------------------------------------------------
func HandleAddGrade(w http.ResponseWriter, r *http.Request) {
	lessonName := r.FormValue("LessonName")
	score, err := strconv.Atoi(r.FormValue("Score"))
	if err != nil {
		http.Error(w, "Invalid Score", http.StatusBadRequest)
		return
	}
	studentId, err := strconv.Atoi(r.FormValue("StudentId"))
	if err != nil {
		http.Error(w, "Invalid StudentId", http.StatusBadRequest)
		return

	}

	newGrade := Grade{
		ID:         len(grades) + 1,
		LessonName: lessonName,
		Score:      float64(score),
		StudentID:  studentId,
	}
	grades = append(grades, newGrade)

	http.Redirect(w, r, "/students", http.StatusSeeOther)
}

// --------------------------------------------------------------------------------------------
func HandleSeeStudentAverageScore(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.PathValue("id"))

	res := StudentAverageView{}

	for _, student := range students {

		if student.ID == id {
			res.StudentName = student.Name
		}
	}

	// var studentGrades []Grade
	for _, grade := range grades {
		if grade.StudentID == id {
			res.Grades = append(res.Grades, grade)
		}
	}

	// // average
	totalScore := 0.0
	count := 0
	for _, grade := range res.Grades {
		totalScore += float64(grade.Score)
		count++
	}

	if count == 0 {
		totalScore = 0
	} else {
		totalScore = totalScore / float64(count)
	}

	res.Average = totalScore

	tmpl := template.Must(template.ParseFiles("templates/average_score.html"))
	tmpl.Execute(w, res)
}

// -----------------------------------------------------------------------------
func HandleSearch(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")

	ok := false
	for _, student := range students {
		if strings.Contains(student.Name, name) {
			ok = true

		}
	}
	studentId, err := strconv.Atoi(r.FormValue("studentId"))
	if err != nil {
		http.Error(w, "Invalid StudentId", http.StatusBadRequest)
	}
	s := 0
	if ok {
		s = studentId
	}
	for _, student := range students {
		if student.ID == s {
			for _, grade := range grades {
				if student.ID == grade.StudentID {

				}
			}
		}
	}

	if !ok {
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/search.html"))
	tmpl.Execute(w, s)
}
