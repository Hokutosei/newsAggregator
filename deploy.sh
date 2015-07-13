#!/bin/bash

GOOS=linux GOARCH=amd64 go build -v -o linux_news_aggregator

# deploy
echo "--->> building container...."
docker build -t jeanepaul/news_aggregator_n .

echo "--->> pushing container"
# docker push jeanepaul/news_aggregator_n:latest
gcloud docker push gcr.io/chat-app-proto01/news_aggregator_n
#
echo "--->> stoping newscrawlers pod"
kubectl stop pod newsaggregator

echo "--->> creating newscrawlers pod"
kubectl create -f "$(pwd)"/kubernetes.yaml

echo "done!"
# run container
# docker run -d -e "COREOS_PRIVATE_IPV4=" jeanepaul/news_crawlers
