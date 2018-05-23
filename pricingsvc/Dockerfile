FROM golang as builder
WORKDIR /go/src/github.com/jemacom/cryptomon/pricingsvc

RUN go get -v github.com/coincircle/go-coinmarketcap

COPY ["cmc.go", "main.go", "./"]

RUN CGO_ENABLED=0 GOOS=linux go build -o pricingsvc .



FROM alpine

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/src/github.com/jemacom/cryptomon/pricingsvc/pricingsvc /

CMD [ "/pricingsvc", "-port=8081", "2>/dev/null &"]

EXPOSE 8080
