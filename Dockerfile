FROM golang:1.12.4-alpine3.9 as build

RUN mkdir -p /go/src/github.com/Tez-Private/jetbook
WORKDIR /go/src/github.com/Tez-Private/jetbook

ADD . .

ENV PATH /go/bin:$PATH

RUN apk add --update --no-cache git make gcc g++ dep  && \
    make setup && \
    go build .

FROM gcr.io/distroless/base-debian10

COPY --from=build /go/src/github.com/Tez-Private/jetbook/jetbook .
ENTRYPOINT ["./jetbook"]