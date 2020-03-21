FROM golang:1.13 as builder

COPY . /app
WORKDIR /app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o monit4g .


FROM alpine

COPY --from=builder /app/monit4g /app/monit4g
WORKDIR /app

CMD "./monit4g"
