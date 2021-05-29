#!/bin/sh
for((;;))
do
git fetch --all
git reset --hard origin/master
sleep 15
done
