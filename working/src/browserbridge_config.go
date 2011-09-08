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




/*

The only thing you need here, is ReadPropertiesFile() (rbrowsercommand string, rport string, rpass string, rip string)

*/

package browserbridge_config

import (
	"os"
	"bytes"
	"strings"
	"bufio"
	"./paths"
)




// reads a given config file
func ReadConfigFile(filename string) (browsercommand string, port string, pass string, ip string) {

	// set defaults
	browsercommand = "%b %u"
	port = "7896"
	pass = "hallo"
	ip = "127.0.0.1"

	// open file
	file, err := os.Open(filename)
	if err != nil {
		os.Stdout.WriteString("Error opening config file. proceeding with standard config...")
		return
	}
	
  // ToDo: close file
	reader := bufio.NewReader(file)

	for {
    // get a string from the next line; leave loop on eof
		part,_,err := reader.ReadLine()
		if err != nil {
			break
		}
		buffer := bytes.NewBuffer(make([]byte,0,2048))
		buffer.Write(part)
		s := strings.ToLower(buffer.String())
		
    // remove everything after a #
		if pos := strings.Index(s,"#"); pos != -1 {
			s = s[:pos]
		}
		
    // if there is no = in this line, assume this is a comment, and skip.
		if pos := strings.Index(s,"="); pos != -1 {
      // split the line to 2 strings, on the first =
			arr := strings.SplitN(s,"=",2)
      
      //TrimSpace: allow some spaces before and after the = or at beginning/end of the line
			key := strings.TrimSpace(arr[0])
			val := strings.TrimSpace(s[pos+1:])
			
      // find out which key is given, and set its value
			switch key {
				case "browsercommand": browsercommand = val
				case "ip": ip = val
				case "password": pass = val
				case "port": port = val
				default: os.Stdout.WriteString("\nUnknown Key in config: \"" + key + "\"")
			}
		}
	}
  // get default browser, replace %b with it.
	browsercommand = strings.Replace(browsercommand,"%b",os.Getenv("BROWSER"),-1)
	return
}


// determinates from where to read the config file, and then reads it
func ReadPropertiesFile() (browsercommand string, port string, pass string, ip string) {
	conffile_etc,conffile_home := paths.GetConfFilenames() // get paths
	
	// Try to open the file in the user home. If that had success (means: file exists), then must be f!=nil and err==nil
	f,err := os.OpenFile(conffile_home,os.O_RDONLY,0)

	// Read out from the needed ConfFile. (success in last command: from home, on failure: from /etc/browserbridge/
	if (f != nil) && (err == nil) {
		os.Stdout.WriteString("Reading config from" + conffile_home + "\n")
                browsercommand,port,pass,ip = ReadConfigFile(conffile_home)
        } else {
                os.Stdout.WriteString("Reading config from" + conffile_etc + "\n")
                browsercommand,port,pass,ip = ReadConfigFile(conffile_etc)
        }
	
	// return all
	return
}
