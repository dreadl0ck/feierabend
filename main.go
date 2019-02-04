/*
 * FEIERABEND - A mite integration for software developers
 * Copyright (c) 2018 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	// yyyy-mm-dd formatting of the magical reference date
	dateFormat = "2006-01-02"

	version = "0.1"
)

var (
	// global user config
	uc = parseUserConfig()
)

func main() {

	// parse flags
	flag.Parse()

	// create a new mite client
	m = miteClient()

	// TODO: add a nice table for the lists
	switch {

	case *flagListProjects:
		listMiteProjects()

	case *flagListCustomers:
		listMiteCustomers()

	case *flagListUsers:
		listMiteUsers()

	// project was set explicitely - ignore projects from user config
	case *flagDir != ".":
		readProject(*flagDir)

	// if there are projects in the user config, use these
	case len(uc.Projects) > 0:

		// iterate over project paths
		for _, path := range uc.Projects {

			// check path
			info, err := os.Stat(path)
			if err != nil {
				exitWith("invalid project path in user config:", err, path)
			}
			if !info.IsDir() {
				exitWith("project path from user config is not a directory:", err, info.Name())
			}

			// check project
			readProject(path)
		}

	// use the flag default: current directory
	default:
		readProject(*flagDir)
	}

	fmt.Println("All done! Enjoy your evening")
	fmt.Println("Bye bye.")
}
