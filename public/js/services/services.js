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
        },

        feedMoreNews: function(content_type, length, callback) {
//            $http.post('/feed_more', { data: { 'ContentType': content_type, 'Skip': length } }, function(data, status) {
//                callback(data, status)
//            })
            $http({method: 'POST', url: '/feed_more', data: { 'ContentType': content_type, 'Skip': length }})
                .success(function(data, status) {
                    callback(data, status)
                })
        }
    }

});