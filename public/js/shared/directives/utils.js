(function() {
  'use strict';

  var log = function(str) { console.log(str) }

  var stage_height = 500;

  app.directive('imageLoader', function() {
    return {
      restrict: 'E',
      scope: {
        data: '='
      },
      templateUrl: 'js/shared/templates/newsitem_index_thumb.html',
      link: function(scope, elem, attrs) {
        scope.news_item = scope.data
      }
    }
  })

  app.factory('adjustStageHeight', function($document) {
    return {
      adjustHeight: function() {
        return ($document.height() + stage_height) + 'px'
      }
    }
  })

}());
