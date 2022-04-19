package problem

import "fmt"

func ReplaceFour(data []string) []string {
	length := len(data)
	if data == nil || length == 0 {
		return nil
	}
	nilNum := 0
	for i := range data {
		if data[i] == " " {
			nilNum++
		}
	}
	tmp := length - 1
	length += 2 * nilNum
	tail := length - 1
	newData := make([]string, tail+1)
	for tmp > 0 && tail > tmp {
		if newData[tmp] == " " {
			tail--
			newData[tail] = "0"
			tail--
			newData[tail] = "2"
			tail--
			newData[tail] = "%"
		} else {
			tail--
			tmp--
			newData[tail] = data[tmp]
		}
		tmp--
	}
	return newData
}

func TestFour() {
	data := []string{"1", "2", " ", "4", " ", "6"}
	fmt.Println(data)
	newData := ReplaceFour(data)
	fmt.Println(newData)
}
