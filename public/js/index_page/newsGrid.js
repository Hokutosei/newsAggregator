(function() {
  'use strict';

  app.directive('newsGrid', function(Utils) {
    return {
      restrict: 'E',
      scope: {
        data: '=',
        mainIndexNews: '='
      },
      templateUrl: 'js/index_page/templates/newsGrid.html',
      link: function(scope) {
        Utils.log(scope.mainIndexNews)
        scope.news_tag_style = function(object) {
          return scope.$parent.news_tag_style(object)
        }

        scope.absURL = function(news_item_id, index_url) {
          Utils.log(news_item_id, index_url)
          return scope.$parent.absURL(news_item_id, index_url)
        }
      }
    }
  })
}());
