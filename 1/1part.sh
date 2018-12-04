#!/usr/bin/env dash

#wget "https://adventofcode.com/2018/day/1/input" -O input

res=0

for op in $(cat input); do
    res=$(echo "${res}${op}" | bc)
done

echo "RESULT:"
echo "$res"
