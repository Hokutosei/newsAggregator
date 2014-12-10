'use strict';

var app = angular.module("newsAggregator", ["ngTouch"]);

var log = function(str) { console.log(str) };

app.controller("HeaderCtrl", ["$scope", function($scope) {
    $scope.project_name = "learnJap";

}]);


app.controller("MainCtrl", ["$scope", "$window", "httpService", function($scope, $window, httpService) {
    httpService.getIndexNews(function(data, status) {
        log(data)
        $scope.main_index_news = data
    })

    $scope.timeToLocal = function(unix_time) {
    	return new Date(unix_time * 1000)
    }
}]);