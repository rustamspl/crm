/**
 * Created by eldars on 15.03.2016.
 */

MetronicApp.service('RestApiService', function($http) {

    var baseHost = "http://dev.beton.bapps.kz:9999/restapi/"
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

