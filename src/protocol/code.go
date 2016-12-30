package protocol

import (
	"fmt"
)

func Encode(data []byte)(out []byte){
	begin, bLen := GetBegin()
	dataLen := len(data)
	startPos := 0
	endPos := 0
	for i := 0; i < dataLen; i++{
		if data[i] == Sharp {
			endPos = i
			out = append(out, data[startPos: endPos])
			out = append(out, Trans)
			out = append(out, Sharp)
			startPos = endPos+1
		}
	}	
	
	if startPos == 0 {
		out = data[0:]
	}else if startPos >= endPos{
		out = append(out, data[startPos:])
	}else{
		//do nothing
	}
	
	return
}

func Decode(data []byte)(out []byte){
	begin, bLen := GetBegin()
	dataLen := len(data)
	startPos := 0
	endPos := 0
	for i := 0; i < dataLen; i++{
		if i < dataLen - 1 && data[i] == Trans && data[i+1] == Sharp {
			endPos = i
			out = append(out, data[startPos: endPos]
			startPos = i+1
		}
	}
	
	if startPos == 0{
		out = data[0:]
	}else if startPos >= endPos {
		out = append(out, data[startPos:])
	}else{
		//do nothing
	}

	return
}
