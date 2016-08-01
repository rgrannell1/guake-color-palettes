
"use strict"



const fs   = require('fs')
const path = require('path')





const constants = {
	paths: {
		palettes: path.resolve(path.join(__dirname, '../..', 'palettes/'))
	},
	options: {

	},
	lightSchemes: [
		'3024-day',
		'belafonte-day',
		'clrs',
		'github',
		'man-page',
		'novel',
		'pencil-light',
		'solarized-light',
		'spring',
		'tango',
		'terminal-basic',
		'tomorrow'
	],
	package: require('../../package')
}





fs
.readdirSync(constants.paths.palettes)
.map(fpath => {

	const names = {
		file:   fpath,
		scheme: fpath.replace(/^guake-/, '').replace(/[.]sh$/, '')
	}

	names.option = `--${names.scheme}`

	constants.options[names.option] = {
		fileName:   names.file,
		schemeName: names.scheme
	}

})





module.exports = constants
