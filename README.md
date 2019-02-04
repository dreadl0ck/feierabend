# FEIERABEND - A mite integration for software developers

[![Go Report Card](https://goreportcard.com/badge/github.com/dreadl0ck/feierabend)](https://goreportcard.com/report/github.com/dreadl0ck/feierabend)
[![License](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://raw.githubusercontent.com/dreadl0ck/feierabend/master/docs/LICENSE)
[![Golang](https://img.shields.io/badge/Go-1.11-blue.svg)](https://golang.org)
![Linux](https://img.shields.io/badge/Supports-Linux-green.svg)
![macOS](https://img.shields.io/badge/Supports-macOS-green.svg)
![windows](https://img.shields.io/badge/Supports-windows-green.svg)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/dreadl0ck/feierabend)

Feierabend is a simple command-line utility to push time entries to the [mite](https://mite.yo.lk) time-tracking service,
that contain a generated note with all commits from one or several [git](https://git-scm.com) projects.

The whole process is interactive and asks the user how long he worked on each project prior to creating the time entries.

Creating time entries for days in the past is also possible.

The user config file *.feierabend.yml* is placed in the home directory and contains information required to authenticate to mite.
Optionally, a list of projects can be supplied, that will be checked on every execution:

```yaml
name: "Your Name"
apiKey: "<API-KEY>"
team: "Your Team"
userName: "you@company.com"
projects:
  - "/Users/you/Developer/awesome-project"
  - "/Users/you/Developer/project-xyz"
```

The project config file *.feierabend.yml* is placed in the root of the repository.
It contains the mite project and customer of the repository:

```yaml
customer: GoodCustomer
project: AwesomeProject
```

The idea here is, that at the end of a long working day, you are running the command

    $ feierabend

in your terminal, it will prompt you for the total amount of time you worked on each project and you are free to go!

To assist with the project configuration, the commandline tool can list all users, projects and customers to the terminal.

## Help

    $ feierabend -h
    Usage of feierabend:
    -customers
            list all available mite customers
    -date string
            set a date
    -debug
            toggle debug mode
    -dir string
            specify a project directory (default ".")
    -projects
            list all available mite projects
    -users
            list all available mite users
    -yesterday
            show yesterday

## How do I even pronounce this weird German word?

https://www.youtube.com/watch?v=WsZJNfqJDM4

It means end of the work day :)

## License

GPLv3