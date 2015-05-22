'use strict';
angular.module('Admin', ['ngCookies']).
config(['$httpProvider', '$interpolateProvider', function($httpProvider, $interpolateProvider) {
  //设置 content-Type 格式
  $httpProvider.defaults.headers.post = {
    'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
  };
  $httpProvider.defaults.headers.delete = {
    'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'
  };
  // 设置 _xsrf 
  $httpProvider.defaults.xsrfCookieName = "_xsrf";
  $httpProvider.defaults.xsrfHanderName = "X-XSRFToken";
  //post数据格式化
  $httpProvider.defaults.transformRequest = function(data) {
    var param = function(obj) {
      var query = '';
      var name, value, fullSubName, subName, subValue, innerObj, i;
      for (name in obj) {
        value = obj[name];
        if (value instanceof Array) {
          for (i = 0; i < value.length; ++i) {
            subValue = value[i];
            fullSubName = name;
            innerObj = {};
            innerObj[fullSubName] = subValue;
            query += param(innerObj) + '&';
          }
        } else if (value instanceof Object) {
          for (subName in value) {
            subValue = value[subName];
            fullSubName = name + '-' + subName;
            innerObj = {};
            innerObj[fullSubName] = subValue;
            query += param(innerObj) + '&';
          }
        } else if (value !== undefined && value !== null) {
          query += encodeURIComponent(name) + '=' + encodeURIComponent(value) + '&';
        }
      }
      return query.length ? query.substr(0, query.length - 1) : query;
    };
    return angular.isObject(data) && String(data) !== '[object File]' ? param(data) : data;
  };
  //覆写templates标签 使用 ${}   已弃用
  $interpolateProvider.startSymbol('${');
  $interpolateProvider.endSymbol('}');
}]).controller('ArticleAddCtrl', ['$scope', '$cookies', '$window', '$http', '$log', function($scope, $cookies, $window, $http, $log) {
    var _xsrf = $cookies['_xsrf']
    $scope.images = [];
    $scope.image_ids = [];
    $scope.is_edit = 0;
    
    $scope.imgPost = function () {
        var data = new FormData();
        data.append('file', $scope.imgData.image)
        data.append('description', $scope.imgData.description)
        data.append('_xsrf', _xsrf)
        
        $http.post('/admin/images/upload', data, {
            transformRequest: angular.identity,
            headers: {'Content-Type': undefined}
        }).success(function (data) {
            $scope.images.push(data)
            if (angular.isString($scope.image_ids)) {
                $scope.image_ids = JSON.parse($scope.image_ids)
            }
            $scope.image_ids.push(data.id)
        })
    }
    $scope.delImage = function(id, index) {
        $http.delete('/admin/images/'+id+'/del',{data:{'_xsrf': _xsrf}}).success(function(data) {
            if (angular.isString($scope.image_ids)) {
                $scope.image_ids = JSON.parse($scope.image_ids)
            }
            $scope.image_ids.shift(data.id);
            $scope.images.shift($scope.images[index]);
        })
    }
    $scope.articlePost = function () {
        var data = $scope.articleData
        data['img_list'] = $scope.image_ids;
        data['_xsrf'] = _xsrf;
        if (angular.isString(data['img_list'])) {
                data['img_list']= JSON.parse(data['img_list'])
            }
        $log.info($scope.is_edit)
        if($scope.is_edit){
            $http.post('/admin/article/'+ $scope.id_no +'/edit', data).success(function(data) {
            $window.location.replace(data.url)
        })
            return
        }
        $http.post('/admin/article/add', data).success(function(data) {
            $window.location.replace(data.url)
        })
    }
}]).
controller('ArticlesCtrl', ['$scope', '$http', '$cookies', '$window', function($scope, $http, $cookies, $window) {
    $scope.delArticle = function(id) {
        $http.delete('/admin/article/'+ id + '/del', {data:{'_xsrf': $cookies['_xsrf']}}).success(function(data) {
            $window.location.reload()
        })
    }
}]).
controller('PageAddCtrl', ['$scope', '$cookies', '$window', '$http', '$log', function($scope, $cookies, $window, $http, $log) {
    var _xsrf = $cookies['_xsrf']
    $scope.images = [];
    $scope.image_ids = [];
    $scope.is_edit = 0;
    
    $scope.imgPost = function () {
        var data = new FormData();
        data.append('file', $scope.imgData.image)
        data.append('description', $scope.imgData.description)
        data.append('_xsrf', _xsrf)
        
        $http.post('/admin/images/upload', data, {
            transformRequest: angular.identity,
            headers: {'Content-Type': undefined}
        }).success(function (data) {
            $scope.images.push(data)
            if (angular.isString($scope.image_ids)) {
                $scope.image_ids = JSON.parse($scope.image_ids)
            }
            $scope.image_ids.push(data.id)
        })
    }
    $scope.delImage = function(id, index) {
        $http.delete('/admin/images/'+id+'/del',{data:{'_xsrf': _xsrf}}).success(function(data) {
            if (angular.isString($scope.image_ids)) {
                $scope.image_ids = JSON.parse($scope.image_ids)
            }
            $scope.image_ids.shift(data.id);
            $scope.images.shift($scope.images[index]);
        })
    }
    $scope.articlePost = function () {
        var data = $scope.articleData
        data['img_list'] = $scope.image_ids;
        data['_xsrf'] = _xsrf;
        if (angular.isString(data['img_list'])) {
                data['img_list']= JSON.parse(data['img_list'])
            }
        $log.info($scope.is_edit)
        if($scope.is_edit){
            $http.post('/admin/page/'+ $scope.articleData._slug, data).success(function(data) {
            $window.location.replace(data.url)
        })
            return
        }
        $http.post('/admin/page/add', data).success(function(data) {
            $window.location.replace(data.url)
        })
    }
}]).
controller('CategoriesCtrl', ['$scope', '$http', '$cookies', '$window', function($scope, $http, $cookies, $window) {
    $scope.delCategory = function(id) {
        $http.delete('/admin/category/'+ id, {data:{'_xsrf': $cookies['_xsrf']}}).success(function(data) {
            $window.location.reload()
        })
    }
}]).
controller('PagesCtrl', ['$scope', '$http', '$cookies', '$window', function($scope, $http, $cookies, $window) {
    $scope.delPage = function(slug) {
        $http.delete('/admin/page/'+ slug, {data:{'_xsrf': $cookies['_xsrf']}}).success(function(data) {
            $window.location.reload()
        })
    }
}]).
directive('fileModel', ['$parse', function($parse) {
  return {
    restrict: 'A',
    link: function(scope, element, attrs) {
      var model = $parse(attrs.fileModel);
      var modelSetter = model.assign;

      element.bind('change', function() {
        scope.$apply(function() {
          modelSetter(scope, element[0].files[0]);
        });
      });
    }
  };
}])
