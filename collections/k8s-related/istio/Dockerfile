FROM hub.whateveradmin.com/prj-uhost-ussg/golang:1.16  AS builder
WORKDIR /root
COPY apigateway.go .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o istio-gateway apigateway.go


FROM hub.whateveradmin.com/public/centos:7
WORKDIR /root/
COPY --from=builder /root/istio-gateway .
EXPOSE 12345
ENTRYPOINT [ "/root/istio-gateway" ]