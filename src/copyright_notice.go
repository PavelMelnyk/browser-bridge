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

package copyright_notice

import "os"

const (
	website = "http://code.google.com/p/browser-bridge/"
	copyright_str = "browser-bridge   Copyright (C)  2011  Tim Gerhard\nThis program comes with ABSOLUTELY NO WARRANTY; for details see "
	redist_info = "\nThis is free software, and you are welcome to redistribute it"
)


func Print() {
	os.Stdout.WriteString(copyright_str + website + redist_info + "\n\n")
}
