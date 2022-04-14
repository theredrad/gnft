package gnft

type NFT721 interface {
	ERC721
	ERC721Metadata
	ERC721TokenReceiver
	ERC721Enumerable
}
