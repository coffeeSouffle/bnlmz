package bnlmz

import (
	"errors"
	// "fmt"
	"regexp"
	"strconv"
	"strings"
)

// func Add(numA, numB string) (string, error) {

// }

// func Sub(numA, numB string) (string, error) {

// }

type BigNumber struct {
	NumA string
	NumB string

	intA   []string
	intB   []string
	floatA []string
	floatB []string
	typed  string
}

func (b *BigNumber) Parse(num string) (string, string, error) {
	str := ""
	re := regexp.MustCompile(`([0-9]+)(.([0-9]+))?`)
	match := re.FindStringSubmatch(num)

	if match[0] != num {
		err := errors.New("input error")
		return str, str, err
	}

	if len(match) != 4 {
		err := errors.New("struct error")
		return str, str, err
	}

	return match[1], match[3], nil
}

func (b *BigNumber) init(numA, numB string) error {

	b.NumA = numA
	b.NumB = numB

	intA, floatA, ok := b.Parse(b.NumA)
	if ok != nil {
		return ok
	}

	intB, floatB, ok := b.Parse(b.NumB)
	if ok != nil {
		return ok
	}

	intDiffLen := len(intA) - len(intB)

	if intDiffLen > 0 {
		for i := 0; i < intDiffLen; i++ {
			intB = "0" + intB
		}
	} else if intDiffLen < 0 {
		for i := 0; i < -1*intDiffLen; i++ {
			intA = "0" + intA
		}
		b.typed = "-"
	} else {
		if intA == intB {
			b.typed = "="
		} else {
			for i := 0; i < len(intA); i++ {
				if intB[i] > intA[i] {
					b.typed = "-"
					break
				}
			}
		}
	}

	floatDiffLen := len(floatA) - len(floatB)

	if floatDiffLen > 0 {
		for i := 0; i < floatDiffLen; i++ {
			floatB = floatB + "0"
		}
	} else if floatDiffLen < 0 {
		for i := 0; i < -1*floatDiffLen; i++ {
			floatA = floatA + "0"
		}
	}

	if b.typed == "=" {
		for i := 0; i < len(floatA); i++ {
			if floatB[i] > floatA[i] {
				b.typed = "-"
				break
			}
		}

		if b.typed == "=" {
			b.typed = ""
		}
	}

	if b.typed == "-" {
		b.intA = strings.Split(intB, "")
		b.intB = strings.Split(intA, "")
		b.floatA = strings.Split(floatB, "")
		b.floatB = strings.Split(floatA, "")
	} else {
		b.intA = strings.Split(intA, "")
		b.intB = strings.Split(intB, "")
		b.floatA = strings.Split(floatA, "")
		b.floatB = strings.Split(floatB, "")
	}

	return nil
}

func (b *BigNumber) Add(numA, numB string) (string, error) {
	if ok := b.init(numA, numB); ok != nil {
		return "", ok
	}

	number, err := b.additional("+")

	if err != nil {
		return "", err
	}

	return number, nil
}

func (b *BigNumber) Sub(numA, numB string) (string, error) {
	if ok := b.init(numA, numB); ok != nil {
		return "", ok
	}

	number, err := b.additional("-")

	if err != nil {
		return "", err
	}

	return b.typed + number, nil
}

func (b *BigNumber) additional(flag string) (string, error) {
	var (
		str    string
		istr   string
		tmpA   int
		tmpB   int
		tmp    int
		err    error
		tmpStr string
	)
	operator := 1
	op := 0
	length := len(b.floatA)

	if flag == "-" {
		operator = -1
	}

	if length > 0 {
		for i := length - 1; i >= 0; i-- {
			tmpA, err = strconv.Atoi(b.floatA[i])
			tmpB, err = strconv.Atoi(b.floatB[i])

			if err != nil {
				return "", err
			}

			tmp = tmpA + (operator * tmpB) + op

			if op == 1 || op == -1 {
				op = 0
			}

			if tmp >= 10 {
				tmp = tmp % 10
				op = 1
			}

			if tmp < 0 {
				tmp = 10 + tmp
				op = -1
			}

			tmpStr = strconv.Itoa(tmp)
			str = tmpStr + str
		}

		if ft, fok := strconv.Atoi(str); ft == 0 || fok != nil {
			str = "0"
		}

		str = "." + str
	}

	length = len(b.intA)

	for i := length - 1; i >= 0; i-- {

		tmpA, err = strconv.Atoi(b.intA[i])
		tmpB, err = strconv.Atoi(b.intB[i])

		if err != nil {
			return "", err
		}

		tmp = tmpA + (operator * tmpB) + op

		if op == 1 || op == -1 {
			op = 0
		}

		if tmp >= 10 {
			tmp = tmp % 10
			op = 1
		}

		if tmp < 0 {
			tmp = 10 + tmp
			op = -1
		}

		tmpStr = strconv.Itoa(tmp)

		istr = tmpStr + istr
	}

	if op == 1 {
		istr = "1" + istr
	}

	if it, iok := strconv.Atoi(istr); it == 0 || iok != nil {
		istr = "0"
	} else {
		istr = strconv.Itoa(it)
	}

	str = istr + str

	return str, nil
}

func Add(numA, numB string) (string, error) {
	bn := BigNumber{}
	return bn.Add(numA, numB)
}

func Sub(numA, numB string) (string, error) {
	bn := BigNumber{}
	return bn.Sub(numA, numB)
}
