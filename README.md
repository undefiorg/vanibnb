# VaniBNB
A Binance Chain vanity address generator written in golang.
For example address ending with `0xkat`

![image](https://user-images.githubusercontent.com/97060/115947168-0cdf5000-a4f0-11eb-8fc6-8b4bdfbc9938.png)

## Raw
- https://github.com/makevoid/vanieth
- https://github.com/octofoxio/stress
- https://github.com/binance-chain/go-sdk

## Installation
```
$ go get -u github.com/undefiorg/vanibnb
```

## Usage
### Parameter
- `-s`: Suffix to find
- `-n`: Network mainnet or testnet

## Example run:
Run
```
$ vanibnb -s cat -n testnet
```
Output
```
🔍 Finding Address end with cat on testnet
🌀 Running task no. 1 
🌀 Running task no. 2 
🌀 Running task no. 3 
🌀 Running task no. 4 
🌀 Running task no. 5 
🌀 Running task no. 6 
🌀 Running task no. 7 
🌀 Running task no. 8 
🔥 Running at 54616 hash per second
💎 ʕ•́ᴥ•̀ʔっ Got it! 
🔒 public address : tbnb1ynmj2u00dsuc7urwulam0sd0qf8vfdxrqspcat
🔑 private key    : 14e28edda2884969e6258ace734f9fc2e4f494e8abc7be5e3ada7dd10da6072d
```

## Support
That's it! If you want to give me a cup of coffee ☕️ here's my wallet. 👇  

```
ETH      : 0xD63eCF4Fa7ef94ea4E9CF7429Af6a4F482B447B7
BNB(bsc) : 0x0c62B13f8116eBA444DeFdE842aBd7d1f68F964f
BNB(bc)  : bnb1uqht090p2fjs9dr2umh5w5ykm697srgr70xkat
```
