'use strict';

var app = angular.module("learnJap", ["ngTouch"]);

console.log(app);

app.controller("HeaderCtrl", ["$scope", function($scope) {
    $scope.project_name = "learnJap";

}])


app.controller("MainCtrl", ["$scope", "$window", function($scope, $window) {
}]);