package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var (
	students []Student
	grades   []Grade
)

func main() {

	// grade
	GradeBytes, err := os.ReadFile("./cmd/student-system-v4/grades.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(GradeBytes, &grades)
	if err != nil {
		panic(err)
	}

	GradeBytes, err = json.Marshal(grades)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./cmd/student-system-v4/grades.json", GradeBytes, 0644)
	if err != nil {
		panic(err)
	}
	// --------------------------------------------------------------------------------------------
	// student
	StudentBytes, err := os.ReadFile("./cmd/student-system-v4/students.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(StudentBytes, &students)
	if err != nil {
		panic(err)
	}

	StudentBytes, err = json.Marshal(students)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("./cmd/student-system-v4/students.json", StudentBytes, 0644)
	if err != nil {
		panic(err)
	}
	// -------------------------------------------------------------------------------------------
	r := http.NewServeMux()

	r.HandleFunc("GET /grades", HandleGetAllGrades)

	// Handle the form submission
	r.HandleFunc("POST /add-student", HandleAddStudent)

	// Render the form to add a new student
	r.HandleFunc("GET /add-student-form", HandleAddStudentForm)

	// Render all students
	r.HandleFunc("/students", HandleGetAllStudents)

	// Handle the form submission
	r.HandleFunc("POST /add-grade", HandleAddGrade)

	//render the form to add a new grade
	r.HandleFunc("GET /add-grade-form", HandleAddGradeForm)

	//Form for student's overall grade point average
	r.HandleFunc("GET /students/{id}/average", HandleSeeStudentAverageScore)

	// Start the server
	log.Println("Server started on http://127.0.0.1:8080")
	log.Println(http.ListenAndServe(":8080", r))
}
