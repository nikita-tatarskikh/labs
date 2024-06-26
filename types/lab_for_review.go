package types

type LabForReview struct {
	LabInfo     Lab
	StudentInfo StudentInfo

	Redo   func()
	Accept func()
}
