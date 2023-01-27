package coin

type CoinCommandHandler struct {
	repository CoinRepository
}

func NewCoinCommandHandler() *CoinCommandHandler {
	return &CoinCommandHandler{}
}
