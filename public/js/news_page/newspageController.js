(function() {
    'use strict';

    var log = function(str) { console.log(str); };

    app.controller('NewsPageCtrl', function($routeParams, newsPageService, $scope, $analytics, $location) {
        newsPageService.getNewsItem($routeParams.id).then(function(resp, status) {
            $scope.news_item = resp.data

            $analytics.eventTrack('news_item', { category: 'news_item', label: $scope.news_item.title });

        })

        $scope.absURL = function(news_item, index_url) {
          var protocol = $location.protocol()
              , port = $location.port()
              , urlString = protocol + '://' + location.host + '/news/' + news_item

          return index_url == true ? news_item : decodeURIComponent(news_item.url);
        }
        $scope.decodeURL = function(url) {
            return decodeURIComponent(url)
        };
        $scope.news_tag_style = function(initial) {
          var initials = {
            'y': '#ffca28',
            'w': '#ff7043',
            'b': '#81c784',
            'p': '#4db6ac',
            'e': '#9575cd',
            's': '#bcaaa4',
            't': '#81d4fa',
            'ir': '#80cbc4',
          }
          return {
            'background-color': initials[initial]
          }
        }

    })

}());
