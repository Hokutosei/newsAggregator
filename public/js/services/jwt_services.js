(function() {
  'use strict';

  app.factory('jwtServices', function(API, auth) {
    return {
      authInterceptor: function(config) {
        return config;
      },

      response: function(res) {
        return res;
      }
    }
  });

  app.service('authService', function($window) {
    var self = this;
  });

  app.service('userService', function($http, API, auth) {
    var self = this;
    self.getQuote = function() {
      return $http.get(API + '/auth/qoute')
    }
  })

})
