(function() {
    'use strict';

    var log = function(str) { console.log(str); };

    app.controller('NewsPageCtrl', function($routeParams, newsPageService, $scope, $analytics) {
        newsPageService.getNewsItem($routeParams.id).then(function(resp, status) {
            $scope.news_item = resp.data

            $analytics.eventTrack('news_item', { category: 'news_item', label: $scope.news_item.title });

        })

    })

}());
