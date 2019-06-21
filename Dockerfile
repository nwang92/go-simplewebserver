FROM golang:1.11

WORKDIR /simple-webserver/
COPY main.go /simple-webserver/
RUN go build && ls

EXPOSE 8001
ENTRYPOINT [ "/simple-webserver/simple-webserver" ]
