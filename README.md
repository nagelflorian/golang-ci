# GoLang-CI [![Travis-CI build status](https://img.shields.io/travis/nagelflorian/golang-ci.svg)](https://travis-ci.org/nagelflorian/golang-ci) [![CircleCI](https://circleci.com/gh/nagelflorian/golang-ci/tree/master.svg?style=svg)](https://circleci.com/gh/nagelflorian/golang-ci/tree/master)

## Run

```
go run hello-world.go
```

### Docker

You can also run the example app using Docker:

```
docker build -t golang-ci .
docker run --name=golang-ci -i -t golang-ci
```

## Testing

```
go test
```

## CircleCI

You can see the build status here: https://circleci.com/gh/nagelflorian/golang-ci

## Docker Hub

New Docker builds for Docker Hub are automatically triggered by Travis-CI:
https://hub.docker.com/r/floriannagel/golang-ci/
