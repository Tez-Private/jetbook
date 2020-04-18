FROM golang:1.12.4-alpine3.9 as build

#WORKDIR go/src/github.com/Tez-Private/jetbook

COPY ["./", "./"]

ENV PATH /go/bin:$PATH

ENV GOPATH $HOME/go

#RUN apk add --update --no-cache git make gcc g++ dep  && \
#    make setup && \
#    go build -o jetbook
RUN go build -o jetbook

FROM gcr.io/distroless/base-debian10

COPY --from=build go/src/github.com/Tez-Private/jetbook/jetbook /
CMD ["jetbook"]