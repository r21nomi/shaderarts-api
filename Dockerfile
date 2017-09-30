FROM golang:1.8

# Set GOPATH/GOROOT environment variables
# RUN mkdir -p /go
# ENV GOPATH /go
# ENV PATH $GOPATH/bin:$PATH
COPY . $GOPATH/src/github.com/r21nomi/arto-api
WORKDIR $GOPATH/src/github.com/r21nomi/arto-api

# go get all of the dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# Set up app
# ADD . /app

EXPOSE 80

# CMD ["/arto-api"]
CMD go run main.go