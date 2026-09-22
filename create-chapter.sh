#!/usr/bin/env bash

chapter=$1
workDir=$(pwd)
echo "creating ${workDir}/${chapter}"
mkdir -p ${workDir}/${chapter}/exercise
echo "done"
