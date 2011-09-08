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

package paths

import "os"

func GetConfFilenames() (conffile_etc string, conffile_home string) {
        conffile_etc = "/etc/browserbridge.conf"
        conffile_home = os.Getenv("HOME") + "/.browserbridge.conf"
	return
}
