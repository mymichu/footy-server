#!/bin/bash
GOROOT=/usr/local/go
export GOROOT
PATH=$GOROOT/bin:$PATH
export PATH
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
CONFIG_FILE_FOLDER=$DIR
export CONFIG_FILE_FOLDER 
