import json
import sys

def replace_salt(obj, new_salt):
    obj["salt"] = new_salt

def main():
    if len(sys.argv) != 3:
        print("Usage: python replace_salt.py <nft_class_file> <duplicated_new_addresses_file>")
        sys.exit(1)

    nft_class_file = sys.argv[1]
    duplicated_new_addresses_file = sys.argv[2]

    with open(nft_class_file, "r") as f:
        nft_classes = json.load(f)

    with open(duplicated_new_addresses_file, "r") as f:
        duplicated_new_addresses = json.load(f)

    for duplicated_new_address in duplicated_new_addresses:
        for address in duplicated_new_address["addresses"]:
            for nft_class in nft_classes:
                if nft_class["address"] == address["old_address"]:
                    replace_salt(nft_class, str(nft_class["id"]))

    json_str = json.dumps(nft_classes, ensure_ascii=False)
    print(json_str)

if __name__ == "__main__":
    main()
