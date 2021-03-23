### Build binary from official Go image
FROM golang:stretch as build
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /whats-api .

### Put the binary onto Heroku image
FROM heroku/heroku:18
COPY --from=build /whats-api /whats-api
CMD ["/whats-api"]