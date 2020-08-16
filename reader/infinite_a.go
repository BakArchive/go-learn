package main

type MyReader struct {

}

func (mr MyReader) Read(b []byte) (int,error){
	b[0] = 'A'
	return 1,nil
}

func main() {
	for ;;{
		MyReader{}.Read(make([]byte,8))
	}
}
