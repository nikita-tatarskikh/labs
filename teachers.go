package main

import (
	"context"
	"exam/types"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Teacher struct {
	info types.TeacherInfo

	review chan types.LabForReview
}

const (
	rejectMessage = "Teacher %s %s rejected lab №%d. \"%s\" of student %s %s\n"
	acceptMessage = "Teacher %s %s accepted lab №%d. \"%s\" of student %s %s\n"
)

type Teachers []Teacher

func NewTeachers(teachersInfo types.TeachersInfo, reviewChan chan types.LabForReview) Teachers {
	teachers := make(Teachers, len(teachersInfo))

	for i, teacherInfo := range teachersInfo {
		teachers[i].info = teacherInfo
		teachers[i].review = reviewChan
	}

	return teachers
}

func (t Teacher) checkLabs(ctx context.Context) {
	for lab := range t.review {
		<-time.After(t.info.ReviewSpeed)

		if rand.Float64() < t.info.Probability {
			fmt.Printf(
				rejectMessage,
				t.info.Name, t.info.LastName,
				lab.LabInfo.Number, lab.LabInfo.Name,
				lab.StudentInfo.Name, lab.StudentInfo.LastName,
			)

			lab.Redo()
		} else {
			fmt.Printf(
				acceptMessage,
				t.info.Name, t.info.LastName,
				lab.LabInfo.Number, lab.LabInfo.Name,
				lab.StudentInfo.Name, lab.StudentInfo.LastName,
			)

			lab.Accept()
		}
	}
}

func (t Teachers) Start(ctx context.Context) {
	wg := sync.WaitGroup{}

	for _, teacher := range t {
		wg.Add(1)
		go func() {
			defer wg.Done()
			teacher.checkLabs(ctx)
		}()
	}

	wg.Wait()
}
