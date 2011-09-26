/*
	This file is part of browser-bridge.

    browser-bridge is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    browser-bridge is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with browser-bridge.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"os"
	"flag"
	"net"
	"./browserbridge_config"
	"strconv"
	"./copyright_notice"
)


// get the URL from execution parameters
func GetURL() (url string) {
	flag.Parse()
	s := ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
		}
		s += flag.Arg(i)
	}
	return s
}





// Send any text to given ip and port through TCP:
func Sendtext(ip string, port string, text string) (err int) {
	_,e := strconv.Atoi(port) // Atoi: StrToInt
   
  // Get a connection (get TCPAddr and dial) 
	targ := ip + ":" + port
	raddr,e := net.ResolveTCPAddr("tcp",targ)
	if e != nil {
		os.Stdout.WriteString(e.String()+"\n")
		return 1
	}
	conn,e := net.DialTCP("tcp",nil,raddr)
	if e != nil {
		os.Stdout.WriteString(e.String()+"\n")
		return 1
	}
  
  // Send text:
	conn.Write([]byte(text))

  // wait for server answer (ack or nak), then close the connection
	mess := make([]byte,1024)
	conn.Read(mess)
	message := string(mess)
	conn.Close()

  // anser was 'ack' or 'nak'. Check if it was ack
	if message[0] == 'a' {
		return 0
	} else {
		return 1
	}
	return 0
}





func main() {
	copyright_notice.Print()

	os.Stdout.WriteString("Will send URL: ")
	url := GetURL()
	os.Stdout.WriteString(url + "\n\n")

  // read program config
	_, port, pass, ip := browserbridge_config.ReadPropertiesFile()

	os.Stdout.WriteString("sending this url to " + ip + ":" + port + "\n")
	message := url + "\n" + pass + "\n" // network message convention between server and client to send url: "url\npassword\n"

	os.Stdout.WriteString("\nsending... ")

  // Send the message
	e := Sendtext(ip, port, message)

  // check if sending was successful, and report this.
	if e != 0 {
		os.Stdout.WriteString("ERROR\n")
		os.Exit(e);
	}
	os.Stdout.WriteString("DONE\n")
}
