# Deployment Notes

We rely on operation to maintains the initial contract address of the evm smart contract.

Operator keys are stored in the `env.operator` file. [Blackbox](https://github.com/StackExchange/blackbox) encrypts the keys.

To decrypt the keys, run `blackbox_decrypt_all_files`, then `make operator-key-link` at the root of the repo.

## Expected procedure

Current encrypted operator wallet: `0xC71fe89e4C0e5458a793fc6548EF6B392417A7Fb`

EkilCoin proxy: `0x6B65396685496035Ce0B03Cbec9b9963793B08a9`
LikeNFT proxy: `0xfF79df388742f248c61A633938710559c61faEF1`

Once the initial deployment is done, we should updated the expected proxy address in the `env.operator` file for reference. On new L2 chain launch, we should operate the following steps to obtain same proxy address.

0. Checkout a specific commit
1. Someone sends ETH to the operator's wallet (!!failed operation becasue of lack of fund will make the address different from the expected one)
2. Use operator's wallet to run the deployment on `likecoin` and `likenft` (!! Order matters)
   - 2.1. `likecoin` should run first, with itself as the owner and minter
   - 2.2. `likenft` come second, itself as the owner
3. Upgrade the implementation to latest implementation
4. Update the owner and minter accordingly

Formula refs:
- newAddressCREATE = `keccak256(deployingAddress ++ nonce)[12:]`
- newAddressCREATE2 = `keccak256(0xff ++ deployingAddress ++ salt ++ initCodeHash)[12:]`