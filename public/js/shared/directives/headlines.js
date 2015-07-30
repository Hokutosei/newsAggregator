(function() {
  'use strict';

  app.directive('headlines', function(httpService) {
    return {
      restrict: 'E',
      templateUrl: 'js/shared/templates/headlines.html',
      link: function(scope) {
        httpService.fetchHeadlines(function(data) {
          scope.headlines = data;
        })
      }
    }
  })

}());
