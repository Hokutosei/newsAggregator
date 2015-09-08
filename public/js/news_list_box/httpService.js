(function() {
  'use strict';

  app.factory('newsListHttp', function($http, Utils) {
    return {
        getApi: function(url, params) {
          return $http.get(url, params)
        }
    }
  });
}());
