/***
Metronic AngularJS App Main Script
***/

/* Metronic App */
var FarabiTestApp = angular.module("FarabiTestApp", [
    "pascalprecht.translate"



]);


/********************************************
 BEGIN: BREAKING CHANGE in AngularJS v1.3.x:
*********************************************/
/**
`$controller` will no longer look for controllers on `window`.
The old behavior of looking on `window` for controllers was originally intended
for use in examples, demos, and toy apps. We found that allowing global controller
functions encouraged poor practices, so we resolved to disable this behavior by
default.

To migrate, register your controllers with modules rather than exposing them
as globals:

Before:

```javascript
function MyController() {
  // ...
}
```

After:

```javascript
angular.module('myApp', []).controller('MyController', [function() {
  // ...
}]);

Although it's not recommended, you can re-enable the old behavior like this:

```javascript
angular.module('myModule').config(['$controllerProvider', function($controllerProvider) {
  // this option might be handy for migrating old apps, but please don't use it
  // in new ones!
  $controllerProvider.allowGlobals();
}]);
**/





/********************************************
 END: BREAKING CHANGE in AngularJS v1.3.x:
*********************************************/

/* Setup global settings */



/***
Layout Partials.
By default the partials are loaded through AngularJS ng-include directive. In case they loaded in server side(e.g: PHP include function) then below partial
initialization can be disabled and Layout.init() should be called on page load complete as explained above.
***/
//alert("vata emes");

$.ajax({
        url: "../../restapi/translates/get",
        dataType: 'json',
        async: false,

        success: function(data) {
            //stuff
            //...

            FarabiTestApp.config(function($translateProvider) {
                $translateProvider.translations('en', data["en"]).translations('ru', data["ru"]).translations('kk', data["kk"]);
                $translateProvider.preferredLanguage(data.lang);
            });
        }
    });

/*
    $.getJSON( "js/i18n/translate.json", function( data ) {
        async: false,

        //console.log(data);

 FarabiTestApp.config(function($translateProvider) {
            $translateProvider.translations('en', data["en"]).translations('ru', data["ru"]);
            $translateProvider.preferredLanguage('en');

        });







});
 */







FarabiTestApp.run(['$rootScope', '$translate', '$log', function ($rootScope, $translate, $log) {


    $rootScope.changeLanguage = function (langKey) {
        $translate.use(langKey);


    $rootScope.login = function login(){
        //alert("test");
    }
    };



}]);

FarabiTestApp.config(function ($translateProvider) {
    $translateProvider.useStorage('myCustomService');
});


FarabiTestApp.controller('LoginCtrls', ['$scope','$translate','$http', function($scope,$translate,$http) {
    $scope.greeting = 'Hola!';
    $scope.changeLang=function(lang){
        $translate.use(lang);
        $scope.selectedLanguage=lang;
    }




}]);

FarabiTestApp.controller('ProcessCtrls', ['$scope','$http', function($scope,$http) {

    $scope.testing=false;
    $scope.currentQ = 1;


    $scope.currentQuestionText=function(){

        if ($scope.selectedLanguage == "kk") {
            return $scope.questions[$scope.currentQ].q_kk;
        }else{
            return $scope.questions[$scope.currentQ].q_ru;
        }
    }

    $scope.next = function (answer){

        if ($scope.questions.length<=$scope.currentQ){
            alert($scope.questions.length);
            location.href ="finish.html";
            return;
        }
        $scope.currentQ ++;




    }

    $scope.back = function (){
        $scope.currentQ --;
    }
    $scope.start=function(){


        $http.get('../../restapi/query/get?code=farabi_random_questions_by_test_id&param1=1').
        success(function(data) {
            $scope.questions = data.items;
            $scope.testing=true;
            if (data.error==2){
                location.href="/auth/logout";
            }
        });
    }

}]);


FarabiTestApp.factory('myCustomService', function ($http) {

    var curr = "";
    return {
        put: function (name, value) {
            // do something with $http to put value


                curr = value;
                if (value != undefined) {
                    //alert(name);
                    $http
                        .get("../../auth/setlanguage?lang=" + value)
                        .success(function (data) {
                            return data.lang;
                        });
                }
            }

        ,
        get: function (name) {
        // do something with $http to get value

        if (curr == "") {
            $http
                .get("../../auth/getlanguage")
                .success(function (data) {
                    curr = data.lang;
                    return data.lang;

                });

        }
    }


    };
});

