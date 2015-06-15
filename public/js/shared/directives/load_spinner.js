(function() {
  'use strict';

  // this directive accepts scope.data[] that
  // bools if data is empty or not to display spinner
  app.directive('loadSpinner', function() {
    return {
      restrict: 'E',
      templateUrl: 'js/shared/templates/load_spinner.html',
      scope: {
        data: '='
      },
      link: function(scope) {
      }
    }
  })

}());
