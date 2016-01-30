
"use strict"



var fs   = require('fs')
var path = require('path')





var constants = {
	paths: {
		palettes: path.resolve('palettes/')
	}
}

constants.options = { }





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
