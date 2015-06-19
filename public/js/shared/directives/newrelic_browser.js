(function() {
  'use strict';

  app.factory('newRelicBrowser', function() {
    return {
      restrict: 'E',
      templateUrl: 'js/shared/templates/newrelic_browser.html',
      link: function() {}
    }
  })

}());
