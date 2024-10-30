echo "FROM golang:1.23.2" >> Dockerfile
echo "WORKDIR /app" >> Dockerfile
echo "COPY go.mod ./" >> Dockerfile
echo "COPY go.sum ./" >> Dockerfile
echo "RUN go mod download" >> Dockerfile
echo "COPY *go ./" >> Dockerfile
echo "RUN go build" >> Dockerfile
echo "EXPOSE 5050" >> Dockerfile
echo 'CMD ["/app/RestApiPract1"]' >> Dockerfile
docker build . -t dockerapigo
docker run -d -p 5050:5050 dockerapigo
docker ps -a