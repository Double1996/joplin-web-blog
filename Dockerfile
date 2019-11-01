FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY ./bin/server_exec ./bin/server_exec
COPY ./templates  ./templates
COPY ./static  ./staic
RUN chmod +x ./bin/server_exec
CMD ["./bin/server_exec"]