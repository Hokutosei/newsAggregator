#!/bin/bash
echo "--->> delete app.min.js"
rm public/js/app.min.js

echo "--->> compile static dependencies"
gulp

GOOS=linux GOARCH=amd64 go build -v -o linux_news_aggregator

# deploy
echo "--->> building container...."
docker build -t jeanepaul/news_aggregator_n .

echo "--->> re-tag container...."
docker tag -f jeanepaul/news_aggregator_n gcr.io/chat-app-proto01/news_aggregator_n:latest

echo "--->> pushing container"
# docker push jeanepaul/news_aggregator_n:latest
gcloud docker push gcr.io/chat-app-proto01/news_aggregator_n
#

echo "--->> stoping newsaggregator pod"
kubectl stop rc newsaggregator

echo "--->> creating newsaggregator pod"
kubectl create -f "$(pwd)"/kubernetes.yaml

# echo "--->> rolling update"
# kubectl rolling-update newsaggregator --image=gcr.io/chat-app-proto01/news_aggregator_n:latest –-update-period=5s

# echo "--->> clean unused images..."
# docker rmi "$(images | grep none | awk '{print $3}')"

echo "done! ctrl+c to stop status!"
kubectl logs -f newsaggregator
while true; do kubectl get pods; sleep 5; done
# run container
# docker run -d -e "COREOS_PRIVATE_IPV4=" jeanepaul/news_crawlers
