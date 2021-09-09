package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Wallet struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Password	string `json:"password"`
	Token string `json:"token"`

}

func (s *SmartContract) creatUser(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 4 {
        return shim.Error("Incorrect number of arguments. Expecting 2")
    }

	// check the password length
    if len(args[2]) < 6 {
        return shim.Error("Password too short")
    }

	// check if theuser is in the network
    v, err: = APIstub.GetState(args[1])
    var tmp Wallet
    json.Unmarshal(v, &tmp)
    if len(tmp.Password) >= 6 {
        return shim.Error("User already created")
    }

	// create user
    user := Wallet {
        args[0], args[1], args[2], args[3]
    }
	userAsBytes, err := json.Marshal(user)
    if err != nil {
        return shim.Error(err.Error())
    }

	// deploy it in the network
	err = APIstub.PutState(user.ID, userAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (s *SmartContract) login(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	// check arguments
    if len(args) != 2 {
        return shim.Error("Incorrect number of arguments. Expecting 2")
    }

	// get user
    value, err := APIstub.GetState(args[0])
    if err != nil {
        return shim.Error("User not found")
    }

	// decrypt user[]byte("Status: " + asset.Status)
    var user Wallet
    json.Unmarshal(value, & user)

    // check if the password is correct
    if user.Password != args[1] {
        return shim.Error("Credentials not correct")
    }

	// get asset info
    return shim.Success(nil)
}


func (s *SmartContract) exist(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	// check arguments
    if len(args) != 1 {
        return shim.Error("Invalid Argument Number")
    }

	// get user
    value, _: = APIstub.GetState(args[0])
	
	// decrypt user
    var user Wallet
    json.Unmarshal(value, & user)
    
	// check if the user exist
    if len(user.Password) < 6 {
        return shim.Error("User not found")
    }
    
	// get asset info
    return shim.Success(nil)
}