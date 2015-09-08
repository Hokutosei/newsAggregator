(function() {
  'use strict';

  app.factory('newsDataAccessor', function() {
    var newsItemsArr = [];
    return {
      getNewsItems: function() {
        return newsItemsArr;
      },
      setNewsItems: function(data) {
        data = newsItemsArr;
      }
    }
  })

  app.directive('newsInit', function(Utils, newsDataAccessor) {
    return {
      restrict: 'A',
      scope: {
        data: '='
      },
      link: function(scope, element, attrs) {
        Utils.log(scope.data)
        Utils.printMe("this!!!")
      }
    }
  })

}());
