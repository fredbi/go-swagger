#!/bin/bash

prjdir=$(git rev-parse --show-toplevel)
cd $prjdir
. ./hack/deploy.sh
prepare
CIRCLE_PROJECT_USERNAME=go-swagger
CIRCLE_PROJECT_REPONAME=go-swagger
CIRCLE_SHA1=`git describe`
CIRCLE_TAG="vFRED-RELEASE"
API_EMAIL=fredbi@yahoo.com

LDFLAGS="-s -w -X github.com/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME/cmd/swagger/commands.Commit=${CIRCLE_SHA1}"
LDFLAGS="$LDFLAGS -X github.com/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME/cmd/swagger/commands.Version=${CIRCLE_TAG-dev}"
go build -ldflags "$LDFLAGS" -tags netgo -o ./dist/bin/swagger_linux_amd64 github.com/${CIRCLE_PROJECT_USERNAME}/${CIRCLE_PROJECT_REPONAME}/cmd/swagger
prepare_linuxpkg
build_linuxpkg deb
debsigs --sign=origin -k 61DCDA59 dist/build/swagger_FRED-RELEASE_amd64.deb


