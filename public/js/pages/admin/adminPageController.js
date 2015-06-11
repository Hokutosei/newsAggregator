(function() {
  'use strict';

  app.controller('AdminController', function($scope) {
    $("#menu-toggle").click(function(e) {
        e.preventDefault();
        $("#wrapper").toggleClass("toggled");
    });
  })
}());
