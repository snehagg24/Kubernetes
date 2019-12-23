FROM golang:latest as builder

WORKDIR /webapp

COPY . .

RUN go get -u github.com/go-sql-driver/mysql
RUN go build template.go
RUN chmod 777 /webapp/

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./template"]