'use client';

import React from 'react';
import { ConnectButton } from '@rainbow-me/rainbowkit';
import { useAccount, useBalance, useSwitchChain } from 'wagmi';
import { primaryChain } from '@/lib/wagmi';
import { formatEtherShort, shortenAddress } from '@/lib/utils';

// ─── WalletButton ─────────────────────────────────────────────────────────────

export function WalletButton() {
  const { address, isConnected, chain } = useAccount();
  const { data: balance } = useBalance({ address, chainId: primaryChain.id });
  const { switchChain } = useSwitchChain();

  const isWrongNetwork = isConnected && chain?.id !== primaryChain.id;

  return (
    <ConnectButton.Custom>
      {({ openAccountModal, openConnectModal, openChainModal, mounted }) => {
        if (!mounted) {
          return (
            <div
              aria-hidden="true"
              style={{ opacity: 0, pointerEvents: 'none', userSelect: 'none' }}
            />
          );
        }

        // ── Wrong Network ──────────────────────────────────────────────────
        if (isWrongNetwork) {
          return (
            <button
              id="wallet-wrong-network"
              onClick={() => {
                try {
                  switchChain({ chainId: primaryChain.id });
                } catch {
                  openChainModal();
                }
              }}
              className={[
                'inline-flex items-center gap-2',
                'h-9 px-4 rounded-lg',
                'border border-[#DC2626]/50',
                'bg-[rgba(220,38,38,0.08)]',
                'text-[#F87171] text-sm font-medium',
                'hover:bg-[rgba(220,38,38,0.15)] hover:border-[#DC2626]',
                'transition-all duration-200',
                'cursor-pointer',
              ].join(' ')}
            >
              {/* Warning icon */}
              <svg
                className="w-3.5 h-3.5 flex-shrink-0"
                viewBox="0 0 16 16"
                fill="none"
                aria-hidden="true"
              >
                <path
                  d="M8 1.5L14.5 13H1.5L8 1.5Z"
                  stroke="currentColor"
                  strokeWidth="1.5"
                  strokeLinejoin="round"
                />
                <path d="M8 6v3" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" />
                <circle cx="8" cy="10.5" r="0.75" fill="currentColor" />
              </svg>
              Wrong Network
              <span className="text-[#DC2626]/70 text-xs">· Switch</span>
            </button>
          );
        }

        // ── Connected ──────────────────────────────────────────────────────
        if (isConnected && address) {
          const ethBalance = balance
            ? `${formatEtherShort(balance.value)} ETH`
            : '…';

          return (
            <button
              id="wallet-connected"
              onClick={openAccountModal}
              className={[
                'inline-flex items-center gap-2.5',
                'h-9 pl-2 pr-3.5 rounded-lg',
                'border border-[#1F2937]',
                'bg-[#111827]',
                'hover:border-[#2563EB]/40 hover:bg-[#1A2235]',
                'hover:shadow-[0_0_16px_rgba(37,99,235,0.12)]',
                'transition-all duration-200',
                'cursor-pointer group',
              ].join(' ')}
            >
              {/* Teal connected dot */}
              <span className="relative flex h-2 w-2 flex-shrink-0">
                <span className="absolute inline-flex h-full w-full rounded-full bg-[#0D9488] opacity-60 animate-ping" />
                <span className="relative inline-flex h-2 w-2 rounded-full bg-[#0D9488]" />
              </span>

              {/* Balance */}
              <span className="text-[#9CA3AF] text-xs font-medium tabular-nums">
                {ethBalance}
              </span>

              {/* Divider */}
              <span className="w-px h-4 bg-[#1F2937]" aria-hidden="true" />

              {/* Address */}
              <span className="text-[#F9FAFB] text-sm font-medium font-mono">
                {shortenAddress(address)}
              </span>

              {/* Chevron */}
              <svg
                className="w-3.5 h-3.5 text-[#4B5563] group-hover:text-[#9CA3AF] transition-colors"
                viewBox="0 0 16 16"
                fill="none"
                aria-hidden="true"
              >
                <path
                  d="M4 6l4 4 4-4"
                  stroke="currentColor"
                  strokeWidth="1.5"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
            </button>
          );
        }

        // ── Disconnected ───────────────────────────────────────────────────
        return (
          <button
            id="wallet-connect"
            onClick={openConnectModal}
            className={[
              'inline-flex items-center gap-2',
              'h-9 px-4 rounded-lg',
              'border border-[#2563EB]',
              'bg-transparent text-[#2563EB]',
              'text-sm font-medium',
              'hover:bg-[rgba(37,99,235,0.08)]',
              'hover:shadow-[0_0_14px_rgba(37,99,235,0.25)]',
              'transition-all duration-200',
              'cursor-pointer',
            ].join(' ')}
          >
            {/* Wallet icon */}
            <svg
              className="w-4 h-4 flex-shrink-0"
              viewBox="0 0 20 20"
              fill="none"
              aria-hidden="true"
            >
              <rect
                x="2"
                y="5"
                width="16"
                height="12"
                rx="2"
                stroke="currentColor"
                strokeWidth="1.5"
              />
              <path
                d="M14 11a1 1 0 1 0 2 0 1 1 0 0 0-2 0Z"
                fill="currentColor"
              />
              <path
                d="M2 8h16"
                stroke="currentColor"
                strokeWidth="1.5"
                strokeLinecap="round"
              />
            </svg>
            Connect Wallet
          </button>
        );
      }}
    </ConnectButton.Custom>
  );
}

export default WalletButton;
