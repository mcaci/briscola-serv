package internal

import (
	"strconv"
)

func Points(args []string) (*cpEP, error) {
	var number string
	number, args = pop(args)
	n, _ := strconv.Atoi(number)
	ep := cpEP{
		number: uint32(n),
	}
	return &ep, nil
}

func Count(args []string) (*pcEP, error) {
	var numbers []uint32
	for _, arg := range args {
		n, _ := strconv.Atoi(arg)
		numbers = append(numbers, uint32(n))
	}
	return &pcEP{
		cardNumbers: numbers,
	}, nil
}

func Compare(args []string) (*ccEP, error) {
	var number string
	number, args = pop(args)
	fcnum, _ := strconv.Atoi(number)
	number, args = pop(args)
	fcseed, _ := strconv.Atoi(number)
	number, args = pop(args)
	scnum, _ := strconv.Atoi(number)
	number, args = pop(args)
	scseed, _ := strconv.Atoi(number)
	number, args = pop(args)
	brseed, _ := strconv.Atoi(number)
	return &ccEP{
		firstCardNumber:  uint32(fcnum),
		firstCardSeed:    uint32(fcseed),
		secondCardNumber: uint32(scnum),
		secondCardSeed:   uint32(scseed),
		briscolaSeed:     uint32(brseed),
	}, nil
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}
