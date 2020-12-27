FROM golang:latest AS build

# Setup
RUN mkdir -p /go/src/route-api
WORKDIR /go/src/route-api

# Copy & build
ADD . /go/src/route-api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -installsuffix nocgo -o /bin/app ./cmd/.
COPY ./cmd/config.yaml /bin/

FROM scratch
LABEL Name=routes-api
COPY --from=build /bin/. /bin/
ENTRYPOINT ["/bin/app"]
