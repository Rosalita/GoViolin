FROM golang:1.18

RUN mkdir /web_go

# Set the Current Working Directory inside the container
WORKDIR /web_go

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . /web_go

# Download all the dependencies
RUN go get -d -v ./...
RUN go mod init github.com/Rosalita/GoViolin
RUN go mod tidy
RUN go mod vendor
RUN go mod verify

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 6060

# Run the executable
CMD ["./GoViolin"]

