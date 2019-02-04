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
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

// get a note for the project at path
// the returned displayNote is separated by newlines
// and intended for displaying it to the user before submission
// this function decides at what date the git history will be checked
func getNoteForProject(dir string) (note string, displayNote string, date time.Time) {

	switch {

	// check yesterday
	case *flagYesterday:
		date = time.Now().AddDate(0, 0, -1)
		note, displayNote = getCommits(dir, date, uc.Name)

	// date was set explicitely
	case *flagDate != "":
		t, err := time.Parse(dateFormat, *flagDate)
		if err != nil {
			exitWith("invalid date:", err, *flagDate)
		}
		date = t
		note, displayNote = getCommits(dir, date, uc.Name)

	// check today
	default:
		date = time.Now()
		note, displayNote = getCommits(dir, date, uc.Name)
	}

	return
}

// returns the commits for the given day in the supplied directory at path
// the displayNote is separated by newlines and better readable at a quick glance
// the assembled git log query will also only display records for the supplied username
func getCommits(dir string, start time.Time, userName string) (note string, displayNote string) {

	// reset the date to the begining of the day
	start = zeroDate(start)

	debug("date:", start)

	// change directory to supplied path
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	// assemble query
	query := []string{"log", "--date=local", "--author=" + userName, "--since=" + start.Format(time.RFC3339) + "", "--until=" + start.AddDate(0, 0, 1).Format(time.RFC3339), "--pretty=format:%s"}

	debug("executing: git", query)

	// execute the desired git log query
	commits, err := exec.Command("git", query...).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	// split data by newlines
	lines := strings.Split(string(commits), "\n")
	if len(lines) == 1 {
		return "", lines[0]
	}

	// assemble human readable notes for mite
	// note is comma separated
	// displayNote is a newline separated list
	lastElem := len(lines) - 1
	for i, line := range lines {
		if i != lastElem {
			note += line + ", "
			displayNote += " - " + line + "\n"
			continue
		}
		note += line
		displayNote += " - " + line
	}

	return
}
