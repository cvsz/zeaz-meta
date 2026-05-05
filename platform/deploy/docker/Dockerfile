# syntax=docker/dockerfile:1.7
FROM golang:1.22-alpine AS build
ARG SERVICE
WORKDIR /src
COPY platform/go.mod platform/go.sum* ./
RUN go mod download
COPY platform/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags='-s -w' -o /out/service ./cmd/${SERVICE}

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/service /service
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/service"]
