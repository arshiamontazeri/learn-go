package main

import (
	"log"
	"net/http"
)

var (
	students []Student
	grades   []Grade
)

func main() {

	// Sample data
	students = []Student{
		{ID: 1, Name: "Alice", Age: 20},
		{ID: 2, Name: "Bob", Age: 22},
	}

	grades = []Grade{
		{ID: 1, StudentID: 1, LessonName: "Math", Score: 95.5},
		{ID: 2, StudentID: 1, LessonName: "Science", Score: 88.0},
		{ID: 3, StudentID: 2, LessonName: "Math", Score: 78.0},
	}

	r := http.NewServeMux()

	r.HandleFunc("GET /grades", HandleGetAllGrades)

	// Handle the form submission
	r.HandleFunc("POST /add-student", HandleAddStudent)

	// Render the form to add a new student
	r.HandleFunc("GET /add-student-form", HandleAddStudentForm)

	// Render all students
	r.HandleFunc("/students", HandleGetAllStudents)

	// Start the server
	log.Println("Server started on http://127.0.0.1:8080")
	log.Println(http.ListenAndServe(":8080", r))
}
