FROM golang:1.23-bookworm AS build

WORKDIR /go/src/app
ADD . /go/src/app
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base-debian12
COPY --from=build /go/bin/app /
CMD ["/app"]
