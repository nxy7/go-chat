from golang:1.21 as build

WORKDIR /app

ADD ./go.mod ./go.sum .
RUN go mod download
ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build

FROM scratch
COPY --from=build /app/go-chat /go-chat

ENTRYPOINT ["/go-chat"]


