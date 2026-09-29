#!/usr/bin/env bash
if [[ "$#" -ne 1 ]]; then
  echo "Need a chapter argument"
  exit 1
fi
chapter=$1
workDir=$(pwd)
echo "creating ${workDir}/${chapter}"
mkdir -p ${workDir}/${chapter}/exercise
touch ${workDir}/${chapter}/.gitkeep
echo "done"
