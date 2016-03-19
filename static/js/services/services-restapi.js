/**
 * Created by eldars on 15.03.2016.
 */

MetronicApp.service('RestApiService', function($http,$rootScope) {

    if ($rootScope.isMobile) {
        var baseHost = $rootScope.uri+"/restapi/"
    }else{
        var baseHost = "/restapi/"
    }
    var post = function (uri,data){
        return  $http.post(baseHost+uri,data);
    }
    var get = function (uri){
        return  $http.get(baseHost+uri);
    }
    return {
        post:post,
        get:get

    };

});

