(function() {
    'use strict';

    var log = function(str) { console.log(str); };

    app.controller('NewsPageCtrl', function($routeParams, newsPageService, $scope) {
        newsPageService.getNewsItem($routeParams.id).then(function(resp, status) {
            $scope.news_item = resp.data
            log($scope.news_item)

        })

    })

}());
