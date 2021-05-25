# pull official base image
FROM golang:alpine3.13

# set working directory
RUN mkdir -p /app
WORKDIR /app

# install app dependencies
RUN go mod init github.com/Rosalita/GoViolin
RUN go mod tidy
RUN go mod vendor
RUN go mod verify

# add app
COPY . /app

# listener port at runtime
EXPOSE 3000

# test
HEALTHCHECK --interval=1m --timeout=20s --start-period=30s --retries=3 \  
    CMD go test || exit 1

# build & run
RUN go build
ENTRYPOINT [ "./GoViolin" ]