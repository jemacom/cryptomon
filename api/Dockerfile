FROM golang as builder

WORKDIR /go/src/github.com/jemacom/cryptomon/api

COPY ["main.go", "./"]

RUN CGO_ENABLED=0 GOOS=linux go build -o api .



FROM alpine

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/src/github.com/jemacom/cryptomon/api/api /

CMD [ "/api", "-port=8080", "-pricingServerURL=http://pricingsvc:8081/v1/pricing", "-rankingServerURL=http://rankingsvc:8082/v1/ranking", "2>/dev/null &"]

EXPOSE 8080
