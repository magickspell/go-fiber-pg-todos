# build
FROM golang:1.24

WORKDIR /go-app-dir

COPY . .

RUN apt-get update && apt-get install -y bash

CMD ["sh", "run.sh"]