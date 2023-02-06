package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Calculator struct {
	Operation       string
	OperationActive string
	CalcActive      string
	Input           string
	Matrix1         [][]int
	Matrix2         [][]int
	Result          [][]int
	Det             int
	Output          string
}

//Todo написать инструкцию по корректной работе с калькулятором

func main() {
	var command string

	fmt.Println(
		"\nHello, a randomly wandering user.\n" +
			"This is a MATRIX calculator.\n\n" +
			"Here's what he can do:\n" +
			"<>Sum two matrices\n" +
			"<>Subtract two matrices\n" +
			"ATTENTION: Matrices can be sumed or subtracted if they are the same size!!!\n" +
			"<>Multiply two matrices\n" +
			"ATTENTION: In order for the matrix K to be multiplied by the matrix N,\n" +
			"it is necessary that the number of columns of the matrix K equals the number of rows of the matrix N!!!\n" +
			"<>Bring the matrix to the standard form\n" +
			"<>Find the determinants of the matrix\n\n" +
			"Instructions for using the calculator :\n" +
			"1. Write \"start\" and press \"enter\" \n" +
			"2. Write the input option in one word(\"file\", \"console\") and press \"enter\"\n" +
			" If you choose input from file, you should write your input in file\n" +
			" that lies in the directory matrix_calculator which names \"input.txt\"\n" +
			"3. Write the action you want to perform in one word(\"sum\", \"subtraction\", \"multiplication\", \"standardType\", \"determinant\")\n" +
			"4. Enter the data to calculate\n" +
			"5. If you want to output the result to a file,choose (\"yes\" or \"no\") write your choice and press \"enter\"\"\n" +
			"6. If you want to continue with this operation write \"continue\" and press \"enter\" \n" +
			"   If you want to stop with this operation write \"stop\" and press \"enter\" \n" +
			"7. If you want calculating farther write \"continue\" and press \"enter\"\n" +
			"   If you want to stop calculating write \"stop\" and press \"enter\" \n" +
			"8. Choose operation again\n" +
			"9. To complete the calculator, write \"stop\"\n" +
			"Example input:(enter the numbers as shown in the example) \n" +
			"3 3 - Size matrix А\n" +
			"Matrix А:\n" +
			"1 1 1\n" +
			"1 1 1\n" +
			"1 1 1\n" +
			"3 3 - Size matrix В\n" +
			"Matrix В:\n" +
			"1 2 3\n" +
			"4 5 6\n" +
			"7 8 9",
	)

	for {
		fmt.Println("\nWrite command...")
		fmt.Scan(&command)

		if command == "stop" {
			return
		}

		if command == "start" {
			calc := NewCalculator()
			for {
				if !StartCalculator(calc) {
					break
				}
			}
			continue
		}

		fmt.Println("Undefined command, see which commands the app supports and try again")
	}
}

func StartCalculator(c *Calculator) bool {
	fmt.Println("\nChoose operation which one you want to use:\n" +
		"sum\n" +
		"subtraction\n" +
		"multiplication\n" +
		"standardType\n" +
		"determinant",
	)

	for {
		_, err := fmt.Scan(&c.Operation)
		if err != nil {
			fmt.Println(err)
		}

		var flag bool

		if c.Operation == "sum" {
			for {
				for {
					if InputTwoMatrix(c) {
						break
					}
				}
				c.Sum()
				fmt.Println("Calculating...")
				OutputData(c)

				fmt.Println("If you want to continue with this operation, write \"continue\", if not write \"stop\"")

				for {
					fmt.Scan(&c.OperationActive)
					if c.OperationActive == "stop" {
						c.OperationActive = ""
						flag = false
						break
					}
					if c.OperationActive == "continue" {
						c.OperationActive = ""
						flag = true
						break
					}
					fmt.Println("Undefined, write your choice again")
				}

				if flag == false {
					break
				}
			}

			if flag == false {
				break
			}
		}

		if c.Operation == "subtraction" {
			for {
				for {
					if InputTwoMatrix(c) {
						break
					}
				}
				c.Subtraction()

				fmt.Println("Calculating...")
				OutputData(c)
				fmt.Println("If you want to continue with this operation, write \"continue\", if not write \"stop\"")

				for {
					fmt.Scan(&c.OperationActive)
					if c.OperationActive == "stop" {
						c.OperationActive = ""
						flag = false
						break
					}
					if c.OperationActive == "continue" {
						c.OperationActive = ""
						flag = true
						break
					}
					fmt.Println("Undefined, write your choice again")
				}

				if flag == false {
					break
				}
			}

			if flag == false {
				break
			}
		}

		if c.Operation == "multiplication" {
			for {
				for {
					if InputTwoMatrix(c) {
						break
					}
				}
				c.Multiplication()

				fmt.Println("Calculating...")
				OutputData(c)
				fmt.Println("If you want to continue with this operation, write \"continue\", if not write \"stop\"")

				for {
					fmt.Scan(&c.OperationActive)
					if c.OperationActive == "stop" {
						c.OperationActive = ""
						flag = false
						break
					}
					if c.OperationActive == "continue" {
						c.OperationActive = ""
						flag = true
						break
					}
					fmt.Println("undefined, write your choice again")
				}

				if flag == false {
					break
				}
			}

			if flag == false {
				break
			}
		}

		if c.Operation == "standardType" {
			for {
				for {
					if InputOneMatrix(c) {
						break
					}
				}

				c.StandardType()
				fmt.Println("Calculating...")
				OutputData(c)
				fmt.Println("If you want to continue with this operation, write \"continue\", if not write \"stop\"")

				for {
					fmt.Scan(&c.OperationActive)
					if c.OperationActive == "stop" {
						c.OperationActive = ""
						flag = false
						break
					}
					if c.OperationActive == "continue" {
						c.OperationActive = ""
						flag = true
						break
					}
					fmt.Println("undefined, write your choice again")
				}

				if flag == false {
					break
				}
			}

			if flag == false {
				break
			}
		}

		if c.Operation == "determinant" {
			for {
				for {
					if InputOneMatrix(c) {
						break
					}
				}

				c.Determinant()
				fmt.Println("Calculating...")
				OutputData(c)
				fmt.Println("If you want to continue with this operation, write \"continue\", if not write \"stop\"")

				for {
					fmt.Scan(&c.OperationActive)
					if c.OperationActive == "stop" {
						c.OperationActive = ""
						flag = false
						break
					}
					if c.OperationActive == "continue" {
						c.OperationActive = ""
						flag = true
						break
					}
					fmt.Println("undefined, write your choice again")
				}

				if flag == false {
					break
				}
			}

			if flag == false {
				break
			}
		}

		fmt.Println("Undefined operation, see how to enter the operations correctly and try again ")
	}
	fmt.Println("If you want to continue calculating, write \"continue\", if not write \"stop\"")

	for {
		fmt.Scan(&c.CalcActive)
		if c.CalcActive == "stop" {
			c.CalcActive = ""
			return false
		}
		if c.CalcActive == "continue" {
			c.CalcActive = ""
			return true
		}
		fmt.Println("undefined, write your choice again")
	}
}

func InputOneMatrix(c *Calculator) bool {
	fmt.Println("\nWrite the input option in one word(\"file\", \"console\") and press \"enter\"")
	for {
		fmt.Scan(&c.Input)
		if c.Input == "file" {
			f, err := ioutil.ReadFile("input.txt")
			if err != nil {
				fmt.Println("read fail", err)
				return false
			}

			var arrOfStr []string
			var newStr, str []byte
			var arrOfDigits []int

			for i := range f {
				if f[i] == 32 || f[i] == 10 {
					continue
				}

				str = append(str, f[i])

				if i != len(f)-1 {
					if f[i+1] == 32 || f[i+1] == 10 {
						arrOfStr = append(arrOfStr, string(str))
						str = newStr
					}
				}
				if i == len(f)-1 {
					arrOfStr = append(arrOfStr, string(str))
				}
			}

			for _, val := range arrOfStr {
				digit, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Неподходящие данные в файле, перезапишите файл и попробуйте снова")
					return false
				}
				arrOfDigits = append(arrOfDigits, digit)
			}

			if len(arrOfDigits) < 3 {
				fmt.Println("insufficient data, try again")
				return false
			}

			n := arrOfDigits[0]
			k := arrOfDigits[1]

			if n < 0 || k < 0 {
				fmt.Println("impossible matrix sizes")
				return false
			}

			if (c.Operation == "determinant") && (n != k) {
				fmt.Println("The determinant can be calculated only for a square matrix, enter the data again")
				return false
			}

			c.Matrix1 = make([][]int, n)
			for i := range c.Matrix1 {
				c.Matrix1[i] = make([]int, k)
			}

			z := 1
			for i := 0; i < n; i++ {
				for j := 0; j < k; j++ {
					z++
					c.Matrix1[i][j] = arrOfDigits[z]
				}
			}
			fmt.Println(c.Matrix1)
			break
		}

		if c.Input == "console" {
			fmt.Println("Write size of matrix")
			var i, j uint32

			if (c.Operation == "determinant") && (i != j) {
				fmt.Println("The determinant can be calculated only for a square matrix, enter the data again")
				return false
			}

			_, err := fmt.Scan(&i, &j)
			if err != nil {
				fmt.Println("wrong input, try again")
				fmt.Println(err)
				return false
			}

			c.Matrix1 = make([][]int, i)
			for i := range c.Matrix1 {
				c.Matrix1[i] = make([]int, j)
			}

			fmt.Println("write values of matrix")
			for i := 0; i < len(c.Matrix1); i++ {
				for j := 0; j < len(c.Matrix1[i]); j++ {
					_, err := fmt.Scan(&c.Matrix1[i][j])
					if err != nil {
						fmt.Println(err)
						return false
					}
				}
			}
			break
		}
		fmt.Println("Such input is not possible, try again")
	}
	return true
}

func InputTwoMatrix(c *Calculator) bool {
	fmt.Println("Write the input option in one word(\"file\", \"console\") and press \"enter\"")
	for {
		fmt.Scan(&c.Input)
		if c.Input == "file" {
			f, err := ioutil.ReadFile("input.txt")
			if err != nil {
				fmt.Println("read fail", err)
				return false
			}

			var arrOfStr []string
			var newStr, str []byte
			var arrOfDigits []int

			for i := range f {
				if f[i] == 32 || f[i] == 10 {
					continue
				}

				str = append(str, f[i])
				if i != len(f)-1 {
					if f[i+1] == 32 || f[i+1] == 10 {
						arrOfStr = append(arrOfStr, string(str))
						str = newStr
					}
				}

				if i == len(f)-1 {
					arrOfStr = append(arrOfStr, string(str))
				}
			}

			for _, val := range arrOfStr {
				digit, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println(err)
					fmt.Println("Неподходящие данные в файле, перезапишите файл и попробуйте снова")
					return false
				}
				arrOfDigits = append(arrOfDigits, digit)
			}

			if len(arrOfDigits) < 6 {
				fmt.Println("insufficient data, try again")
				return false
			}

			n := arrOfDigits[0]
			k := arrOfDigits[1]

			if n < 0 || k < 0 {
				fmt.Println("Impossible matrix sizes")
				return false
			}

			c.Matrix1 = make([][]int, n)
			for i := range c.Matrix1 {
				c.Matrix1[i] = make([]int, k)
			}

			z := 1
			for i := 0; i < n; i++ {
				for j := 0; j < k; j++ {
					z++
					c.Matrix1[i][j] = arrOfDigits[z]
				}
			}

			m := arrOfDigits[n*k+2]
			l := arrOfDigits[n*k+3]

			if m < 0 || l < 0 {
				fmt.Println("Impossible matrix sizes")
				return false
			}

			if (c.Operation == "sum" || c.Operation == "substraction") && (n != m || k != l) {
				fmt.Println("Matrices of different sizes!!!Try again")
				return false
			}

			if (c.Operation == "multiplication") && (k != m) {
				fmt.Println("These matrix cannot be Multiplied!!!Unsuitable size of matrix.Try again")
				return false
			}

			c.Matrix2 = make([][]int, m)
			for i := range c.Matrix2 {
				c.Matrix2[i] = make([]int, l)
			}

			x := n*k + 4
			for i := 0; i < m; i++ {
				for j := 0; j < l; j++ {
					fmt.Println(arrOfDigits[x])
					c.Matrix2[i][j] = arrOfDigits[x]
					x++
				}
			}
			break
		}

		if c.Input == "console" {
			fmt.Println("Write size of first matrix")
			var i, j uint32
			_, err := fmt.Scan(&i, &j)
			if err != nil {
				fmt.Println(err)
				fmt.Println("wrong input, try again")
				return false
			}
			c.Matrix1 = make([][]int, i)
			for i := range c.Matrix1 {
				c.Matrix1[i] = make([]int, j)
			}

			fmt.Println("\nWrite values of first matrix")
			for i := 0; i < len(c.Matrix1); i++ {
				for j := 0; j < len(c.Matrix1[i]); j++ {
					_, err = fmt.Scan(&c.Matrix1[i][j])
					if err != nil {
						fmt.Println(err)
						return false
					}
				}
			}

			fmt.Println("\nWrite size of second matrix")
			var m, n uint32
			_, err = fmt.Scan(&m, &n)
			if err != nil {
				fmt.Println(err)
				return false
			}

			if (c.Operation == "sum" || c.Operation == "substraction") && (i != m || j != n) {
				fmt.Println("Matrices of different sizes!!!Try again")
				return false
			}

			if (c.Operation == "multiplication") && (j != m) {
				fmt.Println("These matrix cannot be Multiplied!!!Unsuitable size of matrix.Try again")
				return false
			}

			c.Matrix2 = make([][]int, m)
			for m := range c.Matrix2 {
				c.Matrix2[m] = make([]int, n)
			}

			fmt.Println("\nWrite values of second matrix")
			for m := 0; m < len(c.Matrix2); m++ {
				for n := 0; n < len(c.Matrix2[m]); n++ {
					_, err = fmt.Scan(&c.Matrix2[m][n])
					if err != nil {
						fmt.Println(err)
						return false
					}
				}
			}
			break
		}
		fmt.Println("Such input is not possible, try again")
	}
	return true
}

func OutputData(c *Calculator) {
	fmt.Println("Result:")

	if c.Operation == "determinant" {
		fmt.Println(c.Det)
		return
	}

	for i := 0; i < len(c.Result); i++ {
		fmt.Println(c.Result[i])
	}

	fmt.Println("Do you want to output the result to a file,choose (\"yes\" or \"no\") write your choice and press \"enter\"")
	for {
		fmt.Scan(&c.Output)
		if c.Output == "yes" {
			fileName := "./result.txt"
			f, err := os.Create(fileName) // Создать файл
			if err != nil {
				fmt.Println("create file failed")
				return
			}

			w := bufio.NewWriter(f) // Создаем новый объект Writer

			if c.Operation == "determinant" {
				_, err = f.WriteString(strconv.Itoa(c.Det))
				if err != nil {
					fmt.Println(err)
				}
				err = w.Flush()
				if err != nil {
					fmt.Println(err)
					return
				}

				err = f.Close()
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			var allStr string
			for i := 0; i < len(c.Result); i++ {
				for j := 0; j < len(c.Result[i]); j++ {
					fmt.Println(string(byte(c.Result[i][j])))
					if j == len(c.Result[i])-1 {
						allStr += strconv.Itoa(c.Result[i][j]) + "\n"
						continue
					}
					allStr += strconv.Itoa(c.Result[i][j]) + " "
				}
				_, err = f.WriteString(allStr)
				if err != nil {
					fmt.Println(err)
				}

				allStr = ""
			}

			err = w.Flush()
			if err != nil {
				fmt.Println(err)
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Result also written in file")
			break
		}

		if c.Output == "no" {
			break
		}
		fmt.Println("undefined, write again \"yes\" or \"no\"")
	}
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Sum() bool {
	m := len(c.Matrix1)
	n := len(c.Matrix1[m-1])
	c.Result = make([][]int, m)
	for i := range c.Result {
		c.Result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c.Result[i][j] = c.Matrix1[i][j] + c.Matrix2[i][j]
		}
	}
	return true
}

func (c *Calculator) Subtraction() {
	m := len(c.Matrix1)
	n := len(c.Matrix1[m-1])
	c.Result = make([][]int, m)
	for i := range c.Result {
		c.Result[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c.Result[i][j] = c.Matrix1[i][j] - c.Matrix2[i][j]
		}
	}
}

func (c *Calculator) Multiplication() {
	//матрица m * n(сначала строки потом столбцы)
	//результатом произведения матрицы A m на k и B k на n будет матрица C m на n
	M := len(c.Matrix1)
	N := len(c.Matrix2[0])
	K := len(c.Matrix2)

	c.Result = make([][]int, M)
	for i := range c.Result {
		c.Result[i] = make([]int, N)
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			c.Result[i][j] = 0
			for k := 0; k < K; k++ {
				c.Result[i][j] += c.Matrix1[i][k] * c.Matrix2[k][j]
			}
		}
	}
}

func (c *Calculator) StandardType() {
	var buf, buf1 int
	n := len(c.Matrix1)
	c.Matrix1 = make([][]int, n)
	for i := range c.Matrix1 {
		c.Matrix1[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if c.Matrix1[i][i] == 0 {
				for z := i; z < n; z++ {
					buf1 = c.Matrix1[j][z]
					c.Matrix1[j][z] = c.Matrix1[i][z]
					c.Matrix1[i][z] = buf1
				}
				continue
			}
			buf = c.Matrix1[j][i] / c.Matrix1[i][i]
			for z := i; z < n; z++ {
				c.Matrix1[j][z] = c.Matrix1[j][z] - buf*c.Matrix1[i][z]
			}
		}
	}
	c.Result = c.Matrix1
}

func (c *Calculator) Determinant() {
	count := -1
	var buf, buf1 int
	c.Det = 1
	n := len(c.Matrix1)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if c.Matrix1[i][i] == 0 {
				for z := i; z < n; z++ {
					buf1 = c.Matrix1[j][z]
					c.Matrix1[j][z] = c.Matrix1[i][z]
					c.Matrix1[i][z] = buf1
					count++
				}
				continue
			}
			buf = c.Matrix1[j][i] / c.Matrix1[i][i]
			for z := i; z < n; z++ {
				c.Matrix1[j][z] = c.Matrix1[j][z] - buf*c.Matrix1[i][z]
			}
		}
		c.Det = c.Det * c.Matrix1[i][i]
	}

	for i := 0; i < count; i++ {
		c.Det = c.Det * (-1)
	}
}
