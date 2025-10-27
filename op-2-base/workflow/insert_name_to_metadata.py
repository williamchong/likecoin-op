import json
import sys


def main():
    if len(sys.argv) != 2:
        print("Usage: python insert_name_to_metadata.py <nft_class_file>")
        sys.exit(1)

    nft_class_file = sys.argv[1]

    with open(nft_class_file, "r") as f:
        nft_classes = json.load(f)
        for nft_class in nft_classes:
            metadata = nft_class["metadata"]
            name = nft_class["name"]
            if "name" not in metadata:
                metadata["name"] = name
    json_str = json.dumps(nft_classes, ensure_ascii=False)
    print(json_str)

if __name__ == "__main__":
    main()