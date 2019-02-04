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
	"io/ioutil"
	"os"
	"runtime"

	yaml "gopkg.in/yaml.v2"
)

var (
	pathSeparator  = "/"
	pathUserConfig = os.Getenv("HOME") + pathSeparator + configFileName
)

const configFileName = ".feierabend.yml"

// config for a single project
// defines
type projectConfig struct {
	CustomerName string `yaml:"customer"`
	ProjectName  string `yaml:"project"`
}

// config for a user
// contains data used to authenticate to mite
// and optionally a list of global projects to check when executed
type userConfig struct {
	Name     string   `yaml:"name"`
	APIKey   string   `yaml:"apiKey"`
	Team     string   `yaml:"team"`
	UserName string   `yaml:"userName"`
	Projects []string `yaml:"projects"`
}

// parse the user config and return an instance
func parseUserConfig() *userConfig {

	// check if we are running on windows at runtime
	// if true adjust the config file path accordingly
	if runtime.GOOS == "windows" {
		pathSeparator = "\\"
		pathUserConfig = os.Getenv("HOME") + pathSeparator + configFileName
	}

	// parse the user config
	c, err := ioutil.ReadFile(pathUserConfig)
	if err != nil {
		exitWith("failed to read user config", err)
	}

	// unmarshal yaml
	var u = new(userConfig)
	err = yaml.Unmarshal(c, &u)
	if err != nil {
		exitWith("failed to unmarshal user config", err)
	}

	return u
}

// parse the project config and return an instance
func parseProjectConfig(path string) *projectConfig {

	// change directory to path
	err := os.Chdir(path)
	if err != nil {
		exitWith("failed to change the current directory to "+path+":", err)
	}

	// parse the project config file
	c, err := ioutil.ReadFile(configFileName)
	if err != nil {
		exitWith(path+": failed to read project config", err)
	}

	// unmarshal yaml
	var p = new(projectConfig)
	err = yaml.Unmarshal(c, &p)
	if err != nil {
		exitWith(path+": failed to unmarshal project config", err)
	}

	return p
}
