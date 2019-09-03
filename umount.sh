#!/bin/bash

unmapandumount() {
    starttime=$(date +%s%3N)
    umount /mnt/test-"$1"k
    rbd unmap test-"$1"k --pool scalable_sgx --name client.scalable
    endtime=$(date +%s%3N)
    diff=$(( $endtime - $starttime ))
    echo "$1, $diff" >> umount-"$max".csv
}

max=100
for (( c=1; c<=$max; c++ ))
do
    unmapandumount $c &
done
