# Use a more recent version of Go
FROM golang:1.23.4-alpine3.20 AS build

WORKDIR /app
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s -extldflags '-static'" -o /app/secret main.go

# Use scratch as the base image
FROM alpine:3.20.3

##  install ca-certificates
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/secret /app/secret

# Copy db/migration folder
COPY --from=build /app/db/migrations /app/db/migrations
COPY --from=build /app/views /app/views

ARG PORT
EXPOSE ${PORT:-8080}

RUN apk add --no-cache curl

## add user and group
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

## change ownership of the app directory to the appuser
RUN chown -R appuser:appgroup /app

## change the user to appuser
USER appuser

ENTRYPOINT ["/app/secret"]