'use client';

import React, { useEffect, useRef, useState } from 'react';

// ─── Types ───────────────────────────────────────────────────────────────────

export interface ProgressBarProps {
  current: number;
  goal: number;
  showLabel?: boolean;
  className?: string;
  id?: string;
}

// ─── ProgressBar ─────────────────────────────────────────────────────────────

export function ProgressBar({
  current,
  goal,
  showLabel = false,
  className = '',
  id,
}: ProgressBarProps) {
  const [rendered, setRendered] = useState(false);
  const barRef = useRef<HTMLDivElement>(null);

  // Clamp between 0 and 100
  const rawPercent = goal > 0 ? (current / goal) * 100 : 0;
  const percent = Math.min(Math.max(rawPercent, 0), 100);
  const displayPercent = Math.round(percent);

  // Animate on mount (slight delay to trigger CSS transition)
  useEffect(() => {
    const timer = setTimeout(() => setRendered(true), 50);
    return () => clearTimeout(timer);
  }, []);

  // Format ETH-like numbers for display
  const formatAmount = (wei: number): string => {
    if (wei >= 1e18) return `${(wei / 1e18).toFixed(2)} ETH`;
    if (wei >= 1e9) return `${(wei / 1e9).toFixed(2)} Gwei`;
    return wei.toLocaleString();
  };

  return (
    <div id={id} className={['w-full', className].filter(Boolean).join(' ')}>
      {/* Label row */}
      {showLabel && (
        <div className="flex items-center justify-between mb-2">
          <span className="text-xs font-medium text-[#9CA3AF]">
            {formatAmount(current)}{' '}
            <span className="text-[#4B5563]">/ {formatAmount(goal)}</span>
          </span>
          <span
            className="text-xs font-semibold tabular-nums"
            style={{
              color: percent >= 100 ? '#2DD4BF' : percent >= 75 ? '#60A5FA' : '#9CA3AF',
            }}
          >
            {displayPercent}%
          </span>
        </div>
      )}

      {/* Track */}
      <div
        ref={barRef}
        className="relative w-full h-2 rounded-full overflow-hidden"
        style={{ backgroundColor: '#1F2937' }}
        role="progressbar"
        aria-valuenow={displayPercent}
        aria-valuemin={0}
        aria-valuemax={100}
        aria-label={`${displayPercent}% funded`}
      >
        {/* Filled bar */}
        <div
          className="h-full rounded-full"
          style={{
            width: rendered ? `${percent}%` : '0%',
            background: 'linear-gradient(90deg, #2563EB 0%, #0D9488 100%)',
            transition: 'width 0.8s cubic-bezier(0.4, 0, 0.2, 1)',
            boxShadow:
              percent > 0 ? '0 0 8px rgba(37, 99, 235, 0.4)' : 'none',
          }}
        />

        {/* Shimmer overlay when in-progress */}
        {percent > 0 && percent < 100 && (
          <div
            className="absolute inset-0 rounded-full"
            style={{
              background:
                'linear-gradient(90deg, transparent 0%, rgba(255,255,255,0.08) 50%, transparent 100%)',
              backgroundSize: '200% 100%',
              animation: 'shimmer 2s linear infinite',
            }}
          />
        )}
      </div>
    </div>
  );
}

export default ProgressBar;
