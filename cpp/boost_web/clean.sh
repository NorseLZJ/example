#!/usr/bin/env bash

if [[ -d CMakeFiles ]]; then
	rm -rf CMakeFiles
fi

if [[ -f cmake_install.cmake ]]; then
	rm cmake_install.cmake
fi

if [[ -f CMakeCache.txt ]]; then
	rm CMakeCache.txt
fi

if [[ -f web ]]; then
	rm web
fi

# $1 debug|release

if [[ $1 == "1" ]]; then
	printf "DEBUG...\n"
	cmake -DCMAKE_BUILD_TYPE=Debug .
elif [[ $1 == "0" ]]; then
	printf "RELEASE...\n"
	cmake -DCMAKE_BUILD_TYPE=Release .
else
	printf "Nothing ...\n"
fi
