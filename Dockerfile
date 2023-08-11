# GO Base image
FROM golang:1.20-bullseye as builder
ENV GOPRIVATE=github.com/vippsas

RUN --mount=type=bind,target=. \
  --mount=type=secret,id=GITHUB_READONLY_TOKEN  \
    git config --global url."https://vipps-mr-readonly-git:$(cat /run/secrets/GITHUB_READONLY_TOKEN)@github.com".insteadOf "https://github.com"

# Set directory
WORKDIR /app

COPY go.* ./

# Download and verify go modules
RUN go mod download && go mod verify

# Copy the source code
COPY . .

# Build the go program
RUN go build -ldflags '-s' -o api-server cmd/main.go

FROM gcr.io/distroless/base:nonroot

USER nonroot

WORKDIR /app

COPY --from=builder /app/api-server /app/api-server

# Set entry point of container
ENTRYPOINT ["/app/api-server"]

