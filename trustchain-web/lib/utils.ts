import { formatUnits } from 'viem';

// ─── ETH Formatting ──────────────────────────────────────────────────────────

/**
 * Formats a wei bigint to a string with 4 decimal places.
 * Example: 1500000000000000000n → "1.5000"
 */
export function formatEther(wei: bigint): string {
  const value = formatUnits(wei, 18);
  const num = parseFloat(value);
  return num.toFixed(4);
}

/**
 * Formats a wei bigint to a string with 2 decimal places.
 * Example: 1500000000000000000n → "1.50"
 */
export function formatEtherShort(wei: bigint): string {
  const value = formatUnits(wei, 18);
  const num = parseFloat(value);
  return num.toFixed(2);
}

// ─── Address Formatting ───────────────────────────────────────────────────────

/**
 * Shortens an Ethereum address to the 0x1234...abcd format.
 * Example: "0x1234567890abcdefabcdef1234567890abcdef12" → "0x1234...ef12"
 */
export function shortenAddress(address: string): string {
  if (!address || address.length < 10) return address;
  return `${address.slice(0, 6)}...${address.slice(-4)}`;
}

// ─── Date Formatting ──────────────────────────────────────────────────────────

const DATE_FORMATTER = new Intl.DateTimeFormat('en-US', {
  month: 'short',
  day: 'numeric',
  year: 'numeric',
});

/**
 * Formats a Unix timestamp (seconds) to a human-readable date string.
 * Example: 1737000000 → "Jan 16, 2025"
 */
export function formatDate(timestamp: number): string {
  // Support both seconds and milliseconds
  const ms = timestamp > 1e12 ? timestamp : timestamp * 1000;
  return DATE_FORMATTER.format(new Date(ms));
}

/**
 * Returns a relative time string from a Unix timestamp (seconds).
 * Example: (timestamp 5 minutes ago) → "5 minutes ago"
 */
export function formatTimeAgo(timestamp: number): string {
  const ms = timestamp > 1e12 ? timestamp : timestamp * 1000;
  const now = Date.now();
  const diffMs = now - ms;
  const diffSeconds = Math.floor(diffMs / 1000);

  if (diffSeconds < 5) return 'just now';
  if (diffSeconds < 60) return `${diffSeconds} seconds ago`;

  const diffMinutes = Math.floor(diffSeconds / 60);
  if (diffMinutes < 60) {
    return diffMinutes === 1 ? '1 minute ago' : `${diffMinutes} minutes ago`;
  }

  const diffHours = Math.floor(diffMinutes / 60);
  if (diffHours < 24) {
    return diffHours === 1 ? '1 hour ago' : `${diffHours} hours ago`;
  }

  const diffDays = Math.floor(diffHours / 24);
  if (diffDays < 7) {
    return diffDays === 1 ? '1 day ago' : `${diffDays} days ago`;
  }

  const diffWeeks = Math.floor(diffDays / 7);
  if (diffWeeks < 5) {
    return diffWeeks === 1 ? '1 week ago' : `${diffWeeks} weeks ago`;
  }

  const diffMonths = Math.floor(diffDays / 30);
  if (diffMonths < 12) {
    return diffMonths === 1 ? '1 month ago' : `${diffMonths} months ago`;
  }

  const diffYears = Math.floor(diffDays / 365);
  return diffYears === 1 ? '1 year ago' : `${diffYears} years ago`;
}

/**
 * Returns a countdown string from a future Unix timestamp (seconds).
 * Example: (timestamp 47h 23m from now) → "47h 23m remaining"
 * Returns "Expired" if the timestamp is in the past.
 */
export function formatCountdown(futureTimestamp: number): string {
  const ms = futureTimestamp > 1e12 ? futureTimestamp : futureTimestamp * 1000;
  const now = Date.now();
  const diffMs = ms - now;

  if (diffMs <= 0) return 'Expired';

  const totalSeconds = Math.floor(diffMs / 1000);
  const totalMinutes = Math.floor(totalSeconds / 60);
  const totalHours = Math.floor(totalMinutes / 60);
  const totalDays = Math.floor(totalHours / 24);

  if (totalDays > 0) {
    const hours = totalHours % 24;
    const mins = totalMinutes % 60;
    if (hours === 0) return `${totalDays}d remaining`;
    if (mins === 0) return `${totalDays}d ${hours}h remaining`;
    return `${totalDays}d ${hours}h ${mins}m remaining`;
  }

  if (totalHours > 0) {
    const mins = totalMinutes % 60;
    if (mins === 0) return `${totalHours}h remaining`;
    return `${totalHours}h ${mins}m remaining`;
  }

  if (totalMinutes > 0) {
    const secs = totalSeconds % 60;
    if (secs === 0) return `${totalMinutes}m remaining`;
    return `${totalMinutes}m ${secs}s remaining`;
  }

  return `${totalSeconds}s remaining`;
}

// ─── Basescan URLs ───────────────────────────────────────────────────────────

type BasescanLinkType = 'tx' | 'address' | 'token';

const MAINNET_CHAIN_ID = '8453';
const MAINNET_BASE_URL = 'https://basescan.org';
const TESTNET_BASE_URL = 'https://sepolia.basescan.org';

/**
 * Returns a Basescan URL for a transaction, address, or token.
 * Automatically selects mainnet vs testnet based on NEXT_PUBLIC_CHAIN_ID.
 */
export function getBasescanUrl(type: BasescanLinkType, value: string): string {
  const chainId =
    typeof process !== 'undefined'
      ? process.env.NEXT_PUBLIC_CHAIN_ID
      : undefined;

  const baseUrl = chainId === MAINNET_CHAIN_ID ? MAINNET_BASE_URL : TESTNET_BASE_URL;

  const pathMap: Record<BasescanLinkType, string> = {
    tx: 'tx',
    address: 'address',
    token: 'token',
  };

  return `${baseUrl}/${pathMap[type]}/${value}`;
}

// ─── Misc Utilities ──────────────────────────────────────────────────────────

/**
 * Classnames helper — filters falsy values and joins with space.
 * Useful alternative to `clsx` without an extra dependency.
 */
export function cn(...classes: (string | undefined | null | false)[]): string {
  return classes.filter(Boolean).join(' ');
}
