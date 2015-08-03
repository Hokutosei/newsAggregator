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
    app.run(function($rootScope, $timeout) {
      $rootScope.$on('$routeChangeSuccess', function($window, $timeout) {
        (function(d, s, id) {
          var initFB = function() {
            var js, fjs = d.getElementsByTagName(s)[0];
            if (d.getElementById(id)) return;
            js = d.createElement(s); js.id = id;
            js.src = "//connect.facebook.net/en_US/sdk.js#xfbml=1&version=v2.4&appId=604252113051244";
            fjs.parentNode.insertBefore(js, fjs);
          }

          initFB();
          // $timeout(function(){console.log("aa")}, 3000)
          setTimeout(function() { initFB(); FB.XFBML.parse()}, 4000)
        }(document, 'script', 'facebook-jssdk'));

      })
    });


}());
