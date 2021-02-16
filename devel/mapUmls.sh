#!/bin/bash

declare -A accents labels
accents=( ["DIAERESIS"]="Dia" ["TILDE"]="Tld" ["GRAVE"]="Grv" ["CIRCUMFLEX"]="Crc" ["ACUTE"]="Agu" ["CEDILLA"]="Ced" ["STROKE"]="Strk")
labels=( ["DIAERESIS"]="a diaresis" ["TILDE"]="a tilde" ["GRAVE"]="an accent grave" ["CIRCUMFLEX"]="a circumflex" ["ACUTE"]="an accent aigu" ["CEDILLA"]="cedilla" ["STROKE"]="strikethrough")

umllines=$(curl "https://utf8-chartable.de/" | sed 's|</tr><tr>|\n|g' | grep "LATIN")

while IFS= read -r line; do
	#echo "... $line ..."
	uml=$(echo $line | grep -o "<td class=\"char\">.</td>" | cut -c 18-19)
	case=$(echo $line | grep -e "SMALL" -e "CAPITAL" -o)
	caseshort=$(echo $case | cut -c -1)
	case=$(echo $case | tr 'A-Z' 'a-z') # don't scream at me
	letter=$(echo $line | grep "LETTER . WITH" -o | awk '{print $2}')
	accent=$(echo $line | grep "WITH .*</td>" -o | sed 's/WITH //g' | sed 's|</td>||g')

	hex=$(echo $line | grep -oE "[0-9a-f]{2} [0-9a-f]{2}" | cut -c 4-)

	accentshort=${accents[$accent]}
	accentdesc=${labels[$accent]}
	varname=$(printf "Uml%s%s%s" $caseshort $letter $accentshort)

	printf "// %s is a %s %s with %s: %s\n" $varname $case $letter "$accentdesc" "$uml"
	printf " %s = 0x%s\n\n" $varname "$hex"

done <<< "$umllines"
