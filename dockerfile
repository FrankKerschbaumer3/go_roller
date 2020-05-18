FROM golang:alpine AS builder
LABEL stage=builder
WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a

FROM alpine AS final
WORKDIR /
COPY --from=builder /workspace/go_roller .
CMD [ "./go_roller" ]