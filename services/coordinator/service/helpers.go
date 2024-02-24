package service

// TODO: deprecate
// func CreateNewErrFromErrors(currErr error, errs []error) error {
// 	var errMsg = strings.Builder{}
// 	// todo: add count to each msg building on the previous msgs
// 	if currErr != nil {
// 		errMsg.WriteString(currErr.Error())
// 		if len(errs) > 0 {
// 			errMsg.WriteString("|")
// 		}
// 	}
// 	for i, err := range errs {
// 		errMsg.WriteString(err.Error())
// 		if i < len(errs)-1 {
// 			errMsg.WriteString("|")
// 		}
// 	}
// 	return errors.New(errMsg.String())
// }
