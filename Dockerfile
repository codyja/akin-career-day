FROM arm64v8/golang:1.20

WORKDIR /app

COPY go.mod ./
COPY static ./static
COPY templates ./templates
COPY *.go ./

RUN go build -o /akin-app

CMD [ "/akin-app" ]