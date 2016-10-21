#!/bin/bash

echo $1 $2 $3 ' -> echo $1 $2 $3'

#store in array
args=("$@")

#echo arguments to the shell
echo ${args[0]} ${args[1]} ${args[2]} ' -> echo args("$@"); echo ${args[0]} ${args[1]} ${args[2]}'

# use $@ to print out all argument at once
echo $@ ' -> echo $@'

#use $# variable to print out number of arguments
echo Number of arguments passed: $# ' -> print len'
