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
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func getNoteForProject(dir string) (note string) {

	switch {
	case *flagYesterday:
		note = dateFilteredByUser(dir, time.Now().AddDate(0, 0, -1), uc.UserName)
	case *flagDate != "":
		t, err := time.Parse(dateFormat, *flagDate)
		if err != nil {
			exitWith("invalid date:", err, *flagDate)
		}
		note = dateFilteredByUser(dir, t, uc.UserName)
	default:
		note = dateFilteredByUser(dir, time.Now(), uc.UserName)
	}

	return
}

func dateFilteredByUser(dir string, start time.Time, userName string) (out string) {

	start = zeroDate(start)

	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	commits, err := exec.Command("git", "log", "--author="+userName, "--since='"+start.Format(time.RFC3339)+"'", "--until='"+start.AddDate(0, 0, 1).Format(time.RFC3339)+"'", "--pretty=format:%s").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(commits), "\n")
	if len(lines) == 1 {
		return lines[0]
	}

	lastElem := len(lines) - 1
	for i, line := range lines {
		if i != lastElem {
			out += line + ", "
			continue
		}
		out += line
	}

	return
}

func todayUnfiltered() (out string) {

	commits, err := exec.Command("git", "log", "--since='"+*flagSince+"'", "--pretty=format:%s").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(commits), "\n")
	if len(lines) == 1 {
		fmt.Println(lines[0])
		os.Exit(0)
	}

	lastElem := len(lines) - 1
	for i, line := range lines {
		if i != lastElem {
			out += line + ", "
			continue
		}
		out += line
	}

	return
}

func yesterdayUnfiltered() (out string) {

	// time.Now().Format(time.RFC3339)

	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	yesterday = zeroDate(yesterday)
	today := yesterday.AddDate(0, 0, 1)

	fmt.Println("yesterday:", yesterday.Format(time.RFC3339))
	fmt.Println("today:", today.Format(time.RFC3339))

	commits, err := exec.Command("git", "log", "--since='"+yesterday.Format(time.RFC3339)+"'", "--until='"+today.Format(time.RFC3339)+"'", "--pretty=format:%s").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(commits), "\n")
	if len(lines) == 1 {
		fmt.Println(lines[0])
		os.Exit(0)
	}

	lastElem := len(lines) - 1
	for i, line := range lines {
		if i != lastElem {
			out += line + ", "
			continue
		}
		out += line
	}

	return
}

func zeroDate(t time.Time) time.Time {
	t = t.Add(-time.Duration(t.Hour()) * time.Hour)
	t = t.Add(-time.Duration(t.Minute()) * time.Minute)
	t = t.Add(-time.Duration(t.Second()) * time.Second)
	t = t.Add(-time.Duration(t.Nanosecond()) * time.Nanosecond)
	return t
}
