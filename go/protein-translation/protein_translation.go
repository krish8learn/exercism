package protein

import "errors"

func FromRNA(rna string) ([]string, error) {
	// panic("Please implement the FromRNA function")
	index := 0
	var ret []string
	for index < len(rna) {
		protien, err := FromCodon(rna[index : index+3])
		if err != nil {
			if err == ErrInvalidBase {
				return nil, ErrInvalidBase
			} else if err == ErrStop {
				break
			}
		}
		ret = append(ret, protien)
		index += 3
	}

	return ret, nil
}

func FromCodon(codon string) (string, error) {
	// panic("Please implement the FromCodon function")
	value, cond := seqToPro[codon]
	if !cond {
		return "", ErrInvalidBase
	}
	if value == "STOP" {
		return "", ErrStop
	}

	return value, nil
}

var seqToPro = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var ErrStop = errors.New("STOP sequence occured")
var ErrInvalidBase = errors.New("sequence not found")

/*
AUG	Methionine
UUU, UUC	Phenylalanine
UUA, UUG	Leucine
UCU, UCC, UCA, UCG	Serine
UAU, UAC	Tyrosine
UGU, UGC	Cysteine
UGG	Tryptophan
UAA, UAG, UGA	STOP
*/
