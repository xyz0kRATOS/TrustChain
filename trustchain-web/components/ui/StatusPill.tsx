import React from 'react';

// ─── Types ───────────────────────────────────────────────────────────────────

// All 8 MilestoneStatus values from TrustChain contracts
export type MilestoneStatus =
  | 'pending'
  | 'submitted'
  | 'under_review'
  | 'approved'
  | 'releasing'
  | 'completed'
  | 'disputed'
  | 'overdue';

export interface StatusPillProps {
  status: MilestoneStatus | string;
  animate?: boolean;
  className?: string;
  id?: string;
}

// ─── Status Config ───────────────────────────────────────────────────────────

interface PillConfig {
  label: string;
  bg: string;
  text: string;
  ring: string;
  shouldPulse: boolean;
}

const statusMap: Record<string, PillConfig> = {
  pending: {
    label: 'Pending',
    bg: '#1F2937',
    text: '#9CA3AF',
    ring: '#374151',
    shouldPulse: false,
  },
  submitted: {
    label: 'Submitted',
    bg: 'rgba(124, 58, 237, 0.12)',
    text: '#A78BFA',
    ring: 'rgba(124, 58, 237, 0.3)',
    shouldPulse: false,
  },
  under_review: {
    label: 'Under Review',
    bg: 'rgba(217, 119, 6, 0.10)',
    text: '#FCD34D',
    ring: 'rgba(217, 119, 6, 0.28)',
    shouldPulse: false,
  },
  approved: {
    label: 'Approved',
    bg: 'rgba(13, 148, 136, 0.12)',
    text: '#2DD4BF',
    ring: 'rgba(13, 148, 136, 0.3)',
    shouldPulse: true,
  },
  releasing: {
    label: 'Releasing',
    bg: 'rgba(37, 99, 235, 0.10)',
    text: '#93C5FD',
    ring: 'rgba(37, 99, 235, 0.28)',
    shouldPulse: true,
  },
  completed: {
    label: 'Completed',
    bg: 'rgba(13, 148, 136, 0.10)',
    text: '#2DD4BF',
    ring: 'rgba(13, 148, 136, 0.25)',
    shouldPulse: false,
  },
  disputed: {
    label: 'Disputed',
    bg: 'rgba(220, 38, 38, 0.10)',
    text: '#FCA5A5',
    ring: 'rgba(220, 38, 38, 0.3)',
    shouldPulse: true,
  },
  overdue: {
    label: 'Overdue',
    bg: 'rgba(220, 38, 38, 0.10)',
    text: '#F87171',
    ring: 'rgba(220, 38, 38, 0.25)',
    shouldPulse: false,
  },
};

const fallbackConfig: PillConfig = {
  label: 'Unknown',
  bg: '#1F2937',
  text: '#4B5563',
  ring: '#374151',
  shouldPulse: false,
};

// ─── StatusPill ──────────────────────────────────────────────────────────────

export function StatusPill({
  status,
  animate = false,
  className = '',
  id,
}: StatusPillProps) {
  const config = statusMap[status.toLowerCase()] ?? fallbackConfig;
  const shouldAnimate = animate && config.shouldPulse;

  return (
    <span
      id={id}
      className={['inline-flex items-center', className].filter(Boolean).join(' ')}
      style={{ display: 'inline-flex' }}
    >
      <span
        style={{
          display: 'inline-flex',
          alignItems: 'center',
          gap: '6px',
          paddingTop: '2px',
          paddingBottom: '2px',
          paddingLeft: '10px',
          paddingRight: '10px',
          borderRadius: '999px',
          backgroundColor: config.bg,
          color: config.text,
          border: `1px solid ${config.ring}`,
          fontSize: '11px',
          fontWeight: 600,
          letterSpacing: '0.03em',
          textTransform: 'uppercase',
          whiteSpace: 'nowrap',
          userSelect: 'none',
        }}
      >
        {/* Indicator dot */}
        <span
          style={{
            width: '5px',
            height: '5px',
            borderRadius: '50%',
            backgroundColor: config.text,
            flexShrink: 0,
            animation: shouldAnimate ? 'pulse-blue 1.6s ease-in-out infinite' : undefined,
          }}
          aria-hidden="true"
        />
        {config.label}
      </span>
    </span>
  );
}

export default StatusPill;
