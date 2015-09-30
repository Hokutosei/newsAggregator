[![Stories in Ready](https://badge.waffle.io/Hokutosei/newsAggregator.png?label=ready&title=Ready)](https://waffle.io/Hokutosei/newsAggregator)
### docker start

`docker run -d --name news_aggregator_new -v $(pwd):/go/src/web_apps/news_aggregator -e COREOS_PRIVATE_IPV4=${COREOS_PRIVATE_IPV4}  -p 49154:3000 news_aggregator`
