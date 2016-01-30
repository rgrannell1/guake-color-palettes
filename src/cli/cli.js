#!/usr/bin/env node

"use strict"




const setScheme = require('../app/set-scheme')




const docs = `
Name:
	guake-cl - Command line binding for changing Guake's colour scheme.

Usage:
	guake-cl (-h | --help | --version)

Options:
	-h, --help    Display this documentation.
`

const docopt = require('docopt').docopt

require('babel-polyfill')





docopt(docs)
