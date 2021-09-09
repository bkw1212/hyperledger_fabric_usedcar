#!/bin/bash
set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

docker-compose -f docker-compose.yaml down
docker-compose -f docker-compose.yaml up -d

# wait for Hyperledger Fabric to start
# incase of errors when running later commands, issue export FABRIC_START_TIMEOUT=<larger number>
export FABRIC_START_TIMEOUT=20
echo ${FABRIC_START_TIMEOUT}
sleep ${FABRIC_START_TIMEOUT}

# Create the channel1
docker exec cli1 peer channel create -o orderer1.hub.com:7050 -c channelsales1 -f /etc/hyperledger/configtx/channel1.tx

# Join peer0.sales1.hub.com to the channel and Update the Anchor Peers in Channel1
docker exec cli1 peer channel join -b channelsales1.block
docker exec cli1 peer channel update -o orderer1.hub.com:7050 -c channelsales1 -f /etc/hyperledger/configtx/Sales1Organchors.tx

# Join peer1.sales1.hub.com to the channel
docker exec -e "CORE_PEER_ADDRESS=peer1.sales1.hub.com:7051" cli1 peer channel join -b channelsales1.block

# Join peer0.customer.hub.com to the channel and update the Anchor Peers in Channel1
docker exec -e "CORE_PEER_LOCALMSPID=CustomerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.hub.com/users/Admin@customer.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.customer.hub.com:7051" cli1 peer channel join -b channelsales1.block
sleep 3
docker exec -e "CORE_PEER_LOCALMSPID=CustomerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.hub.com/users/Admin@customer.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.customer.hub.com:7051" cli1 peer channel update -o orderer1.hub.com:7050 -c channelsales1 -f /etc/hyperledger/configtx/CustomerOrganchorsChannel1.tx

# Join peer1.customer.hub.com to the channel
docker exec -e "CORE_PEER_LOCALMSPID=CustomerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.hub.com/users/Admin@customer.hub.com/msp" -e "CORE_PEER_ADDRESS=peer1.customer.hub.com:7051" cli1 peer channel join -b channelsales1.block

# Join peer0.insurance1.hub.com to the channel and update the Anchor Peers in Channel1
#docker exec -e "CORE_PEER_LOCALMSPID=Insurance1Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insurance1.hub.com/users/Admin@insurance1.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.insurance1.hub.com:7051" cli1 peer channel join -b channelsales1.block
#docker exec -e "CORE_PEER_LOCALMSPID=Insurance1Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insurance1.hub.com/users/Admin@insurance1.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.insurance1.hub.com:7051" cli1 peer channel update -o orderer1.hub.com:7050 -c channelsales1 -f /etc/hyperledger/configtx/Insurance1Organchors.tx

# Join peer1.insurance1.hub.com to the channel
#docker exec -e "CORE_PEER_LOCALMSPID=Insurance1Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insurance1.hub.com/users/Admin@insurance1.hub.com/msp" -e "CORE_PEER_ADDRESS=peer1.insurance1.hub.com:7051" cli1 peer channel join -b channelsales1.block

# Join peer0.repair1.hub.com to the channel and update the Anchor Peers in Channel1
#docker exec -e "CORE_PEER_LOCALMSPID=Repair1Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/repair1.hub.com/users/Admin@repair1.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.repair1.hub.com:7051" cli1 peer channel join -b channelsales1.block
#docker exec -e "CORE_PEER_LOCALMSPID=Repair1Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/repair1.hub.com/users/Admin@repair1.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.repair1.hub.com:7051" cli1 peer channel update -o orderer1.hub.com:7050 -c channelsales1 -f /etc/hyperledger/configtx/Repair1Organchors.tx

# Join peer1.repair1.hub.com to the channel
#docker exec -e "CORE_PEER_LOCALMSPID=Repair1Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/repair1.hub.com/users/Admin@repair1.hub.com/msp" -e "CORE_PEER_ADDRESS=peer1.repair1.hub.com:7051" cli1 peer channel join -b channelsales1.block



# Create the channel2
docker exec cli2 peer channel create -o orderer1.hub.com:7050 -c channelsales2 -f /etc/hyperledger/configtx/channel2.tx

# Join peer0.sales2.hub.com to the channel and Update the Anchor Peers in Channel1
docker exec cli2 peer channel join -b channelsales2.block
docker exec cli2 peer channel update -o orderer1.hub.com:7050 -c channelsales2 -f /etc/hyperledger/configtx/Sales2Organchors.tx

# Join peer1.sales2.hub.com to the channel
docker exec -e "CORE_PEER_ADDRESS=peer1.sales2.hub.com:7051" cli2 peer channel join -b channelsales2.block

# Join peer0.customer.hub.com to the channel and update the Anchor Peers in Channel2
docker exec -e "CORE_PEER_LOCALMSPID=CustomerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.hub.com/users/Admin@customer.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.customer.hub.com:7051" cli2 peer channel join -b channelsales2.block
docker exec -e "CORE_PEER_LOCALMSPID=CustomerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.hub.com/users/Admin@customer.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.customer.hub.com:7051" cli2 peer channel update -o orderer1.hub.com:7050 -c channelsales2 -f /etc/hyperledger/configtx/CustomerOrganchorsChannel2.tx

# Join peer0.customer.hub.com to the channel
docker exec -e "CORE_PEER_LOCALMSPID=CustomerOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/customer.hub.com/users/Admin@customer.hub.com/msp" -e "CORE_PEER_ADDRESS=peer1.customer.hub.com:7051" cli2 peer channel join -b channelsales2.block

# Join peer0.insurance2.hub.com to the channel and update the Anchor Peers in Channel2
#docker exec -e "CORE_PEER_LOCALMSPID=Insurance2Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insurance2.hub.com/users/Admin@insurance2.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.insurance2.hub.com:7051" cli2 peer channel join -b channelsales2.block
#docker exec -e "CORE_PEER_LOCALMSPID=Insurance2Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insurance2.hub.com/users/Admin@insurance2.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.insurance2.hub.com:7051" cli2 peer channel update -o orderer1.hub.com:7050 -c channelsales2 -f /etc/hyperledger/configtx/Insurance2Organchors.tx

# Join peer1.insurance2.hub.com to the channel
#docker exec -e "CORE_PEER_LOCALMSPID=Insurance2Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/insurance2.hub.com/users/Admin@insurance2.hub.com/msp" -e "CORE_PEER_ADDRESS=peer1.insurance2.hub.com:7051" cli2 peer channel join -b channelsales2.block

#sleep 3
# Join peer0.repair2.hub.com to the channel and update the Anchor Peers in Channel2
#docker exec -e "CORE_PEER_LOCALMSPID=Repair2Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/repair2.hub.com/users/Admin@repair2.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.repair2.hub.com:7051" cli2 peer channel join -b channelsales2.block
#sleep 3
#docker exec -e "CORE_PEER_LOCALMSPID=Repair2Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/repair2.hub.com/users/Admin@repair2.hub.com/msp" -e "CORE_PEER_ADDRESS=peer0.repair2.hub.com:7051" cli2 peer channel update -o orderer1.hub.com:7050 -c channelsales2 -f /etc/hyperledger/configtx/Repair2Organchors.tx

# Join peer1.repair2.hub.com to the channel
#docker exec -e "CORE_PEER_LOCALMSPID=Repair2Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/repair2.hub.com/users/Admin@repair2.hub.com/msp" -e "CORE_PEER_ADDRESS=peer1.repair2.hub.com:7051" cli2 peer channel join -b channelsales2.block