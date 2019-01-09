package main


import (
	"fmt"
	"net"
	"os"
	"time"
	"strconv"
)


/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}


const UDP_PORT string = "1122"


func main() {

	ListenerAddr, err := net.ResolveUDPAddr("udp", "127.17.0.1:1122")
	CheckError(err)

	MyAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)
	
	Conn, err := net.DialUDP("udp", MyAddr, ListenerAddr)
	CheckError(err)
	defer Conn.Close()

	i := 0
	for {
		msg := strconv.Itoa(i)
		buf := []byte(msg)

		_, err := Conn.Write(buf)
		CheckError(err)
		
		i++
		time.Sleep(time.Second * 1)
	}
}

