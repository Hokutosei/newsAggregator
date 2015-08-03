(function() {
  'use strict';

  app.directive('headlines', function(httpService, $location, strServices) {
    return {
      restrict: 'E',
      templateUrl: 'js/shared/templates/headlines.html',
      link: function(scope) {
        httpService.fetchHeadlines(function(data) {
          scope.headlines = data;
        })
        scope.truncateTitle = function(str) {
          return strServices.truncateStr(str, 35);
        }
        scope.absURL = function(news_item_id, index_url) {
            var protocol = $location.protocol()
                , port = $location.port()
                , urlString = protocol + '://' + location.host + '/news/' + news_item_id

            return index_url == true ? urlString : urlString + '#disqus_thread';
        }

      }
    }
  })

}());
