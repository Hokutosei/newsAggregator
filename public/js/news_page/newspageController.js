(function() {
    'use strict';

    var log = function(str) { console.log(str); };

    app.controller('NewsPageCtrl', function($routeParams, newsPageService) {
        log($routeParams)
        log("load news page")

        newsPageService.getNewsItem($routeParams.id).then(function(data, status) {
            log(data)
            log(status)
        })

    })

}());
