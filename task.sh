#!/bin/bash -exu

REPO=/tmp/diego-release
SPEC_FILE_REPO=~/workspace/github.com/sks/specs/github_com_cloudfoundry_incubator_diego_release

mkdir -p $SPEC_FILE_REPO

for tag in $(git -C $REPO tag --list | grep "^v\d" | grep -v "-" | gsort -V  )
do
  if [ -f $SPEC_FILE_REPO/$tag.spec ]
  then
    continue
  fi
  echo "Checking the tag $tag"
  pushd $REPO
    git checkout $tag
    git clean -xffd
    if [ -f ./scripts/update ]
    then
      ./scripts/update
    else
      git submodule update --init --recursive
    fi
  popd
  go run main.go -d $REPO > $SPEC_FILE_REPO/$tag.spec

done
