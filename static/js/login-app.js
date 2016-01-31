/***
Metronic AngularJS App Main Script
***/

/* Metronic App */
var LoginApp = angular.module("LoginApp", [
    "pascalprecht.translate",


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

/*
    $.getJSON( "js/i18n/translate.json", function( data ) {
        async: false,

        //console.log(data);

        LoginApp.config(function($translateProvider) {
            $translateProvider.translations('en', data["en"]).translations('ru', data["ru"]);
            $translateProvider.preferredLanguage('en');

        });







});
 */







LoginApp.run(['$rootScope', '$translate', '$log', function ($rootScope, $translate, $log) {

    $rootScope.changeLanguage = function (langKey) {
        $translate.use(langKey);

    $rootScope.login = function login(){
        alert("test");
    }
    };
}]);





