FROM golang:alpine AS builder

WORKDIR /build

#ADD go.mod .

COPY . .

RUN go build -o gmb .

FROM alpine

WORKDIR /build

COPY --from=builder /build/gmb /build/gmb

COPY 2023.json /build
COPY family.json /build
COPY advice.txt /build

CMD ["./gmb"]