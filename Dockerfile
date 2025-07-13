FROM golang:1.10 AS build
WORKDIR /go/src
COPY server/server ./server/server
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o server .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/server ./
EXPOSE 8080/tcp
ENTRYPOINT ["./server"]
