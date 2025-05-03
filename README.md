# asnbcli

A cli app to simplify buying asnb funds so you (hopefully) don't have to wake up at 2 am. The cli takes your username and password and tries to buy asm1, asm2 and asm3 funds, if successful it will return a payment link so you can pay directly. You can use this with a cron & a cli app that texts yourself (maybe coming soon I'm interested in making something like that)

If you would like to test out if it works, you could test it with ASN funds, as those can be bought directly. ASM still requires a bit of luck.

## READ FIRST

This is your reminder to not download code from the internet and run it willy nilly. 

I am not responsible for anything that happens if you use this code, PLEASE AUDIT THE CODE YOURSELF BEFORE USING IT.

I WILL NEVER RELEASE A BINARY OF IT.

You should NEVER use any distributed binaries that claim to be from this repo.

If you know how to run this, you know how to run this.

## DISCLAIMER

I am in no way affiliated with Asnb, I am just frustruated that I have to stay up to 2 in the morning to buy these funds.
All code and requests used in this repo can be retrieved from publically available sources.

## Usage

Flags:
```
  -a, --amount string           Amount to buy
      --fpx-bank string         Fpx bank to use (ex. HLB0224)
  -f, --funds strings           The funds to try, defaults to ASM1, ASM2 and ASM3 (default [asm1,asm2,asm3])
  -h, --help                    help for asnbcli
  -p, --password string         Password for your account
  -m, --payment-method string   Payment method to use, accepted values: tngd,boost,fpx
  -u, --username string         Username for your account
```

Expected output: the tool should return a payment link that you can copy / open to pay directly
