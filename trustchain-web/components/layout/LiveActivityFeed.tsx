'use client';

import React, { useState } from 'react';
import { formatTimeAgo } from '@/lib/utils';

// ─── Types ───────────────────────────────────────────────────────────────────

type ActivityEventType =
  | 'donation'
  | 'milestone_approved'
  | 'funds_released'
  | 'campaign_frozen';

interface ActivityEvent {
  id: string;
  type: ActivityEventType;
  description: string;
  timestamp: number; // Unix seconds
}

// ─── Mock Data ────────────────────────────────────────────────────────────────
// Replace with WebSocket/GraphQL subscription later

const now = Math.floor(Date.now() / 1000);

const MOCK_EVENTS: ActivityEvent[] = [
  {
    id: '1',
    type: 'donation',
    description: '0x1234...abcd donated 0.05 ETH to Clean Water Nairobi',
    timestamp: now - 45,
  },
  {
    id: '2',
    type: 'milestone_approved',
    description: 'Milestone 2 approved for School Rebuild Kenya',
    timestamp: now - 3 * 60,
  },
  {
    id: '3',
    type: 'funds_released',
    description: '0.8 ETH released to School Rebuild Kenya',
    timestamp: now - 12 * 60,
  },
  {
    id: '4',
    type: 'donation',
    description: '0xdead...beef donated 0.2 ETH to Medical Aid Uganda',
    timestamp: now - 28 * 60,
  },
  {
    id: '5',
    type: 'campaign_frozen',
    description: 'Campaign Medical Aid Sudan paused — donors being refunded',
    timestamp: now - 55 * 60,
  },
];

// ─── Event Config ─────────────────────────────────────────────────────────────

const eventConfig: Record<
  ActivityEventType,
  { emoji: string; color: string; bg: string }
> = {
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
  campaign_frozen: {
    emoji: '🔴',
    color: '#F87171',
    bg: 'rgba(220, 38, 38, 0.08)',
  },
};

// ─── FeedItem ─────────────────────────────────────────────────────────────────

function FeedItem({ event, index }: { event: ActivityEvent; index: number }) {
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
          className="text-xs leading-relaxed text-[#9CA3AF] break-words"
          style={{ color: config.color === '#60A5FA' ? undefined : undefined }}
        >
          <span style={{ color: config.color, fontWeight: 500 }}>
            {/* Highlight the token/amount if present */}
          </span>
          {event.description}
        </p>
        <time
          className="block text-[10px] text-[#4B5563] mt-1 tabular-nums"
          dateTime={new Date(event.timestamp * 1000).toISOString()}
        >
          {formatTimeAgo(event.timestamp)}
        </time>
      </div>
    </div>
  );
}

// ─── LiveActivityFeed ─────────────────────────────────────────────────────────

export function LiveActivityFeed() {
  const [collapsed, setCollapsed] = useState(false);

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
          <span className="text-xs text-[#4B5563] font-normal">· mock data</span>
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
          {MOCK_EVENTS.map((event, i) => (
            <FeedItem key={event.id} event={event} index={i} />
          ))}
        </ol>

        {/* Empty state / bottom note */}
        <div className="px-4 py-5 border-t border-[#1F2937]">
          <p className="text-[11px] text-[#4B5563] text-center leading-relaxed">
            Showing last {MOCK_EVENTS.length} events.{' '}
            <span className="text-[#2563EB]">WebSocket live feed coming soon.</span>
          </p>
        </div>
      </div>
    </aside>
  );
}

export default LiveActivityFeed;
