#!/bin/bash
export EXTRAE_HOME=`pwd`/extrae_build/
source ${EXTRAE_HOME}/etc/extrae.sh
export EXTRAE_CONFIG_FILE=share/extrae.xml
export LD_LIBRARY_PATH=${EXTRAE_HOME}/lib

go run examples/taskloop.go
