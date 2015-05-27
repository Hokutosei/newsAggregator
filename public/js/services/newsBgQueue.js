(function() {
    'use strict';

    var log = function(str) { console.log(str); };

    app.service('NewsBgQueue', function($http) {
        log("loaded")

        return {
            connectSSE: function() {
                var source = new EventSource('/events/');

                source.onmessage = function(e) {
                    log(e)
                }
            }
        }

    })

}())
