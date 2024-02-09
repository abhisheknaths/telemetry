FROM golang:1.22-alpine AS BUILDER

RUN /sbin/apk update && /sbin/apk --no-cache add ca-certificates \
git tzdata && /usr/sbin/update-ca-certificates

RUN adduser -D -g '' builduser

WORKDIR /home/telemetry

COPY ./app1/main.go /home/telemetry/main.go
COPY ./go.mod /home/telemetry/go.mod
COPY ./go.sum /home/telemetry/go.sum
COPY ./internal /home/telemetry/internal
COPY ./handler /home/telemetry/handler
COPY ./router /home/telemetry/router

ARG VERSION

RUN CGO_ENABLED=0 go build -a -tags netgo,osusergo \
--ldflags "-extldflags '-static' -s -w" \
--ldflags "-X main.version=$VERSION" -o telemetry .

FROM busybox:musl

COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=BUILDER /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=BUILDER /etc/passwd /etc/passwd
COPY --from=BUILDER /home/telemetry/telemetry /home/telemetry

USER builduser
WORKDIR /home
EXPOSE 3000

ENTRYPOINT ["/home/telemetry"]