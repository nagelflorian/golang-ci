# GoLang-CI [![Travis-CI build status](https://img.shields.io/travis/nagelflorian/golang-ci.svg)](https://travis-ci.org/nagelflorian/golang-ci)

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

## Travis-CI

You can see the build status here: https://travis-ci.org/nagelflorian/golang-ci

## Docker Hub

New Docker builds for Docker Hub are automatically triggered by Travis-CI:
https://hub.docker.com/r/floriannagel/golang-ci/
