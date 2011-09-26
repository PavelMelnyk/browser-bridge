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
	"strings"
	"net"
	"exec"
	"./browserbridge_config"
	"./copyright_notice"
)


// Opens a url in browser
func OpenURL(url string, browsercommand string) {
	comm := strings.Replace(browsercommand,"%u",url,1) // replace %u with received url
	os.Stdout.WriteString("executing:  " + comm + "\n\n")
   
  // split command from arguments
	comms := strings.Split(comm," ")
	args := ""
	for i := 1; i < len(comms); i++ {
		if i > 1 {
			args += " "
		}
		args += comms[i]
	}
   
  // execute
	cm := exec.Command(comms[0],args)
	cm.Start()
}


// React on incoming TCP connection
func HandleClient(conn *net.TCPConn, password string, browsercommand string) {

  // prepare for message, and wait for it
	mess := make([]byte,1024)
	conn.Read(mess)
   
  // get message
	message := string(mess)
	os.Stdout.WriteString("\n\ngot this url:  ")
	s := strings.Split(message,"\n") // network message convention between server and client: "url\npassword\n"
   
  // s should now contain 2 elements. s[0]=url, s[1]=password
  if len(s) < 2 {
		os.Stdout.WriteString("ERROR: can't understand client.")
		conn.Write([]byte("nak")) //send nak if url wasn't opened
		return
	}
  url := s[0]
	os.Stdout.WriteString(url + "\n")
	pass := strings.Replace(s[1],"\n","",-1)
   
  //check password, reply
	if pass != password {
		os.Stdout.WriteString("invalid password!\n\n")
			conn.Write([]byte("nak"))
	} else {
		os.Stdout.WriteString("password seems correct. Opening URL...\n")
		conn.Write([]byte("ack"))
		OpenURL(url,browsercommand)
	}
  
	conn.Close()
}




func main() {
	copyright_notice.Print()

	os.Stdout.WriteString("Starting Server Interface...\n")
	browsercommand, port, password, _ := browserbridge_config.ReadPropertiesFile()

  // start listener
	tcpad,_ := net.ResolveTCPAddr("tcp",":"+port)
	list,_ := net.ListenTCP(tcpad.Network(), tcpad)
	defer list.Close()
   
  // enter endless loop, waiting for clients
	for {
		os.Stdout.WriteString("waiting for clients on port "+port+"...\n")
		conn,_ := list.AcceptTCP()

		os.Stdout.WriteString("Client connected: "+ conn.RemoteAddr().String())

		HandleClient(conn,password,browsercommand)
	}
}
