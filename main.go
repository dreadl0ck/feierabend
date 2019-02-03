/*
 * FEIERABEND
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
	"os"
)

const (
	// yyyy-mm-dd formatting of the magical reference date
	dateFormat = "2006-01-02"
)

var (
	uc = parseUserConfig()
)

func main() {

	flag.Parse()

	m = miteClient()

	var (
		note string
	)

	// TODO: add a nice table for the lists
	switch {
	case *flagListProjects:
		listMiteProjects()
		return
	case *flagListCustomers:
		listMiteCustomers()
		return
	case *flagListUsers:
		listMiteUsers()
		return
	case *flagDir != ".":
		note = getNoteForProject(*flagDir)
	case len(uc.Projects) > 0:

		for _, path := range uc.Projects {

			info, err := os.Stat(path)
			if err != nil {
				exitWith("invalid project path in user config:", err, path)
			}
			if !info.IsDir() {
				exitWith("project path from user config is not a directory:", err, info.Name())
			}

			note = getNoteForProject(path)

			postNoteToMite(note, path)
		}
	default:
		note = getNoteForProject(*flagDir)
		postNoteToMite(note, *flagDir)
	}
}
