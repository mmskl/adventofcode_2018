<?php

$handler = fopen('input.txt', 'r');

$overLap = 0;
$overlaping = [];
$all = [];
$map = [];

while (($line = fgets($handler)) !== false) {
    
    $line = trim($line);
    preg_match_all('/#(\d+)\s@\s(\d+),(\d+):\s(\d+)x(\d+).*/', $line, $subject);

    $id = (integer) $subject[1][0];
    $all[] = $id;
    $fromLeft = $subject[2][0];
    $fromTop = $subject[3][0];

    $width = $subject[4][0];
    $height = $subject[5][0];


    for ($y = $fromTop; $y < $fromTop + $height; $y++) {
        for ($x = $fromLeft; $x < $fromLeft + $width; $x++) {

            if (!isset($map[$x][$y])) {
                $map[$x][$y] = $id;
            } else {
                $overlaping[$map[$x][$y]] = true;
                $map[$x][$y] = $id;
                $overlaping[$id] = true;
                $overLap++;
            }


        }
    }

}

fclose($handler);
print_r($overLap);
var_dump(array_diff($all, array_keys($overlaping)));
