package main

import "strings"

func searchGradesByLesson(LessonName string) []Grade {
	var result []Grade
	for _, grade := range grades {
		if strings.Contains(strings.ToLower(grade.LessonName), strings.ToLower(LessonName)) {
			result = append(result, grade)
		}
	}
	return result
}

func getGradesByStudentId(StudentID int) []Student {
	var result []Student
	for _, student := range students {
		if student.ID == StudentID {
			result = append(result, student)
		}
	}
	return result
}
