FROM golang:1.22
WORKDIR /app/

COPY /src .

RUN go get
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]