FROM golang:alpine as build 
RUN apk add --no-cache ca-certificates git openssh-client

ARG SSH_PRIVATE_KEY

WORKDIR /go/src/github.com/RapidCodeLab/fakedsp
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags '-extldflags "-static"' -o server ./cmd/server

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt \
     /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/src/github.com/RapidCodeLab/fakedsp/server /server


ENTRYPOINT ["/server"]