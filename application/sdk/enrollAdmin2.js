'use strict';

const FabricCAServices = require('fabric-ca-client');
const { FileSystemWallet, X509WalletMixin } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const express = require('express');
const router = express.Router();

const ccpPath = path.resolve(__dirname, '..', 'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

const caInfo = ccp.certificateAuthorities['ca.sales1.hub.com'];
const caTLSCACerts = caInfo.tlsCACerts.pem;
const ca = new FabricCAServices(caInfo.url, { trustedRoots: caTLSCACerts, verify: false }, caInfo.caName);

const walletPath = path.join(process.cwd(), '..', 'wallet');
const wallet = new FileSystemWallet(walletPath);
console.log(`Wallet path: ${walletPath}`);

async function getCA(req, res){

    try{
        const adminExists = await wallet.exists('admin');
        if (adminExists) {
            console.log('An identity for the admin user "admin" already exists in the wallet');
            return;
        }

        const enrollment = await ca.enroll({ enrollmentID: 'admin', enrollmentSecret: 'adminpw' });
        const identity = X509WalletMixin.createIdentity('Sales1Org', enrollment.certificate, enrollment.key.toBytes());
        await wallet.import('admin', identity);
        console.log('Successfully enrolled admin user "admin" and imported it into the wallet');

        const userExists = await wallet.exists('user1');
        if (!userExists) {
                
            const gateway = new Gateway();
            await gateway.connect(ccpPath, { wallet, identity: 'admin', discovery: { enabled: true, asLocalhost: true } });
            const ca = gateway.getClient().getCertificateAuthority();
            const adminIdentity = gateway.getCurrentIdentity();

            const secret = await ca.register({ affiliation: 'org1.department1', enrollmentID: 'user1', role: 'client' }, adminIdentity);
            const enrollment = await ca.enroll({ enrollmentID: 'user1', enrollmentSecret: secret });
            const userIdentity = X509WalletMixin.createIdentity('Sales1Org', enrollment.certificate, enrollment.key.toBytes());
            await wallet.import('user1', userIdentity);
            console.log('Successfully registered and enrolled admin user "user1" and imported it into the wallet');
        }

        res.json({"msg":"ok"});
    }catch(e){
        console.loe(e);
        res.json({"msg":"connect error"});
    }
}

module.exports = {
    send:send
}
    




/*router.get('/get', async (req, res, next) => {
    try{
        console.log("query..."+req.body.id);
        const userExists = await wallet.exists('user1');
        if(!userExists){
            console.log('An identity for th user "user1" does not exist in the wallet');
            await res.json({'msg':'연결부터 해주세요'});
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'admin', discovery: { enabled: true, asLocalhost: true } });

        const network = await gateway.getNetwork()
        const network = await gateway.getNetwork('channelsales1');
        const contract = network.getContract('car-cc-ch1');

    }
})
*/