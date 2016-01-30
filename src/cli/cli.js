#!/usr/bin/env node

"use strict"




const constants = require('../commons/constants')
const setScheme = require('../app/set-scheme')





const schemeOptions =
	Object.keys(constants.options)
	.join('|')

const optionDocs    =
	Object.keys(constants.options)
	.map(option => `	${option}`)
	.join('\n')


const docs = `
Name:
	guake-cl - Command line binding for changing Guake's colour scheme.

Usage:
	guake-cl (${schemeOptions})
	guake-cl (-h | --help | --version)

Options:
	-h, --help    Display this documentation.
${optionDocs}
`

const docopt = require('docopt').docopt

require('babel-polyfill')





setScheme(docopt(docs))
