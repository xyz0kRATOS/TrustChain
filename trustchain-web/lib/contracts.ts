import { CAMPAIGN_FACTORY_ABI } from '@/lib/abis/CampaignFactory.abi';
import { CAMPAIGN_ABI } from '@/lib/abis/Campaign.abi';
import { REPUTATION_REGISTRY_ABI } from '@/lib/abis/ReputationRegistry.abi';
import { DONATION_NFT_ABI } from '@/lib/abis/DonationNFT.abi';
import deployedAddresses from '../../trustchain-contracts/deployed-addresses.json';

export const CONTRACTS = {
  campaignFactory: {
    address: (
      process.env.NEXT_PUBLIC_FACTORY_ADDRESS || deployedAddresses.campaignFactory
    ) as `0x${string}`,
    abi: CAMPAIGN_FACTORY_ABI,
  },
  reputationRegistry: {
    address: (
      process.env.NEXT_PUBLIC_REGISTRY_ADDRESS || deployedAddresses.reputationRegistry
    ) as `0x${string}`,
    abi: REPUTATION_REGISTRY_ABI,
  },
  donationNFT: {
    address: (process.env.NEXT_PUBLIC_NFT_ADDRESS || deployedAddresses.donationNFT) as `0x${string}`,
    abi: DONATION_NFT_ABI,
  },
} as const;

export { CAMPAIGN_ABI };
