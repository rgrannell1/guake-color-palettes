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

Version:
	v${constants.package.version}

Description:
	guake-cl is a tool for choosing a guake colour-scheme.

	schemes can be picked by name, or picked at random from light or dark
	colour schemes.

Usage:
	guake-cl [-s | --set] (${schemeOptions})
	guake-cl [-s | --set] (-r | --random) [(-l | --light) | (-d | --dark)]
	guake-cl (-h | --help | --version)

Options:
	-h, --help      Display this documentation.
	-r, --random    Choose a random colour scheme.
	-l, --light     Display a light colour scheme.
	-d, --dark      Display a dark colour scheme.
	-s, --set       Set the chosen scheme as the current colour-scheme. Otherwise just
	                  print the required code to the console.
	--version       Display the package version.
${optionDocs}
`





const docopt = require('docopt').docopt

require('babel-polyfill')





setScheme(docopt(docs))
