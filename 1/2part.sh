#!/usr/bin/env dash

fre=0
> "result.txt"

while true; do
    for op in $(cat input.txt); do

        fre=$(echo "${fre}${op}" | bc)

        if grep -Fx -l --quiet  -- "$fre" "result.txt"; then
            echo "$fre"
            \rm "result.txt"
            exit
        fi

        echo $fre >> "result.txt"

    done
done
