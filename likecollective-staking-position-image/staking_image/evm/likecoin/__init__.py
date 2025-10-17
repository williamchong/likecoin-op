from pydantic import RootModel
from web3 import Web3
from web3.eth import Contract

likecoin_contract_abi = [
    {
        "inputs": [],
        "name": "decimals",
        "outputs": [{"internalType": "uint8", "name": "", "type": "uint8"}],
        "stateMutability": "view",
        "type": "function",
    },
]


class LikeCoinDecimals(RootModel[int]):
    pass


def get_likecoin_contract(w3: Web3, likecoin_contract_address: str) -> Contract:
    return w3.eth.contract(address=likecoin_contract_address, abi=likecoin_contract_abi)
