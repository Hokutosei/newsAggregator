(function() {
  'use strict';

  app.factory('userSession', function($cookies, APP_CONFIG, Utils, httpService) {
    return {
      userSessionId: function() {
        Utils.log(APP_CONFIG.APP_NAME)
        var id = $cookies.get[APP_CONFIG.APP_NAME.toLowerCase()]
        return id;
      },

      setUserSession: function() {
        httpService.fetchUniqueSessionId(function(data) {
          Utils.log(data)
        })
        // $cookies.putObject(APP_CONFIG.APP_NAME, )
      }
    }
  });
}());
