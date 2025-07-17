#!/bin/bash

set -e

pushd $(dirname $0) > /dev/null
SCRIPTPATH=$(pwd -P)
popd > /dev/null

#=========================================

cd ${SCRIPTPATH}/../..

echo "Please put the code here to run unit test"