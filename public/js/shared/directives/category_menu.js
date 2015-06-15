(function() {
  'use strict';

  var log = function(str) { console.log(str) };

  app.directive('categoryMenu', function(httpService, $rootScope, $location) {
    return {
      restrict: 'E',
      templateUrl: 'js/shared/templates/category_menu.html',
      controller: function($scope) {
        $(".button-collapse").sideNav({
          closeOnClick: true
        });

        $scope.category_lists = []
        $scope.categories = function() {
          httpService.fetchHeaderCategories(function(data) {
            $scope.category_lists = _.map(data, function(val) {
              return val
            })
          })
        }

        $scope.category_get = function(category) {
          var location = $location.path();

          // redirect if request is not from index
          if(location.match('/(.*)$')[1] == '' || location.match('/news')) {
            window.location.href = '/' + category.initial;
            return false;
          }

          $rootScope.$emit('empty_main_index_news')
          $rootScope.$emit('update_current_news_cat', category)
          $rootScope.current_news_cat_name = category;
          httpService.fetchCategoryNews(category.initial).success(function(data) {
            // emit categorized data
            $rootScope.$emit('update_main_index_news', data)
            $location.path('/' + category.initial)
          })
        }
      }
    }
  })
}());
