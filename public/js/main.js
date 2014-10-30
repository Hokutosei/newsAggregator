'use strict';

var app = angular.module("learnJap", ["ngTouch"]);

console.log(app);

app.controller("HeaderCtrl", ["$scope", function($scope) {
    $scope.project_name = "learnJap";

}])


app.controller("MainCtrl", ["$scope", "$window", function($scope, $window) {
    var log = function(str) { console.log(str) };
    $scope.project_name = "learnJap";

    $scope.window_height = $window.innerHeight + 'px';
}]);