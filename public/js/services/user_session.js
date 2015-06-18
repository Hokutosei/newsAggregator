(function() {
  'use strict';

  app.factory('userSession', function($cookies) {
    return {
      userSessionId: function() {
        return $cookies.getObject('kedoyoId')
      }
    }
  })

}());
