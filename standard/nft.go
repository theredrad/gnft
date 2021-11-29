package standard

type NFT interface {
	ERC721
	ERC721Metadata
	ERC721TokenReceiver
	ERC721Enumerable
}
