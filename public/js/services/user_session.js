(function() {
  'use strict';

  app.factory('userSession', function($cookies) {
    return {
      userSessionId: function() {
        var id = $cookies.get['kedoyoId']
        return id;
      }
    }
  })

}());
