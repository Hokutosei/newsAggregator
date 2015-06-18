(function() {
    'use strict';

    app = angular.module("newsAggregator", ["ngTouch", 'ngCookies', "angulartics", "angulartics.google.analytics", 'ngResource', 'ngRoute']);

    var log = function(str) { console.log(str) };

    app.config(function($routeProvider, $locationProvider) {

	$routeProvider
		.when('/', {
			templateUrl: 'js/index_page/' + 'index.html',
			controller: 'MainCtrl'
		})
    .when('/:q', {
			templateUrl: 'js/index_page/' + 'index.html',
			controller: 'MainCtrl'
		})
    .when('/admin', {
      templateUrl: 'js/pages/admin/templates/admin.html',
      controller: 'AdminController'
    })
    .when('/news/:id', {
        templateUrl: 'js/news_page/template/news.html',
        controller: 'NewsPageCtrl'
    })
		.otherwise({
			redirectTo: '/'
		})

	$locationProvider.html5Mode({
		enabled: true,
		requireBase: true
	})
})


}());
