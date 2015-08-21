(function() {
  'use strict';

  var log = function(str) { console.log(str); };

  app.directive('globalHeader', function(APP_CONFIG, $rootScope) {
    return {
      restrict: 'E',
      templateUrl: 'js/shared/templates/global_header.html',
      link: function(scope) {
        scope.app_name = "test";
      }
    }
  })
}());
