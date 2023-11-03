###########################################
# BASE IMAGE
###########################################

FROM ubuntu:latest

# Copy your installation script into the container
COPY install.sh /install.sh

# Make the script executable
RUN chmod +x /install.sh

# Run the installation script during image build
RUN /install.sh

ENV GO111MODULE=off

COPY . .

RUN CGO_ENABLED=0 go build -o /app .

############################################
# HERE STARTS THE MAGIC OF MULTI STAGE BUILD
############################################

FROM scratch

# Copy the compiled binary from the build stage
COPY --from=build /app /app

# Set the entrypoint for the container to run the binary
ENTRYPOINT ["/app"] 

#next