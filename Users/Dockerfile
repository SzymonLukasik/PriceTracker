FROM go-base:1.0.0
COPY main.go ./Users/
ADD internal/ ./Users/internal/
RUN go build -o users ./Users
EXPOSE 8081
ENTRYPOINT [ "./users" ]