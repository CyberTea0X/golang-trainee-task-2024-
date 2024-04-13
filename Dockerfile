FROM alpine

WORKDIR banners
COPY gobanner .

ENV PORT=$8080
CMD ["./gobanner"]
EXPOSE 8080
