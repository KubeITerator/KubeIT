# 1. Build
FROM golang:alpine AS builder
# Install git.
RUN apk update && apk add --no-cache git
# Add certs
RUN apk --no-cache add ca-certificates

# Create appuser.
ENV USER=kubeituser
ENV UID=10001
# Statically link c-libs
ENV CGO_ENABLED=0
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR /kubeit-build/
COPY . .
# Get dependencies
RUN go get -d -v
# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /kubeit/kubeit.bin
# 2. Build small image
FROM scratch
# Copy certificates for S3
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Copy our static executable.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /kubeit/ /kubeit/
COPY --from=builder /kubeit-build/default-settings/ /kubeit/default-settings/
# Run the hello binary.
USER kubeituser:kubeituser
ENTRYPOINT ["/kubeit/kubeit.bin"]