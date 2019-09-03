#!/bin/bash

mapandmount() {
    #rbd map test-"$1"k --pool scalable_sgx --name client.scalable
    #mount /dev/rbd/scalable_sgx/test-"$1"k /mnt/test-"$1"k
}


for i in {1..100};
do
    starttime=$(date +%s%3N)
    mapandmount $i
    endtime=$(date +%s%3N)
    diff=$(( $endtime - $starttime ))
    echo "$i, $diff" >> out
done