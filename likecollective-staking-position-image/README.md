# Staking Image

## Start development server

```bash
make dev
```

## Add dependencies

```bash
make poetry add {pkg_name}
```

## Local dev

To test, use browsers to visit the following url with desire params,

```
http://localhost:8000/?book_nft_address=0x2fDDD0872563d10317f38ECE3dd0900EF1e21536&staked_amount=123&reward_index=234&initial_staker=0x146736dd2Ccb1dbDF266130a117609E00ad566b1
```

To test with base64 encoded abi packed request,

```
http://localhost:8000/49tc2YXYaaGMAAy01RGAsctFDowAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAcJl5cMUYEtw6AQx9AbUODRfcecg=
```

