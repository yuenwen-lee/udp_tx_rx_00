package main

import(
	"fmt"
	"net"
	"os"
)


/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}


func main() {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":1122")
	CheckError(err)
	fmt.Println(ServerAddr.Network(), ServerAddr.IP, ServerAddr.Port)
	
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Data: ", string(buf[0:n]), ", From ", addr)
		if (err != nil) {
			fmt.Println("Error: ", err)
		}
	}
}


