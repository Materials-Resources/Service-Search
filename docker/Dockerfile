FROM golang:1.19.5-alpine as build

WORKDIR /usr/src/app/build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/app ./server


FROM alpine as production
EXPOSE 50057
COPY --from=build /usr/local/bin/app /usr/local/bin/app

CMD ["app"]