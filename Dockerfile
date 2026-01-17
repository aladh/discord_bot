FROM golang:1.25.6-trixie AS build

WORKDIR /go/src/app
ADD . /go/src/app
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base-debian12
COPY --from=build /go/bin/app /
CMD ["/app"]
