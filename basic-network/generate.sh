#!/bin/sh

export PATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin:${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}
CHANNEL_NAME1=channelsales1
CHANNEL_NAME2=channelsales2

# remove previous crypto material and config transactions
rm -fr config/*
rm -fr crypto-config/*

# generate crypto material
./bin/cryptogen generate --config=./crypto-config.yaml
if [ "$?" -ne 0 ]; then
  echo "Failed to generate crypto material..."
  exit 1
fi

# generate genesis block for orderer
mkdir config
./bin/configtxgen -profile OrdererGenesis -outputBlock ./config/genesis.block
if [ "$?" -ne 0 ]; then
  echo "Failed to generate orderer genesis block..."
  exit 1
fi

# generate channel 1 configuration transaction
./bin/configtxgen -profile Channel1 -outputCreateChannelTx ./config/channel1.tx -channelID $CHANNEL_NAME1
if [ "$?" -ne 0 ]; then
  echo "Failed to generate channel configuration transaction..."
  exit 1
fi

# generate channel 2 configuration transaction
./bin/configtxgen -profile Channel2 -outputCreateChannelTx ./config/channel2.tx -channelID $CHANNEL_NAME2
if [ "$?" -ne 0 ]; then
  echo "Failed to generate channel configuration transaction..."
  exit 1
fi

# generate anchor peer transaction
./bin/configtxgen -profile Channel1 -outputAnchorPeersUpdate ./config/Sales1Organchors.tx -channelID $CHANNEL_NAME1 -asOrg Sales1Org
if [ "$?" -ne 0 ]; then
  echo "Failed to generate anchor peer update for Sales1Org... Channel1"
  exit 1
fi

#./bin/configtxgen -profile Channel1 -outputAnchorPeersUpdate ./config/Insurance1Organchors.tx -channelID $CHANNEL_NAME1 -asOrg Insurance1Org
#if [ "$?" -ne 0 ]; then
#  echo "Failed to generate anchor peer update for Insurance1Org... Channel1"
#  exit 1
#fi

#./bin/configtxgen -profile Channel1 -outputAnchorPeersUpdate ./config/Repair1Organchors.tx -channelID $CHANNEL_NAME1 -asOrg Repair1Org
#if [ "$?" -ne 0 ]; then
#  echo "Failed to generate anchor peer update for Repair1Org... Channel1"
#  exit 1
#fi

# generate anchor peer transaction
./bin/configtxgen -profile Channel2 -outputAnchorPeersUpdate ./config/Sales2Organchors.tx -channelID $CHANNEL_NAME2 -asOrg Sales2Org
if [ "$?" -ne 0 ]; then
  echo "Failed to generate anchor peer update for Sales2Org... Channel2"
  exit 1
fi

# generate anchor peer transaction
#./bin/configtxgen -profile Channel2 -outputAnchorPeersUpdate ./config/Insurance2Organchors.tx -channelID $CHANNEL_NAME2 -asOrg Insurance2Org
#if [ "$?" -ne 0 ]; then
#  echo "Failed to generate anchor peer update for Insurance2Org... Channel2"
#  exit 1
#fi

# generate anchor peer transaction
#./bin/configtxgen -profile Channel2 -outputAnchorPeersUpdate ./config/Repair2Organchors.tx -channelID $CHANNEL_NAME2 -asOrg Repair2Org
#if [ "$?" -ne 0 ]; then
#  echo "Failed to generate anchor peer update for Repair2Org... Channel2"
#  exit 1
#fi

# generate anchor peer transaction
./bin/configtxgen -profile Channel1 -outputAnchorPeersUpdate ./config/CustomerOrganchorsChannel1.tx -channelID $CHANNEL_NAME1 -asOrg CustomerOrg
if [ "$?" -ne 0 ]; then
  echo "Failed to generate anchor peer update for CustomerOrg... Channel1"
  exit 1
fi

# generate anchor peer transaction
./bin/configtxgen -profile Channel2 -outputAnchorPeersUpdate ./config/CustomerOrganchorsChannel2.tx -channelID $CHANNEL_NAME2 -asOrg CustomerOrg
if [ "$?" -ne 0 ]; then
  echo "Failed to generate anchor peer update for CustomerOrg... Channel2"
  exit 1
fi