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
# RUN ls
# Copy the code into the container
ADD cmd/ cmd/
ADD database/ database/
ADD env.yaml env.yaml
# Build the application
RUN go build -o main cmd/api/v1/address/main.go

FROM alpine:latest
# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist/
# Copy binary from build to dist folder
COPY --from=builder /build .
# Export necessary port
EXPOSE 3000
# Run
CMD ["./main"]