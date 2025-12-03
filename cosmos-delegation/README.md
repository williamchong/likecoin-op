# Cosmos Delegation

## Setup

```
cp .env.example .env
```

```
npm i
```

## Compute delegate cmd

```
npx tsx \
    --env-file .env \
    index.ts \
    compute-delegate-cmd \
    like15aef08e9skf5stlahddntsn87uuwuvqf4ngr6s \
    --selected-validator-address \
        likevaloper1knfp5qgzgr29k00nhjp2qzxlezrlalhyh77e4c \
        likevaloper1jxpfche2386a6m0kvfpj6xq9zlrjtuqwz2rnug \
        likevaloper1mztweu8y2lazpapfgtqmadxaqaapyasv7nhexk \
        likevaloper1f90nyyptfajz58m9sa9dygwajwuxkv838n72k6 \
        likevaloper1r4sv5ea8mhd7q2cp566sh5zvkwg8xf3xwgw6uw \
        likevaloper18kzunt5rj3gzqmhda5fae2zasqqj6w4vs0gk8a \
        likevaloper1x79u5q9eldwx4cp8h93p65jx3q6fzyfp48uzgz \
        likevaloper1zqe23l0vwxwlg6hx3sqmvtuw0h46evpecap636 \
        likevaloper1t3v9clrvqeudph4gf9fthcn4csalahfuhwzcfy \
        likevaloper1g2dpslkge0wmhgpdegeg0wq549syz8tjrxd4y2 \
    2>/dev/null
```

Add `--fees 2000000000nanolike` for tx fee;


Ref for community validator confirmed will stay in transition period:
oursky likevaloper1knfp5qgzgr29k00nhjp2qzxlezrlalhyh77e4c
civic liker likevaloper1jxpfche2386a6m0kvfpj6xq9zlrjtuqwz2rnug
liker.social likevaloper1mztweu8y2lazpapfgtqmadxaqaapyasv7nhexk
oldcat likevaloper1f90nyyptfajz58m9sa9dygwajwuxkv838n72k6
Yoitsu likevaloper1r4sv5ea8mhd7q2cp566sh5zvkwg8xf3xwgw6uw
Yasu likevaloper18kzunt5rj3gzqmhda5fae2zasqqj6w4vs0gk8a
UD likevaloper1x79u5q9eldwx4cp8h93p65jx3q6fzyfp48uzgz
HK startup dao likevaloper1zqe23l0vwxwlg6hx3sqmvtuw0h46evpecap636
Quebec likevaloper1t3v9clrvqeudph4gf9fthcn4csalahfuhwzcfy
Leafwind likevaloper1g2dpslkge0wmhgpdegeg0wq549syz8tjrxd4y2