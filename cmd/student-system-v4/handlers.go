package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

// --------------------------------------------------------------------------------------------
func HandleGetAllGrades(w http.ResponseWriter, r *http.Request) {

	var studentGradeViews []StudentGradeView
	for _, grade := range grades {
		for _, student := range students {
			if grade.StudentID == student.ID {
				studentGradeViews = append(studentGradeViews, StudentGradeView{
					ID:          grade.ID,
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
	StudentBytes, err := json.Marshal(students)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("./cmd/student-system-v4/students.json", StudentBytes, 0644)
	if err != nil {
		panic(err)
	}
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
	GradeBytes, err := json.Marshal(grades)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("./cmd/student-system-v4/grades.json", GradeBytes, 0644)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/students", http.StatusSeeOther)
}

// --------------------------------------------------------------------------------------------
func HandleSeeStudentAverageScore(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Error(w, "Invalid Score", http.StatusBadRequest)
		return
	}

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

	// average
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
	err = tmpl.Execute(w, res)
	if err != nil {
		http.Error(w, "<****>error1<****>", http.StatusBadRequest)
	}
}
func HandleAverageform(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/average_score.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "<****>error2<****>", http.StatusBadRequest)
	}
}

// -----------------------------------------------------------------------------
func HandleSearch(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	matchedStudents := searchStudentsByName(name)

	var studentsWithGrades []StudentWithGrades
	for _, student := range matchedStudents {
		grades := getGradesByStudentID(student.ID)
		studentsWithGrades = append(studentsWithGrades, StudentWithGrades{
			Student: student,
			Grades:  grades,
		})
	}

	tmpl := template.Must(template.ParseFiles("templates/student_search.html"))
	tmpl.Execute(w, TemplateData{Students: studentsWithGrades})

}

func HandleSearchForm(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/student_search.html"))
	tmpl.Execute(w, nil)

}

// ---------------------------------------------------------------------------------
func HandlesearchLessonName(w http.ResponseWriter, r *http.Request) {

	LessonName := r.URL.Query().Get("LessonName")
	matchedGrades := searchGradesByLesson(LessonName)
	var gradesWithStudents []GradeWithStudents
	for _, grade := range matchedGrades {
		Students := getGradesByStudentId(grade.StudentID)
		gradesWithStudents = append(gradesWithStudents, GradeWithStudents{
			Students: Students,
			Grade:    grade,
		})
	}
	tmpl := template.Must(template.ParseFiles("templates/grade_search.html"))
	err := tmpl.Execute(w, GradeTemplateDate{Grades: gradesWithStudents})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
func HandlesearchLessonNameForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/grade_search.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template form", http.StatusInternalServerError)
	}
}

// ----------------------------------------------------------------------------------------------------
func HandelGetHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

// ----------------------------------------------------------------------------------------------------
func HandUpdateForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/update_student.html"))
	tmpl.Execute(w, nil)
}

// ----------------------------------------------------------------------------------------
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	updated := false
	for i, student := range students {
		if student.ID == id {
			students[i].Name = name
			students[i].Age = age
			updated = true
			break
		}
	}
	if !updated {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	studentBytes, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		http.Error(w, "Failed to serialize students", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("./cmd/student-system-v4/students.json", studentBytes, 0644)
	if err != nil {
		http.Error(w, "Failed to save students data", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

// ---------------------------------------------------------------------------------------
func HandUpdateGradeForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/update_grade.html"))
	tmpl.Execute(w, nil)
}

// ---------------------------------------------------------------------------------------
func UpdateGrade(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}
	score, err := strconv.Atoi(r.FormValue("score"))
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}
	lessonName := r.FormValue("lessonName")
	studentId, err := strconv.Atoi(r.FormValue("studentId"))
	if err != nil {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}
	update := false
	for i, grade := range grades {
		if grade.ID == id {
			grades[i].ID = id
			grades[i].LessonName = lessonName
			grades[i].Score = float64(score)
			grades[i].StudentID = studentId
			update = true
			break
		}
	}
	if !update {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	gradeBytes, err := json.MarshalIndent(grades, "", "  ")
	if err != nil {
		http.Error(w, "Failed to serialize students", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile("./cmd/student-system-v4/grades.json", gradeBytes, 0644)
	if err != nil {
		http.Error(w, "Failed to save students data", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
