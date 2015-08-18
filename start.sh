#!/bin/bash

# rm public/js/app.min.js

rm public/js/app.min.js

gulp

go build -v && ./news_aggregator
