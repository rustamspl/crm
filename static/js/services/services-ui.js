/**
 * Created by eldars on 04.01.2016.
 */

MetronicApp.service('UIService', function($http,$rootScope) {

   var editPage = function(pageCode){
        $http.get('../restapi/query/get?code=pageByCode&param1='+pageCode).
        success(function(data) {
            location.href="#/settings/pagedetails/"+data.items[0].id;
        });
    };

    var editQuery = function(query){
        $http.get('../restapi/query/get?code=queryByCode&param1='+query).
        success(function(data) {
            location.href="#/settings/querydetails/"+data.items[0].id;
        });
    }

    var bindUITools = function($theScope){
        if ($rootScope.session_roles.admin) {
            $theScope.editPage = editPage;
            $theScope.editQuery = editQuery;
        }
    }

    return {
        bindUITools:bindUITools
    };

});

