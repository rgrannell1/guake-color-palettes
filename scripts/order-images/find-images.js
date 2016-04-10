
"use strict"



const fs       = require('fs')
const path     = require('path')
const pngparse = require('pngparse')





const asyncMap = (fn, data, callback, acc) => {

	if (!acc) {
		acc = [ ]
	}

	if (data.length === 0) {
		callback(acc)
	} else {

		fn(data.shift( ), mapped => {
			asyncMap( fn, data, callback, acc.concat([mapped]) )
		})

	}

}





const collectColourHistogram = image => {

	var colours = { }

	for (var ith = 0; ith < image.width; ++ith)	{
		for (var jth = 0; jth < image.width; ++jth)	{

			const pixel = image.getPixel(ith, jth)

			if (colours.hasOwnProperty(pixel)) {
				colours[pixel]++
			} else {
				colours[pixel] = 1
			}

		}
	}

	const data = Object.keys(colours)
		.map(colour => {

			return {
				colour,
				count:  colours[colour]
			}

		})
		.sort((pair0, pair1) => {
			return pair0.count - pair1.count
		})

	return {
		data,
		width:  image.width,
		height: image.height
	}

}





const minimiseDistances = imageData => {

	const histograms = imageData.map(info => {

		const histogram = collectColourHistogram(info.data)

		return {
			fpath:     info.fpath,
			data:      histogram.data,
			width:     histogram.width,
			height:    histogram.height
		}
	})

	const vectors = histograms.map(data => {

		console.log(data.data)
		throw 'xxx'

		return {
			fpath:  data.fpath,
			width:  data.width,
			height: data.height
		}

	})


}




const displaySorted = sorted => {
	console.log(sorted)
}





const findImageOrder = fpath => {

	const files = fs.readdirSync(fpath).map(fname => {
		return path.resolve(path.join(fpath, fname))
	})

	const schemaColours = files.filter(file => {
		return /[.]png$/.test(file) && /^guake/.test(path.basename(file))
	})

	asyncMap(
		(fpath, callback) => {
			pngparse.parseFile(fpath, (err, data) => {

				if (err) {
					throw err
				}

				callback({fpath, data})

			})
		},
		schemaColours,
		collected => {
			displaySorted(minimiseDistances(collected))
		}
	)

}





findImageOrder('../../images')
