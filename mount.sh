#!/bin/bash

mapandmount {
    starttime=$(date +%s%3N)
    rbd map test-"$1"k --pool scalable_sgx --name client.scalable
    mount /dev/rbd/scalable_sgx/test-"$1"k /mnt/test-"$1"k
    endtime=$(date +%s%3N)
    diff=$(( $endtime - $starttime ))
    echo "$i, $diff" >> out-"$max".csv
}


$max = 1000
for i in {1..$max};
do
    mapandmount $i &
done