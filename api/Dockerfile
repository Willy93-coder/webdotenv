FROM golang:1.24.5-alpine3.22 AS build

WORKDIR /src

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /out/app ./main.go

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /out/app /app

EXPOSE 5000

ENTRYPOINT ["/app"]
