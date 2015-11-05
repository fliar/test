package main
import(
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp" , "192.168.0.78:8013")
	if err != nil {
		panic(err)
	}
	localAddr := conn.LocalAddr().String()
	conn.Close()
	// i got the local port, what to do?
	fmt.Println(localAddr)
}
