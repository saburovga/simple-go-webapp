FROM golang
ADD app /
ENTRYPOINT ["/app"]