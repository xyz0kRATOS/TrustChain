import { network } from 'hardhat';
import { writeFileSync } from 'node:fs';
import { join } from 'node:path';

async function main() {
  const { viem } = await network.connect('baseSepolia');

  const publicClient = await viem.getPublicClient();
  const [deployer] = await viem.getWalletClients();

  // NOTE: use 172800 (48h) before mainnet deployment.
  const minDelay = 300n;

  const timelock = await viem.deployContract('TestTimelockController', [
    minDelay,
    [deployer.account.address],
    [],
    deployer.account.address,
  ]);
  const reputationRegistry = await viem.deployContract('ReputationRegistry', []);
  const donationNFT = await viem.deployContract('DonationNFT', []);
  const campaignFactory = await viem.deployContract('CampaignFactory', [
    deployer.account.address,
    reputationRegistry.address,
    donationNFT.address,
    timelock.address,
  ]);

  const reputationWrite = await viem.getContractAt('ReputationRegistry', reputationRegistry.address);
  const nftWrite = await viem.getContractAt('DonationNFT', donationNFT.address);

  const transferRegHash = await reputationWrite.write.transferOwnership([campaignFactory.address], {
    account: deployer.account,
  });
  await publicClient.waitForTransactionReceipt({ hash: transferRegHash });

  const transferNftHash = await nftWrite.write.transferOwnership([campaignFactory.address], {
    account: deployer.account,
  });
  await publicClient.waitForTransactionReceipt({ hash: transferNftHash });

  const payload = {
    network: 'baseSepolia',
    chainId: 84532,
    deployedAt: new Date().toISOString(),
    timelockController: timelock.address,
    reputationRegistry: reputationRegistry.address,
    donationNFT: donationNFT.address,
    campaignFactory: campaignFactory.address,
    deployerWallet: deployer.account.address,
  };

  const outPath = join(process.cwd(), 'deployed-addresses.json');
  writeFileSync(outPath, JSON.stringify(payload, null, 2));

  console.log('\nDeployment complete.');
  console.log(JSON.stringify(payload, null, 2));
  console.log('\nCopy these addresses to:');
  console.log('- trustchain-web/.env.local');
  console.log('- trustchain-backend/.env');
  console.log('Then run: npm run verify:sepolia');
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
