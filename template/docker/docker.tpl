FROM 8.134.142.155:8888/ops/golang:1.23-builder AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
{{if .Chinese}}ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
{{end}}
WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN for i in $(seq 1 3); do go mod download && break || sleep 5; done
COPY . .
RUN go build -ldflags="-s -w" -o /app/{{.ExeFile}} {{.GoMainFrom}}


FROM registry.cn-guangzhou.aliyuncs.com/dbinggo-docker/alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
{{if .HasTimezone}}COPY --from=builder /usr/share/zoneinfo/{{.Timezone}} /usr/share/zoneinfo/{{.Timezone}}
ENV TZ {{.Timezone}}
{{end}}
WORKDIR /app
COPY --from=builder /app/{{.ExeFile}} /app/{{.ExeFile}}
{{if .HasPort}}
EXPOSE {{.Port}}
{{end}}
CMD ["./{{.ExeFile}}"{{.Argument}}]
