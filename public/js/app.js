(function() {
    'use strict';

    app = angular.module("newsAggregator", ["ngTouch", "angulartics", "angulartics.google.analytics", 'ngResource', 'ngRoute']);

    var log = function(str) { console.log(str) };

    app.config(function($routeProvider, $locationProvider) {

	$routeProvider
		.when('/', {
			templateUrl: 'js/index_page/' + 'index.html',
			controller: 'MainCtrl'
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
