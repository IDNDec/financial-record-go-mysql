FROM golang:1.25.5-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app .
EXPOSE 8000
CMD ["./app"]

# # --- Stage 1: Build the Go app ---
# FROM golang:1.24-alpine AS builder

# WORKDIR /app

# # Copy go.mod & go.sum first (better caching)
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the rest of the source code
# COPY . .

# # Build the Go application
# RUN go build -o app .

# --- Stage 2: Run the built binary ---
# FROM alpine

# WORKDIR /app

# # Install timezone (optional)
# RUN apk add --no-cache tzdata

# # Copy config file & built app from builder stage
# COPY --from=builder /app/app .
# COPY --from=builder /app/app.conf.json .

# # Expose the port used by your Go server
# EXPOSE 8000

# # Start the application
# CMD ["./app"]
