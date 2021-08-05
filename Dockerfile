### stage 1 (build) ###
FROM golang:alpine3.13 as build

# set working directory
RUN mkdir -p /app
WORKDIR /app

# install app dependencies
RUN go mod init github.com/Rosalita/GoViolin

# add app
COPY . /app

# listener port at runtime
EXPOSE 3000

# build
RUN go build -o go

# test
HEALTHCHECK --interval=1m --timeout=20s --start-period=30s --retries=3 \  
    CMD go test || exit 1

### stage 2 (run) ###
FROM alpine as production

WORKDIR /app
COPY --from=build /app .

# run
ENTRYPOINT [ "./go" ]