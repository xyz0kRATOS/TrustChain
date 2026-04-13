'use client';

import React, { useMemo, useState } from 'react';
import Link from 'next/link';
import { CampaignCard } from '@/components/campaign/CampaignCard';
import type { Campaign } from '@/components/campaign/CampaignCard';
import { Button } from '@/components/ui/Button';
import type { BadgeVariant } from '@/components/ui/Badge';

// ─── Types ───────────────────────────────────────────────────────────────────

type CampaignStatus = 'ALL' | 'ACTIVE' | 'COMPLETED' | 'FROZEN';
type SortOption = 'most_raised' | 'newest' | 'ending_soon' | 'pct_goal';

interface MockCampaign {
  id: string;
  name: string;
  description: string;
  goalAmount: bigint;
  totalRaised: bigint;
  status: 'ACTIVE' | 'COMPLETED' | 'FROZEN';
  donorCount: number;
  milestoneCount: number;
  completedMilestones: number;
  creatorAddress: string;
  daysRemaining?: number;
}

// ─── Mock Data ────────────────────────────────────────────────────────────────

const MOCK_CAMPAIGNS: MockCampaign[] = [
  {
    id: '1',
    name: 'Clean Water Project — Nairobi',
    description: 'Providing clean water access to 3 rural villages through borehole drilling.',
    goalAmount: BigInt('5000000000000000000'),
    totalRaised: BigInt('3200000000000000000'),
    status: 'ACTIVE',
    donorCount: 47,
    milestoneCount: 3,
    completedMilestones: 1,
    creatorAddress: '0x4a2f8c3d9e1b7f2a',
    daysRemaining: 24,
  },
  {
    id: '2',
    name: 'School Rebuild — Rural Kenya',
    description: 'Rebuilding primary school damaged by flooding. 4 classrooms for 200 students.',
    goalAmount: BigInt('8000000000000000000'),
    totalRaised: BigInt('8000000000000000000'),
    status: 'COMPLETED',
    donorCount: 124,
    milestoneCount: 4,
    completedMilestones: 4,
    creatorAddress: '0x7b3e9f1a4c2d8e5f',
  },
  {
    id: '3',
    name: 'Medical Aid — Gaza',
    description: 'Emergency medical supplies and equipment for a field hospital.',
    goalAmount: BigInt('10000000000000000000'),
    totalRaised: BigInt('6100000000000000000'),
    status: 'ACTIVE',
    donorCount: 89,
    milestoneCount: 3,
    completedMilestones: 0,
    creatorAddress: '0x9c5d2e4b1a7f3c8d',
    daysRemaining: 11,
  },
  {
    id: '4',
    name: 'Flood Relief — Bangladesh',
    description: 'Emergency shelter materials and food packages for 500 displaced families.',
    goalAmount: BigInt('6000000000000000000'),
    totalRaised: BigInt('5800000000000000000'),
    status: 'ACTIVE',
    donorCount: 203,
    milestoneCount: 2,
    completedMilestones: 1,
    creatorAddress: '0x2e8f1a9c4d7b3e5f',
    daysRemaining: 5,
  },
  {
    id: '5',
    name: 'Solar Panels — Rural Uganda',
    description: 'Installing solar power for 3 off-grid schools, lighting 800 students.',
    goalAmount: BigInt('4000000000000000000'),
    totalRaised: BigInt('1200000000000000000'),
    status: 'ACTIVE',
    donorCount: 34,
    milestoneCount: 3,
    completedMilestones: 0,
    creatorAddress: '0x5f2c7a1e9b4d8f3a',
    daysRemaining: 45,
  },
  {
    id: '6',
    name: 'Orphanage Renovation — Manila',
    description: 'Renovation and safety upgrades for a 60-child orphanage in Metro Manila.',
    goalAmount: BigInt('3000000000000000000'),
    totalRaised: BigInt('3000000000000000000'),
    status: 'COMPLETED',
    donorCount: 67,
    milestoneCount: 2,
    completedMilestones: 2,
    creatorAddress: '0x8a3d5f2c1e9b4a7f',
  },
  {
    id: '7',
    name: 'Reforestation — Amazon Basin',
    description: 'Planting 10,000 native trees in degraded Amazon areas with local communities.',
    goalAmount: BigInt('7000000000000000000'),
    totalRaised: BigInt('2100000000000000000'),
    status: 'ACTIVE',
    donorCount: 58,
    milestoneCount: 4,
    completedMilestones: 0,
    creatorAddress: '0x1b9f4a7c2e5d8a3f',
    daysRemaining: 60,
  },
  {
    id: '8',
    name: 'Veteran Housing — Detroit',
    description: 'Emergency housing assistance for 15 homeless veterans in Detroit.',
    goalAmount: BigInt('9000000000000000000'),
    totalRaised: BigInt('900000000000000000'),
    status: 'FROZEN',
    donorCount: 28,
    milestoneCount: 3,
    completedMilestones: 0,
    creatorAddress: '0x6d2a8f1c4e7b9a5f',
  },
  {
    id: '9',
    name: 'Dialysis Centre — Lagos',
    description: 'Equipment and setup costs for a free dialysis centre serving low-income patients.',
    goalAmount: BigInt('12000000000000000000'),
    totalRaised: BigInt('4800000000000000000'),
    status: 'ACTIVE',
    donorCount: 112,
    milestoneCount: 5,
    completedMilestones: 2,
    creatorAddress: '0x3f7a1d9c5b2e4f8a',
    daysRemaining: 30,
  },
];

// ─── Helpers ──────────────────────────────────────────────────────────────────

const WEI = 1_000_000_000_000_000_000n;

function weiToEth(wei: bigint): number {
  return Number((wei * 10_000n) / WEI) / 10_000;
}

function statusToBadgeVariant(status: MockCampaign['status']): BadgeVariant {
  const map: Record<MockCampaign['status'], BadgeVariant> = {
    ACTIVE: 'active',
    COMPLETED: 'completed',
    FROZEN: 'frozen',
  };
  return map[status];
}

function mockToCampaign(m: MockCampaign): Campaign {
  return {
    id: m.id,
    name: m.name,
    description: m.description,
    goalAmount: weiToEth(m.goalAmount),
    totalRaised: weiToEth(m.totalRaised),
    status: statusToBadgeVariant(m.status),
    donorCount: m.donorCount,
    milestoneCount: m.milestoneCount,
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
  const [statusFilter, setStatusFilter] = useState<CampaignStatus>('ALL');
  const [searchQuery, setSearchQuery] = useState('');
  const [sortOption, setSortOption] = useState<SortOption>('most_raised');

  const statusPills: { label: string; value: CampaignStatus }[] = [
    { label: 'All', value: 'ALL' },
    { label: 'Active', value: 'ACTIVE' },
    { label: 'Completed', value: 'COMPLETED' },
    { label: 'Frozen', value: 'FROZEN' },
  ];

  const filteredCampaigns = useMemo<Campaign[]>(() => {
    let results = [...MOCK_CAMPAIGNS];

    // Status filter
    if (statusFilter !== 'ALL') {
      results = results.filter((c) => c.status === statusFilter);
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
          return Number(b.totalRaised - a.totalRaised);

        case 'newest':
          return Number(b.id) - Number(a.id);

        case 'ending_soon': {
          // ACTIVE campaigns with daysRemaining first (ascending), others last
          const aActive = a.status === 'ACTIVE' && a.daysRemaining !== undefined;
          const bActive = b.status === 'ACTIVE' && b.daysRemaining !== undefined;
          if (aActive && bActive) return (a.daysRemaining ?? 0) - (b.daysRemaining ?? 0);
          if (aActive) return -1;
          if (bActive) return 1;
          return 0;
        }

        case 'pct_goal': {
          const aPct = a.goalAmount > 0n ? Number((a.totalRaised * 10000n) / a.goalAmount) : 0;
          const bPct = b.goalAmount > 0n ? Number((b.totalRaised * 10000n) / b.goalAmount) : 0;
          return bPct - aPct;
        }

        default:
          return 0;
      }
    });

    return results.map(mockToCampaign);
  }, [statusFilter, searchQuery, sortOption]);

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

              {filteredCampaigns.length === 0 ? (
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
                      Showing {filteredCampaigns.length} of 12 campaigns
                    </p>
                    <Button
                      variant="ghost"
                      size="md"
                      id="campaigns-load-more"
                    >
                      Load more
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
