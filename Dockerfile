## Run
FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add --no-cache git mercurial \
    && go get github.com/fachryansyah/test-golang-solusi-menghitung-huruf \
    && apk del git mercurial 
CMD ["go", "run", "/app/server.go"]
