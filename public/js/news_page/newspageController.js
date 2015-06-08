(function() {
    'use strict';

    var log = function(str) { console.log(str); };

    app.controller('NewsPageCtrl', function($routeParams, newsPageService, $scope, $analytics) {
        newsPageService.getNewsItem($routeParams.id).then(function(resp, status) {
            $scope.news_item = resp.data

            $analytics.eventTrack('news_item', { category: 'news_item', label: $scope.news_item.title });

        })

        $scope.absURL = function(news_item_id, index_url) {
            var protocol = $location.protocol()
                , port = $location.port()
                , urlString = protocol + '://' + location.host + '/news/' + news_item_id

            return index_url == true ? urlString : urlString + '#disqus_thread';
        }
        $scope.decodeURL = function(url) {
            return decodeURIComponent(url)
        };
    })

}());
