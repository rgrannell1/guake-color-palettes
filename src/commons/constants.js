
"use strict"



var fs = require('fs')





var constants = {
	palettes: 'palettes/'
}

constants.themes = fs
	.readdirSync(constants.palettes)
	.map(fpath => {
		console.log(fpath)
	})
