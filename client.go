package main
import (
	"github.com/pkg/errors"
	"github.com/ypapax/tcp_ddos_golang/hashcash2"
	"log"
	"net"
	"os"
)
const PORT_CLIENT = "9001"
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := func() error {
		h := hashcash2.NewStd() // or .New(bits, saltLength, extra)
		// Mint a new stamp
		stamp, err := h.Mint("client_id")
		if err != nil {
			return errors.WithStack(err)
		}
		strEcho := stamp
		servAddr := "localhost:"+PORT_CLIENT
		tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
		if err != nil {
			println("ResolveTCPAddr failed:", err.Error())
			os.Exit(1)
		}

		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			println("Dial failed:", err.Error())
			os.Exit(1)
		}

		_, err = conn.Write([]byte(strEcho))
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		println("write to server = ", strEcho)

		reply := make([]byte, 1024)

		_, err = conn.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}

		println("reply from server=", string(reply))

		conn.Close()
		return nil
	}(); err != nil {
		log.Printf("error: %+v", err)
	}

}