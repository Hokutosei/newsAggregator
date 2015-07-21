(function() {
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
                $http({method: 'POST', url: '/feed_more', data: { 'ContentType': content_type, 'Skip': length }})
                    .success(function(data, status) {
                        callback(data, status)
                    })
            },

            incrementNewsItemScore: function(news_item) {
                log(news_item);
                $http({ method: 'POST', url: '/increment_news', data: {
                    'Id': news_item._id
                    }
                }).success(function(data, status) {
                    log(status)
                })
            },

            fetchHeaderCategories: function(callback) {
              $http.get('/header_categories').success(function(data, status) {
                callback(data)
              })
            },

            fetchCategoryNews: function(initial) {
              return $http({
                        url: '/fetch_category_news',
                        method: 'GET',
                        params: { initial: initial }
                      })
            },
            fetchTopRankingNews: function(callback) {
              $http.get('/top_ranking_news').success(function(data, status) {
                callback(data)
              })
            }
        }

    });
}());
