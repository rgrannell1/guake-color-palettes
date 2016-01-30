
"use strict"




var exec      = require('child_process').exec
var constants = require('../commons/constants')
var path      = require('path')





var setScheme = args => {

	Object.keys(args).forEach(option => {

		if (args[option] && constants.options.hasOwnProperty(option)) {

			const fpath      = constants.options[option].fileName
			const schemePath = path.join(constants.paths.palettes, fpath)

			exec(schemePath)

		}

	})

}




module.exports = setScheme
