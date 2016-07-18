// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)

// UserRecoverProceed mapping
type UserRecoverProceed struct {
  Password  *string  `json:"password,omitempty"`
}


// GetRecoveryDetails gets details on a recovery keypair. Useful to check validity of recovery keypair.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-recover-get
func (service *UserService) GetRecoveryDetails(userID string, recoverIdentifier string, recoverKey string) (*Response, error) {
  url := fmt.Sprintf("user/%s/account/recover/%s/%s", userID, recoverIdentifier, recoverKey)
  req, _ := service.client.NewRequest("GET", url, nil)

  return service.client.Do(req, nil)
}


// SendRecoveryPassword submits new password and recover account.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-recover-put
func (service *UserService) SendRecoveryPassword(userID string, recoverIdentifier string, recoverKey string, password string) (*Response, error) {
  url := fmt.Sprintf("user/%s/account/recover/%s/%s", userID, recoverIdentifier, recoverKey)
  req, _ := service.client.NewRequest("PUT", url, UserRecoverProceed{Password: &password})

  return service.client.Do(req, nil)
}


// DeleteRecoveryKeypair deletes a recovery keypair. Useful to invalidate keys if you ignore recovery and never use the keys to recover password.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-recover-delete
func (service *UserService) DeleteRecoveryKeypair(userID string, recoverIdentifier string, recoverKey string) (*Response, error) {
  url := fmt.Sprintf("user/%s/account/recover/%s/%s", userID, recoverIdentifier, recoverKey)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
