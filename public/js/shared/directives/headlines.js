(function() {
  'use strict';

  app.directive('headlines', function(httpService, $location) {
    return {
      restrict: 'E',
      templateUrl: 'js/shared/templates/headlines.html',
      link: function(scope) {
        httpService.fetchHeadlines(function(data) {
          scope.headlines = data;
        })
        console.log($location.path())
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
