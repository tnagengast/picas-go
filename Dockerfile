from golang:latest

WORKDIR /app 

ADD . /app 

RUN go get -d -v ./... 

RUN go install -v ./... 

EXPOSE 8080

CMD ["go", "run", "src/main.go"]

