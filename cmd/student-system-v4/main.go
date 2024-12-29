package main

import (
	"html/template"
	"net/http"
	"strconv"
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

	// Handle the homepage with grades
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// Render the form to add a new student
	http.HandleFunc("/add-student-form", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/add_student.html"))
		tmpl.Execute(w, nil)
	})

	// Handle the form submission
	http.HandleFunc("POST /add-student", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	// Render all students
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/students.html"))
		tmpl.Execute(w, students)
	})

	// Start the server
	http.ListenAndServe(":8080", nil)
}
