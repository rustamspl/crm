/***
Metronic AngularJS App Main Script
***/

var $stateProviderRef = null;

/* Metronic App */
var MetronicApp = angular.module("MetronicApp", [
    "ui.router", 
    "ui.bootstrap",
    "oc.lazyLoad",
    "ngSanitize",
    "pascalprecht.translate",
    'smart-table',
    'angularFileUpload',
    "ui.mask",
    "tmh.dynamicLocale",
    "ui.checkbox",
    "flowChart",
    "angular.css.injector",
    "ui.ace"
])
    .filter('split', function() {
        return function(input, splitChar, splitIndex) {
            // do some bounds checking here to ensure it has that index
            return input.split(splitChar)[splitIndex];
        }
    })
    .filter('tel', function () {
        return function (tel) {

            if (tel) {
                if (tel.length == 11) {
                    return "8" + tel.substring(1);
                }
                else {
                    return tel;
                }
            }
    }});

/* Configure ocLazyLoader(refer: https://github.com/ocombe/ocLazyLoad) */
MetronicApp.config(['$ocLazyLoadProvider', function($ocLazyLoadProvider) {
    $ocLazyLoadProvider.config({
        // global configs go here
    });
}]);

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



//AngularJS v1.3.x workaround for old style controller declarition in HTML
MetronicApp.config(['$controllerProvider', function($controllerProvider) {
  // this option might be handy for migrating old apps, but please don't use it
  // in new ones!
  $controllerProvider.allowGlobals();
}]);

/********************************************
 END: BREAKING CHANGE in AngularJS v1.3.x:
*********************************************/

/* Setup global settings */
MetronicApp.factory('settings', ['$rootScope', function($rootScope) {
    // supported languages
    var settings = {
        layout: {
            pageSidebarClosed: false, // sidebar menu state
            pageBodySolid: false, // solid body color state
            pageAutoScrollOnLoad: 1000 // auto scroll to top on page load
        },
        layoutImgPath: Metronic.getAssetsPath() + 'admin/layout/img/',
        layoutCssPath: Metronic.getAssetsPath() + 'admin/layout/css/'
    };

    $rootScope.settings = settings;
    $rootScope.mainuri = "/static";

    return settings;
}]);

/* Setup App Main Controller */
MetronicApp.controller('AppController', ['$scope', '$rootScope', function($scope, $rootScope) {
    $scope.$on('$viewContentLoaded', function() {
        Metronic.initComponents(); // init core components
        //Layout.init(); //  Init entire layout(header, footer, sidebar, etc) on page load if the partials included in server side instead of loading with ng-include directive 
    });
    //$scope.changeLanguage("ru");
}]);

/***
Layout Partials.
By default the partials are loaded through AngularJS ng-include directive. In case they loaded in server side(e.g: PHP include function) then below partial 
initialization can be disabled and Layout.init() should be called on page load complete as explained above.
***/

/* Setup Layout Part - Header */
MetronicApp.controller('HeaderController', ['$scope', function($scope) {
    $scope.$on('$includeContentLoaded', function() {
        Layout.initHeader(); // init header
    });
}]);

/* Setup Layout Part - Sidebar */
MetronicApp.controller('SidebarController', ['$scope', function($scope) {
    $scope.$on('$includeContentLoaded', function() {
        Layout.initSidebar(); // init sidebar
    });
}]);

/* Setup Layout Part - Quick Sidebar */
MetronicApp.controller('QuickSidebarController', ['$scope', function($scope) {    
    $scope.$on('$includeContentLoaded', function() {
        setTimeout(function(){
            QuickSidebar.init(); // init quick sidebar        
        }, 2000)
    });
}]);

/* Setup Layout Part - Theme Panel */
MetronicApp.controller('ThemePanelController', ['$scope', function($scope) {    
    $scope.$on('$includeContentLoaded', function() {
        Demo.init(); // init theme panel
    });
}]);

/* Setup Layout Part - Footer */
MetronicApp.controller('FooterController', ['$scope', function($scope) {



    $scope.$on('$includeContentLoaded', function() {
        Layout.initFooter(); // init footer
    });
}]);

/* Setup Rounting For All Pages */
MetronicApp.config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider) {
    // Redirect any unmatched url
    $urlRouterProvider.deferIntercept();
    $urlRouterProvider.otherwise("/dashboard.html");
    $stateProviderRef = $stateProvider;
}]);




    $.ajax({
        url: "../restapi/translates/get",
        dataType: 'json',
        async: false,

        success: function(data) {
            //stuff
            //...
            MetronicApp.config(function($translateProvider) {
                $translateProvider.translations('en', data["en"]).translations('ru', data["ru"]).translations('kk', data["kk"]);
                $translateProvider.preferredLanguage('ru');

            });
        }
    });
/*
    $.getJSON( "js/i18n/translate.json", function( data ) {
        async: false,

        //console.log(data);

        MetronicApp.config(function($translateProvider) {
            $translateProvider.translations('en', data["en"]).translations('ru', data["ru"]);
            $translateProvider.preferredLanguage('en');

        });







});
 */







MetronicApp.run(['$rootScope', '$translate', '$log','tmhDynamicLocale', function ($rootScope, $translate, $log,tmhDynamicLocale) {

    $rootScope.changeLanguage = function (langKey) {
        $translate.use(langKey);
        tmhDynamicLocale.set(langKey);
    };
}]);




///* Init global settings and run the app */
MetronicApp.run(['$urlRouter',"$rootScope", '$http',"settings", "$state","cssInjector", "$timeout", function($urlRouter,$rootScope, $http, settings, $state,cssInjector,$timeout) {





    //Open ACcount

    $rootScope.createAccount= function(caller_id){
        $rootScope.waitCall.show = false;
        window.location.href="#/crm/accountdetails/0?set_caller_id="+caller_id;
        location.reload();
    }

    $rootScope.openAccount= function(accountId){
        $rootScope.waitCall.show = false;
        window.location.href="#/crm/accountdetails/"+accountId;
        location.reload();
    }
    $rootScope.waitcall = function(){
            $timeout(function() {
                $http
                    .get("../restapi/waitcall")
                    .success(function(data) {

                        //$rootScope.waitCall = {};
                        $rootScope.waitCall = data;
                        //$rootScope.waitCall.show = true;
                        if ((data.ok) && (data.answer_state=="ringing")){
                            $rootScope.waitCall.show = true;
                            //if (confirm("Звонок "+data.caller_id+". Открыть карточку "+data.account_name)){
                              //  location.href="#/crm/accountdetails/"+data.account_id;
                            //}
                        }else if (data.ok){
                            $rootScope.waitCall.show = false;

                        }
                        $rootScope.waitcall();
                        //$rootScope.$apply();
                    }).error(function(data) {
                    console.log("fail");

                    $rootScope.waitcall();
                });
            }, 5000);

        }
    $rootScope.waitcall();

        $rootScope.$state = $state; // state to be accessed from view
        //$rootScope.$state = $state; // state to be accessed from view
        //var $state = $rootScope.$state;


    $http
        .get("../restapi/query/get?code=sessioninfo")
        .success(function(data) {

            //alert(data.items);
            if (data.items==null)
            {
                location.href="/auth/logout";
                //alert('Exit');
            }
            $rootScope.sessioninfo=data.items[0];
            cssInjector.add($rootScope.sessioninfo.company_css);
            <!--<link href="theme/assets/admin/layout/css/themes/default.css" rel="stylesheet" type="text/css" id="style_color"/>-->



        });

    $http
        .get("../restapi/query/get?code=session_parameters")
        .success(function(data) {

            $rootScope.session_parameters=[];
            angular.forEach(data.items, function(value, key) {
                $rootScope.session_parameters[value.code]=value.value;
                console.log(value.code);
                console.log(value.value);
            });
            <!--<link href="theme/assets/admin/layout/css/themes/default.css" rel="stylesheet" type="text/css" id="style_color"/>-->



        });

    $http
            .get("../restapi/pages/get")
            .success(function(data) {
                angular.forEach(data, function(value, key) {


                    //console.log(value);
                    var getExistingState = $state.get(value.name)

                    if(getExistingState !== null){
                        return;
                    }





                    var files = [];
                    angular.forEach(value["files"], function(value, key) {
                       // console.log(value);
                        files.push(value);
                    });

                data={
                    url: value.url,
                    templateUrl: value.templateurl,
                    data: {pageTitle: value.title},
                    controller: value.controller,
                    resolve: {
                        deps: ['$ocLazyLoad', function($ocLazyLoad) {
                            return $ocLazyLoad.load({
                                name: 'MetronicApp',
                                insertBefore: '#ng_load_plugins_before', // load the above css files before a LINK element with this ID. Dynamic CSS files must be loaded between core and theme css files
                                files: files
                            });
                        }]
                    }
                };

                dataPage={
                    url: value.url+"/?p[]",
                    templateUrl: value.templateurl,
                    data: {pageTitle: value.title},
                    controller: value.controller,
                    resolve: {
                        deps: ['$ocLazyLoad', function($ocLazyLoad) {
                            return $ocLazyLoad.load({
                                name: 'MetronicApp',
                                insertBefore: '#ng_load_plugins_before', // load the above css files before a LINK element with this ID. Dynamic CSS files must be loaded between core and theme css files
                                files: files
                            });
                        }]
                    }
                };


                    $stateProviderRef

                    // Dashboard
                        .state(value.name, data)
                        .state("syspage_"+value.name, dataPage)

                });
                // Configures $urlRouter's listener *after* your custom listener

                $urlRouter.sync();
                $urlRouter.listen();



            });
    }
]);




MetronicApp.config( [
    '$compileProvider',
    function( $compileProvider )
    {
        $compileProvider.aHrefSanitizationWhitelist(/^\s*(https?|sip|mailto|chrome-extension):/);
        // Angular before v1.2 uses $compileProvider.urlSanitizationWhitelist(...)
    }
]);