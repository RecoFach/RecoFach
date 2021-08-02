FROM golang:1.15 AS build

ENV GO111MODULE=on
# Set the Current Working Directory inside the container
WORKDIR /go/src/relatDB/db_initialiser

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Install the package
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o init_mariadb *.go
# This container exposes port 8080 to the outside world

FROM scratch
WORKDIR /go/src/relatDB/db_initialiser
COPY --from=build /go/src/relatDB/db_initialiser/init_mariadb /go/src/relatDB/db_initialiser
COPY --from=build /go/src/relatDB/db_initialiser/head.csv /go/src/relatDB/db_initialiser
#EXPOSE 3306
# Run the executable
CMD ["./init_mariadb"]