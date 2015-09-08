(function() {
  'use strict';
  app.directive('newsTitle', function(Utils) {
    return {
      restrict: 'E',
      scope: {
        title: '@'
      },
      template: '<h2 class="title title--preview">{{ news_item_title }}</h2>',
      link: function(scope) {
        scope.news_item_title = scope.title.replace(/&#39/g, "'")
      }
    }
  })
}());
