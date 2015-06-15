(function() {
    'use strict';

    var log = function(str) { console.log(str); };

    app.controller("HeaderCtrl", ["$scope", function($scope) {
        $scope.project_name = "learnJap";

    }]);


    app.controller("MainCtrl", ["$scope", "$window", "httpService", '$analytics', '$location', '$rootScope',
        function($scope, $window, httpService, $analytics, $location, $rootScope) {
        $analytics.pageTrack('/');
        //$analytics.eventTrack('index');
        $analytics.eventTrack('index', { category: 'index_main', label: 'index_label' });
        //$analytics.eventTrack('news_item_clicked', { category: 'news_clicks', label: 'news_item_clicked' })

        // hold init var for conten_type
        $scope.news_content_type = 'latest_news';
        $rootScope.content_type = $scope.news_content_type;
        $scope.main_index_news = [];


        httpService.getNewsContent($rootScope.content_type, function(data, status) {
            $scope.main_index_news = data;
        });

        $scope.news_content = function(content_type) {
            $rootScope.content_type = content_type;
            $analytics.eventTrack('index', { category: 'index_main', label: content_type });

            if($location.path() != "/") {
                window.location.href = "/";
            } else {
                httpService.getNewsContent($rootScope.content_type, function(data, status) {
                    $scope.news_content_type = $rootScope.content_type;
            		$scope.main_index_news = data;
            	})
            }
        };


        $scope.feed_more = function(length) {
            $analytics.eventTrack('feed_more', { category: 'system_func', label: 'feed_more_data' });
            httpService.feedMoreNews($scope.news_content_type, length, function(data, status) {
                for(var i = 0; i < data.length; i++) {
                    $scope.main_index_news.push(data[i])
                }
            })
        };

        //update_main_index_news
        $rootScope.$on("update_main_index_news", function(event, data) {
          $scope.main_index_news = data
        })

        $rootScope.$on('empty_main_index_news', function() {
          $scope.main_index_news = []
        })

        $scope.ga_event = function(news_item) {
            httpService.incrementNewsItemScore(news_item);
            $analytics.eventTrack('news_item_' + news_item.title , { category: 'news_clicks', label: 'news_item_clicked' })
        };

        $scope.decodeURL = function(url) {
            return decodeURIComponent(url)
        };

        $scope.timeToLocal = function(unix_time) {
        	return new Date(unix_time * 1000)
        };

        $scope.absURL = function(news_item_id, index_url) {
            var protocol = $location.protocol()
                , port = $location.port()
                , urlString = protocol + '://' + location.host + '/news/' + news_item_id

            return index_url == true ? urlString : urlString + '#disqus_thread';
        }
    }]);
}());
