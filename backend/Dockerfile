FROM golang:1.17-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o go_auth ./main.go

CMD [ "./go_auth" ]