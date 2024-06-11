FROM golang:1.22.3

#Set destination for COPY
WORKDIR /app

#Copy go.mod and execute to download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

#Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /sederhana

#To bind to a TCP port, runtime parameters must be supplied to the docker command
EXPOSE 80

#Run
CMD ["/sederhana"]