FROM docker.io/library/golang:1.16-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux
WORKDIR /src

RUN apk add --no-cache git
COPY . .
RUN go build -a -o /server ./cmd/server

FROM busybox:latest

ENV PORT=8000
EXPOSE $PORT
CMD [ "/server" ]

COPY --from=build /server /server
USER 10000
