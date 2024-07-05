FROM golang:1.22.4-alpine AS build_base

RUN apk --no-cache add bash git make gcc gettext musl-dev

WORKDIR /usr/local/src

COPY ["./go.mod", "./go.sum", "./"]
RUN go mod download


COPY . ./
RUN go build -o ./bin/app cmd/weatherapp/main.go


FROM alpine as runner

COPY --from=build_base /usr/local/src/bin/app /
COPY ./.env /
COPY config/local.yaml /config/local.yaml
COPY ./migrations /migrations

CMD ["/app"]