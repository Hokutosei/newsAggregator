'use strict';

var log = function(str) { console.log(str) };

app.factory('httpService', function($http) {
    return {
        getIndexNews: function(callback) {
            $http.get('/get_index_news').success(function(data, status) {
                callback(data, status)
            })
        },

        getNewsContent: function(content_type, callback) {
        	$http.get('/' + content_type).success(function(data, status) {
        		callback(data, status)
        	})
        }
    }

});