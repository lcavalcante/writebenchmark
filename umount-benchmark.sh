#!/bin/bash
set -e

function usage {
    echo "Run bash rbd benchmark - unmap and umount"
    echo ""
    echo "Usage: `basename ${BASH_SOURCE[0]}` [OPTIONS]"
    echo ""
    echo "-h --help   Print this help message"
    echo "-n --num    Number of rbds to mount [default: 10]"
    echo ""
    echo "Example:   `basename ${BASH_SOURCE[0]}` --num 10"
    exit 1
}

while [ "$#" -gt 0 ]; do
    case "$1" in
        -n|--num)
            if (($# > 1)); then
                NUM=$2
                shift
            else
                usage
            fi ;;
        -h|--help)
            usage ;;
        *)
            usage ;;
    esac
    shift
done

unmapandumount() {
    starttime=$(date +%s%3N)
    umount /mnt/test-"$1"k
    rbd unmap test-"$1"k --pool scalable_sgx --name client.scalable
    endtime=$(date +%s%3N)
    diff=$(( $endtime - $starttime ))
    echo "$1, $diff" >> output-umount-"$NUM".csv
}

NUM=10
for (( c=1; c<=$NUM; c++ ))
do
    unmapandumount $c &
done
