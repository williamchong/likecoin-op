from pydantic import RootModel
from web3 import Web3
from web3.eth import Contract

likenft_contract_abi = [
    {
        "inputs": [],
        "name": "name",
        "outputs": [{"internalType": "string", "name": "", "type": "string"}],
        "stateMutability": "view",
        "type": "function",
    }
]


class LikeNFTName(RootModel[str]):
    pass


def get_likenft_contract(w3: Web3, likenft_contract_address: str) -> Contract:
    return w3.eth.contract(address=likenft_contract_address, abi=likenft_contract_abi)
