FROM golang:1.16 as build
WORKDIR /
ADD ./ ./
RUN go test -v
RUN go build

FROM scratch
COPY --from=build ./golang-ci /golang-ci
ENTRYPOINT ["/golang-ci"]
