#!/bin/bash -e

wget https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz -O /tmp/go-sdk.tar.gz
sha256sum /tmp/go-sdk.tar.gz > /tmp/checksum.txt
grep fb26c30e6a04ad937bbc657a1b5bba92f80096af1e8ee6da6430c045a8db3a5b /tmp/checksum.txt
if [ "$?" != "0" ]; then
    echo -n SHA checksum of Go distro does not match. Expected fb26c30e6a04ad937bbc657a1b5bba92f80096af1e8ee6da6430c045a8db3a5b, got
    cat /tmp/checksum.txt
    exit 1
fi

tar -C /usr/local -xzf /tmp/go-sdk.tar.gz

rm -f /tmp/go-sdk.tar.gz
rm -f /tmp/checksum.txt
