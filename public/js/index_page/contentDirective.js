(function() {
  'use strict';

  app.directive('contentDisplay', function(Utils, APP_CONFIG) {
    return {
      restrict: 'E',
      template: '<div class="newsContent">{{ newsContent() }}</div>',
      scope: {
        content: '='
      },
      link: function(scope) {
        scope.content_limit = APP_CONFIG.content_limit;
        scope.newsContent = function() {
          var str = scope.content.replace(/&nbsp;/gi,'');
          return str.length > APP_CONFIG.content_limit ? str.substring(0, APP_CONFIG.content_limit) + '...' : str;
        }
      }
    }
  })
}());
