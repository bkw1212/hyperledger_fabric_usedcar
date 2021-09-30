package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct{}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()

	if function == "initWallet" {
		return s.initWallet(APIstub)
	} else if function == "getWallet" {
		return s.getWallet(APIstub, args)
	} else if function == "setWallet" {
		return s.setWallet(APIstub, args)
	} else if function == "getCar" {
		return s.getCar(APIstub, args)
	} else if function == "setCar" {
		return s.setCar(APIstub, args)
	} else if function == "getAllCar" {
		return s.getAllCar(APIstub)
	} else if function == "purchaseCar" {
		return s.purchaseCar(APIstub, args)
	} else if function == "deleteCar" {
		return s.deleteCar(APIstub, args)
	} else if function == "setRepair" {
		return s.setRepair(APIstub, args)
	} else if function == "getRepair" {
		return s.getRepair(APIstub, args)
	} else if function == "getInsurance" {
		return s.getInsurance(APIstub, args)
	} else if function == "setInsurance" {
		return s.setInsurance(APIstub, args)
	} else if function == "setRenewal" {
		return s.setInsurance(APIstub, args)
	} else if function == "getAllRepair" {
		return s.getAllRepair(APIstub)
	} else if function == "getAllInsurance" {
		return s.getAllInsurance(APIstub)
	} else if function == "login" {
		return s.login(APIstub)
	} else if function == "creatUser" {
		return s.creatUser(APIstub,args)
	} else if function == "exist" {
		return s.exist(APIstub)
	}
	
	fmt.Println("Please check your function : " + function)
	return shim.Error("Unknown function")
}

func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

type Wallet struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Password	string `json:"password"`
	Token string `json:"token"`
}

func (s *SmartContract) initWallet(APIstub shim.ChaincodeStubInterface) pb.Response {

	//Declare wallets
	seller := Wallet{Name: "Byun", ID: "bkw1212", Token: "100"}
	customer := Wallet{Name: "Lee", ID: "lmj1234", Token: "200"}

	// Convert seller to []byte
	SellerasJSONBytes, _ := json.Marshal(seller)
	err := APIstub.PutState(seller.ID, SellerasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + seller.Name)
	}
	// Convert customer to []byte
	CustomerasJSONBytes, _ := json.Marshal(customer)
	err = APIstub.PutState(customer.ID, CustomerasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + customer.Name)
	}

	return shim.Success(nil)
}
func (s *SmartContract) getWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	walletAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}

	wallet := Wallet{}
	json.Unmarshal(walletAsBytes, &wallet)

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}
	buffer.WriteString("{\"Name\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.Name)
	buffer.WriteString("\"")

	buffer.WriteString(", \"ID\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.ID)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.Token)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())

}

type Car struct {
	Model      string `json:"model"`
	Maker      string `json:"maker"`
	Price      string `json:"price"`
	WalletID   string `json:"walletid"`
	PurchaseCount string `json:"purchasecount"`
	RepairCount string `json:"repaircount"`
}

type CarKey struct {
	Key string
	Idx int
}

func generateKey(APIstub shim.ChaincodeStubInterface, key string) []byte {

	var isFirst bool = false

	carkeyAsBytes, err := APIstub.GetState(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	carkey := CarKey{}
	json.Unmarshal(carkeyAsBytes, &carkey)
	var tempIdx string
	tempIdx = strconv.Itoa(carkey.Idx)
	fmt.Println(carkey)
	fmt.Println("Key is " + strconv.Itoa(len(carkey.Key)))
	if len(carkey.Key) == 0 || carkey.Key == "" {
		isFirst = true
		carkey.Key = "CAR"
	}
	if !isFirst {
		carkey.Idx = carkey.Idx + 1
	}

	fmt.Println("Last CarKey is " + carkey.Key + " : " + tempIdx)

	returnValueBytes, _ := json.Marshal(carkey)

	return returnValueBytes
}

type RepairKey struct {
	R_Key string
	R_Idx int
}

func generateKeyRepair(APIstub shim.ChaincodeStubInterface, key string) []byte {

	var isFirst bool = false

	repairkeyAsBytes, err := APIstub.GetState(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	repairkey := RepairKey{}
	json.Unmarshal(repairkeyAsBytes, &repairkey)
	var tempIdx string
	tempIdx = strconv.Itoa(repairkey.R_Idx)
	fmt.Println(repairkey)
	fmt.Println("Key is " + strconv.Itoa(len(repairkey.R_Key)))
	if len(repairkey.R_Key) == 0 || repairkey.R_Key == "" {
		isFirst = true
		repairkey.R_Key = "REPAIR"
	}
	if !isFirst {
		repairkey.R_Idx = repairkey.R_Idx + 1
	}

	fmt.Println("Last RepairKey is " + repairkey.R_Key + " : " + tempIdx)

	returnValueBytesr, _ := json.Marshal(repairkey)

	return returnValueBytesr
}

type InsuranceKey struct {
	I_Key string
	I_Idx int
}

func generateKeyInsurance(APIstub shim.ChaincodeStubInterface, key string) []byte {

	var isFirst bool = false

	insurancekeyAsBytes, err := APIstub.GetState(key)
	if err != nil {
		fmt.Println(err.Error())
	}

	insurancekey := InsuranceKey{}
	json.Unmarshal(insurancekeyAsBytes, &insurancekey)
	var tempIdxs string
	tempIdxs = strconv.Itoa(insurancekey.I_Idx)
	fmt.Println(insurancekey)
	fmt.Println("Key is " + strconv.Itoa(len(insurancekey.I_Key)))
	if len(insurancekey.I_Key) == 0 || insurancekey.I_Key == "" {
		isFirst = true
		insurancekey.I_Key = "INSURANCE"
	}
	if !isFirst {
		insurancekey.I_Idx = insurancekey.I_Idx + 1
	}

	fmt.Println("Last InsuranceKey is " + insurancekey.I_Key + " : " + tempIdxs)

	returnValueBytesrs, _ := json.Marshal(insurancekey)

	return returnValueBytesrs
}


func (s *SmartContract) setWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	var wallet = Wallet{Name: args[0], ID: args[1], Token: args[2]}

	WalletasJSONBytes, _ := json.Marshal(wallet)
	err := APIstub.PutState(wallet.ID, WalletasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + wallet.Name)
	}
	return shim.Success(nil)
}

func (s *SmartContract) setCar(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var carkey = CarKey{}
	json.Unmarshal(generateKey(APIstub, "latestKey"), &carkey)
	keyidx := strconv.Itoa(carkey.Idx)
	fmt.Println("Key : " + carkey.Key + ", Idx : " + keyidx)

	var car = Car{Model: args[0], Maker: args[1], Price: args[2], WalletID: args[3], PurchaseCount:"0", RepairCount: "0"}
	carAsJSONBytes, _ := json.Marshal(car)

	var keyString = carkey.Key + keyidx
	fmt.Println("carkey is " + keyString)

	err := APIstub.PutState(keyString, carAsJSONBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record car catch: %s", carkey))
	}

	carkeyAsBytes, _ := json.Marshal(carkey)
	APIstub.PutState("latestKey", carkeyAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getAllCar(APIstub shim.ChaincodeStubInterface) pb.Response {

	// Find latestKey
	carkeyAsBytes, _ := APIstub.GetState("latestKey")
	carkey := CarKey{}
	json.Unmarshal(carkeyAsBytes, &carkey)
	idxStr := strconv.Itoa(carkey.Idx + 1)

	var startKey = "CAR0"
	var endKey = carkey.Key + idxStr
	fmt.Println(startKey)
	fmt.Println(endKey)

	resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIter.HasNext() {
		queryResponse, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")

		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]\n")
	return shim.Success(buffer.Bytes())
}
func (s *SmartContract) purchaseCar(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	var tokenFromKey, tokenToKey int // Asset holdings
	var carprice int                 // Transaction value
	var purchasecount int
	var err error


	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	carAsBytes, err := APIstub.GetState(args[2])
	if err != nil {
		return shim.Error(err.Error())
	}
	car := Car{}
	json.Unmarshal(carAsBytes, &car)
	carprice, _ = strconv.Atoi(car.Price)
	purchasecount, _ = strconv.Atoi(car.PurchaseCount)
	
	car.PurchaseCount = strconv.Itoa(purchasecount + 1)

	SellerAsBytes, err := APIstub.GetState(args[1])
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if SellerAsBytes == nil {
		return shim.Error("Entity not found")
	}

	seller := Wallet{}
	json.Unmarshal(SellerAsBytes, &seller)    
	tokenToKey, _ = strconv.Atoi(seller.Token)    

	CustomerAsBytes, err := APIstub.GetState(args[0])  // 인자 반환
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if CustomerAsBytes == nil {
		return shim.Error("Entity not found")
	}

	customer := Wallet{}
	json.Unmarshal(CustomerAsBytes, &customer)
	tokenFromKey, _ = strconv.Atoi(string(customer.Token))

	json.Unmarshal(CustomerAsBytes, &car)

	customer.Token = strconv.Itoa(tokenFromKey - carprice)
	seller.Token = strconv.Itoa(tokenToKey + carprice) 
	car.WalletID = args[0]
	updatedCustomerAsBytes, _ := json.Marshal(customer)
	updatedSellerAsBytes, _ := json.Marshal(seller)
	updatedCarAsBytes, _ := json.Marshal(car)
	APIstub.PutState(args[0], updatedCustomerAsBytes)
	APIstub.PutState(args[1], updatedSellerAsBytes)
	APIstub.PutState(args[2], updatedCarAsBytes)
	
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	buffer.WriteString("{\"Customer Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(customer.Token)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Seller Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(seller.Token)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Car Owner\":")
	buffer.WriteString("\"")
	buffer.WriteString(car.WalletID)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) getCar(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	carAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}

	car := Car{}
	json.Unmarshal(carAsBytes, &car)

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}
	buffer.WriteString("{\"Model\":")
	buffer.WriteString("\"")
	buffer.WriteString(car.Model)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Maker\":")
	buffer.WriteString("\"")
	buffer.WriteString(car.Maker)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Price\":")
	buffer.WriteString("\"")
	buffer.WriteString(car.Price)
	buffer.WriteString("\"")

	buffer.WriteString(", \"WalletID\":")
	buffer.WriteString("\"")
	buffer.WriteString(car.WalletID)
	buffer.WriteString("\"")


	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}


func (s *SmartContract) deleteCar(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := APIstub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}


type Repair struct {
	Engineer    string `json:"engineer"`
	Date        string `json:"date"`
	Information	string `json:"infomation"`
	Rcar		string `json:"rcar"`
}


func (s *SmartContract) setRepair(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	var repaircount int

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	var repairkey = RepairKey{}
	json.Unmarshal(generateKeyRepair(APIstub, "r_latestKey"), &repairkey)
	keyidx := strconv.Itoa(repairkey.R_Idx)
	fmt.Println("Key : " + repairkey.R_Key + ", Idx : " + keyidx)

	var repair = Repair{Engineer: args[0], Date: args[1], Infomation: args[2], Rcar: args[3]}

	RepairJSONBytes, _ := json.Marshal(repair)

	var keyString = repairkey.R_Key + keyidx
	fmt.Println("repairkey is " + keyString)

	err := APIstub.PutState(keyString, RepairJSONBytes)
	if err != nil {
			return shim.Error("Failed to create asset " + repair.Engineer)
	}

	repairkeyAsBytes, _ := json.Marshal(repairkey)
	APIstub.PutState("r_latestKey", repairkeyAsBytes)

	repairAsBytes, err := APIstub.GetState(args[3])
	if err != nil {
		return shim.Error(err.Error())
	}

	car := Car{}
	json.Unmarshal(repairAsBytes, &car)

	repaircount, _ = strconv.Atoi(car.RepairCount)

	car.RepairCount = strconv.Itoa(repaircount + 1)

	updatedRepairAsBytes, _ := json.Marshal(car)

	APIstub.PutState(args[2], updatedRepairAsBytes)


	return shim.Success(nil)
}



func (s *SmartContract) getRepair(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	repairAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
			fmt.Println(err.Error())
	}

	repair := Repair{}
	json.Unmarshal(repairAsBytes, &repair)

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
	}
	buffer.WriteString("{\"Engineer\":")
	buffer.WriteString("\"")
	buffer.WriteString(repair.Engineer)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Date\":")
	buffer.WriteString("\"")
	buffer.WriteString(repair.Date)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Car\":")
	buffer.WriteString("\"")
	buffer.WriteString(repair.Rcar)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) getAllRepair(APIstub shim.ChaincodeStubInterface) pb.Response {

	// Find latestKey
	repairkeyAsBytes, _ := APIstub.GetState("r_latestKey")
	repairkey := RepairKey{}
	json.Unmarshal(repairkeyAsBytes, &repairkey)
	idxStr := strconv.Itoa(repairkey.R_Idx + 1)

	var startKey = "REPAIR0"
	var endKey = repairkey.R_Key + idxStr
	fmt.Println(startKey)
	fmt.Println(endKey)

	resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIter.HasNext() {
		queryResponse, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")

		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]\n")
	return shim.Success(buffer.Bytes())
}


type Insurance struct {
	Icar    string `json:"icar"`
	Turm string `json:"turm"`
}


func (s *SmartContract) setInsurance(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	var insurancekey = InsuranceKey{}
	json.Unmarshal(generateKeyInsurance(APIstub, "i_latestKey"), &insurancekey)
	keyidx := strconv.Itoa(insurancekey.I_Idx)
	fmt.Println("Key : " + insurancekey.I_Key + ", Idx : " + keyidx)

	var insurance = Insurance{Icar: args[0], Turm: args[1]}

	InsuranceasJSONBytes, _ := json.Marshal(insurance)

	var keyString = insurancekey.I_Key + keyidx
	fmt.Println("insurancekey is " + keyString)

	err := APIstub.PutState(keyString, InsuranceasJSONBytes)
	if err != nil {
			return shim.Error("Failed to create asset " + insurance.Icar)
	}

	insurancekeyAsBytes, _ := json.Marshal(insurancekey)
	APIstub.PutState("i_latestKey", insurancekeyAsBytes)

	return shim.Success(nil)

}

func (s *SmartContract) getInsurance(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	insuranceAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}

	insurance := Insurance{}
	json.Unmarshal(insuranceAsBytes, &insurance)

	var buffer bytes.Buffer
	buffer.WriteString("[") 
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}
	buffer.WriteString("{\"CARNUM\":")
	buffer.WriteString("\"")
	buffer.WriteString(insurance.Icar)
	buffer.WriteString("\"")

	buffer.WriteString(", \"TURM\":")
	buffer.WriteString("\"")
	buffer.WriteString(insurance.Turm)
	buffer.WriteString("\"")


	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())

}

func (s *SmartContract) getAllInsurance(APIstub shim.ChaincodeStubInterface) pb.Response {

	// Find latestKey
	insurancekeyAsBytes, _ := APIstub.GetState("i_latestKey")
	insurancekey := InsuranceKey{}
	json.Unmarshal(insurancekeyAsBytes, &insurancekey)
	idxStr := strconv.Itoa(insurancekey.I_Idx + 1)

	var startKey = "INSURANCE0"
	var endKey = insurancekey.I_Key + idxStr
	fmt.Println(startKey)
	fmt.Println(endKey)

	resultsIter, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false
	for resultsIter.HasNext() {
		queryResponse, err := resultsIter.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")

		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]\n")
	return shim.Success(buffer.Bytes())
}


func (s *SmartContract) setRenewal(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	insurancebytes, err := APIstub.GetState(args[0])
	if err != nil {
		return shim.Error("Could not locate insurance")
	}

	insurance := Insurance{}
    json.Unmarshal(insurancebytes, &insurance)

	insurance.Turm = args[1]
    insurancebytes, _ = json.Marshal(insurance)
    err2 := APIstub.PutState(args[0], insurancebytes)


	if err2 != nil {
        return shim.Error(fmt.Sprintf("Failed to change insurance turm: %s", args[0]))
    }
    return shim.Success(nil)

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

    var user Wallet
    json.Unmarshal(value, & user)

    // check if the password is correct
    if user.Password != args[1] {
        return shim.Error("Credentials not correct")
    }

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