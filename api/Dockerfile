FROM golang:latest AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY ./ ./
# roda o gerador dos docs tbm
RUN swag init --parseDependency
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-api github.com/Leinadium/puc-disciplinas/api

FROM gcr.io/distroless/base-debian11 AS release
WORKDIR /
COPY --from=build /docker-api /docker-api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/docker-api"]
