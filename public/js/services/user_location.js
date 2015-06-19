(function() {
  'use strict';

  app.factory('userLocation', function() {

    return {
      getLocation: function() {
        console.log(navigator.geolocation)
        if (navigator.geolocation) {
            navigator.geolocation.getCurrentPosition(function(position){
              console.log(position)
              // $scope.$apply(function(){
              //   $scope.position = position;
              // });
            });
          }
      }
    }
  })
}());
