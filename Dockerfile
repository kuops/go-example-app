FROM centos:7

WORKDIR /app
COPY build/app /app/server
COPY build/templates /app/templates
CMD ["/app/server"]
