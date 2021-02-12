FROM golang:1.15.8-alpine as builder
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
# Move to working directory /build
WORKDIR /build
# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download
# Copy the code into the container
COPY . .
# Build the application
RUN go build -o main .

FROM alpine:latest
# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist/
# Copy binary from build to dist folder
COPY --from=builder /build .
# Export necessary port
EXPOSE 3000
# Run
CMD ["./main"]