# Stage 1: Build the Nuxt frontend
FROM node:22-alpine AS frontend-build

WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm ci

COPY frontend/ ./
RUN npm run generate

# Stage 2: Build the Go binary (with embedded frontend)
FROM golang:1.23.4-alpine3.21 AS go-build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY --from=frontend-build /app/frontend/dist ./frontend/dist

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags "-w -s -extldflags '-static'" -o /app/ghost-send main.go

# Stage 3: Minimal runtime image
FROM alpine:3.21.2

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=go-build /app/ghost-send /app/ghost-send
COPY --from=go-build /app/db/migrations /app/db/migrations

RUN addgroup -S appgroup && adduser -S appuser -G appgroup && \
    chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080
ENV PORT=8080

ENTRYPOINT ["/app/ghost-send"]
