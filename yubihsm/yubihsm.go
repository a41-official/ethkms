package yubihsm

import (
	"errors"

	"github.com/certusone/yubihsm-go"
	"github.com/certusone/yubihsm-go/commands"
	"github.com/certusone/yubihsm-go/connector"
)

var (
	sm *yubihsm.SessionManager
)

func Init(authKeyID uint16, password string) {
	c := connector.NewHTTPConnector("localhost:12345")
	sessionManager, err := yubihsm.NewSessionManager(c, authKeyID, password)
	if err != nil {
		panic(err)
	}
	sm = sessionManager
}

func GetOpaque(objID uint16) (*commands.GetOpaqueResponse, error) {
	command, err := commands.CreateGetOpaqueCommand(objID)
	if err != nil {
		return nil, err
	}
	resp, err := sm.SendEncryptedCommand(command)
	if err != nil {
		return nil, err
	}
	parsedResp, matched := resp.(*commands.GetOpaqueResponse)
	if !matched {
		panic("invalid response type")
	}
	return parsedResp, nil
}

func PutOpaque(objID uint16, label []byte, domains uint16, capabilities uint64, algorithm commands.Algorithm, data []byte) (*commands.PutOpaqueResponse, error) {
	command, err := commands.CreatePutOpaqueCommand(objID, label, domains, capabilities, algorithm, data)
	if err != nil {
		return nil, err
	}
	resp, err := sm.SendEncryptedCommand(command)
	if err != nil {
		return nil, err
	}
	parsedResp, matched := resp.(*commands.PutOpaqueResponse)
	if !matched {
		return nil, errors.New("invalid response type")
	}
	return parsedResp, nil
}
