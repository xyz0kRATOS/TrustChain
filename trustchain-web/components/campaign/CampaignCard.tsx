import React from 'react';
import Link from 'next/link';
import { Badge, BadgeVariant } from '@/components/ui/Badge';
import { ProgressBar } from '@/components/ui/ProgressBar';

// ─── Types ───────────────────────────────────────────────────────────────────

export interface Campaign {
  id: string;
  name: string;
  description: string;
  /** Goal in ETH (plain number, e.g. 5.0) */
  goalAmount: number;
  /** Total raised in ETH (plain number, e.g. 3.2) */
  totalRaised: number;
  status: BadgeVariant;
  donorCount: number;
  milestoneCount: number;
  imageUrl?: string;
}

export interface CampaignCardProps {
  campaign: Campaign;
  className?: string;
}

// ─── Gradient palettes per status ────────────────────────────────────────────

const gradientByStatus: Record<string, string> = {
  active:    'from-[#0F1D3A] via-[#111827] to-[#0A1628]',
  completed: 'from-[#0A1F1E] via-[#111827] to-[#091918]',
  frozen:    'from-[#1F0A0A] via-[#111827] to-[#180909]',
  pending:   'from-[#131722] via-[#111827] to-[#0F1420]',
};

// ─── Donor + Milestone Icons ──────────────────────────────────────────────────

function DonorIcon() {
  return (
    <svg width="13" height="13" viewBox="0 0 16 16" fill="none" aria-hidden="true">
      <circle cx="8" cy="5" r="3" stroke="currentColor" strokeWidth="1.4" />
      <path d="M2 14c0-3.314 2.686-6 6-6s6 2.686 6 6" stroke="currentColor" strokeWidth="1.4" strokeLinecap="round" />
    </svg>
  );
}

function MilestoneIcon() {
  return (
    <svg width="13" height="13" viewBox="0 0 16 16" fill="none" aria-hidden="true">
      <path d="M3 8l3.5 3.5L14 4" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
    </svg>
  );
}

// ─── CampaignCard ─────────────────────────────────────────────────────────────

export function CampaignCard({ campaign, className = '' }: CampaignCardProps) {
  const {
    id,
    name,
    description,
    goalAmount,
    totalRaised,
    status,
    donorCount,
    milestoneCount,
    imageUrl,
  } = campaign;

  const gradient = gradientByStatus[status] ?? gradientByStatus.pending;
  const percentFunded = goalAmount > 0 ? Math.min((totalRaised / goalAmount) * 100, 100) : 0;

  return (
    <article
      id={`campaign-card-${id}`}
      className={[
        'group flex flex-col',
        'bg-[#111827] rounded-xl',
        'border border-[#1F2937]',
        'overflow-hidden',
        'transition-all duration-300 ease-in-out',
        'hover:border-[#2563EB]/35',
        'hover:shadow-[0_0_28px_rgba(37,99,235,0.13),0_4px_16px_rgba(0,0,0,0.4)]',
        'hover:-translate-y-0.5',
        className,
      ].join(' ')}
    >
      {/* ── Image / Gradient Banner ───────────────────────────────────── */}
      <div className="relative h-44 flex-shrink-0 overflow-hidden">
        {imageUrl ? (
          // eslint-disable-next-line @next/next/no-img-element
          <img
            src={imageUrl}
            alt={name}
            className="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
          />
        ) : (
          /* Gradient placeholder — never a grey box */
          <div
            className={`w-full h-full bg-gradient-to-br ${gradient} transition-all duration-500 group-hover:opacity-80`}
            aria-hidden="true"
          >
            {/* Subtle grid pattern overlay */}
            <div
              className="absolute inset-0 opacity-[0.04]"
              style={{
                backgroundImage:
                  'linear-gradient(#fff 1px, transparent 1px), linear-gradient(90deg, #fff 1px, transparent 1px)',
                backgroundSize: '28px 28px',
              }}
            />
            {/* Centre icon watermark */}
            <div className="absolute inset-0 flex items-center justify-center">
              <svg width="52" height="52" viewBox="0 0 24 24" fill="none" opacity="0.12" aria-hidden="true">
                <path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71" stroke="#2563EB" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
                <path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71" stroke="#60A5FA" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
              </svg>
            </div>
          </div>
        )}

        {/* Status badge — top-right overlay */}
        <div className="absolute top-3 right-3">
          <Badge variant={status} />
        </div>

        {/* Gradient fade at bottom of image into card body */}
        <div
          className="absolute bottom-0 left-0 right-0 h-12"
          style={{ background: 'linear-gradient(to bottom, transparent, #111827)' }}
          aria-hidden="true"
        />
      </div>

      {/* ── Card Body ─────────────────────────────────────────────────── */}
      <div className="flex flex-col flex-1 px-5 pt-3 pb-5 gap-3">

        {/* Name */}
        <h3 className="text-[#F9FAFB] font-semibold text-base leading-snug line-clamp-2 group-hover:text-white transition-colors">
          {name}
        </h3>

        {/* Description */}
        <p className="text-[#9CA3AF] text-sm leading-relaxed line-clamp-2 flex-1">
          {description}
        </p>

        {/* Progress + label */}
        <div className="space-y-2">
          <ProgressBar
            current={totalRaised}
            goal={goalAmount}
            showLabel={false}
          />
          <div className="flex items-baseline justify-between">
            <span className="text-sm font-semibold text-[#F9FAFB] tabular-nums">
              {totalRaised.toFixed(2)}{' '}
              <span className="text-xs font-normal text-[#9CA3AF]">ETH raised</span>
            </span>
            <span className="text-xs text-[#4B5563] tabular-nums">
              of {goalAmount.toFixed(2)} ETH
            </span>
          </div>
        </div>

        {/* Separator */}
        <div className="h-px bg-[#1F2937]" aria-hidden="true" />

        {/* Meta row */}
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-3">
            <span className="inline-flex items-center gap-1.5 text-xs text-[#9CA3AF]">
              <DonorIcon />
              {donorCount.toLocaleString()} donors
            </span>
            <span className="inline-flex items-center gap-1.5 text-xs text-[#9CA3AF]">
              <MilestoneIcon />
              {milestoneCount} milestones
            </span>
          </div>

          {/* Funded percent chip */}
          <span
            className="text-xs font-semibold tabular-nums px-1.5 py-0.5 rounded"
            style={{
              color: percentFunded >= 100 ? '#2DD4BF' : '#60A5FA',
              backgroundColor:
                percentFunded >= 100
                  ? 'rgba(13,148,136,0.1)'
                  : 'rgba(37,99,235,0.1)',
            }}
          >
            {Math.round(percentFunded)}%
          </span>
        </div>

        {/* CTA */}
        <Link
          href={`/campaigns/${id}`}
          id={`campaign-card-cta-${id}`}
          className={[
            'inline-flex items-center gap-1.5',
            'text-sm font-medium text-[#2563EB]',
            'hover:text-[#60A5FA] transition-colors duration-200',
            'group/link',
          ].join(' ')}
        >
          View Campaign
          <svg
            className="w-3.5 h-3.5 transition-transform duration-200 group-hover/link:translate-x-0.5"
            viewBox="0 0 14 14"
            fill="none"
            aria-hidden="true"
          >
            <path d="M2.5 7h9M8 3.5L11.5 7 8 10.5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
          </svg>
        </Link>
      </div>
    </article>
  );
}

export default CampaignCard;
