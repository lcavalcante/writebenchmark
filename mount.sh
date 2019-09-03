#!/bin/bash

mapandmount() {
    local starttime=$(date +%s%3N)
    rbd map test-"$1"k --pool scalable_sgx --name client.scalable
    mount /dev/rbd/scalable_sgx/test-"$1"k /mnt/test-"$1"k
    local endtime=$(date +%s%3N)
    local diff=$(( $endtime - $starttime ))
    echo "$1, $diff" >> mount-"$max".csv
}

max=100
for (( c=1; c<=$max; c++ ))
do
    mapandmount $c &
done
