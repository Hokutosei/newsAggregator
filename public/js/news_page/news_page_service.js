(function() {
    'use strict';

    app.service('newsPageService', function($http) {
        return {
            getNewsItem: function(news_id) {
                return $http.get('/news_item', {params: { news_id: news_id }})
            }
        }
    })

}())
