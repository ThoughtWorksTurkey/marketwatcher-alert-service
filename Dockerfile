FROM debian:jessie

RUN mkdir -p /opt/alert-service/conf

COPY marketwatcher-alert-service /opt/alert-service/service
COPY conf/ /opt/alert-service/conf/
COPY db/ /data
COPY scripts/wait-for-it.sh /usr/local/bin/wait-for-it.sh
COPY scripts/start.sh /usr/local/bin/start.sh

ENTRYPOINT ["/opt/alert-service/service"]
