(function() {
    'use strict';

    var log = function(str) { console.log(str); };

    app.directive('overallHeader', function($rootScope, $analytics, httpService, $location, APP_CONFIG) {
        return {
            restrict: 'E',
            templateUrl: 'js/shared/templates/overallheader.html',
            controller: function($scope) {
              $scope.app_name = APP_CONFIG.APP_NAME;

              $scope.current_news_cat = $rootScope.current_news_cat_name || '';
              $rootScope.$on('update_current_news_cat', function(event, data) {
                $scope.current_news_cat = data;
              })

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

              $scope.category_get = function(category) {
                httpService.fetchCategoryNews(category).success(function(data) {
                  // emit categorized data
                  $rootScope.$emit('update_main_index_news', data)
                })
              }
            }
        }
    })

}());
