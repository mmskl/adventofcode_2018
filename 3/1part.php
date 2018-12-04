<?php

$handler = fopen('input.txt', 'r');

$overlap = 0;
$i=0;
$map = [];

while (($line = fgets($handler)) !== false) {

    $line = trim($line);
    preg_match_all('/.*@\s(\d+),(\d+):\s(\d+)x(\d+).*/', $line, $subject);

    $fromLeft = $subject[1][0];
    $fromTop = $subject[2][0];

    $width = $subject[3][0];
    $height = $subject[4][0];


    for ($y = $fromTop; $y < $fromTop + $height; $y++) {
        for ($x = $fromLeft; $x < $fromLeft + $width; $x++) {

            if (!isset($map[$x][$y])) {
                $map[$x][$y] = false;
            } elseif(($map[$x][$y]) === false) {
                $map[$x][$y] = true;
                $overlap++;
            }
        }
    }

}

fclose($handler);
print_r($overlap);
