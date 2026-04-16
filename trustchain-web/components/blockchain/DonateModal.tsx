'use client';

import React, { useMemo, useState } from 'react';
import { formatEther, parseEther } from 'viem';
import { useAccount, useBalance } from 'wagmi';

import { useDonate } from '@/hooks/useDonate';
import { Button } from '@/components/ui/Button';

interface DonateModalProps {
  campaignId: string;
  campaignAddress: `0x${string}`;
  campaignName: string;
  isOpen: boolean;
  onClose: () => void;
}

export function DonateModal({
  campaignId,
  campaignAddress,
  campaignName,
  isOpen,
  onClose,
}: DonateModalProps) {
  const [amountEth, setAmountEth] = useState('');

  const { address } = useAccount();
  const { data: balance } = useBalance({ address });

  const { donate, status, txHash, error, reset } = useDonate({
    campaignId,
    campaignAddress,
  });

  const amountWei = useMemo(() => {
    try {
      return parseEther(amountEth || '0');
    } catch {
      return 0n;
    }
  }, [amountEth]);

  const isInvalidAmount = !amountEth || amountWei <= 0n;
  const isInsufficient = !!balance && amountWei > balance.value;

  const balanceDisplay = balance
    ? `${parseFloat(formatEther(balance.value)).toFixed(4)} ETH`
    : '...';

  const basescanUrl = txHash ? `https://sepolia.basescan.org/tx/${txHash}` : null;

  if (!isOpen) return null;

  const closeAndReset = () => {
    setAmountEth('');
    reset();
    onClose();
  };

  const handleConfirm = async () => {
    if (isInvalidAmount || isInsufficient) return;
    await donate(amountEth);
  };

  const showInput = status === 'idle' || status === 'preparing';

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-[rgba(0,0,0,0.65)] px-4">
      <div className="w-full max-w-md rounded-2xl border border-[#1F2937] bg-[#111827] p-6">
        <div className="flex items-start justify-between gap-4">
          <div>
            <h3 className="text-white text-xl font-semibold">Donate to {campaignName}</h3>
            <p className="text-[#9CA3AF] text-sm mt-1">Balance: {balanceDisplay}</p>
          </div>
          <button
            type="button"
            className="text-[#9CA3AF] hover:text-white"
            onClick={closeAndReset}
            aria-label="Close donation modal"
          >
            x
          </button>
        </div>

        <div className="mt-5">
          {showInput && (
            <div className="space-y-3">
              <label className="text-sm text-[#9CA3AF] block" htmlFor="donation-amount">
                Amount (ETH)
              </label>
              <input
                id="donation-amount"
                value={amountEth}
                onChange={(e) => setAmountEth(e.target.value)}
                placeholder="0.01"
                className="w-full bg-[#0D1117] border border-[#1F2937] rounded-lg px-3 py-2 text-white"
              />
              {isInsufficient && (
                <p className="text-xs text-[#FCA5A5]">
                  Insufficient balance for this donation plus gas.
                </p>
              )}
              {error && <p className="text-xs text-[#FCA5A5]">{error}</p>}
            </div>
          )}

          {status === 'awaiting-wallet' && (
            <p className="text-[#9CA3AF] text-sm">Check your wallet to confirm the transaction.</p>
          )}

          {status === 'pending' && (
            <div className="space-y-2">
              <p className="text-[#9CA3AF] text-sm">Transaction submitted. Waiting for confirmation...</p>
              {basescanUrl && (
                <a
                  href={basescanUrl}
                  target="_blank"
                  rel="noreferrer"
                  className="text-[#60A5FA] text-sm hover:text-[#93C5FD]"
                >
                  View on Basescan
                </a>
              )}
            </div>
          )}

          {status === 'success' && (
            <div className="space-y-2">
              <p className="text-[#6EE7B7] text-sm">Donation confirmed successfully.</p>
              {basescanUrl && (
                <a
                  href={basescanUrl}
                  target="_blank"
                  rel="noreferrer"
                  className="text-[#60A5FA] text-sm hover:text-[#93C5FD]"
                >
                  View transaction
                </a>
              )}
            </div>
          )}

          {status === 'error' && error && (
            <p className="text-[#FCA5A5] text-sm">{error}</p>
          )}
        </div>

        <div className="mt-6 flex items-center justify-end gap-2">
          <Button variant="ghost" size="sm" onClick={closeAndReset}>
            Close
          </Button>
          {status === 'error' ? (
            <Button variant="secondary" size="sm" onClick={reset}>
              Retry
            </Button>
          ) : null}
          {showInput ? (
            <Button
              variant="primary"
              size="sm"
              onClick={handleConfirm}
              disabled={isInvalidAmount || isInsufficient}
              loading={status === 'preparing'}
            >
              Confirm Donation
            </Button>
          ) : null}
        </div>
      </div>
    </div>
  );
}

export default DonateModal;
