# Deployment Notes

We rely on operation to maintains the initial contract address of the evm smart contract.

Operator keys are stored in the `env.operator` file. [Blackbox](https://github.com/StackExchange/blackbox) encrypts the keys.

To decrypt the keys, run `blackbox_decrypt_all_files`, then `make operator-key-link` at the root of the repo.

## Expected procedure

Current encrypted operator wallet: `0x6bE6f226D024b3B7022d18978F41300b96263C1b`

Once the initial deployment is done, we should updated the expected proxy address in the `env.operator` file for reference. On new L2 chain launch, we should operate the following steps to obtain same proxy address.

0. Checkout a specific commit
1. Someone sends ETH to the operator's wallet
2. Use operator's wallet to run the deployment on `likecoin` and `likenft` (!! Order matters)
   - 2.1. `likecoin` should run first, with itself as the owner and minter
   - 2.2. `likenft` come second, itself as the owner
3. Upgrade the implementation to latest implementation
4. Update the owner and minter accordingly

Formula refs:
- newAddressCREATE = `keccak256(deployingAddress ++ nonce)[12:]`
- newAddressCREATE2 = `keccak256(0xff ++ deployingAddress ++ salt ++ initCodeHash)[12:]`