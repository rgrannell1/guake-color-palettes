
"use strict"




const exec      = require('child_process').exec
const constants = require('../commons/constants')
const path      = require('path')





const setScheme = args => {

	if (args['--version']) {

		console.log(constants.package.version)
		process.exit(0)

	}

	if (args['--random']) {

		const availableSchemes = Object.keys(constants.options).filter(option => {

			const optionName = option.replace('--', '')

			if (args['-l'] || args['--light']) {

				return constants.lightSchemes.indexOf(optionName) !== -1

			} else if (args['-d'] || args['--dark']) {

				return constants.lightSchemes.indexOf(optionName) === -1

			} else {

				return true

			}

		})

		const chosenScheme     = availableSchemes[Math.floor(Math.random( ) * availableSchemes.length)]
		args[chosenScheme]     = true

	}

	Object.keys(args).forEach(option => {

		if (args[option] && constants.options.hasOwnProperty(option)) {

			const fpath      = constants.options[option].fileName
			const schemePath = path.join(constants.paths.palettes, fpath)

			exec(schemePath)
			process.exit(0)

		}

	})

}




module.exports = setScheme
