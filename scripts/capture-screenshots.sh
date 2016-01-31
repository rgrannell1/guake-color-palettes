
function example {

	echo '{"name":"guake-color-palettes","version":"0.1.0","description":"command-line wrapper for varamenous guake colour pallete generator.","preferGlobal":true,"engines":{"node":">=4.0.0"},"contributors":[{"name":"Ryan Grannell","email":"r.grannell@gmail.com"}],"main":"cli.js","scripts":{"postinstall":""},"repository":{"type":"git","url":"git+https://github.com/rgrannell1/guake-color-palettes.git"},"bin":{"guake-cl":"src/cli/cli.js"},"author":"Ryan Grannell","license":"MIT","bugs":{"url":"https://github.com/rgrannell1/guake-color-palettes/issues"},"homepage":"https://github.com/rgrannell1/guake-color-palettes#readme","dependencies":{"babel-eslint":"^5.0.0-beta8","babel-polyfill":"^6.3.14","docopt":"^0.6.2","eslint":"^1.10.3"}}' | jq .

}






for file in $( ls palettes )
do

	option=--$(echo "$file" | sed 's/^guake-//' | sed 's/[.]sh//' )
	outpath=$(echo $file | sed 's/[.]sh/.png/')



	# thank you, terrible bash security!! (this is why '--' exists!)
	guake-cl $option

	clear
	echo "$option"
	printf "\n\n"

	example

	sleep 1.5
	gnome-screenshot --file=images/$outpath
	convert -crop 300x200+0+30 images/$outpath images/$outpath

done
