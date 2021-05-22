# build layer
FROM  golang:1.16.3-alpine3.13 as build
WORKDIR /build
ENV CGO_ENABLED=0
ENV GOOS=linux

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags="-w -s" -o backend

# release image
FROM alpine:3.13 as release
COPY --from=build /build/backend /backend
EXPOSE 8090
ENTRYPOINT [ "/backend" ]
