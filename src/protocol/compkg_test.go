package protocol

import (
	"fmt"
	"testing"
)

func Test_GetBegin(t *testing.T){
	b, size := GetBegin()
	fmt.Println(b, size)
}

func Test_WriteHeader(t *testing.T){
	buffer := WriteHeader(10)
	fmt.Println(buffer.Bytes())
}

func Test_WriteBytes(t *testing.T){
	b := []byte{0x01, 0x02, 0x03, 0x04}
	buffer := WriteHeader(4)
	c, err := WriteBytes(buffer, b)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(c, buffer.Bytes())
}


//func Test_WriteString(t *testing.T){
//	s := "Hello"
//	buffer := WriteHeader(5)
//	c, err := WriteString(buffer, s)
//	if err != nil{
//		fmt.Println(err)
//	}
//	fmt.Println(c, buffer.Bytes())
//}

func Test_Write(t *testing.T){
	in := []byte{0x01, 0x02, 0x03, 0x04}
	b := Write(in)
	fmt.Println(b)
}

func Test_WriteStrData(t *testing.T){
	str := "How are you?"
	b := WriteStrData(str)
	fmt.Println(b)
}

func Test_Read(t *testing.T){
	in := []byte{35, 35, 35, 49, 46, 48, 0x0, 0x0, 0x0, 0x4, 0x01, 0x02, 0x03, 0x04}
	out, _ := Read(in)
	fmt.Println(out)
	//fmt.Println(string(out[0:]))
}

