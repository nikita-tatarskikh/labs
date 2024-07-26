package types

import (
	"fmt"
	"time"
)

type TeacherInfo struct {
	Name        string
	LastName    string
	ReviewSpeed time.Duration
	Probability float64
}

type TeachersInfo []TeacherInfo

func GetTeachersInfo() TeachersInfo {
	teachersInfo := TeachersInfo{
		{
			Name:        "Dmitry",
			LastName:    "Popov",
			ReviewSpeed: 50 * time.Millisecond,
			Probability: 0.75,
		},
		{
			Name:        "Ivan",
			LastName:    "Holopov",
			ReviewSpeed: 150 * time.Millisecond,
			Probability: 0.65,
		},
	}

	fmt.Printf("Teachers:\n")
	for _, teacherInfo := range teachersInfo {
		fmt.Printf("Teacher: \"%s %s\"\n", teacherInfo.Name, teacherInfo.LastName)
	}

	return teachersInfo
}
