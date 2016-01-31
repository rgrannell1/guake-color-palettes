
for file in $( ls palettes )
do

	option=--$(echo "$file" | sed 's/^guake-//' | sed 's/[.]sh//' )
	outpath=$(echo $file | sed 's/[.]sh/.png/')



	# thank you, terrible bash security!!
	guake-cl $option

	clear
	echo "$option"
	printf "\n\n"
	echo " Lorem ipsum dolor sit amet, urna est congue non sed aenean in, orci sapien id lectus pede elit, sit nec nibh lacinia justo pede, lacus quam semper interdum neque lobortis nibh, erat duis dolore est. Pellentesque sed et velit in. In nunc hac tortor, praesent fusce amet dui, tincidunt pellentesque quam nonummy penatibus, nibh dictum lorem eros, turpis eu purus. Laoreet tortor nunc. Parturient ut, nec aliquam egestas conubia taciti hendrerit, eget pulvinar pede, officiis morbi eget facilisis. Phasellus lacus fermentum posuere felis, hendrerit quis nec. Auctor praesent wisi natoque amet, fringilla vestibulum interdum odio porro facilisi, lorem est pretium pede suspendisse arcu. Ac hendrerit eum nulla. Elit dolor orci, commodo massa eleifend eget lectus. Sit massa ut. Ultricies etiam, tellus lacus, orci mattis enim magna in. Mattis purus maecenas orci aliquet, vivamus pellentesque, vestibulum gravida libero sed leo pede, pulvinar at sed dignissim ultricies facilisis, auctor tempor metus eu pharetra nulla."


	gnome-screenshot --file=images/$outpath
	convert -crop 300x200+0+30 images/$outpath images/$outpath

done
