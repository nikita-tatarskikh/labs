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

	//fmt.Printf("Labs:\n")
	//for _, lab := range labs {
	//	fmt.Printf("Lab: \"%d. %s\"\n", lab.Number, lab.Name)
	//}

	labsChan := make(chan Lab, len(labs))

	for _, lab := range labs {
		labsChan <- lab
	}

	return labsChan
}
