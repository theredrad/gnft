package geth

import (
	"math/big"
	"testing"
)

const cfeth = "https://cloudflare-eth.com/"

func TestGETH_NewERC721Contract(t *testing.T) {
	type fields struct {
		rawURL string
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ERC721Contract
		wantErr bool
	}{
		{
			name:   "new contract",
			fields: fields{rawURL: "https://cloudflare-eth.com/"},
			args: args{
				address: "0xba033D82c64DD514B184e2d1405cD395dfE6e706",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := New(tt.fields.rawURL)
			if err != nil {
				t.Fatalf("failed to create geth client: %v", err)
			}
			got, err := NewERC721Contract(c, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewContract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("NewContract() got = %v", got)
			}
		})
	}
}

func TestERC721Contract_BalanceOf(t *testing.T) {
	type fields struct {
		contract string
	}
	type args struct {
		address string
	}

	geth, err := New("https://cloudflare-eth.com/")
	if err != nil {
		t.Fatalf("failed to create geth client: %v", err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "balance",
			fields: fields{
				contract: "0xe1d8dda05beb37603a608cc21dee6fd92a0a727e",
			},
			args: args{
				address: "0x2473e4f5111c11c63680977581ae11be4fd6982b",
			},
			want:    big.NewInt(3),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewERC721Contract(geth, tt.fields.contract)
			if err != nil {
				t.Fatalf("failed to create the contract: %v", err)
			}
			got, err := c.BalanceOf(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("BalanceOf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(tt.want) != 0 {
				t.Errorf("BalanceOf() got = %v, want %v", got.Uint64(), tt.want.Uint64())
			}
		})
	}
}

func TestERC721Contract_OwnerOf(t *testing.T) {
	type fields struct {
		rawURL   string
		contract string
	}
	type args struct {
		tokenID *big.Int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "token owner",
			fields: fields{
				rawURL:   cfeth,
				contract: "0xba033D82c64DD514B184e2d1405cD395dfE6e706",
			},
			args: args{
				tokenID: big.NewInt(2826),
			},
			want:    "0xb9F0111abecC02d3d42541D6D2957Af2BE96DeA8",
			wantErr: false,
		},
		/*{
			name: "token owner 1155",
			fields: fields{
				contract: "0x76be3b62873462d2142405439777e971754e8e77",
				abi:      ABIERC1155,
			},
			args: args{
				tokenID: big.NewInt(10632),
			},
			want:    "0x5c01B7C87c09f85a63F69772e5C381EB01a0924C",
			wantErr: false,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			geth, err := New(tt.fields.rawURL)
			if err != nil {
				t.Fatalf("failed to create geth client: %v", err)
			}
			c, err := NewERC721Contract(geth, tt.fields.contract)
			if err != nil {
				t.Fatalf("failed to create the contract: %v", err)
			}
			got, err := c.OwnerOf(tt.args.tokenID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OwnerOf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OwnerOf() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestERC721Contract_Name(t *testing.T) {
	type fields struct {
		rawURL   string
		contract string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "OnChainMask",
			fields: fields{
				rawURL:   cfeth,
				contract: "0xB4B55Cb5C7d3C59c69DcAAc83eD067BB3AbfA2D3",
			},
			want:    "OnChainMask",
			wantErr: false,
		},
		{
			name: "not found",
			fields: fields{
				rawURL:   cfeth,
				contract: "somecontract",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			geth, err := New(tt.fields.rawURL)
			if err != nil {
				t.Fatalf("failed to create geth: %v", err)
			}
			c, err := NewERC721Contract(geth, tt.fields.contract)
			if err != nil {
				t.Fatalf("failed to create ERC721 contract: %v", err)
			}
			got, err := c.Name()
			if (err != nil) != tt.wantErr {
				t.Errorf("Name() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Name() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestERC721Contract_Symbol(t *testing.T) {
	type fields struct {
		rawURL   string
		contract string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "OnChainMask",
			fields: fields{
				rawURL:   cfeth,
				contract: "0xB4B55Cb5C7d3C59c69DcAAc83eD067BB3AbfA2D3",
			},
			want:    "OnChainMask",
			wantErr: false,
		},
		{
			name: "not found",
			fields: fields{
				rawURL:   cfeth,
				contract: "somecontract",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			geth, err := New(tt.fields.rawURL)
			if err != nil {
				t.Fatalf("failed to create geth: %v", err)
			}
			c, err := NewERC721Contract(geth, tt.fields.contract)
			if err != nil {
				t.Fatalf("failed to create ERC721 contract: %v", err)
			}
			got, err := c.Symbol()
			if (err != nil) != tt.wantErr {
				t.Errorf("Symbol() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Symbol() got = %v, want %v", got, tt.want)
			}
		})
	}
}
