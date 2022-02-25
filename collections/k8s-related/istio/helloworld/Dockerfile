FROM hub.whateveradmin.com/prj-uhost-ussg/golang:1.16  AS builder
WORKDIR /root
COPY helloworld.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o helloworld helloworld.go


FROM hub.whateveradmin.com/public/centos:7
WORKDIR /root/
COPY --from=builder /root/helloworld .
EXPOSE 12345
ENTRYPOINT [ "/root/helloworld" ]