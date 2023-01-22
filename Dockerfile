FROM go-base:1.0.0
COPY app.go ./
ADD assets/ ./assets
ADD templates ./templates/
RUN go build -o app ./app.go
EXPOSE 8080
ENTRYPOINT [ "./app" ]
