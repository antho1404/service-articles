FROM golang:1.10.4
WORKDIR /go/src/github.com/ilgooz/service-articles
COPY . .
RUN go install ./...
CMD articles --mongoAddr mongodb://mongo:27017 --dbName articles