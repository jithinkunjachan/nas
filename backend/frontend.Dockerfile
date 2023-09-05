# syntax=docker/dockerfile:1
FROM ubuntu:latest

# Set destination for COPY
WORKDIR /app


# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY cmd/api/main /app

EXPOSE 8081

# Run
CMD ["/app/main"]