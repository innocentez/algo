package sol_1

import (
	"fmt"
	"strconv"
)

func Solution() {
	testCases := [][]string{
		//{},                     // 4. пустой список
		//{"Hello", "World"},     // 1. обычный случай
		//{"Hello,World", "Foo"}, // 2. запятая внутри данных
		// {""}, // 3. одна пустая строка (из условия!)
		//{"", "", ""}, // 5. несколько пустых строк подряд
		//{"a#b:c", "d|e"},       // 6. куча "опасных" символов-разделителей
		//{"3", "5#abc"},         // 7. данные, похожие на служебную разметку
		//{"Привет", "мир 🌍"},    // 8. многобайтовые символы
	}

	for _, testCase := range testCases {
		code := encode(testCase)
		res := decode(code)

		fmt.Println("Code:", code)
		fmt.Println("Res", res)
	}
}

func encode(strs []string) string {
	var encodedData string
	for _, str := range strs {
		strLen := len(str)
		encodedData += strconv.Itoa(strLen) + "#" + str
	}

	return encodedData
}

func decode(encoded string) []string {
	var decodedData []string

	var strLen string
	var length int
	for i := 0; i < len(encoded); i++ {
		if length == 0 {
			switch encoded[i] {
			case '#':
				num, err := strconv.Atoi(strLen)
				if err != nil {
					panic(err)
				}
				length = num
			default:
				strLen += string(encoded[i])
				if int(strLen[0]-'0') == 0 {
					length = -1
				}
			}
			continue
		} else if length == -1 {
			decodedData = append(decodedData, "")
		} else {
			decodedData = append(decodedData, encoded[i:i+length])
			i = i + length - 1
		}

		length = 0
		strLen = ""
	}

	return decodedData
}
