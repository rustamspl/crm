'use strict';
MetronicApp.controller('SimpleDetailTableController', function($rootScope, $scope, $http, $timeout,$stateParams) {


    $scope.selectRecordsState = false;
    $scope.rowCollection = [];
    $scope.pageCount = 0;
    $scope.id=$stateParams.id;


    $scope.currentPage = 1;
    $scope.perPage = 25;
    $scope.table_name = "accounts";

    $scope.init = function(opt){
        $scope.currentPage = opt.currentPage;
        $scope.perPage = opt.perPage;
        $scope.table_name = opt.table_name;
        $scope.selectRecordsState =  (opt.selectRecordsState) ? opt.selectRecordsState : false;
        $scope.bind();
    }
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

    $scope.bind = function (){
        $http.get('../restapi/query/get?code='+$scope.table_name+'&page='+$scope.currentPage+'&perpage='+$scope.perPage+"&param1="+$stateParams.id).
        success(function(data) {
            $scope.rowCollection = data.items;
            $scope.pageCount = data.pageCount;
            //alert(pageCount);
            //console.log(data.items);
        });
    }

    $scope.deleteSelectedRecord = function (){
        //alert("test");

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

});