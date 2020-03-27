package contracts
import (
    "github.com/binhnt-teko/test_loyalty/app/server/account"
    "github.com/binhnt-teko/test_loyalty/app/server/connection"
    "github.com/binhnt-teko/test_loyalty/app/server/config"
    // "math/big"
    // "strings"
    "fmt"
    // "encoding/json"
    "errors"
    "strings"
    // "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/common"
    // "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "encoding/hex"
    // "sync"
    // "math"
    // "strconv"
    // "bytes"
    "context"
      "io/ioutil"
    "gopkg.in/yaml.v2"
    "math/big"
)


type  Contract struct {
    Version string `yaml:"version" json:"version"`
    Released config.SpecialDate `yaml:"released" json:"released"`
    Update bool `yaml:"update" json:"update"`
    GasPrice string `yaml:"gasprice" json:"gasprice"`
    GasLimitDefault uint64 `yaml:"gaslimitdefault" json:"gaslimitdefault"`
    GasLimit map[string]uint64 `yaml:"gaslimit" json:"gaslimit"`
    Owner string `yaml:"owner" json:"owner"`
    Address string `yaml:"address" json:"address"`
    Content struct {
        ContractABI string `yaml:"contractABI" json:"contractABI"`
        ContractBIN string `yaml:"contractBIN" json:"contractBIN"`
    } `yaml:"content" json:"content"`
}

var LContract *Contract
var Loaded bool = false

func Init() {
    LContract = &Contract{}
}

func (fw *Contract) ContractInstance(address string , backend bind.ContractBackend) (*Contracts, error) {
     fmt.Println("Load contract instance from address: ", address)
     addr := common.HexToAddress(address)
     return NewContracts(addr, backend)
}

func (fw *Contract) Deploy() (string,string, error) {
      wallet := account.Pool.GetWallet(fw.Owner)
      if wallet == nil {
         return "","", errors.New("Cannot find owner")
      }
      conn := connection.Pool.GetConnection()

      auth := wallet.NewTransactor()

      gasPrice := new(big.Int)
      gasPrice.SetString(fw.GasPrice,10)
      auth.GasPrice = gasPrice

      gasLimit,ok := fw.GasLimit["deploy"]
      if !ok {
        gasLimit = fw.GasLimitDefault
      }
      fmt.Println("Set GasLimit: ", gasLimit)
      auth.GasLimit = gasLimit

      // parsed, err := abi.JSON(strings.NewReader(fw.Content.Contract))
      parsed, err := abi.JSON(strings.NewReader(ContractsABI))
      if err != nil {
         fmt.Println("Cannot parse contract ABI", err)
         return "","",err
      }
      // bytecode := common.FromHex(strings.TrimRight(fw.Content.ContractBIN, "\r\n"))
      bytecode := common.FromHex(ContractsBin)
      fmt.Println("Bytecode:",bytecode)
      // bytecode := []byte(fw.Content.ContractBIN)

      address, tx, _, err := bind.DeployContract(auth, parsed, bytecode, conn.Client)
      if err != nil {
        fmt.Println("Cannot deploy contract ", err)
        return "","", err
      }
      fmt.Println("Transaction: ", tx.Hash())
      fmt.Println("Contract address deploy:", address.Hex())
      fw.Address = address.Hex()
      return fw.Address, tx.Hash().Hex() , nil
}

func (fw *Contract) Save(config_file string) error {
    newcfg, err := yaml.Marshal(LContract)
    if err != nil {
        fmt.Println("yaml.Marshal error: %v", err)
        return err
    }
    err1 := ioutil.WriteFile(config_file, newcfg, 0644)
    if err1 != nil {
      fmt.Println("Write file error:",err1)
      return err1
    }
    return nil
}

func (fw *Contract) Load(file string) error {
    var contract = &Contract{}

    yamlFile, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println("yamlFile.Get err: ", err)
        return err
    }

    err = yaml.Unmarshal(yamlFile, contract)
    if err != nil {
       fmt.Println("Unmarshal error: ", err)
        return err
    }
    fmt.Println("Load config: ",contract.Version, ", Released: ", contract.Released)
    LContract = contract
    Loaded = true
    return nil
}

func (fw *Contract)  CodeAt() (string, error) {
      conn := connection.Pool.GetConnection()
      fmt.Println("Current contract address: ", LContract.Address)
      contractAddress := common.HexToAddress(LContract.Address)
      bytecode, err := conn.Client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
      if err != nil {
          fmt.Println("Error in get bytecode: ", err)
          return "", err
      }
      if len(bytecode) > 0 {
           fmt.Println("Data:",bytecode)
           data := hex.EncodeToString(bytecode)
           // fmt.Println(data) // 60806...10029
           return data, nil
      } else {
           fmt.Println("is account, not smartcontract")
      }
      return "", nil
}
