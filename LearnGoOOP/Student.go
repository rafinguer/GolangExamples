package main

import "fmt"

type Student struct {
	user      User // Relation One to One (one student = one user)
	studentId int
	courses   []Course // One to Many (one student = meny courses)
}

func (s *Student) String() string {
	return fmt.Sprintf("Student: %s, id: %d", s.user.String(), s.studentId)
}
