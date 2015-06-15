(function() {
  'use strict';

  var log = function(str) { console.log(str) };

  app.directive('categoryMenu', function(httpService, $rootScope) {
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
          $rootScope.$emit('empty_main_index_news')
          $rootScope.$emit('update_current_news_cat', category)
          $rootScope.current_news_cat_name = category;
          httpService.fetchCategoryNews(category.initial).success(function(data) {
            // emit categorized data
            $rootScope.$emit('update_main_index_news', data)
          })
        }
      }
    }
  })
}());
