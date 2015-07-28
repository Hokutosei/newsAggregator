(function() {
  'use strict';

  var log = function(str) { console.log(str) }

  function twoDigitFmt(int) {
    if((int + "").length > 1) return int;
    return ("0" + int).slice(-2);
  }

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

  app.directive('popularIndexThumb', function() {
    return {
      restrict: 'E',
      scope: {
        data: '='
      },
      templateUrl: 'js/shared/templates/popular_index_thumb.html',
      link: function(scope) {
        scope.item = scope.data;
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

  app.directive('dateFormatDayMonth', function() {
    return {
      restrict: 'E',
      scope: {
        data: '='
      },
      template: '{{ dateStr }}',
      link: function(scope) {
        var date = new Date(scope.data.created_at)
            , day = date.getDate()
            , month = date.getMonth() + 1
            , hours = twoDigitFmt(date.getHours())
            , minutes = twoDigitFmt(date.getMinutes())
        scope.dateStr = day + '日 ' + month + '月' + hours + ':' + minutes
      }
    }
  })

}());
