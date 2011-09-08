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






func Sendtext(ip string, port string, text string) (err int) {
	_,e := strconv.Atoi(port)
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

	conn.Write([]byte(text))

	mess := make([]byte,1024)
	conn.Read(mess)
	message := string(mess)

	conn.Close()

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

	_, port, pass, ip := browserbridge_config.ReadPropertiesFile()

	os.Stdout.WriteString("sending this url to " + ip + ":" + port + "\n")
	message := url + "\n" + pass + "\n"

	os.Stdout.WriteString("\nsending... ")

	e := Sendtext(ip, port, message)

	if e != 0 {
		os.Stdout.WriteString("ERROR\n")
		os.Exit(e);
	}
	os.Stdout.WriteString("DONE\n")
}
