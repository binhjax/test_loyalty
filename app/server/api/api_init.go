package api
//
// func (fw *WalletHandler) RegisterBatchEthToContract(requestTime int64) []string {
//     ret := []string{}
//     list := fw.GetAccountList()
//     sublist :=  []common.Address{}
//     for i,item := range list {
//       sublist = append(sublist,item)
//       if i > 0 && i % 5 == 0 {
//         if len(sublist) > 0 {
//           fmt.Println("Start register sublist: ", i/5)
//           tx,err := fw.RegisterAccETH(requestTime,sublist)
//           if err != nil {
//              ret = append(ret, err.Error())
//           } 	else {
//              ret = append(ret, tx.Hash().Hex())
//           }
//           sublist = []common.Address{}
//         }
//       }
//     }
//     if len(sublist) > 0 {
//       fmt.Println("Start register last sublist")
//       tx,err := fw.RegisterAccETH(requestTime,sublist)
//       if err != nil {
//          ret = append(ret, err.Error())
//       } 	else {
//          ret = append(ret, tx.Hash().Hex())
//       }
//     }
//     return ret
// }
