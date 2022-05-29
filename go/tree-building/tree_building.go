package tree

import "fmt"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	// panic("Please implement the Build function")

	//empty input
	if len(records) == 0 {
		return nil, nil
	}

	//sort the input : []Record{}
	position := make([]int, len(records))
	for index, record := range records {
		if record.ID < 0 || record.ID >= len(records) {
			return nil, fmt.Errorf("out of bound record %d", record.ID)
		}
		position[record.ID] = index
	}

	fmt.Println("Positions of input records", position)

	//node making
	nodes := make([]Node, len(records))
	for index := range records {
		record := records[position[index]]
		if record.ID != index {
			return nil, fmt.Errorf("non-contiguous node %d (want %d)", record.ID, index)
		}

		validParentforChild := (record.ID > record.Parent) || (record.ID == 0 && record.Parent == 0)

		if !validParentforChild {
			return nil, fmt.Errorf("node %d has impossible parent %d", record.ID, record.Parent)
		}

		//inital node
		nodes[index].ID = index
		if index != 0 {
			p := &nodes[record.Parent]
			p.Children = append(p.Children, &nodes[index])
		}
	}

	fmt.Println("->", nodes)
	fmt.Println("->", nodes[0])

	return &nodes[0], nil

}
