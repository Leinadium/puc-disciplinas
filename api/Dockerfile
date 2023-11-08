FROM golang:latest AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-api

FROM grc.io/distroless/base-debian11 AS release
WORKDIR /
COPY --from=build /docker-api /docker-api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/docker-api"]
