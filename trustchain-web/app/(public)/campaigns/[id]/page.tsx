'use client';

import React, { useCallback, useEffect, useMemo, useState } from 'react';
import Link from 'next/link';
import { useParams } from 'next/navigation';
import { ProgressBar } from '@/components/ui/ProgressBar';
import { Badge } from '@/components/ui/Badge';
import { StatusPill } from '@/components/ui/StatusPill';
import { Button } from '@/components/ui/Button';
import { DonateModal } from '@/components/blockchain/DonateModal';
import { getBasescanUrl, shortenAddress } from '@/lib/utils';
import type { ApiEnvelope, CampaignResponse } from '@/types/campaign';

const WEI = 1_000_000_000_000_000_000n;

function weiToEth(wei: string | null): number {
  if (!wei) return 0;
  const value = BigInt(wei);
  return Number((value * 10_000n) / WEI) / 10_000;
}

function statusToBadgeVariant(status: string) {
  const s = status.toLowerCase();
  if (s === 'live' || s === 'active') return 'active' as const;
  if (s === 'completed') return 'completed' as const;
  if (s === 'frozen') return 'frozen' as const;
  return 'pending' as const;
}

function LoadingState() {
  return (
    <div className="min-h-screen bg-[#0A0F1E] px-6 py-24">
      <div className="max-w-5xl mx-auto space-y-6 animate-pulse">
        <div className="h-10 w-2/3 bg-[#111827] border border-[#1F2937] rounded-xl" />
        <div className="h-20 bg-[#111827] border border-[#1F2937] rounded-xl" />
        <div className="h-28 bg-[#111827] border border-[#1F2937] rounded-xl" />
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div className="h-40 bg-[#111827] border border-[#1F2937] rounded-xl" />
          <div className="h-40 bg-[#111827] border border-[#1F2937] rounded-xl" />
          <div className="h-40 bg-[#111827] border border-[#1F2937] rounded-xl" />
        </div>
      </div>
    </div>
  );
}

function NotFoundState() {
  return (
    <div className="min-h-screen bg-[#0A0F1E] flex items-center justify-center px-6">
      <div className="text-center">
        <h1 className="text-2xl text-white font-semibold">Campaign not found</h1>
        <p className="text-[#9CA3AF] mt-2">This campaign may have been removed or the link is invalid.</p>
        <Link href="/campaigns" className="inline-block mt-6">
          <Button variant="secondary" size="md">Browse all campaigns -&gt;</Button>
        </Link>
      </div>
    </div>
  );
}

function ErrorState({ onRetry }: { onRetry: () => void }) {
  return (
    <div className="min-h-screen bg-[#0A0F1E] flex items-center justify-center px-6">
      <div className="text-center">
        <h1 className="text-2xl text-white font-semibold">Unable to load campaign</h1>
        <p className="text-[#9CA3AF] mt-2">Please check your backend connection and try again.</p>
        <div className="mt-6 flex items-center justify-center gap-3">
          <Button variant="primary" size="md" onClick={onRetry}>Retry</Button>
          <Link href="/campaigns"><Button variant="ghost" size="md">Back to campaigns</Button></Link>
        </div>
      </div>
    </div>
  );
}

export default function CampaignDetailPage() {
  const params = useParams<{ id: string }>();
  const id = params?.id ?? '';
  const [campaign, setCampaign] = useState<CampaignResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [notFound, setNotFound] = useState(false);
  const [errored, setErrored] = useState(false);
  const [isDonateOpen, setIsDonateOpen] = useState(false);

  const fetchCampaign = useCallback(async () => {
    if (!id) return;
    setLoading(true);
    setNotFound(false);
    setErrored(false);
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/campaigns/${id}`);
      const json = (await res.json()) as ApiEnvelope<CampaignResponse | null>;
      if (res.status === 404) {
        setNotFound(true);
        setCampaign(null);
        return;
      }
      if (!res.ok || !json.data) {
        setErrored(true);
        return;
      }
      setCampaign(json.data);
    } catch {
      setErrored(true);
    } finally {
      setLoading(false);
    }
  }, [id]);

  useEffect(() => {
    fetchCampaign();
  }, [fetchCampaign]);

  const progress = useMemo(() => {
    if (!campaign) return { current: 0, goal: 0 };
    return {
      current: weiToEth(campaign.totalRaisedWei),
      goal: weiToEth(campaign.goalAmountWei),
    };
  }, [campaign]);

  if (loading) return <LoadingState />;
  if (notFound) return <NotFoundState />;
  if (errored || !campaign) return <ErrorState onRetry={fetchCampaign} />;

  return (
    <div className="min-h-screen bg-[#0A0F1E] px-6 py-24">
      <div className="max-w-5xl mx-auto">
        <p className="text-sm text-[#4B5563]">
          <Link href="/">Home</Link>
          <span className="mx-1.5 text-[#374151]">/</span>
          <Link href="/campaigns">Campaigns</Link>
          <span className="mx-1.5 text-[#374151]">/</span>
          <span className="text-[#6B7280]">{campaign.name}</span>
        </p>

        <div className="mt-3 flex items-center gap-3">
          <h1 className="text-4xl font-bold text-white tracking-tight">{campaign.name}</h1>
          <Badge variant={statusToBadgeVariant(campaign.status)} />
        </div>

        <p className="mt-4 text-[#9CA3AF] leading-relaxed max-w-4xl">{campaign.description}</p>

        <div className="mt-8 bg-[#111827] border border-[#1F2937] rounded-2xl p-6">
          <div className="flex items-center justify-between mb-3">
            <p className="text-sm text-[#9CA3AF]">Funding Progress</p>
            <p className="text-sm text-[#9CA3AF]">Donors: <span className="text-white">{campaign.donorCount}</span></p>
          </div>
          <ProgressBar current={progress.current} goal={progress.goal} showLabel={false} />
          <div className="mt-3 flex items-center justify-between text-sm">
            <span className="text-white font-semibold">{progress.current.toFixed(4)} ETH raised</span>
            {campaign.goalAmountWei ? (
              <span className="text-[#9CA3AF]">Goal: {progress.goal.toFixed(4)} ETH</span>
            ) : (
              <span className="text-[#9CA3AF]">Goal: ${Number(campaign.goalAmountUsd || '0').toLocaleString()}</span>
            )}
          </div>
        </div>

        <div className="mt-6 bg-[#111827] border border-[#1F2937] rounded-2xl p-6">
          <p className="text-sm text-[#9CA3AF]">Contract</p>
          {campaign.contractAddress ? (
            <div className="mt-2 flex flex-wrap items-center gap-3">
              <a
                href={getBasescanUrl('address', campaign.contractAddress)}
                target="_blank"
                rel="noreferrer"
                className="inline-flex items-center gap-2 text-[#60A5FA] hover:text-[#93C5FD]"
              >
                <span className="font-mono">{shortenAddress(campaign.contractAddress)}</span>
                <span aria-hidden="true">↗</span>
              </a>
              <Button variant="primary" size="sm" onClick={() => setIsDonateOpen(true)}>
                Donate
              </Button>
            </div>
          ) : (
            <div className="mt-2 flex flex-wrap items-center gap-3">
              <p className="text-[#6B7280]">Pending deployment</p>
              <button
                disabled
                className="opacity-50 cursor-not-allowed rounded-lg border border-[#374151] px-3 py-1.5 text-xs text-[#9CA3AF]"
              >
                Donate (Pending Deployment)
              </button>
            </div>
          )}
        </div>

        <div className="mt-8">
          <h2 className="text-white text-xl font-semibold">Milestones</h2>
          <div className="mt-4 space-y-4">
            {campaign.milestones.map((m) => (
              <div key={m.id} className="bg-[#111827] border border-[#1F2937] rounded-2xl p-5">
                <div className="flex items-center justify-between gap-3">
                  <h3 className="text-white font-medium">{m.sequenceIndex + 1}. {m.name}</h3>
                  <StatusPill status={m.status} />
                </div>
                <p className="mt-2 text-[#9CA3AF] text-sm">{m.description}</p>
                <div className="mt-3 text-xs text-[#6B7280] flex flex-wrap gap-4">
                  <span>Amount: ${Number(m.amountUsd || '0').toLocaleString()}</span>
                  <span>Deadline: {new Date(m.deadline).toLocaleDateString()}</span>
                  <span>Evidence: {m.requiredEvidence}</span>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>

      {campaign.contractAddress ? (
        <DonateModal
          campaignId={campaign.id}
          campaignAddress={campaign.contractAddress as `0x${string}`}
          campaignName={campaign.name}
          isOpen={isDonateOpen}
          onClose={() => setIsDonateOpen(false)}
        />
      ) : null}
    </div>
  );
}
