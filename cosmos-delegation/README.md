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
    like1cerk8zu2ehsxsf6namnu84z9z58dlwurztla9l \
    --selected-validator-address \
        likevaloper17yqg3way0slmf4lj8vtmhlcd73006hhpljd5d0 \
        likevaloper1pcw47y66m5clzn9cxx5lq0lrqnwzkeq86krsfj \
        likevaloper1etpln3zmzk2qv834qy7llgw3t98rztj7e3f7cq \
        likevaloper1dl4s8xxpmt0u6yrrlh5tnmrlgsq0t70uwzyar2 \
        likevaloper10qz95ywphsmq8dpgk8xvfpehvnd0y5kc98x484 \
        likevaloper1c6ch7qx2gvluc9uw85qlzx949z8ckyhud7hm6r \
    2>/dev/null
```
