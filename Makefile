MKFILE_PATH = $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR=$(dir $(MKFILE_PATH))

export GO11MODULE=on
export GOPROXY=https://goproxy.io
export GONOSUMDB=
export GOSUMDB=
export GOSUMDB=off

COMMIT = $(git log --format="%H" -n 1)

