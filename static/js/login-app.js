/***
Metronic AngularJS App Main Script
***/

/* Metronic App */
var LoginApp = angular.module("LoginApp", [
    "pascalprecht.translate",


]);

$.ajax({
        url: "../../restapi/translates/get",
        dataType: 'json',
        async: false,

        success: function(data) {
            //stuff
            //...
            LoginApp.config(function($translateProvider) {
                $translateProvider.translations('en', data["en"]).translations('ru', data["ru"]).translations('kk', data["kk"]);
                $translateProvider.preferredLanguage('ru');

            });
        }
    });


LoginApp.run(['$rootScope', '$translate', '$log', function ($rootScope, $translate, $log) {

    $rootScope.changeLanguage = function (langKey) {
        $translate.use(langKey);
    };
}]);






LoginApp.controller('LoginController', function($scope, $http, $timeout) {

 $scope.doLogin = function(){
     var loginData = {login: $scope.login, password: $scope.password, system: "browser"};
     $scope.loginIncorrect = false;
     $http
         .post("../../restapi/login",loginData)
         .success(function(data) {
             if (data.Result === "ok")
             {
                 location.href = data.RedirectURL;
             }
             else {
                $scope.loginIncorrect = true;
             }


         });
 }

});