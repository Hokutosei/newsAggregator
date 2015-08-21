(function() {
  'use strict';

  app.factory('strServices', function() {
    return {
      truncateStr: function(str, limit) {
        if(str.length > limit) {
          return str.substring(0, limit) + '...';
        }
        return str;
      }
    }
  })

}());
