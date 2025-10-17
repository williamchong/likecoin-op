from web3 import Web3

from staking_image.evm.booknft import LikeNFTName, get_likenft_contract
from staking_image.evm.likecoin import LikeCoinDecimals, get_likecoin_contract
from staking_image.utils.methodtools import lru_cache


class EVMClient:
    def __init__(
        self,
        rpc_url: str,
        likecoin_contract_address: str,
    ):
        self.w3 = Web3(Web3.HTTPProvider(rpc_url))
        self.likecoin_contract = get_likecoin_contract(
            self.w3, likecoin_contract_address
        )

    @lru_cache(maxsize=1)
    def get_likecoin_decimals(self) -> int:
        return LikeCoinDecimals.model_validate(
            self.likecoin_contract.functions.decimals().call()
        ).root

    @lru_cache(maxsize=1000)
    def get_likenft_name(self, contract_address: str) -> str:
        likenft_contract = get_likenft_contract(
            self.w3, Web3.to_checksum_address(contract_address)
        )
        return LikeNFTName.model_validate(likenft_contract.functions.name().call()).root
