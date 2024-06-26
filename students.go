package main

import (
	"context"
	"exam/types"
	"fmt"
	"sync"
	"time"
)

type Student struct {
	lock *sync.Mutex

	Info types.StudentInfo

	labs   chan types.Lab
	review chan types.LabForReview

	labsDone  int
	labsTotal int
}

const sentMessage = "Student %s %s sent lab \"№%d. %s\" for review\n"

func (s *Student) doLabs(ctx context.Context) {
	for lab := range s.labs {
		s.doLab(ctx, lab)
	}
}

func (s *Student) doLab(ctx context.Context, lab types.Lab) {
	<-time.After(s.Info.Speed)

	s.review <- types.LabForReview{
		LabInfo: types.Lab{
			Number: lab.Number,
			Name:   lab.Name,
		},
		StudentInfo: types.StudentInfo{
			Name:     s.Info.Name,
			LastName: s.Info.LastName,
		},
		Redo: func() {
			s.labs <- lab
		},
		Accept: func() {
			s.lock.Lock()
			defer s.lock.Unlock()

			if s.labsTotal-s.labsDone == 1 {
				s.labsDone++
				close(s.labs)
			} else {
				s.labsDone++
			}
		},
	}

	fmt.Printf(sentMessage, s.Info.Name, s.Info.LastName, lab.Number, lab.Name)
}

type Students struct {
	students []Student
	review   chan types.LabForReview
}

func (s Students) Start(ctx context.Context) {
	wg := sync.WaitGroup{}

	go func() {
		wg.Wait()
		close(s.review)
	}()

	for _, student := range s.students {
		wg.Add(1)
		go func() {
			defer wg.Done()
			student.doLabs(ctx)
		}()
	}
}

func NewStudents(studentsInfo types.StudentsInfo, labsInfo func() types.LabsTodo, reviewChan chan types.LabForReview) Students {
	students := make([]Student, len(studentsInfo))

	for i, studentInfo := range studentsInfo {
		labs := labsInfo()
		students[i].Info = studentInfo
		students[i].labs = labs
		students[i].review = reviewChan
		students[i].labsTotal = len(labs)
		students[i].lock = new(sync.Mutex)
	}

	return Students{students: students, review: reviewChan}
}
