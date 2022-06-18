package solution

import "fmt"

func intToRoman(num int) string {
	if num < 1 || 3999 < num {
		return ""
	}

	var roman string

	for i := 0; num != 0; i++ {
		mod := num % 10
		v, err := romanVal(mod, i)
		if err != nil {
			return ""
		}
		roman = v + roman
		num /= 10
	}

	return roman
}

func romanVal(mod, i int) (string, error) {
	base, err := getBase(i)
	if err != nil {
		return "", err
	}

	if mod < 4 {
		var val string
		for idx := 0; idx < mod; idx++ {
			val += base
		}
		return val, nil
	}

	fiveBase, err := getFiveBase(i)
	if err != nil {
		return "", err
	}

	if mod == 4 {
		return base + fiveBase, nil
	}

	if mod < 9 {
		val := fiveBase
		for idx := 5; idx < mod; idx++ {
			val += base
		}

		return val, nil
	}

	tenBase, err := getTenBase(i)
	if err != nil {
		return "", err
	}
	return base + tenBase, nil
}

func getBase(i int) (string, error) {
	if i == 0 {
		return "I", nil
	}

	if i == 1 {
		return "X", nil
	}

	if i == 2 {
		return "C", nil
	}

	if i == 3 {
		return "M", nil
	}

	return "", fmt.Errorf("getBase can't handle %v", i)
}

func getFiveBase(i int) (string, error) {
	if i == 0 {
		return "V", nil
	}

	if i == 1 {
		return "L", nil
	}

	if i == 2 {
		return "D", nil
	}

	return "", fmt.Errorf("getFiveBase can't handle %v", i)
}

func getTenBase(i int) (string, error) {
	if i == 0 {
		return "X", nil
	}

	if i == 1 {
		return "C", nil
	}

	if i == 2 {
		return "M", nil
	}

	return "", fmt.Errorf("getTenBase can't handle %v", i)

}
