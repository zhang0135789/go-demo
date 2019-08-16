package main

import "fmt"

/*
Go 错误处理
Go 语言通过内置的错误接口提供了非常简单的错误处理机制。
*/

//定义一个DivideError结构
type DivideError struct {
	dividee int
	divider int
}

//实现 error 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	// 正常情况
	//这种写法是go 的一种简写
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	result, errorMsg := Divide(100, 0)
	if errorMsg != "" {
		fmt.Printf("%d", result)
		fmt.Println("errorMsg is: ", errorMsg)
	}
}
