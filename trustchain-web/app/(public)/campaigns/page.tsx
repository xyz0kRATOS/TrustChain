'use client';

import React, { useEffect, useMemo, useState } from 'react';
import Link from 'next/link';
import { CampaignCard } from '@/components/campaign/CampaignCard';
import type { Campaign } from '@/components/campaign/CampaignCard';
import { Button } from '@/components/ui/Button';
import type { BadgeVariant } from '@/components/ui/Badge';
import type { ApiEnvelope, CampaignResponse } from '@/types/campaign';

// ─── Types ───────────────────────────────────────────────────────────────────

type CampaignStatus = 'ALL' | 'ACTIVE' | 'COMPLETED' | 'FROZEN' | 'PENDING';
type SortOption = 'most_raised' | 'newest' | 'ending_soon' | 'pct_goal';

// ─── Helpers ──────────────────────────────────────────────────────────────────

const WEI = 1_000_000_000_000_000_000n;

function weiToEth(wei: bigint): number {
  return Number((wei * 10_000n) / WEI) / 10_000;
}

function statusToBadgeVariant(status: string): BadgeVariant {
  const normalized = status.toLowerCase();
  const map: Record<string, BadgeVariant> = {
    live: 'active',
    active: 'active',
    completed: 'completed',
    frozen: 'frozen',
    pending: 'pending',
    rejected: 'disputed',
  };
  return map[normalized] ?? 'pending';
}

function responseToCampaign(m: CampaignResponse): Campaign {
  const goalWei = m.goalAmountWei ? BigInt(m.goalAmountWei) : 0n;
  const raisedWei = m.totalRaisedWei ? BigInt(m.totalRaisedWei) : 0n;
  return {
    id: m.id,
    name: m.name,
    description: m.description,
    goalAmount: goalWei > 0n ? weiToEth(goalWei) : 0,
    totalRaised: weiToEth(raisedWei),
    status: statusToBadgeVariant(m.status),
    donorCount: m.donorCount,
    milestoneCount: m.milestones.length,
    imageUrl: m.imageUrl ?? undefined,
  };
}

// ─── Filter Pill ──────────────────────────────────────────────────────────────

interface FilterPillProps {
  label: string;
  active: boolean;
  onClick: () => void;
  id: string;
}

function FilterPill({ label, active, onClick, id }: FilterPillProps) {
  return (
    <button
      id={id}
      onClick={onClick}
      className={[
        'text-sm px-4 py-1.5 rounded-full cursor-pointer transition-all duration-200',
        'border whitespace-nowrap',
        active
          ? 'bg-[#2563EB] text-white border-[#2563EB]'
          : 'bg-[#111827] border-[#1F2937] text-[#9CA3AF] hover:border-[#374151] hover:text-white',
      ].join(' ')}
    >
      {label}
    </button>
  );
}

// ─── Empty State ──────────────────────────────────────────────────────────────

function EmptyState({ onClear }: { onClear: () => void }) {
  return (
    <div className="flex flex-col items-center justify-center mt-20 text-center px-4">
      {/* Magnifying glass with X illustration */}
      <svg
        width="80"
        height="80"
        viewBox="0 0 80 80"
        fill="none"
        aria-hidden="true"
      >
        <circle cx="34" cy="34" r="22" stroke="#374151" strokeWidth="3" />
        <line x1="50" y1="50" x2="68" y2="68" stroke="#374151" strokeWidth="3" strokeLinecap="round" />
        <line x1="27" y1="27" x2="41" y2="41" stroke="#374151" strokeWidth="2.5" strokeLinecap="round" />
        <line x1="41" y1="27" x2="27" y2="41" stroke="#374151" strokeWidth="2.5" strokeLinecap="round" />
      </svg>
      <p className="text-xl font-medium text-white mt-6">No campaigns found</p>
      <p className="text-[#9CA3AF] mt-2 text-sm">Try adjusting your search or filters</p>
      <Button
        variant="secondary"
        size="sm"
        onClick={onClear}
        className="mt-6"
        id="empty-state-clear-filters"
      >
        Clear filters
      </Button>
    </div>
  );
}

// ─── Sidebar Cards ────────────────────────────────────────────────────────────

function PlatformStatsCard() {
  const stats: { label: string; value: string; highlight?: boolean }[] = [
    { label: 'Total Raised', value: '$48,200' },
    { label: 'Active Campaigns', value: '7' },
    { label: 'Total Donors', value: '847' },
    { label: 'Completed', value: '3', highlight: true },
  ];

  return (
    <div className="bg-[#111827] rounded-xl border border-[#1F2937] p-5">
      <p className="text-sm font-semibold text-[#9CA3AF] tracking-wider uppercase mb-1">
        Platform Stats
      </p>
      <div className="divide-y divide-[#1F2937]">
        {stats.map((stat) => (
          <div key={stat.label} className="flex items-center justify-between py-3">
            <span className="text-sm text-[#6B7280]">{stat.label}</span>
            <span
              className="text-sm font-semibold tabular-nums"
              style={{ color: stat.highlight ? '#0D9488' : '#F9FAFB' }}
            >
              {stat.value}
            </span>
          </div>
        ))}
      </div>
    </div>
  );
}

function FundProtectionCard() {
  const points = [
    'Locked until milestones complete',
    '48-hour release delay — always',
    'Smart contract audited',
  ];

  return (
    <div
      className="rounded-r-xl p-5"
      style={{
        backgroundColor: '#111827',
        borderLeft: '4px solid #2563EB',
        borderTop: '1px solid #1F2937',
        borderRight: '1px solid #1F2937',
        borderBottom: '1px solid #1F2937',
        borderRadius: '0 0.75rem 0.75rem 0',
      }}
    >
      <p className="text-sm font-semibold text-white mb-4">How funds are protected</p>
      <ul className="space-y-3">
        {points.map((point) => (
          <li key={point} className="flex items-start gap-3">
            <span
              className="flex-shrink-0 w-5 h-5 rounded-full flex items-center justify-center mt-0.5"
              style={{ backgroundColor: 'rgba(37,99,235,0.15)' }}
            >
              <svg width="10" height="10" viewBox="0 0 10 10" fill="none" aria-hidden="true">
                <path
                  d="M1.5 5l2.5 2.5 4.5-4.5"
                  stroke="#2563EB"
                  strokeWidth="1.5"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
            </span>
            <span className="text-sm text-[#9CA3AF] leading-relaxed">{point}</span>
          </li>
        ))}
      </ul>
    </div>
  );
}

function ApplyCtaCard() {
  return (
    <div
      className="rounded-xl p-5"
      style={{
        background: 'linear-gradient(135deg, #1E3A5F 0%, #111827 100%)',
        border: '1px solid rgba(37,99,235,0.3)',
      }}
    >
      <p className="text-base font-semibold text-white leading-snug">
        Running a verified campaign?
      </p>
      <p className="text-sm text-[#9CA3AF] mt-2 leading-relaxed">
        Apply for funding on TrustChain. Manual review, on-chain accountability.
      </p>
      <Link href="/apply">
        <Button
          variant="secondary"
          size="sm"
          className="mt-4"
          id="sidebar-apply-now"
        >
          Apply Now →
        </Button>
      </Link>
    </div>
  );
}

// ─── Page ─────────────────────────────────────────────────────────────────────

export default function CampaignsPage() {
  const [campaigns, setCampaigns] = useState<CampaignResponse[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [statusFilter, setStatusFilter] = useState<CampaignStatus>('ALL');
  const [searchQuery, setSearchQuery] = useState('');
  const [sortOption, setSortOption] = useState<SortOption>('most_raised');

  const statusPills: { label: string; value: CampaignStatus }[] = [
    { label: 'All', value: 'ALL' },
    { label: 'Active', value: 'ACTIVE' },
    { label: 'Completed', value: 'COMPLETED' },
    { label: 'Frozen', value: 'FROZEN' },
    { label: 'Pending', value: 'PENDING' },
  ];

  useEffect(() => {
    let mounted = true;
    async function fetchCampaigns() {
      try {
        setIsLoading(true);
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/campaigns`);
        const json = (await res.json()) as ApiEnvelope<CampaignResponse[]>;
        if (!mounted) return;
        setCampaigns(Array.isArray(json.data) ? json.data : []);
      } catch {
        if (!mounted) return;
        setCampaigns([]);
      } finally {
        if (mounted) setIsLoading(false);
      }
    }
    fetchCampaigns();
    return () => {
      mounted = false;
    };
  }, []);

  const filteredCampaigns = useMemo<Campaign[]>(() => {
    let results = [...campaigns];

    // Status filter
    if (statusFilter !== 'ALL') {
      if (statusFilter === 'ACTIVE') {
        results = results.filter((c) => c.status.toLowerCase() === 'live' || c.status.toLowerCase() === 'active');
      } else {
        results = results.filter((c) => c.status.toLowerCase() === statusFilter.toLowerCase());
      }
    }

    // Search filter (case-insensitive, name + description)
    if (searchQuery.trim()) {
      const q = searchQuery.toLowerCase();
      results = results.filter(
        (c) =>
          c.name.toLowerCase().includes(q) ||
          c.description.toLowerCase().includes(q),
      );
    }

    // Sort
    results.sort((a, b) => {
      switch (sortOption) {
        case 'most_raised':
          return Number(BigInt(b.totalRaisedWei || '0') - BigInt(a.totalRaisedWei || '0'));

        case 'newest':
          return new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime();

        case 'ending_soon': {
          return 0;
        }

        case 'pct_goal': {
          const aGoal = a.goalAmountWei ? BigInt(a.goalAmountWei) : 0n;
          const bGoal = b.goalAmountWei ? BigInt(b.goalAmountWei) : 0n;
          const aRaised = BigInt(a.totalRaisedWei || '0');
          const bRaised = BigInt(b.totalRaisedWei || '0');
          const aPct = aGoal > 0n ? Number((aRaised * 10_000n) / aGoal) : 0;
          const bPct = bGoal > 0n ? Number((bRaised * 10_000n) / bGoal) : 0;
          return bPct - aPct;
        }

        default:
          return 0;
      }
    });

    return results.map(responseToCampaign);
  }, [campaigns, statusFilter, searchQuery, sortOption]);

  function clearFilters() {
    setStatusFilter('ALL');
    setSearchQuery('');
    setSortOption('most_raised');
  }

  return (
    <>
      {/* ── Page Header ────────────────────────────────────────────────── */}
      <div
        className="pt-24 pb-10"
        style={{ backgroundColor: '#0A0F1E' }}
      >
        <div className="max-w-7xl mx-auto px-6">
          {/* Breadcrumb */}
          <p className="text-sm text-[#4B5563]">
            Home{' '}
            <span className="mx-1.5 text-[#374151]">/</span>
            <span className="text-[#6B7280]">Campaigns</span>
          </p>

          {/* Title */}
          <h1 className="text-4xl font-bold text-white mt-2 tracking-tight">
            Browse Campaigns
          </h1>

          {/* Subtitle */}
          <p className="text-[#9CA3AF] mt-2 text-base">
            12 verified campaigns — milestone-locked, audited, and tracked on-chain
          </p>
        </div>
      </div>

      {/* ── Sticky Filters Bar ─────────────────────────────────────────── */}
      <div
        className="sticky z-40"
        style={{
          top: '64px',
          backgroundColor: 'rgba(10,15,30,0.9)',
          backdropFilter: 'blur(12px)',
          WebkitBackdropFilter: 'blur(12px)',
          borderBottom: '1px solid #1F2937',
          paddingTop: '1rem',
          paddingBottom: '1rem',
        }}
      >
        <div className="max-w-7xl mx-auto px-6 flex flex-wrap items-center justify-between gap-3">
          {/* Left — Status pills */}
          <div className="flex items-center gap-2 flex-wrap">
            {statusPills.map((pill) => (
              <FilterPill
                key={pill.value}
                id={`filter-pill-${pill.value.toLowerCase()}`}
                label={pill.label}
                active={statusFilter === pill.value}
                onClick={() => setStatusFilter(pill.value)}
              />
            ))}
          </div>

          {/* Right — Search + Sort */}
          <div className="flex items-center gap-3 flex-wrap">
            {/* Search */}
            <div className="relative">
              <span className="absolute left-3 top-1/2 -translate-y-1/2 text-[#4B5563] pointer-events-none">
                <svg width="15" height="15" viewBox="0 0 20 20" fill="none" aria-hidden="true">
                  <circle cx="9" cy="9" r="6" stroke="currentColor" strokeWidth="1.6" />
                  <path d="M13.5 13.5L18 18" stroke="currentColor" strokeWidth="1.6" strokeLinecap="round" />
                </svg>
              </span>
              <input
                id="campaigns-search"
                type="search"
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                placeholder="Search campaigns..."
                className={[
                  'w-64 pl-9 pr-4 py-2 text-sm rounded-lg',
                  'bg-[#111827] text-white placeholder-[#4B5563]',
                  'border border-[#1F2937]',
                  'focus:outline-none focus:border-[#2563EB]',
                  'transition-colors duration-200',
                ].join(' ')}
              />
            </div>

            {/* Sort */}
            <select
              id="campaigns-sort"
              value={sortOption}
              onChange={(e) => setSortOption(e.target.value as SortOption)}
              className={[
                'px-3 py-2 text-sm rounded-lg',
                'bg-[#111827] text-[#9CA3AF]',
                'border border-[#1F2937]',
                'focus:outline-none focus:border-[#2563EB]',
                'cursor-pointer transition-colors duration-200',
              ].join(' ')}
            >
              <option value="most_raised">Most Raised</option>
              <option value="newest">Newest First</option>
              <option value="ending_soon">Ending Soon</option>
              <option value="pct_goal">% to Goal</option>
            </select>
          </div>
        </div>
      </div>

      {/* ── Main Content ───────────────────────────────────────────────── */}
      <div
        className="min-h-screen"
        style={{ backgroundColor: '#0A0F1E' }}
      >
        <div className="max-w-7xl mx-auto px-6 py-10">
          <div className="flex gap-8 items-start">
            {/* ── Main Grid ─────────────────────────────── */}
            <div className="flex-1 min-w-0">
              {/* Results count */}
              <p className="text-sm text-[#9CA3AF] mb-6">
                Showing{' '}
                <span className="text-white font-medium">{filteredCampaigns.length}</span>
                {' '}campaign{filteredCampaigns.length !== 1 ? 's' : ''}
              </p>

              {isLoading ? (
                <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                  {Array.from({ length: 6 }).map((_, i) => (
                    <div key={i} className="h-[380px] rounded-xl border border-[#1F2937] bg-[#111827] animate-pulse" />
                  ))}
                </div>
              ) : filteredCampaigns.length === 0 ? (
                <EmptyState onClear={clearFilters} />
              ) : (
                <>
                  <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                    {filteredCampaigns.map((campaign) => (
                      <CampaignCard key={campaign.id} campaign={campaign} />
                    ))}
                  </div>

                  {/* Load More */}
                  <div className="mt-12 text-center">
                    <p className="text-sm text-[#9CA3AF] mb-4">
                      Showing {filteredCampaigns.length} campaigns
                    </p>
                    <Button
                      variant="ghost"
                      size="md"
                      id="campaigns-load-more"
                      disabled
                    >
                      End of results
                    </Button>
                  </div>
                </>
              )}
            </div>

            {/* ── Sidebar ───────────────────────────────── */}
            <aside
              className="hidden lg:flex flex-col gap-5 flex-shrink-0"
              style={{ width: '280px' }}
              aria-label="Campaign sidebar"
            >
              <PlatformStatsCard />
              <FundProtectionCard />
              <ApplyCtaCard />
            </aside>
          </div>
        </div>
      </div>
    </>
  );
}
