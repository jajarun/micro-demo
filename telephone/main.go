package main

func letterCombinations(digits string) []string {
	var arr []string
	var temArr []string
	var phoneNums = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	for index, digit := range digits {
		numLetter := phoneNums[string(digit)]
		if index == 0 {
			for _, letter := range numLetter {
				arr = append(arr, string(letter))
			}
		} else {
			temArr = []string{}
			for _, letter := range numLetter {
				for _, exitLetter := range arr {
					temArr = append(temArr, exitLetter+string(letter))
				}
			}
			arr = temArr
		}
	}
	return arr
}

func smallestRepunitDivByK(k int) int {
	var v int

	return v
}

func main() {
	//arr := letterCombinations("2345")
	//fmt.Println(arr)
}
