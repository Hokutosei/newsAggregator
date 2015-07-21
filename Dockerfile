FROM ubuntu:wily

# USAGE
# docker run -d -p 3001:3000 -e "COREOS_PRIVATE_IPV4=" jeanepaul/news_aggregator_n
# docker run --rm -it --entrypoint=/bin/bash -p 3001:3000 -e "COREOS_PRIVATE_IPV4=" jeanepaul/news_aggregator_n

MAINTAINER jeanepaul@gmail.com

RUN apt-get install -y ca-certificates

RUN mkdir -p /web_apps/news_aggregator/public/

COPY linux_news_aggregator /web_apps/news_aggregator/

COPY public /web_apps/news_aggregator/public/

WORKDIR /web_apps/news_aggregator/

ENTRYPOINT ./linux_news_aggregator
