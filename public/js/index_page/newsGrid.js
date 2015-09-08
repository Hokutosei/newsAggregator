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
      }
    }
  })
}());
