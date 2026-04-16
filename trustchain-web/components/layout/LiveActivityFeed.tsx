'use client';

import React, { useEffect, useMemo, useState } from 'react';
import { formatTimeAgo, getBasescanUrl, shortenAddress } from '@/lib/utils';
import type { ActivityResponse, ApiEnvelope } from '@/types/campaign';

// ─── Types ───────────────────────────────────────────────────────────────────

type ActivityEventType = ActivityResponse['type'];

// ─── Event Config ─────────────────────────────────────────────────────────────

const eventConfig: Record<
  ActivityEventType,
  { emoji: string; color: string; bg: string }
> = {
  campaign_applied: {
    emoji: '📝',
    color: '#9CA3AF',
    bg: 'rgba(156, 163, 175, 0.08)',
  },
  campaign_approved: {
    emoji: '✅',
    color: '#2DD4BF',
    bg: 'rgba(13, 148, 136, 0.08)',
  },
  campaign_live: {
    emoji: '🚀',
    color: '#60A5FA',
    bg: 'rgba(37, 99, 235, 0.08)',
  },
  donation: {
    emoji: '💙',
    color: '#60A5FA',
    bg: 'rgba(37, 99, 235, 0.08)',
  },
  milestone_approved: {
    emoji: '✅',
    color: '#2DD4BF',
    bg: 'rgba(13, 148, 136, 0.08)',
  },
  funds_released: {
    emoji: '💸',
    color: '#A78BFA',
    bg: 'rgba(124, 58, 237, 0.08)',
  },
};

// ─── FeedItem ─────────────────────────────────────────────────────────────────

function eventText(event: ActivityResponse): string {
  switch (event.type) {
    case 'campaign_applied':
      return `${event.campaignName} applied for funding`;
    case 'campaign_approved':
      return `${event.campaignName} approved - now live`;
    case 'campaign_live':
      return `${event.campaignName} is now accepting donations`;
    case 'donation':
      return `${event.wallet || 'A donor'} donated ${event.amount || '0'} ETH to ${event.campaignName}`;
    case 'milestone_approved':
      return `Milestone approved for ${event.campaignName}`;
    case 'funds_released':
      return `${event.amount || '0'} ETH released to ${event.campaignName}`;
    default:
      return event.campaignName;
  }
}

function FeedItem({ event, index }: { event: ActivityResponse; index: number }) {
  const config = eventConfig[event.type];

  return (
    <div
      className="flex gap-3 px-4 py-3 border-b border-[#1F2937] last:border-0 animate-fade-in"
      style={{
        animationDelay: `${index * 60}ms`,
        animationFillMode: 'both',
      }}
      role="listitem"
    >
      {/* Emoji icon */}
      <div
        className="flex-shrink-0 w-8 h-8 rounded-lg flex items-center justify-center text-sm"
        style={{ backgroundColor: config.bg }}
        aria-hidden="true"
      >
        {config.emoji}
      </div>

      {/* Content */}
      <div className="flex-1 min-w-0">
        <p
          className="text-sm leading-relaxed text-[#9CA3AF] break-words truncate"
        >
          {eventText(event)}
        </p>
        <div className="mt-1 flex items-center gap-2">
          <time className="block text-xs text-[#4B5563] tabular-nums" dateTime={event.timestamp}>
            {formatTimeAgo(Date.parse(event.timestamp))}
          </time>
          {event.txHash && (
            <a
              href={getBasescanUrl('tx', event.txHash)}
              target="_blank"
              rel="noreferrer"
              className="text-xs text-[#6B7280] hover:text-[#9CA3AF]"
              aria-label="View transaction on Basescan"
            >
              ↗
            </a>
          )}
        </div>
      </div>
    </div>
  );
}

// ─── LiveActivityFeed ─────────────────────────────────────────────────────────

export function LiveActivityFeed() {
  const [events, setEvents] = useState<ActivityResponse[]>([]);
  const [collapsed, setCollapsed] = useState(false);
  const [secondsAgo, setSecondsAgo] = useState(0);

  const hasOnchainTx = useMemo(() => events.some((e) => !!e.txHash), [events]);

  useEffect(() => {
    let mounted = true;

    async function pull() {
      try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/activity/recent`);
        const json = (await res.json()) as ApiEnvelope<ActivityResponse[]>;
        if (!mounted) return;
        const next = Array.isArray(json.data)
          ? json.data.map((e) => ({ ...e, wallet: e.wallet ? shortenAddress(e.wallet) : null }))
          : [];
        setEvents(next);
        setSecondsAgo(0);
      } catch {
        if (!mounted) return;
      }
    }

    pull();
    const poll = setInterval(pull, 10000);
    const ticker = setInterval(() => setSecondsAgo((v) => v + 1), 1000);

    return () => {
      mounted = false;
      clearInterval(poll);
      clearInterval(ticker);
    };
  }, []);

  return (
    <aside
      id="live-activity-feed"
      aria-label="Live activity feed"
      className={[
        // Hidden on mobile
        'hidden xl:flex flex-col',
        // Fixed right panel
        'fixed top-16 right-0 bottom-0 z-40',
        'w-80',
        'bg-[#111827]',
        'border-l border-[#1F2937]',
        'transition-all duration-300 ease-in-out',
        'overflow-hidden',
      ].join(' ')}
    >
      {/* ── Header ───────────────────────────────────────────────────── */}
      <button
        id="activity-feed-toggle"
        type="button"
        onClick={() => setCollapsed((v) => !v)}
        aria-expanded={!collapsed}
        aria-controls="activity-feed-list"
        className={[
          'flex items-center justify-between',
          'w-full px-4 py-3',
          'border-b border-[#1F2937]',
          'hover:bg-[#1A2235] transition-colors duration-200',
          'cursor-pointer select-none flex-shrink-0',
        ].join(' ')}
      >
        <div className="flex items-center gap-2.5">
          {/* Pulsing green live dot */}
          <span className="relative flex h-2 w-2 flex-shrink-0" aria-hidden="true">
            <span className="absolute inline-flex h-full w-full rounded-full bg-[#22C55E] opacity-60 animate-ping" />
            <span className="relative inline-flex h-2 w-2 rounded-full bg-[#22C55E]" />
          </span>
          <span className="text-sm font-semibold text-[#F9FAFB]">Live Activity</span>
          <span className="text-xs text-[#4B5563] font-normal">Updated {secondsAgo}s ago</span>
        </div>

        {/* Collapse chevron */}
        <svg
          width="14"
          height="14"
          viewBox="0 0 14 14"
          fill="none"
          aria-hidden="true"
          className={`text-[#4B5563] transition-transform duration-300 ${collapsed ? 'rotate-180' : ''}`}
        >
          <path
            d="M3 5l4 4 4-4"
            stroke="currentColor"
            strokeWidth="1.5"
            strokeLinecap="round"
            strokeLinejoin="round"
          />
        </svg>
      </button>

      {/* ── Feed List ─────────────────────────────────────────────────── */}
      <div
        id="activity-feed-list"
        aria-hidden={collapsed}
        className={[
          'flex-1 overflow-y-auto overscroll-contain',
          'transition-all duration-300 ease-in-out',
          collapsed ? 'max-h-0 opacity-0' : 'max-h-full opacity-100',
        ].join(' ')}
      >
        <ol className="flex flex-col" role="list" aria-label="Activity events">
          {events.map((event, i) => (
            <FeedItem key={event.id} event={event} index={i} />
          ))}
        </ol>

        {events.length === 0 && (
          <div className="px-4 py-10 text-center border-t border-[#1F2937]">
            <p className="text-sm text-[#4B5563]">No activity yet</p>
            <p className="text-xs text-[#4B5563] mt-1">Activity appears here as campaigns are submitted and approved</p>
          </div>
        )}

        {!hasOnchainTx && (
          <div className="px-4 py-4 border-t border-[#1F2937]">
            <p className="text-[11px] text-[#4B5563] text-center">
              On-chain events will appear here after contract deployment
            </p>
          </div>
        )}
      </div>
    </aside>
  );
}

export default LiveActivityFeed;
