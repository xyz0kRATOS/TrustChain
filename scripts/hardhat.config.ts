networks: {
  hardhat: {},          // local testing — always start here
  baseSepolia: {        // testnet — deploy here before mainnet
    url: "https://sepolia.base.org",
    accounts: [process.env.DEPLOYER_PRIVATE_KEY]
  },
  base: {               // mainnet — touch this last
    url: "https://mainnet.base.org",
    accounts: [process.env.DEPLOYER_PRIVATE_KEY]
  }
}