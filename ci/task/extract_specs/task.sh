#!/bin/bash -exu

ROOT="${PWD}"

go get github.com/sks/spec_extractor

spec_extractor -d "${ROOT}"/bosh_release
