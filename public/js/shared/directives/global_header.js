(function() {
  'use strict';

  var log = function(str) { console.log(str); };

  app.directive('globalHeader', function(APP_CONFIG) {
    return {
      restrict: 'E',
      templateUrl: 'js/shared/templates/global_header.html',
      link: function(scope) {
        scope.app_name = APP_CONFIG.APP_NAME;
      }
    }
  })
}());
