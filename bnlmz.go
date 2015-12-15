package bnlmz

func Cmp(numA, numB string) int {
	bn := BigNumber{}
	return bn.Cmp(numA, numB)
}

func Abs(num string) string {
	bn := BigNumber{}
	return bn.Abs(num)
}

func Add(numA, numB string) string {
	bn := BigNumber{}
	return bn.Add(numA, numB)
}

func Sub(numA, numB string) string {
	bn := BigNumber{}
	return bn.Sub(numA, numB)
}
