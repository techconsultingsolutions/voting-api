FROM golang as runtime

WORKDIR /app

COPY . ./

CMD [ "go build ." ]

ENTRYPOINT ["./voting-api"]