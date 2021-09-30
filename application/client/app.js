'use strict';

var app = angular.module('application', []);
        
app.controller('AppCtrl', function($scope, appFactory){
        $("#success_setcar").hide();
        $("#success_getallcar").hide();
        $("#success_getcar").hide();
        $("#success_getwallet").hide();
        $("#success_deletecar").hide();
        $("#success_setrepair").hide();
        $("#success_getrepair").hide();
        $("#success_setinsurance").hide();
        $("#success_getinsurance").hide();
        $("#success_setrenewal").hide();
        $("#success_getallrepair").hide();
        $("#success_getallinsurance").hide();

        $scope.getCA = function(){

        }


        $scope.getWallet = function(){
                appFactory.getWallet($scope.walletid, function(data){
                        $scope.search_wallet = data;
                        $("#success_getwallet").show();
                });
        }
       $scope.getAllCar = function(){
                appFactory.getAllCar(function(data){
                        var array = [];
                        for (var i = 0; i < data.length; i++){
                                parseInt(data[i].Key);
                                data[i].Record.Key = data[i].Key;
                                array.push(data[i].Record);
                                $("#success_getallcar").hide();
                        }
                        array.sort(function(a, b) {
                            return parseFloat(a.Key) - parseFloat(b.Key);
                        });
                        $scope.allCar = array;
                });
        }
        $scope.getCar = function(){
                appFactory.getCar($scope.carkey, function(data){
                        $("#success_getcar").show();
                        var array = [];
                        for (var i = 0; i < data.length; i++){
                                data[i].Key = $scope.carkey;
                                data[i].model = data[i].Model;
                                data[i].maker = data[i].Maker;
                                data[i].price = data[i].Price;
                                data[i].walletid = data[i].WalletID;
                                data[i].purchasecount = data[i].PurchaseCount;
                                data[i].repaircount = data[i].RepairCount;
                                array.push(data[i]);
                        }
                        $scope.allCar = array;
                });
        }
        $scope.setCar = function(){
            appFactory.setCar($scope.car, function(data){
                        $scope.create_car = data;
                        $("#success_setcar").show();
            });
        }
        $scope.purchaseCar = function(key){
                appFactory.purchaseCar(key, function(data){
                        var array = [];
                        for (var i = 0; i < data.length; i++){
                                parseInt(data[i].Key);
                                data[i].Record.Key = data[i].Key;
                                array.push(data[i].Record);
                                $("#success_getallcar").hide();
                        }
                        array.sort(function(a, b) {
                            return parseFloat(a.Key) - parseFloat(b.Key);
                        });
                        $scope.allCar = array;
                });
        }
        
        $scope.deleteCar = function(){
                appFactory.deleteCar($scope.carkeydelete, function(data){
                        $scope.delete_car = data;
                        $("#success_deletecar").show();
                });
        }

        $scope.setRepair = function(){
                appFactory.setRepair($scope.repair, function(data){
                            $scope.create_repair = data;
                            $("#success_setrepair").show();
                });
        }

        $scope.getRepair = function(){
                appFactory.getRepair($scope.repairkey, function(data){
                        $("#success_getrepair").show();
                        var array = [];
                        for (var i = 0; i < data.length; i++){
                                data[i].Key = $scope.repairkey;
                                data[i].engineer = data[i].Engineer;
                                data[i].date = data[i].Date;
                                data[i].car = data[i].Rcar;
                                data[i].information = data[i].Information;
                                array.push(data[i]);
                        }
                        $scope.allRepair = array;
                });
        }

        $scope.getAllRepair = function(){
                appFactory.getAllRepair(function(data){
                        var array = [];
                        for (var i = 0; i < data.length; i++){
                                parseInt(data[i].Key);
                                data[i].Record.Key = data[i].Key;
                                array.push(data[i].Record);
                                $("#success_getallrepair").hide();
                        }
                        array.sort(function(a, b) {
                            return parseFloat(a.Key) - parseFloat(b.Key);
                        });
                        $scope.allRepair = array;
                });
        }

/*        $scope.setInsurance = function(){
                appFactory.setInsurance($scope.insurance, function(data){
                            $scope.create_insurance = data;
                            $("#success_setinsurance").show();
                });
        }

        $scope.getInsurance = function(){
                appFactory.getInsurance($scope.insurancekey, function(data){
                        $("#success_getinsurance").show();
                        var array = [];
                        for (var i = 0; i < data.length; i++){
                                data[i].Key = $scope.insurancekey;
                                data[i].car = data[i].Icar;
                                data[i].turm = data[i].Turm;
                                
                                array.push(data[i]);
                        }
                        $scope.allInsurance = array;
                });
        }

        $scope.getAllInsurance = function(){
                appFactory.getAllInsurance(function(data){
                        var array = [];
                        for (var i = 0; i < data.length; i++){
                                parseInt(data[i].Key);
                                data[i].Record.Key = data[i].Key;
                                array.push(data[i].Record);
                                $("#success_getallinsurance").hide();
                        }
                        array.sort(function(a, b) {
                            return parseFloat(a.Key) - parseFloat(b.Key);
                        });
                        $scope.allInsurance = array;
                });
        }

        $scope.setRenewal = function(){
                appFactory.setRenewal($scope.renewal, function(data_n){
                            $scope.create_renewal = data_n;
                            $("#success_setrenewal").show();
                });
        }
*/        
        
});


 app.factory('appFactory', function($http){
        var factory = {};

        factory.getWallet = function(key, callback){
            $http.get('/api/getWallet?walletid='+key).success(function(output){
                        callback(output)
                });
        }
        factory.getAllCar = function(callback){
            $http.get('/api/getAllCar/').success(function(output){
                        callback(output)
                });
        }
        factory.getCar = function(key, callback){
            $http.get('/api/getCar?carkey='+key).success(function(output){
                        callback(output)
                });
        }
        factory.setCar = function(data, callback){
            $http.get('/api/setCar?model='+data.model+'&maker='+data.maker+'&price='+data.price+'&walletid='+data.walletid).success(function(output){
                        callback(output)
                });
        }
        factory.purchaseCar = function(key, callback){
            $http.get('/api/purchaseCar?walletid=lmj1212&walletid_1=bkw1212&carkey='+key).success(function(output){
                $http.get('/api/getAllCar/').success(function(output){
                        callback(output)
                });
            });
        }
        factory.deleteCar = function(key, callback){
            $http.get('/api/deleteCar?carkey='+key).success(function(output){
                        callback(output)
                });
        }
        factory.setRepair = function(data, callback){
                $http.get('/api/setRepair?engineer='+data.engineer+'&date='+data.date+'&rcar='+data.rcar+'&information='+data.information).success(function(output){
                            callback(output)
                    });
        }
        factory.getRepair = function(key, callback){
                $http.get('/api/getRepair?repairkey='+key).success(function(output){
                            callback(output)
                    });
        }

        factory.getAllRepair = function(callback){
                $http.get('/api/getAllRepair/').success(function(output){
                            callback(output)
                    });
        }

        /*factory.setInsurance = function(data, callback){
                $http.get('/api/setInsurance?icar='+data.icar+'&turm='+data.turm).success(function(output){
                            callback(output)
                    });
        }
        factory.getInsurance = function(key, callback){
                $http.get('/api/getInsurance?insurancekey='+key).success(function(output){
                            callback(output)
                    });
        }

        factory.getAllInsurance = function(callback){
                $http.get('/api/getAllInsurance/').success(function(output){
                            callback(output)
                    });
        }

        factory.setRenewal = function(data_n, callback){
                $http.get('/api/setRenewal?insurancekey='+data_n.insurancekey+'&turm='+data_n.turm).success(function(output){
                            callback(output)
                    });
        }*/
        return factory;
});