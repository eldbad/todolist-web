FROM golang:1.23.4-bookworm AS base

WORKDIR .
COPY . .
RUN go mod download
EXPOSE 8080
CMD ["make"]
