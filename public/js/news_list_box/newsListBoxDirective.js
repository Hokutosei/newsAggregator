(function() {
  'use strict';
  app.directive('newsListBox', function(Utils, newsListHttp, httpService, APP_CONFIG) {
    return {
      restrict: 'E',
      templateUrl: 'js/news_list_box/template/newsListBox.html',
      scope: {
        data: '=',
        api: '@',
        title: '@'
      },
      link: function(scope, element) {
        scope.data = {}
        scope.getData = function() {
          newsListHttp.getApi(scope.api).success(function(data) {
            scope.data = data
          }).error(function(err) {
            Utils.log(err)
          })
        }


        // right_panel_item design distinguish if item has no image.url
        scope.right_panel_item = function(item) {
          if(item.image.url == '') {
            return 's12';
          }
          return 's9 offset-s2';
        }

        // news_tag style
        scope.news_tag_style = function(initial) {
          var initials = {
            'y': '#ffca28',
            'w': '#ff7043',
            'b': '#81c784',
            'p': '#4db6ac',
            'e': '#9575cd',
            's': '#afb42b',
            't': '#81d4fa',
            'ir': '#80cbc4',
          }
          return {
            'background-color': initials[initial]
          }
        }

        scope.absURL = function(news_item_id, is_index) {
          var url = httpService.absURL(news_item_id, is_index)
          Utils.log("absurl")
          Utils.log(url)
          return url
        }

        scope.title_limit = APP_CONFIG.title_limit;

        scope.getData()
      }
    }
  })
}());
