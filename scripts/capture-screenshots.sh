
function example {

	echo '{"name":"guake-color-palettes","description":"command-line wrapper for varamenous guake colour palette generator.","bin":{"guake-cl":"src/cli/cli.js"},"homepage":"https://github.com/rgrannell1/guake-color-palettes#readme"}' | jq .

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
