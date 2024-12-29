package main

import (
	"html/template"
	"net/http"
	"strconv"
)

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

func HandleAddStudentForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/add_student.html"))
	tmpl.Execute(w, nil)
}

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

func HandleGetAllStudents(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/students.html"))
	tmpl.Execute(w, students)
}

func HandleAddGrades(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/add_grade.html"))
	tmpl.Execute(w, nil)
}

func HandleAddGradesForm(w http.ResponseWriter, r *http.Request) {
	lessonName := r.FormValue("lessonName")
	score, err := strconv.Atoi(r.FormValue("score"))
	if err != nil {
		http.Error(w, "Invalid Score", http.StatusBadRequest)
	}
	studentId, err := strconv.Atoi(r.FormValue("studentId"))
	if err != nil {
		http.Error(w, "Invalid StudentId", http.StatusBadRequest)
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