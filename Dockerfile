FROM golang:latest
WORKDIR /app
COPY . .
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh
RUN go mod download
RUN go build -o /godocker
EXPOSE 8080
CMD [ "/godocker" ]
