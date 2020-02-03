# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.13-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="j3ns789@gmail.com"

WORKDIR /app

ENV GOPATH="/app/vendor"
ENV PATH="${GOPATH}/bin:${PATH}"
