FROM golang:1.13.8 as golang1
WORKDIR src
COPY src/product.go .
COPY src/productType.go .
COPY src/main.go .
RUN go mod init go-product.com/go-product/v1
RUN rm go.mod
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/jinzhu/gorm/dialects/postgres
RUN go get -u github.com/lib/pq
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /go-product

FROM scratch
COPY --from=golang1 /go-product /go-product
ENTRYPOINT ["/go-product", "passphrase"]
