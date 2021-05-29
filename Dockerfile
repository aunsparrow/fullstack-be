FROM golang:1.15.6 AS build
COPY ./ /app/
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/go_be

FROM alpine:3.12
ENV TZ=Asia/Bangkok
COPY --from=build /bin/go_be /

EXPOSE 7000 
CMD ["/go_be"]