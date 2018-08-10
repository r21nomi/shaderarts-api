FROM golang:1.8

# Set GOPATH/GOROOT environment variables
# RUN mkdir -p /go
# ENV GOPATH /go
# ENV PATH $GOPATH/bin:$PATH
COPY . $GOPATH/src/github.com/r21nomi/shaderarts-api
WORKDIR $GOPATH/src/github.com/r21nomi/shaderarts-api

# go get all of the dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build -o shaderarts-api

# Set up app
# ADD . /app

EXPOSE 5000

# CMD go run main.go  // This doesn't work on EB
CMD ["./shaderarts-api"]