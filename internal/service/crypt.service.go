package service

import "context"

type CryptoService interface {
	Encrypt(ctx context.Context, plain string) (string, error)
	Decrypt(ctx context.Context, cipher string) (string, error)
}

type cryptoService struct {
}

func (s *cryptoService) Encrypt(ctx context.Context, plain string) (string, error) {
	return plain, nil
}

func (s *cryptoService) Decrypt(ctx context.Context, cipher string) (string, error) {
	return cipher, nil
}

func NewCryptoService() CryptoService {
	return &cryptoService{}
}
