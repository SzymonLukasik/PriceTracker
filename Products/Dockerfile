FROM go-base:1.0.0
COPY main.go ./Products/
ADD internal/ ./Products/internal/
RUN go build -o products ./Products
EXPOSE 8083
ENTRYPOINT [ "./products" ]