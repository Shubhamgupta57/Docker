###########################################
# BASE IMAGE
###########################################

# Use an official Go image as the base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy your installation script into the container
COPY install.sh /install.sh

# Make the script executable
RUN chmod +x /install.sh

# Run the installation script during image build
RUN /install.sh

# Copy the Go project source code from your local directory into the container
COPY ./go /app

# Build the Go project
RUN CGO_ENABLED=0 go build -o /app/app

############################################
# HERE STARTS THE MAGIC OF MULTI STAGE BUILD
############################################

# Start a new stage with a minimal base image
FROM scratch

# Copy the compiled binary from the build stage
COPY --from=0 /app/app /app

# Set the entrypoint for the container to run the binary
ENTRYPOINT ["/app"]
