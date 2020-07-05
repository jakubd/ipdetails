[![Actions Status](https://github.com/jakubd/referenceProject/workflows/Test/badge.svg)](https://github.com/jakubd/referenceProject/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/jakubd/referenceProject)](https://goreportcard.com/report/github.com/jakubd/referenceProject)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

# referenceProject

A very simple reference project for golang.  Used to bootstrap new projects but not have too 
much as to make implementation decisions about libraries used.  Basically meant to start
new projects and avoid repeating the same setup process each time.

Contains the following:

* simple main.go
* basic go.mod
* logging via logsrus
* directory template for additional packages
* testing (with asserts by testify lib)
* .gitignore
* github actions running tests against current active golang versions and on mac/linux. 
* makefile
