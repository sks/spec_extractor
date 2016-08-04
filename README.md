[![Build Status](https://travis-ci.org/sks/spec_extractor.svg?branch=develop)](https://travis-ci.org/sks/spec_extractor)

### Spec extractor
Given a bosh release, the command can print out all the specs

### Usage

##### Using source code
```go
go run main.go -d PATH_TO_BOSH_RELEASE
```
##### Using concourse
A Sample [concourse](https://concourse.ci) [pipeline.yml](https://github.com/sks/spec_extractor/blob/develop/ci/pipeline.yml) is provided which can be used as a reference
##### Using [piper](https://github.com/ryanmoran/piper)
```unix
piper -c ./ci/task/extract_specs/task.yml  -i spec_extractor=$PWD -i bosh_release=PATH_TO_BOSH_RELEASE
```
