FROM docker.io/library/golang:1.12-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux
WORKDIR /src
RUN apk add --no-cache git
COPY . .
RUN go build -a -o /server ./cmd/server

FROM scratch

ENV PORT=8000
CMD [ "/server" ]
ENTRYPOINT [ "/server" ]
COPY --from=build /server /server
