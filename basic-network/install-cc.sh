#!/bin/bash
set -ev

# install chaincode for channelsales1
docker exec cli1 peer chaincode install -n car-cc-ch1 -v 1.0.7 -p chaincode/go
sleep 1
# instantiate chaincode for channelsales1
docker exec cli1 peer chaincode instantiate -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -v 1.0.7 -c '{"Args":[""]}' -P "OR ('Sales1Org.member','CustomerOrg.member', 'Insurance1Org.member', 'Repair1.member')"
sleep 6
# invoke chaincode for channelsales1
#docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"initWallet","Args":[""]}'
#docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"setWallet","Args":["Byun", "bkw1212", "200"]}'
#sleep 2
#docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"setWallet","Args":["Lee", "lmj1212", "400"]}'
#sleep 3

docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"creatUser","Args":["Lee", "lmj1212","1234567", "400"]}'
sleep 3

docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"creatUser","Args":["Byun", "bkw1212","2345678", "200"]}'
sleep 2

docker exec cli1 peer chaincode query -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"login","Args":["bkw1212","2345678"]}'
sleep 2


docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"setCar","Args":["G90", "Genesis", "40", "bkw1212"]}'
sleep 2
docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"setCar","Args":["K9", "KIA", "30", "bkw1212"]}'
sleep 2
docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"setCar","Args":["Sonata", "Hyundai", "20", "bkw1212"]}'
sleep 3

# query chaincode for channelsales1
docker exec cli1 peer chaincode query -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"getCar","Args":["CAR1"]}'
sleep 2

docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"purchaseCar","Args":["lmj1212", "bkw1212", "CAR0"]}'
sleep 2

docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"setRepair","Args":["baek", "2021/06/14", "CAR1"]}'
sleep 2

docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"setRepair","Args":["baek", "2021/09/29", "CAR1"]}'
sleep 2

docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"getRepair","Args":["REPAIR0"]}'
sleep 2

docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"getAllRepair","Args":[""]}'
sleep 2

docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"setInsurance","Args":["CAR1", "2021/06/14 ~ 2022/06/14"]}'
sleep 2

docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"setInsurance","Args":["CAR0", "2021/01/14 ~ 2022/01/14"]}'
sleep 2

docker exec cli1 peer chaincode query -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"getInsurance","Args":["INSURANCE0"]}'
sleep 3

docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"setRenewal","Args":["INSURANCE0", "2022/06/14 ~ 2023/06/14"]}'
sleep 2

#docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"setRenewal","Args":["INSURANCE1", "2022/06/14 ~ 2023/06/14"]}'
#sleep 2

docker exec cli1 peer chaincode query -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"getInsurance","Args":["INSURANCE0"]}'
sleep 2

#docker exec cli1 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"creatUser","Args":["bkw1212", "k970307r"]}'
#sleep 2

#docker exec cli1 peer chaincode query -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"login","Args":["bkw1212", "k970307r"]}'
#sleep 2

#docker exec cli1 peer chaincode query -o orderer1.hub.com:7050 -C channelsales1 -n car-cc-ch1 -c '{"function":"exist","Args":["bkw1212"]}'
#sleep 2

# install chaincode for channelsales2
#docker exec cli2 peer chaincode install -n car-cc-ch2 -v 1.0.1 -p chaincode/go
#sleep 1
# install chaincode for channelsales2
#docker exec cli2 peer chaincode instantiate -o orderer1.hub.com:7050 -C channelsales2 -n car-cc-ch2 -v 1.0.1 -c '{"Args":[""]}' -P "OR ('Sales2Org.member','CustomerOrg.member', 'Insurance2Org.member', 'Repair2Org.member')"
#sleep 3
# invoke chaincode for channelsales2
#docker exec cli2 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales2 -n music-cc-ch2 -c '{"function":"initWallet","Args":[""]}'
#docker exec cli2 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales2 -n car-cc-ch2 -c '{"function":"setWallet","Args":["Baek", "Baek1212", "300"]}'
#sleep 2
#docker exec cli2 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales2 -n car-cc-ch2 -c '{"function":"setWallet","Args":["Han", "Han1212", "500"]}'
#sleep 3

#docker exec cli2 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales2 -n car-cc-ch2 -c '{"function":"setCar","Args":["G70", "Genesis", "30", "Han1212"]}'
#sleep 2
#docker exec cli2 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales2 -n car-cc-ch2 -c '{"function":"setCar","Args":["K7", "KIA", "20", "Baek1212"]}'
#sleep 2
#docker exec cli2 peer chaincode invoke -o orderer1.hub.com:7050 -C channelsales2 -n car-cc-ch2 -c '{"function":"setCar","Args":["Sonata", "Hyundai", "20", "Baek1212"]}'

#sleep 3
# query chaincode for channelsales2
#docker exec cli2 peer chaincode query -o orderer1.hub.com:7050 -C channelsales2 -n car-cc-ch2 -c '{"function":"getCar","Args":["CAR0"]}'

