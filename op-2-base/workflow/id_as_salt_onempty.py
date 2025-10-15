import json
import sys


def main():
    if len(sys.argv) != 2:
        print("Usage: python id_as_salt_onempty.py <nft_class_file>")
        sys.exit(1)

    nft_class_file = sys.argv[1]

    with open(nft_class_file, "r") as f:
        nft_classes = json.load(f)
        for nft_class in nft_classes:
            if not nft_class["salt"] and not nft_class["salt2"] :
                # print(nft_class["address"])
                # print(nft_class["id"])
                nft_class["salt"] = str(nft_class["id"])

    json_str = json.dumps(nft_classes, ensure_ascii=False)
    print(json_str)

if __name__ == "__main__":
    main()
