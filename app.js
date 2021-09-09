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
                appFactory.setRepair($scope.repair, function(data_r){
                            $scope.create_repair = data_r;
                            $("#success_setrepair").show();
                });
        }

        $scope.getRepair = function(){
                appFactory.getRepair($scope.repairkey, function(data){
                        $("#success_getrepair").show();
                        var array = [];
                        for (var i = 0; i < data.length; i++){
                                data[i].R_Key = $scope.repairkey;
                                data[i].engineer = data[i].Engineer;
                                data[i].date = data[i].Date;
                                data[i].rcar = data[i].Rcar;
                                array.push(data[i]);
                        }
                        $scope.allRepair = array;
                });
        }

        $scope.getAllRepair = function(){
                appFactory.getAllRepair(function(data_r){
                        var array_r = [];
                        for (var i = 0; i < data_r.length; i++){
                                parseInt(data_r[i].R_Key);
                                data_r[i].Record.R_Key = data_r[i].R_Key;
                                array_r.push(data_r[i].Record);
                                $("#success_getallrepair").hide();
                        }
                        array_r.sort(function(a, b) {
                            return parseFloat(a.R_Key) - parseFloat(b.R_Key);
                        });
                        $scope.allRepair = array_r;
                });
        }

        $scope.setInsurance = function(){
                appFactory.setInsurance($scope.insurance, function(data_i){
                            $scope.create_insurance = data_i;
                            $("#success_setinsurance").show();
                });
        }

        $scope.getInsurance = function(){
                appFactory.getInsurance($scope.insurancekey, function(data_i){
                        $("#success_getinsurance").show();
                        var array_i = [];
                        for (var i = 0; i < data_i.length; i++){
                                data_i[i].I_Key = $scope.insurancekey;
                                data_i[i].icar = data_i[i].Icar;
                                data_i[i].turm = data_i[i].Turm;
                                array_i.push(data_i[i]);
                        }
                        $scope.allInsurance = array_i;
                });
        }

        $scope.getAllInsurance = function(){
                appFactory.getAllInsurance(function(data_i){
                        var array_i = [];
                        for (var i = 0; i < data_i.length; i++){
                                parseInt(data_i[i].I_Key);
                                data_i[i].Record.I_Key = data_i[i].I_Key;
                                array_i.push(data_i[i].Record);
                                $("#success_getallinsurance").hide();
                        }
                        array_i.sort(function(a, b) {
                            return parseFloat(a.I_Key) - parseFloat(b.I_Key);
                        });
                        $scope.allInsurance = array_i;
                });
        }

        $scope.setRenewal = function(){
                appFactory.setRenewal($scope.renewal, function(data_n){
                            $scope.create_renewal = data_n;
                            $("#success_setrenewal").show();
                });
        }
        
        
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
                $http.get('/api/setRepair?engineer='+data.engineer+'&date='+data.date+'&rcar='+data.rcar).success(function(output){
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

        factory.setInsurance = function(data_i, callback){
                $http.get('/api/setInsurance?icar='+data_i.icar+'&turm='+data_i.turm).success(function(output){
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
        }
        return factory;
});