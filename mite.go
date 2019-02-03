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
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	mite "github.com/gosticks/go-mite"
)

var m *mite.Mite
var pc *projectConfig

const appName = "feierabend/v0.1"

// create a mite api instance
// appName should be a discriptive string for you application (e.g. "my-app/v0.1")
func miteClient() *mite.Mite {
	return mite.NewMiteAPI(uc.UserName, uc.Team, uc.APIKey, appName)
}

func postNoteToMite(note, path string) {

	// TODO: show current entries for the day
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("project:", path)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println(note)
	fmt.Println("-----------------------------------------------------------------")

	fmt.Println("How long did you work on the project? Hit [Enter] to continue")
	fmt.Print("> ")

	b, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
	if err != nil {
		panic(err)
	}

	d, err := time.ParseDuration(strings.TrimSpace(string(b)))
	if err != nil {
		panic(err)
	}

	fmt.Println("time:", d)

	pc = parseProjectConfig(path)
	// createMiteEntry(note)
}

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

func getProjectByName(customerName, name string) *mite.Project {

	projects, errProjects := m.GetAllProjects()
	if errProjects != nil {
		panic(errProjects)
	}

	for _, p := range projects {
		fmt.Println(p.CustomerName)
		if p.Name == name && customerName == p.CustomerName {
			return p
		}
	}
	return nil
}

func listMiteUsers() {

	users, errUsers := m.GetUsers()
	if errUsers != nil {
		panic(errUsers)
	}

	for _, u := range users {
		fmt.Println(u)
	}
}

func listMiteProjects() {

	projects, errProjects := m.GetAllProjects()
	if errProjects != nil {
		panic(errProjects)
	}

	for _, p := range projects {
		fmt.Println(p)
	}
}

func listMiteCustomers() {

	customers, errCustomers := m.GetAllCustomers()
	if errCustomers != nil {
		panic(errCustomers)
	}

	for _, c := range customers {
		fmt.Println(c)
	}
}

func createMiteEntry(note string) {

	u := getUserByName(uc.UserName)
	if u == nil {
		panic("invalid user")
	}

	p := getProjectByName(pc.CustomerName, pc.ProjectName)
	if p == nil {
		panic("invalid project")
	}

	// create a time entry instance
	entry := &mite.TimeEntry{
		Minutes: 480, // uint64
		DateAt: mite.MiteTime{
			Time: time.Now(),
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
