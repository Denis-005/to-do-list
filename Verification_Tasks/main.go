package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

Поиск пересечения массивов
Напишите программу, которая находит пересечение двух массивов чисел и выводит только уникальные числа, входящие в оба массива.

Программа должна:

Прочитать два массива целочисленных чисел, введённых пользователем через пробел.
Найти пересечение массивов.
Вывести числа в порядке их появления во втором массиве.
Формат входных данных:
На вход подаётся две строки, каждая из которых содержит массив чисел, разделённых пробелами.

Формат выходных данных:
В одной строке через пробел: уникальные числа, входящие в оба массива, в порядке их появления во втором массиве.

*/

func conversionStr(str []string) []int {
	slice := []int{}
	for i := range str {
		num, _ := strconv.Atoi(str[i])
		slice = append(slice, num)

	}

	return slice
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	i := scanner.Text()

	tempSliceStr1 := strings.Split(i, " ")
	firstSliceInt := conversionStr(tempSliceStr1)

	scanner.Scan()
	n := scanner.Text()

	tempSliceStr2 := strings.Split(n, " ")
	secondSliceInt := conversionStr(tempSliceStr2)

	m := make(map[int]bool)
	for _, v := range firstSliceInt {
		m[v] = true
	}

	res := []int{}
	for _, v := range secondSliceInt {
		if m[v] {
			res = append(res, v)
		}
	}

	for i := range res {
		fmt.Print(res[i], " ")
	}
}

/*

Удаление дубликатов чисел
Напишите программу, которая удаляет все дублирующиеся числа из массива и выводит только уникальные числа в порядке возрастания.

Программа должна:

Прочитать массив чисел, введённый пользователем через пробел.
Удалить все дубликаты чисел.
Вывести массив уникальных чисел.
Формат входных данных:
На вход подаётся массив чисел, разделённых пробелами.

Формат выходных данных:
Все уникальные числа, разделённые пробелами в порядке возрастания.
*/

/*
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()

    tempSlice := strings.Split(input," ")

    slice := []int{}
    for i := range tempSlice {
        num,_ := strconv.Atoi(tempSlice[i])
        slice = append(slice,num)
    }

    m := make(map[int]bool)
    for _,v := range slice {
        m[v] = true
    }

    arr := []int{}
    for i := range m {
        arr = append(arr,i)
    }

    sort.Ints(arr)

    for i, v := range arr {
        if i > 0 {
            fmt.Print(" ")
        }
        fmt.Print(v)
    }
}

*/
