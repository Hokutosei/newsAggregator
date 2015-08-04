(function() {
  'use strict';

  app.factory('jwtServices', function(api, authService) {
    return {
      authInterceptor: function(config) {
        return config;
      },

      response: function(res) {
        return res;
      }
    }
  });
}());


(function() {
  'use strict';
  app.service('authService', function($window) {
    var self = this;
  });


}());
//
(function() {
  'use strict';

  app.service('userService', function($http, authService) {
    var self = this;
    self.getQuote = function(api) {
      return $http.get(api + '/auth/qoute')
    }
  })

}());
