#!/bin/bash

(

echo "package engine";
echo;
echo "var (";
echo -n "	keyStrings = "\";
	mode=false;
	declare -a indexes;
	num=0;
	total=0;
	while read line; do
		if $mode; then
			if [ "$line" = ")" ]; then
				mode=false;
			else
				name="$(echo "$line" | cut -d'/' -f3 | tr -d "\n")";
				let "total += ${#name}";
				indexes[$num]=$total;
				echo -n "$name";
				let "num++";
			fi;
		else
			if [ "$line" = "const (" ]; then
				mode=true;
			fi;
		fi;
	done < keys.go;

	echo "\"";
	echo -n "	keyIndexes = [...]uint16{0";
	for index in ${indexes[@]}; do
		echo -n ", $index";
	done;
	echo "}";
	echo ")";
	echo;
	echo "func (k Key) String() string {";
	echo "	if k >= $num {";
	echo "		return \"UNKNOWN\"";
	echo "	}";
	echo "	return keyStrings[keyIndexes[k]:keyIndexes[k+1]]";
	echo "}";
) > keys_string.go
