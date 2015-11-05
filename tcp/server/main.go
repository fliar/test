package main
import(
	"fmt"
	"net"
)

func main(){
	ln, err := net.Listen("tcp", ":8013")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println(conn.RemoteAddr().String(), "connected")
		conn.Close()
	}
}
