# ベースとなるDockerイメージ指定
FROM golang:1.17-alpine
# コンテナ内に作業ディレクトリを作成
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src
COPY ./src .

RUN apk upgrade --update && \
  apk --no-cache add git alpine-sdk
RUN go get github.com/cosmtrek/air 
RUN go get github.com/99designs/gqlgen
RUN go run github.com/99designs/gqlgen init
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest