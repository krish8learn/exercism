package strand

func ToRNA(dna string) string {
	// panic("Please implement the ToRNA function")
	if len(dna) == 0 {
		return ""
	}

	rnaMap := map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}
	rna := ""
	for _, value := range dna {
		if val, present := rnaMap[string(value)]; present {
			rna += val
		}
		// rna += string(value)
	}

	return rna
}
