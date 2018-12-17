FROM alpine:latest

RUN apk add --no-cache curl
COPY client /usr/bin/client
RUN chmod +x /usr/bin/client
CMD /usr/bin/client
