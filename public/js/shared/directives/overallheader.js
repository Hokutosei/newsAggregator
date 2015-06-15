(function() {
    'use strict';

    var log = function(str) { console.log(str); };

    app.directive('overallHeader', function($rootScope, $analytics, httpService, $location) {
        return {
            restrict: 'E',
            templateUrl: 'js/shared/templates/overallheader.html',
            controller: function($scope) {
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
