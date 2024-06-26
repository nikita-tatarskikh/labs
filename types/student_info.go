package types

import (
	"fmt"
	"time"
)

type StudentInfo struct {
	Name     string
	LastName string
	Speed    time.Duration
}

type StudentsInfo []StudentInfo

func GetStudentsInfo() StudentsInfo {
	studentsInfo := StudentsInfo{
		{
			Name:     "Nikita",
			LastName: "Tatarskikh",
			Speed:    5 * time.Second,
		},
		{
			Name:     "Ivan",
			LastName: "Virolainen",
			Speed:    2 * time.Second,
		},
		{
			Name:     "Max",
			LastName: "Stolyarov",
			Speed:    3 * time.Second,
		},
	}

	fmt.Printf("Students:\n")
	for _, studentInfo := range studentsInfo {
		fmt.Printf("Student: \"%s %s\"\n", studentInfo.Name, studentInfo.LastName)
	}

	return studentsInfo
}
