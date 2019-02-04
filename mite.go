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
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	mite "github.com/gosticks/go-mite"
)

var (
	m  *mite.Mite
	pc *projectConfig
)

// Thats us!
const appName = "Feierabend/v" + version

// create a mite api instance
// appName should be a discriptive string for you application (e.g. "my-app/v0.1")
func miteClient() *mite.Mite {
	return mite.NewMiteAPI(uc.UserName, uc.Team, uc.APIKey, appName)
}

// read a single project path, parse git and create note
// then initiate the upload process to mite
func readProject(path string) {
	note, displayNote, date := getNoteForProject(path)
	postNoteToMite(note, displayNote, path, date)
}

// asks the user how long he worked on the project
// time is parsed in golang duration format, e.g: 12h30m0s = 12 hours 30 minutes 0 seconds
func postNoteToMite(note, displayNote, path string, date time.Time) {

	// TODO: show current entries for the day
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("project:", path)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println(displayNote)
	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("How long did you work on the project?")
	fmt.Println("Enter a time the format xhxm (e.g. 3h45m = 3 hours 45 minutes). Hit [Enter] to continue")
	fmt.Print("> ")

	// Enter or it didn't happen.
	b, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	// Parse supplied value as duration
	d, err := time.ParseDuration(strings.TrimSpace(string(b)))
	if err != nil {
		panic(err)
	}

	debug("time:", d)

	// parse and set config for the current project
	pc = parseProjectConfig(path)

	// create the entry in mite
	createMiteEntry(note, d, date)
}

// fetch mite users and return the named one or nil
func getUserByName(name string) *mite.User {

	users, errUsers := m.GetUsers()
	if errUsers != nil {
		panic(errUsers)
	}

	for _, u := range users {
		if u.Name == name {
			return u
		}

	}
	return nil
}

// fetch available mite projects and return the one matching the supplied attributes
// or nil
func getProjectByName(customerName, name string) *mite.Project {

	projects, errProjects := m.GetAllProjects()
	if errProjects != nil {
		panic(errProjects)
	}

	for _, p := range projects {
		if p.Name == name && customerName == p.CustomerName {
			return p
		}
	}
	return nil
}

// list all available mite users to stdout
func listMiteUsers() {

	users, errUsers := m.GetUsers()
	if errUsers != nil {
		panic(errUsers)
	}

	for _, u := range users {
		fmt.Println(u)
	}
}

// list all available mite projects to stdout
func listMiteProjects() {

	projects, errProjects := m.GetAllProjects()
	if errProjects != nil {
		panic(errProjects)
	}

	for _, p := range projects {
		fmt.Println(p)
	}
}

// list all available mite customers
func listMiteCustomers() {

	customers, errCustomers := m.GetAllCustomers()
	if errCustomers != nil {
		panic(errCustomers)
	}

	for _, c := range customers {
		fmt.Println(c)
	}
}

// create a mite entry with the given note for the given duration and date
func createMiteEntry(note string, d time.Duration, date time.Time) {

	u := getUserByName(uc.Name)
	if u == nil {
		panic("invalid user")
	}

	p := getProjectByName(pc.CustomerName, pc.ProjectName)
	if p == nil {
		panic("invalid project")
	}

	// create a time entry instance
	entry := &mite.TimeEntry{
		Minutes: uint64(d.Minutes()), // uint64
		DateAt: mite.MiteTime{
			Time: date,
		}, // mite.MiteTime
		Note:         note,           // string
		Billable:     true,           // bool
		Locked:       false,          // bool
		HourlyRate:   30,             // uint64
		UserID:       u.ID,           // uint64
		UserName:     uc.UserName,    // string
		ProjectID:    p.ID,           // uint64
		ProjectName:  p.Name,         // string
		CustomerID:   p.CustomerID,   // uint64
		CustomerName: p.CustomerName, // string
		CreatedAt:    time.Now(),     // time
		UpdatedAt:    time.Now(),     // time
	}

	// pass that instance to the miteAPI
	// mite will create the item in the provided userID.
	// if the userID inside cannot be written to (coworker to coworker) the entry will be created in the apiKey user.
	resp, err := m.CreateTimeEntry(entry)
	if err != nil {
		panic(err)
	}

	fmt.Println("done:", resp)
}
