FROM golang:1.20-alpine 
WORKDIR /horikita
COPY . .
RUN go mod download
RUN go build -o /horikita
CMD ./horikita