var sdk = require('./sdk.js');
module.exports = function(app){
  app.get('/api/getWallet', function (req, res) {
    var walletid = req.query.walletid;
    let args = [walletid];
    sdk.send(false, 'getWallet', args, res);
  });
  app.get('/api/setWallet', function(req, res){
    var name = req.query.name;
		var id = req.query.id;
    var coin = req.query.coin;
    let args = [name, id, coin];
    sdk.send(true, 'setWallet', args, res);
  });
  app.get('/api/getCar', function(req, res){
    var carkey = req.query.carkey;
    let args = [carkey];
    sdk.send(false, 'getCar', args, res);
  });
  app.get('/api/setCar', function (req, res) {
    var model = req.query.model;
    var maker = req.query.maker;
    var price = req.query.price;
    var walletid = req.query.walletid;
    //var owner = req.query.owner;
    let args = [model, maker, price, walletid];
    sdk.send(true, 'setCar', args, res);
  });
  
  app.get('/api/getAllCar', function (req, res) {
    let args = [];
    sdk.send(false, 'getAllCar', args, res);
  });
  app.get('/api/purchaseCar', function (req, res) {
    var walletid = req.query.walletid;
    var carkey = req.query.carkey;
    var walletid_1 = req.query.walletid_1;
    let args = [walletid, walletid_1, carkey];
    sdk.send(true, 'purchaseCar', args, res);
  });
  app.get('/api/deleteCar', function(req, res){
    var carkey = req.query.carkey;
    let args = [carkey];
    sdk.send(true, 'deleteCar', args, res);
  });

  app.get('/api/setRepair', function (req, res) {
    var engineer = req.query.engineer;
    var date = req.query.date;
    var rcar = req.query.rcar;
    let args = [engineer, date, rcar];
    sdk.send(true, 'setRepair', args, res);
  });
  app.get('/api/getRepair', function(req, res){
    var repairkey = req.query.repairkey;
    let args = [repairkey];
    sdk.send(false, 'getRepair', args, res);
  });

  app.get('/api/setInsurance', function (req, res) {
    var icar = req.query.icar;
    var turm = req.query.turm;
    let args = [icar, turm];
    sdk.send(true, 'setInsurance', args, res);
  });
  app.get('/api/getInsurance', function(req, res){
    var insurancekey = req.query.insurancekey;
    let args = [insurancekey];
    sdk.send(false, 'getInsurance', args, res);
  });

  app.get('/api/setRenewal', function (req, res) {
    var insurancekey = req.query.insurancekey;
    var turm = req.query.turm;
    let args = [insurancekey, turm];
    sdk.send(true, 'setRenewal', args, res);
  });

  app.get('/api/getAllRepair', function (req, res) {
    let args = [];
    sdk.send(false, 'getAllRepair', args, res);
  });

  app.get('/api/getAllInsurance', function (req, res) {
    let args = [];
    sdk.send(false, 'getAllInsurance', args, res);
  });
}