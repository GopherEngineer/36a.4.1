FROM node:20.15.1-alpine AS web
WORKDIR /web
ADD web .
RUN yarn install --frozen-lockfile
RUN yarn build

FROM golang:alpine AS compiling
RUN apk --no-cache add ca-certificates
RUN mkdir -p /go/src/aggregator
WORKDIR /go/src/aggregator
ADD cmd cmd
ADD pkg pkg
ADD go.mod .
ADD go.sum .
RUN CGO_ENABLED=0 go install cmd/aggregator/aggregator.go
 
FROM scratch
LABEL version="1.0.0"
LABEL maintainer="Gopher Engineer<gopher@skillfactory>"
WORKDIR /root/

COPY --from=compiling /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=web /web/public web/public
COPY --from=compiling /go/bin/aggregator .

ADD config.json .

EXPOSE 80

CMD [ "./aggregator" ]