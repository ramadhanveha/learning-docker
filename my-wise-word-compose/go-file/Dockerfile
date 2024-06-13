FROM golang:1.21

#Set destination for COPY
WORKDIR /myapp

RUN go version
#Copy go.mod and execute to download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

COPY AUTHORS.md LINKS.md ./
#Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /my-go-app

#To bind to a TCP port, runtime parameters must be supplied to the docker command
EXPOSE 78

#Run
CMD ["/my-go-app"]