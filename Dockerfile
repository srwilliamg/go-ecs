FROM golang:1.24 as BUILDER

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -o app ./main.go

FROM golang:1.24

COPY --from=BUILDER /app/app /app/app

EXPOSE 8080
ENTRYPOINT /app/app
