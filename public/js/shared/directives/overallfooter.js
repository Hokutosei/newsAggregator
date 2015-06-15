(function() {
  'use strict';

  app.directive('overallFooter', function($analytics) {
    return {
      restrict: 'E',
      templateUrl: 'js/shared/templates/overallfooter.html',
      controller: function($scope) {
        $scope.luxiar_co = function() {
          // insert google analytics
          $analytics.eventTrack('luxiar_homepage', { category: 'luxiar_co', label: 'footer link' });
          window.open('http://www.luxiar.com/', '_blank')
        }
      }
    }
  })
}());
