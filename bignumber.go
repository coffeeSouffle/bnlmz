package bnlmz

import (
	// "errors"
	"fmt"
	"regexp"
	"strconv"
	// "strings"
)

type BigNumber struct {
}

func (b *BigNumber) checkZero(num string) string {
	flag := false
	str := ""
	for i, lens := 0, len(num); i < lens; i++ {
		if string(num[i]) == "0" && !flag {
			continue
		}

		flag = true
		str = str + string(num[i])
	}

	if str == "" {
		return "0"
	}

	return str
}

func (b *BigNumber) Parse(num string) (string, string, string) {
	typed := ""
	re := regexp.MustCompile(`([-|+])?([0-9]+)(.([0-9]+))?`)
	match := re.FindStringSubmatch(num)

	if match[0] != num {
		panic("input data error")
	}

	if len(match) != 5 {
		panic("input data error")
	}

	if match[1] == "+" {
		typed = ""
	} else {
		typed = match[1]
	}

	return typed, match[2], match[4]
}

func (b *BigNumber) set(num string) string {
	typed, ints, floats := b.Parse(num)

	ints = b.checkZero(ints)

	if len(floats) > 0 {
		ints = ints + "." + floats
	}

	return typed + ints
}

func (b *BigNumber) Cmp(numA, numB string) int {
	typeA, intA, floatA := b.Parse(numA)
	typeB, intB, floatB := b.Parse(numB)

	if typeA != typeB {
		if typeA == "-" {
			return -1
		}

		return 1
	}

	if intA == intB && floatA == floatB {
		return 0
	}

	if len(intA) > len(intB) {
		if typeA == "-" {
			return -1
		}

		return 1
	}

	if len(intA) < len(intB) {
		if typeA == "-" {
			return 1
		}

		return -1
	}

	if intA > intB {
		if typeA == "-" {
			return -1
		}

		return 1
	}

	if intA < intB {
		if typeA == "-" {
			return 1
		}

		return -1
	}

	if floatA > floatB {
		if typeA == "-" {
			return -1
		}

		return 1
	}

	if typeA == "-" {
		return 1
	}

	return -1
}

func (b *BigNumber) Abs(num string) string {
	_, ints, floats := b.Parse(num)

	ints = b.checkZero(ints)

	if len(floats) > 0 {
		ints = ints + "." + floats
	}

	return ints
}

func (b *BigNumber) IsMinus(num string) bool {
	typed, _, _ := b.Parse(num)
	return (typed == "-")
}

func (b *BigNumber) Add(numA, numB string) string {
	var (
		result string
		typeA  string
		typeB  string
	)

	if b.IsMinus(numA) == true {
		typeA = "-"
	}

	if b.IsMinus(numB) == true {
		typeB = "-"
	}

	if typeA == typeB {
		result = b.additional(numA, numB, true)
		return typeA + result
	}

	absA := b.Abs(numA)
	absB := b.Abs(numB)
	cmp := b.Cmp(absA, absB)

	fmt.Println("Cmp => ", cmp)

	if cmp == 0 {
		return "0"
	}

	if cmp == 1 {
		return typeA + b.additional(numA, numB, false)
	}

	fmt.Println("TypeB => ", typeB)
	fmt.Println("TypeA => ", typeA)
	return typeB + b.additional(numB, numA, false)
}

func (b *BigNumber) Sub(numA, numB string) string {
	if b.IsMinus(numB) == true {
		return b.Add(numA, b.Abs(numB))
	}

	return b.Add(numA, "-"+numB)
}

func (b *BigNumber) additional(numA, numB string, flag bool) string {
	var (
		tmp  int
		tmpA int
		tmpB int

		iRet string
		fRet string

		ilen int
		flen int

		ok error
	)
	_, intA, floatA := b.Parse(numA)
	_, intB, floatB := b.Parse(numB)
	ilenA := len(intA)
	ilenB := len(intB)
	flenA := len(floatA)
	flenB := len(floatB)
	op := 0
	typed := 1

	if !flag {
		typed = -1
	}

	fmt.Println("ilenA =>", ilenA)
	fmt.Println("ilenB =>", ilenB)
	if ilenA > ilenB {
		for i, lens := 0, ilenA-ilenB; i < lens; i++ {
			intB = "0" + intB
		}
		ilen = ilenA
	} else {
		for i, lens := 0, ilenB-ilenA; i < lens; i++ {
			intA = "0" + intA
		}
		ilen = ilenB
	}

	if flenA > 0 || flenB > 0 {
		if flenA > flenB {
			for i, lens := 0, flenA-flenB; i < lens; i++ {
				floatB = floatB + "0"
			}
			flen = flenA
		} else {
			for i, lens := 0, flenB-flenA; i < lens; i++ {
				floatA = floatA + "0"
			}
			flen = flenB
		}

		for i := flen - 1; i >= 0; i-- {
			tmpA, ok = strconv.Atoi(string(floatA[i]))
			tmpB, ok = strconv.Atoi(string(floatB[i]))

			if ok != nil {
				panic("computing error")
			}

			tmp = tmpA + (typed * tmpB) + op

			if op == 1 || op == -1 {
				op = 0
			}

			if tmp >= 10 {
				tmp = tmp % 10
				op = 1
			}

			if tmp < 0 {
				tmp = tmp + 10
				op = -1
			}

			fRet = strconv.Itoa(tmp) + fRet
		}

		if num, _ := strconv.Atoi(fRet); num == 0 {
			fRet = "0"
		}
	}

	for i := ilen - 1; i >= 0; i-- {
		tmpA, ok = strconv.Atoi(string(intA[i]))
		tmpB, ok = strconv.Atoi(string(intB[i]))

		if ok != nil {
			panic("computing error")
		}

		tmp = tmpA + (typed * tmpB) + op

		if op == 1 || op == -1 {
			op = 0
		}

		if tmp >= 10 {
			tmp = tmp % 10
			op = 1
		}

		if tmp < 0 {
			tmp = tmp + 10
			op = -1
		}

		iRet = strconv.Itoa(tmp) + iRet
	}

	if op == 1 {
		iRet = "1" + iRet
	}

	iRet = b.checkZero(iRet)

	if len(fRet) > 0 {
		return iRet + "." + fRet
	}

	return iRet
}
