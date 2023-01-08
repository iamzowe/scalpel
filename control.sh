#!/bin/bash

function check_code() {
  go fmt
  check_code_by_dir .
}

function check_code_by_dir() {
  for fn in `ls $1`
  do
    if [ -d $1"/"$fn ];then
      if [ "$fn" != "vendor" ];then
        go fmt $1"/"$fn
        go vet $1"/"$fn
        check_code_by_dir $1"/"$fn
      fi
    fi
  done
}

function usage() {
  echo "Usage: $0 {check}"
  exit 1
}

if [ $# != 1 ]; then
  usage
fi

case "$1" in
    check)
        check_code
        ;;
    *)
    usage
esac
