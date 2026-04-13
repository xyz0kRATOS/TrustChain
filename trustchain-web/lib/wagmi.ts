import { getDefaultConfig } from '@rainbow-me/rainbowkit';
import { base, baseSepolia } from 'wagmi/chains';

// ─── Chain Selection ──────────────────────────────────────────────────────────
// Determined at runtime from NEXT_PUBLIC_CHAIN_ID env var.
// 8453 → Base Mainnet | 84532 → Base Sepolia (default / dev)

const chainId = Number(process.env.NEXT_PUBLIC_CHAIN_ID ?? '84532');
const isMainnet = chainId === 8453;

export const wagmiConfig = getDefaultConfig({
  appName: 'TrustChain',
  projectId: process.env.NEXT_PUBLIC_WALLET_CONNECT_ID ?? '',
  chains: isMainnet ? [base, baseSepolia] : [baseSepolia, base],
  ssr: true, // Next.js App Router requires SSR: true
});

// Export chain constants for use elsewhere
export { base, baseSepolia };
export const primaryChain = isMainnet ? base : baseSepolia;
