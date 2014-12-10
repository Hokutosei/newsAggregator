'use strict';

var app = angular.module("newsAggregator", ["ngTouch"]);

var log = function(str) { console.log(str) };

app.controller("HeaderCtrl", ["$scope", function($scope) {
    $scope.project_name = "learnJap";

}]);


app.controller("MainCtrl", ["$scope", "$window", "httpService", function($scope, $window, httpService) {
    httpService.getIndexNews(function(data, status) {
        $scope.main_index_news = data;
    });

    $scope.news_content_type = 'latest_news';
    $scope.news_content = function(content_type) {
    	httpService.getNewsContent(content_type, function(data, status) {
            $scope.news_content_type = content_type;
    		$scope.main_index_news = data;
    	})
    };

    $scope.feed_more = function(length) {
        log(length);
        httpService.feedMoreNews($scope.news_content_type, length, function(data, status) {
            log(data)
            for(var i = 0; i < data.length; i++) {
                $scope.main_index_news.push(data[i])
            }

        })
    };

    $scope.timeToLocal = function(unix_time) {
    	return new Date(unix_time * 1000)
    }
}]);