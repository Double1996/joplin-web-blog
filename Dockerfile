FROM golang:1.12
WORKDIR /app
COPY ./bin/server_exec ./bin/server_exec
COPY ./templates  ./templates
COPY ./static  ./static
RUN chmod +x ./bin/server_exec
#CMD ["./bin/server_exec"]
