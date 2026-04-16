'use client';

import { useEffect, useState } from 'react';
import { useWaitForTransactionReceipt, useWriteContract } from 'wagmi';
import { parseEther } from 'viem';

import { CAMPAIGN_ABI } from '@/lib/contracts';

export type DonateStatus =
  | 'idle'
  | 'preparing'
  | 'awaiting-wallet'
  | 'pending'
  | 'success'
  | 'error';

interface UseDonateOptions {
  campaignId: string;
  campaignAddress: `0x${string}`;
}

interface DonateResult {
  donate: (amountEth: string) => Promise<void>;
  status: DonateStatus;
  txHash: `0x${string}` | undefined;
  error: string | null;
  reset: () => void;
}

export function useDonate({ campaignId, campaignAddress }: UseDonateOptions): DonateResult {
  const [status, setStatus] = useState<DonateStatus>('idle');
  const [error, setError] = useState<string | null>(null);

  const { writeContractAsync, data: txHash, reset: resetWrite } = useWriteContract();
  const { isSuccess: isConfirmed } = useWaitForTransactionReceipt({ hash: txHash });

  useEffect(() => {
    if (isConfirmed) {
      setStatus('success');
    }
  }, [isConfirmed]);

  const donate = async (amountEth: string) => {
    setError(null);
    setStatus('preparing');

    try {
      const apiUrl = process.env.NEXT_PUBLIC_API_URL;
      const res = await fetch(`${apiUrl}/api/campaigns/${campaignId}/prepare-donation`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          amountWei: parseEther(amountEth).toString(),
        }),
      });

      if (!res.ok) {
        console.warn('prepare-donation failed; proceeding with empty CID');
      }

      const json = res.ok ? await res.json() : { data: { ipfsCid: '' } };
      const ipfsCid: string = json?.data?.ipfsCid ?? '';

      setStatus('awaiting-wallet');

      await writeContractAsync({
        address: campaignAddress,
        abi: CAMPAIGN_ABI,
        functionName: 'donate',
        args: [ipfsCid],
        value: parseEther(amountEth),
      });

      setStatus('pending');
    } catch (err: unknown) {
      setStatus('error');
      setError(mapWagmiError(err));
    }
  };

  const reset = () => {
    setStatus('idle');
    setError(null);
    resetWrite();
  };

  return { donate, status, txHash, error, reset };
}

function mapWagmiError(err: unknown): string {
  if (!(err instanceof Error)) return 'Transaction failed';
  const msg = err.message;

  if (msg.includes('User rejected') || msg.includes('user rejected')) {
    return 'You cancelled the transaction.';
  }
  if (msg.includes('insufficient funds')) {
    return 'Insufficient ETH balance for this donation plus gas.';
  }
  if (msg.includes('use donate()')) {
    return 'Please use the Donate button.';
  }
  if (msg.includes('not active') || msg.includes('Campaign: not active')) {
    return 'This campaign is no longer accepting donations.';
  }
  if (msg.includes('network') || msg.includes('fetch')) {
    return 'Network error. Check your connection and try again.';
  }

  return 'Transaction failed. Please try again.';
}
