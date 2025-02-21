FROM golang:1.14-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql, apt-utils, curl
RUN apt-get update
RUN apt-get -y install postgresql-client
RUN apt-get -y install apt-utils

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh
# make init-letsencrypt.sh
RUN chmod +x init-letsencrypt.sh

# build go app
RUN go mod download
RUN go build -o web-app ./cmd/main.go

CMD ["./web-app"]