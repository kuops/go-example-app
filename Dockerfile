FROM centos:7

COPY build/app /app
CMD ["/app"]
