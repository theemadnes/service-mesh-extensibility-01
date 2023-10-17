FROM cgr.dev/chainguard/go:latest as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 go build -v -o server
#RUN go build -v -o server

FROM cgr.dev/chainguard/static:latest

# Copy the binary to the production image from the builder stage.
#COPY --from=builder /app/server /app/server
COPY --from=builder /app/server .

# Run the web service on container startup.
#CMD ["/app/server"]
CMD ["/server"]