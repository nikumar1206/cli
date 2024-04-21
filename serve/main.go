package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

func toColorBlue(s string) string {
	return fmt.Sprintf("\033[34m%s\033[0m", s)
}

func getLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
	dir := flag.String("d", ".", "directory in which to spin up local server")
	port := flag.String("p", "8000", "port")
	flag.Parse()

	handler := http.FileServer(http.Dir(*dir))
	http.Handle("/", handler)

	addr := ":" + *port

	localhostLink := toColorBlue("http:[::1]" + addr)
	localNetworkLink := toColorBlue("http://" + getLocalIP().String() + addr)

	fmt.Printf("Serving on port %s\n", localhostLink)
	fmt.Printf("Serving on local network: %s\n", localNetworkLink)

	http.ListenAndServe(addr, nil)
}
