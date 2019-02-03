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
	"io/ioutil"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type projectConfig struct {
	CustomerName string `yaml:"customer"`
	ProjectName  string `yaml:"project"`
}

type userConfig struct {
	Name     string   `yaml:"name"`
	APIKey   string   `yaml:"apiKey"`
	Team     string   `yaml:"team"`
	UserName string   `yaml:"userName"`
	Projects []string `yaml:"projects"`
}

func parseUserConfig() *userConfig {

	c, err := ioutil.ReadFile(os.Getenv("HOME") + "/.feierabend.yml")
	if err != nil {
		exitWith("failed to read user config", err)
	}

	var u = new(userConfig)

	err = yaml.Unmarshal(c, &u)
	if err != nil {
		exitWith("failed to unmarshal user config", err)
	}

	return u
}

func parseProjectConfig(path string) *projectConfig {

	err := os.Chdir(path)
	if err != nil {
		exitWith("failed to change the current directory to "+path+":", err)
	}

	c, err := ioutil.ReadFile(".feierabend.yml")
	if err != nil {
		exitWith(path+": failed to read project config", err)
	}

	var p = new(projectConfig)

	err = yaml.Unmarshal(c, &p)
	if err != nil {
		exitWith(path+": failed to unmarshal project config", err)
	}

	return p
}

func exitWith(message string, err error, values ...string) {

	var arr = []string{message, err.Error()}

	fmt.Println(strings.Join(append(arr, values...), " "))

	os.Exit(1)
}
