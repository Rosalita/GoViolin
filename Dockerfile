FROM golang:1.18-alpine3.15 as builder

WORKDIR /violin-app

COPY . .

ENV GO111MODULE=on

RUN go mod init \
    && go build -o violin-app 


FROM alpine:3.15.4 as runner

RUN addgroup -S appgroup && adduser -S appuser -G appgroup 

ENV PORT ${PORT:+8080}

WORKDIR /violin-app/

COPY ./templates "./templates"

COPY ./css "./css"

COPY ./img "./img"

COPY ./mp3 "./mp3"

COPY --from=builder /violin-app/violin-app ./

EXPOSE "$PORT"

USER appuser

CMD ["./violin-app"]