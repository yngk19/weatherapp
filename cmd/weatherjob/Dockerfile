FROM golang:1.22.4-alpine AS build_base

RUN apk --no-cache add bash git make gcc gettext musl-dev

WORKDIR /usr/local/src

COPY ["./go.mod", "./go.sum", "./"]
RUN go mod download


COPY . ./
RUN go build -o ./bin/job cmd/weatherjob/main.go


FROM alpine as runner

COPY --from=build_base /usr/local/src/bin/job /
COPY ./.env /
COPY config/local.yaml /config/local.yaml

CMD ["/job"]