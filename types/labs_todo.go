package types

type Lab struct {
	Number int
	Name   string
}

type Labs []Lab

type LabsTodo chan Lab

func GetLabsTodo() LabsTodo {
	labs := Labs{
		{
			Number: 1,
			Name:   "Channels",
		},
		{
			Number: 2,
			Name:   "Goroutines",
		},
		{
			Number: 3,
			Name:   "Atomics",
		},
	}

	labsChan := make(chan Lab, len(labs))

	for _, lab := range labs {
		labsChan <- lab
	}

	return labsChan
}
