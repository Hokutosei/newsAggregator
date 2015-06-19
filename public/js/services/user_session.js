(function() {
  'use strict';

  app.factory('userSession', function($cookies) {
    return {
      userSessionId: function() {
        $cookies.put("tests", "ktest")
        console.log($cookies)
        var id = $cookies.kedoyoId
        return id;
      }
    }
  })

}());
