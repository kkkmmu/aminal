#!/bin/bash

echo "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
echo "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
echo "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
echo "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"
echo "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"
echo "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"
echo -n "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"

echo -ne "\x1b[A" # up
echo -ne "\x1b[D" # left
echo -ne "\x1b[A" # up
echo -ne "\x1b[D" # left
echo -ne "\x1b[A" # up
echo -ne "\x1b[D" # left
echo -ne "\x1b[0J"
