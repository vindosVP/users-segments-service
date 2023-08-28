package clients

import (
	"fmt"
	"github.com/TwiN/go-pastebin"
)

type PastebinClient struct {
	cl *pastebin.Client
}

func NewPastebinClient(login string, password string, token string) (*PastebinClient, error) {
	client, err := pastebin.NewClient(login, password, token)
	if err != nil {
		return nil, err
	}
	return &PastebinClient{
		cl: client,
	}, nil
}

func (pc *PastebinClient) Post(title string, content string) (pasteLink string, err error) {
	req := pastebin.NewCreatePasteRequest(title, content, pastebin.ExpirationTenMinutes, pastebin.VisibilityUnlisted, "")
	pasteKey, err := pc.cl.CreatePaste(req)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://pastebin.com/%s", pasteKey), nil
}
