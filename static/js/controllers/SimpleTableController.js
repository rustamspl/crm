'use strict';
MetronicApp.controller('SimpleTableController', function($rootScope, $scope, $http, $timeout,UIService, DMLService,$filter,$location,RestApiService) {




    UIService.bindUITools($scope);



    $scope.getQueryParams = function (qs) {
        qs = qs.split('+').join(' ');

        var params = {},
            tokens,
            re = /[?&]?([^=]+)=([^&]*)/g;

        while (tokens = re.exec(qs)) {
            params[decodeURIComponent(tokens[1])] = decodeURIComponent(tokens[2]);
        }

        return params;
    }

    //$scope.filterTitle = query.flt$title;
    $scope.searchURL = "";


    $scope.selectRecordsState = false;
    $scope.rowCollection = [];
    $scope.pageCount = 0;


    $scope.currentPage = 1;
    $scope.perPage = 25;
    $scope.table_name = "accounts";

    $scope.init = function(opt){
        $scope.pageCode = opt.pageCode;
        //alert($scope.pageCode);
        $scope.currentPage = opt.currentPage;
        $scope.perPage = opt.perPage;
        $scope.table_name = opt.table_name;
        $scope.selectRecordsState =  (opt.selectRecordsState) ? opt.selectRecordsState : false;
        $scope.bind();

    }




    $scope.filter = ($location.search()).filter;




    $scope.selectRecords = function (){
        $scope.selectRecordsState = true;
    }

    $scope.deselectRecords = function (){
        $scope.selectRecordsState = false;
    }

    $scope.getSelectRecordState = function (){
        return $scope.selectRecordsState;
    }

    $scope.bindPage = function (inPage,inPerpage){
        $scope.currentPage = inPage;
        $scope.perPage = inPerpage;
        $scope.bind();
    }


    $scope.searchByTitle = function(text){
        //alert($scope.filterText);
        $scope.searchURL="flt$title="+text;
        $scope.bind();
    }

    $scope.searchByName = function(text){
        //alert($scope.filterText);
        $scope.searchURL="flt$name="+text;
        $scope.bind();
    }
    $scope.bind = function (){







        //var hostName = "http://dev.beton.bapps.kz:9999";

        Metronic.startPageLoading();
        RestApiService.get('query/get?code='+$scope.table_name+'&page='+$scope.currentPage+'&perpage='+$scope.perPage+"&"+$scope.searchURL).
        success(function(data) {
            Metronic.stopPageLoading();
            $scope.rowCollection = data.items;
            $scope.pageCount = data.pageCount;
            if (data.error==2){
                location.href="/auth/logout";
            }
            //alert(pageCount);
            //console.log(data.items);
        });
    }

    $scope.deleteSelectedRecord = function (){
        //alert("test");

        if (!confirm($filter('translate')('Delete Selected Records?'))){
            return;
        }
        var deleteValues = [];
        $scope.rowCollection.forEach(function(item, i, arr) {
            if (item.selected) {
                deleteValues.push({id: item.id});
                //console.log("Deleted " + item.id);
            }
        });
        $http.post('../restapi/update_v_1_1', {items: [ {table_name:$scope.table_name, action:"delete",values:deleteValues}    ]}).
        success(function (data) {
            $scope.bind();
        });
    }


    $scope.exportAll=function(){
        alert("export +"+$scope.table_name);
    }

    $scope.removeAll=function(tableName){
        if (confirm($filter('translate')('Delete All Records?'))){
            DMLService.removeAll(tableName).success(function (data) {
                alert(data.ok_text);
            })
        }
    }


});