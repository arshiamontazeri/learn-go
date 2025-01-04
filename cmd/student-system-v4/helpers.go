package main

import "strings"

func searchStudentsByName(name string) []Student {
	var result []Student
	for _, student := range students {
		if strings.Contains(strings.ToLower(student.Name), strings.ToLower(name)) {
			result = append(result, student)
		}
	}
	return result
}

func getGradesByStudentID(studentID int) []Grade {
	var result []Grade
	for _, grade := range grades {
		if grade.StudentID == studentID {
			result = append(result, grade)
		}
	}
	return result
}
