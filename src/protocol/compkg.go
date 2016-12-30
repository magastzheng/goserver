package protocol

import(
	//"fmt"
	"bytes"
	"encoding/binary"
)

const(
	Begin = "###"
	Version = "1.0"
	VerLen = 3
	SizeLen = 4
)

func GetBegin() (b []byte, length int){
	b = []byte(Begin)
	length = len(b)
	
	return
}

func WriteHeader(length int)(buffer *bytes.Buffer){
	buffer = bytes.NewBufferString(Begin)
	buffer.WriteString(Version)
	
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(length))
//	binary.Write(buffer, binary.BigEndian, length)
	buffer.Write(bs)
	
	return
}

func WriteBytes(buffer *bytes.Buffer, data []byte)(n int, err error){
	return buffer.Write(data)
}

//func WriteString(buffer *bytes.Buffer, data string)(n int, err error){
//	return buffer.WriteString(data)
//}

func Write(data []byte)(b []byte){
	n := len(data)
	buffer := WriteHeader(n)
	c, err := WriteBytes(buffer, data)
	if err != nil || n != c {
		panic(err)
	} 
	
	b = buffer.Bytes()
	return
}

func WriteStrData(data string)(b []byte){
	in := []byte(data)
	
	b = Write(in)
	return
}

func Read(in []byte)(out []byte, n int){
	inLen := len(in)
	dataLen := 0
	begin, bLen := GetBegin()
	headerStart := 0
	isHeader := false
	isData := false
	//fmt.Println("====Read===")
	//fmt.Println(in)
	//fmt.Println(begin)
	for i := 0; i < inLen; i++{
		if !isHeader{
			if i+2 < inLen && in[i] == begin[0] && in[i+1] == begin[1] && in[i+2] == begin[2] {
				isHeader = true	
				headerStart = i
			}else{
				break
			}
		} else if isData {
			if i+dataLen < inLen{
				out = in[i: i+dataLen]
				n = i+dataLen
			}else{
				out = in[i:]
				n = inLen
			}
			return
		} else if isHeader && (i == headerStart + bLen + VerLen) {
			if i + SizeLen < inLen{
				//get the size length
				sizeB := in[i: i+SizeLen]
				var size int32
				binary.Read(bytes.NewReader(sizeB), binary.BigEndian, &size)
				dataLen = int(size)
				isData = true
				//fmt.Println("DataLen: ", dataLen)
				//fmt.Println(sizeB)
				i += VerLen
			}else{
				break
			}
		} else{
			//do nothing
		}
	}
	
	return
}
