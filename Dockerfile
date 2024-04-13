FROM alpine

WORKDIR banners
COPY gobanner .

CMD ["./goauth"]
EXPOSE 8080
