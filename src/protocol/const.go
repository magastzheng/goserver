package protocol

const(
	Sharp = '#'
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
