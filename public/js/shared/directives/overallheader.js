(function() {
    'use strict';

    var log = function(str) { console.log(str); };

    app.directive('overallHeader', function() {
        return {
            restrict: 'E',
            templateUrl: 'js/shared/templates/overallheader.html',
            link: function(scope) {
                log("header")
            }
        }
    })

}());
