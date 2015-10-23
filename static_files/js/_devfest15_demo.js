/**
* Main AngularJS
* DevFest 2015 Demo
*/

var devfest15_demo = angular.module('devFest15_demo', ['ngRoute', 'contenteditable']);

/**
* Configuration
*/
devfest15_demo.config(['$routeProvider', '$locationProvider', function ($routeProvider, $locationProvider) {
    $locationProvider.html5Mode(true);
    $locationProvider.hashPrefix('!');

    $routeProvider
    // Home
    .when("/", { templateUrl: "partials/blog.html", controller: "BlogCtrl" })
    .when("/blog/:post_id", { templateUrl: "partials/blog_item.html", controller: "BlogCtrl" })
    // Static Pages
    .when("/blog/", { templateUrl: "partials/blog_add.html", controller: "BlogCtrl" })
    
    

    .when("/error-404", { templateUrl: "partials/error-404.html", controller: "PageCtrl" })
    .otherwise("/error-404");
}]);

/**
* Controllers
*/

var BASE_API = "//devfest-demo-1104.appspot.com/";
var API_BLOG_ITEM_QUERY     = BASE_API + "get/";
var API_BLOG_ITEM_GET       = BASE_API + "get/{post_id}";
var API_BLOG_ITEM_UPDATE    = BASE_API + "update/{post_id}";
var API_BLOG_ITEM_ADD       = BASE_API + "post/"//"{post_id}";
devfest15_demo.controller('BlogCtrl', ['$scope', '$location', '$http', '$log', '$routeParams', '$rootScope', function ($scope, $location, $http, $log, $routeParams, $rootScope) {
    $log.debug("Blog Controller");
    $scope.$error = null;
    $log.info($location.search())
    if ($location.search().action == "new" || $location.search() == "update") {
        $scope.$dirty = $location.search().action == "new";

        var $watch_cancel = $scope.$watch("blog_item", function (n, o) {
            if (n && o && (
                 (n.title != "" && n.title && n.title != o.title) ||
                 (n.content != "" && n.content && n.content != o.content)
                )) {
                $scope.$dirty = true;
            }
        }, true);
        $scope.$on("$destroy", $watch_cancel);

        $scope.Update = function ($event) {
            $http
            .post(API_BLOG_ITEM_UPDATE.replace(/{post_id}/g, $routeParams.post_id), $scope.blog_item)
            .then(function success(response) {
                $log.debug(response);
            }, function error(response) {
                $log.error(response);
                $scope.$error = response;
            });
        };
        $scope.Add = function ($event) {
            $http
            .post(API_BLOG_ITEM_ADD, $scope.blog_item)
            .then(function success(response) {
                $log.debug(response);
            }, function error(response) {
                $log.error(response);
                $scope.$error = response;
            });
           
        };
        $scope.Cancel = function ($event) {
            if(_getPost)
                _getPost()
            if (_getPosts)
                _getPosts();
        };

    }
    if ($routeParams.post_id) {
        var _getPost = function () {
            $scope.$error = null;
            $http
            .get(API_BLOG_ITEM_GET.replace(/{post_id}/g, $routeParams.post_id))
            .then(function success(response) {
                $log.debug(response);
                $scope.blog_item = response.data;
                $rootScope.$broadcast("gdg-page-loaded");
            }, function error(response) {
                $log.error(response);
                $scope.$error = response;
                $rootScope.$broadcast("gdg-page-loaded")
            });
        }
        _getPost();

    } else {
        var _getPosts = function () {
            $scope.$error = null;
            $http
            .get(API_BLOG_ITEM_QUERY)
            .then(function success(response) {
                $log.debug(response);
                $scope.blog = response.data;
                $rootScope.$broadcast("gdg-page-loaded")
            }, function error(response) {
                $log.error(response);
                $scope.$error = response;
                $rootScope.$broadcast("gdg-page-loaded")
            });
        }
        _getPosts();
    }
}]);
devfest15_demo.controller('PageCtrl', ['$scope', '$location', '$http', '$log', function ($scope, $location, $http, $log) {
    $log.debug("Page Controller");
    $rootScope.$broadcast("gdg-page-loaded");
}]);

devfest15_demo.directive('preloader', ['$timeout', '$log', function ($timeout, $log) {
    return {
        restrict: 'AE',
        scope: {
            isLoaded: '='
        },
        template:   '<div id="preloader" ng-class=" { \'loaded\': isLoaded} ">' +
                    '   <div id="loader-wrapper">' +
                    '       <div id="loader"></div>' +
                    '       <div class="loader-section section-left"></div>' + 
                    '       <div class="loader-section section-right"></div>' +
                    '   </div>' +
                    '</div>',
        link: function (scope) {
            var _setLoaded = function () {
                scope.isLoaded = true;
            }
            scope.$on("gdg-page-loaded", _setLoaded);
            $timeout(_setLoaded, 5000);
        }
    };
}]);