# likenft-migration

## Build Setup

```bash
# install dependencies
$ npm install

# serve with hot reload at localhost:3000
$ npm run dev

# build for production and launch server
$ npm run build
$ npm run start

# generate static project
$ npm run generate
```

## Program flow

```mermaid
sequenceDiagram 
    Browser->>Browser: Get user's cosmos address and/or LikerID
    Browser->>Like.co: Get user's full LikerID info, by LikerID or cosmos address
    Browser->>Browser: Get user's evm address
    Browser->>Backend: Send the migration info, LikerID and evm address to backend
    Backend->>Browser: Request user using cosmos wallet to signin
    Browser->>Backend: Send the signed cosmos message to backend
    Backend->>Backend: Verify the cosmos message
    Backend->>Like.co: Pass on the signed cosmos message for LikerID migration
    Backend->>Browser: Return the token for trigger LikeNFT drop on OP
    Backend->>Browser: (Optional) Request user using EVM wallet do a personal sign, for verifying wallet ownership
    Browser->>Cosmos: Querying the user's LikeNFTClass and LikeNFT
    Browser->>Backend: Send the NFT migration info
    Backend->>Cosmos: Verify the NFT migration info is valid.<br> (NFT ownership, not migrated before)
    Backend->>EVM: Mint the NFT on OP
    Backend->>Browser: Return the EVM TX hashes
    Browser->>Browser: Show the migration result
```

## Note
- How backend mint the NFT on OP is not a single TX, because it can be involve minting NFTClass. Details please refer to the [migration-backend](./migration-backend/README.md).

## Frameworks

For detailed explanation on how things work, check out the [documentation](https://nuxtjs.org).

## Special Directories

You can create the following extra directories, some of which have special behaviors. Only `pages` is required; you can delete them if you don't want to use their functionality.

### `assets`

The assets directory contains your uncompiled assets such as Stylus or Sass files, images, or fonts.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/assets).

### `components`

The components directory contains your Vue.js components. Components make up the different parts of your page and can be reused and imported into your pages, layouts and even other components.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/components).

### `layouts`

Layouts are a great help when you want to change the look and feel of your Nuxt app, whether you want to include a sidebar or have distinct layouts for mobile and desktop.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/layouts).

### `pages`

This directory contains your application views and routes. Nuxt will read all the `*.vue` files inside this directory and setup Vue Router automatically.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/get-started/routing).

### `plugins`

The plugins directory contains JavaScript plugins that you want to run before instantiating the root Vue.js Application. This is the place to add Vue plugins and to inject functions or constants. Every time you need to use `Vue.use()`, you should create a file in `plugins/` and add its path to plugins in `nuxt.config.js`.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/plugins).

### `static`

This directory contains your static files. Each file inside this directory is mapped to `/`.

Example: `/static/robots.txt` is mapped as `/robots.txt`.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/static).

### `store`

This directory contains your Vuex store files. Creating a file in this directory automatically activates Vuex.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/store).
