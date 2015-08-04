(function() {
    'use strict';

    var log = function(str) { console.log(str); };

    app.controller("HeaderCtrl", ["$scope", function($scope) {
        $scope.project_name = "learnJap";

    }]);

    // MainCtrl main application controller
    app.controller("MainCtrl", function(
      $scope
      , $window
      , httpService
      , $analytics
      , $location
      , $rootScope
      , userLocation
      , $routeParams
      , userSession
      , $timeout
      , adjustStageHeight
      , APP_CONFIG) {
        $analytics.pageTrack('/');
        $analytics.eventTrack('index', { category: 'index_main', label: 'index_label' });

        // hold init var for conten_type
        $scope.news_content_type = 'latest_news';
        $rootScope.page_title = 'newsInstance'
        $scope.title_limit = APP_CONFIG.title_limit;
        $rootScope.content_type = $scope.news_content_type;
        $scope.main_index_news = [];
        $scope.main_index_topranks = [];

        var services = function() {
          httpService.fetchTopRankingNews(function(data, status) {
            $scope.main_index_topranks = data;
          })
        }

        // main init func
        var init = function() {
          // log(userSession.userSessionId())

          $('.materialboxed').materialbox();
          $(".dropdown-button").dropdown();
          // $('.parallax').parallax();

          $scope.news_category_style = {
            // disable overflow for index news
            // height: adjustStageHeight.adjustHeight(),
            // 'overflow-y': 'auto'
          }
          // $('.index_right_wrap').pushpin({ top: $('.index_right_wrap').offset().top });

          if(_.has($routeParams, 'q') == true && $routeParams.q != 'registration') {
              httpService.fetchCategoryNews($routeParams.q).success(function(data) {
                $scope.main_index_news = data;
              })
          } else {
            // main news initializer in index, needs refactoring
            httpService.getNewsContent($rootScope.content_type, function(data, status) {
                $scope.main_index_news = data.reverse();
            });
          }
          services()
        }


        // get news_content from conten_type string
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

        // feed_more
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

        $scope.setIndexThumbBorder = function(news_item) {
          var color = '#fff'
          if(news_item.image.url != '') { color = '#ededed'; };
          return { 'border': '1px solid ' + color };
        }

        // right_panel_item design distinguish if item has no image.url
        $scope.right_panel_item = function(toprank) {
          if(toprank.image.url == '') {
            return 's12';
          }
          return 's9 offset-s2';
        }

        // news_tag style
        $scope.news_tag_style = function(initial) {
          var initials = {
            'y': '#ffca28',
            'w': '#ff7043',
            'b': '#81c784',
            'p': '#4db6ac',
            'e': '#9575cd',
            's': '#afb42b',
            't': '#81d4fa',
            'ir': '#80cbc4',
          }
          return {
            'background-color': initials[initial]
          }
        }

        // // AUTH BLOCK, REFACTOR THIS!;
        // var self = this;
        //
        // function handleRequest(res) {
        //   var token = res.data ? res.data.token : null;
        //   if(token) { console.log('JWT:', token); }
        //   self.message = res.data.message;
        // }
        //
        // self.login = function() {
        //   user.login(self.username, self.password)
        //     .then(handleRequest, handleRequest)
        // }
        // self.register = function() {
        //   user.register(self.username, self.password)
        //     .then(handleRequest, handleRequest)
        // }
        // self.getQuote = function() {
        //   user.getQuote()
        //     .then(handleRequest, handleRequest)
        // }
        // self.logout = function() {
        //   auth.logout && auth.logout()
        // }
        // self.isAuthed = function() {
        //   return auth.isAuthed ? auth.isAuthed() : false
        // }
        //


        // $rootScope.page_title;
        // disable getting user location
        // userLocation.getLocation()

        // call init func
        init()
    });
}());
