package main

import (
	"context"
	"exam/types"
	"sync"
)

func main() {
	ctx := context.Background()
	studentsInfo := types.GetStudentsInfo()
	teachersInfo := types.GetTeachersInfo()

	reviewChan := make(chan types.LabForReview, len(studentsInfo))

	students := NewStudents(studentsInfo, types.GetLabsTodo, reviewChan)
	teachers := NewTeachers(teachersInfo, reviewChan)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		students.Start(ctx)
	}()

	go func() {
		defer wg.Done()
		teachers.Start(ctx)
	}()

	wg.Wait()
}
