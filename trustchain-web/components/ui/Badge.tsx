import React from 'react';

// ─── Types ───────────────────────────────────────────────────────────────────

export type BadgeVariant =
  | 'active'
  | 'completed'
  | 'frozen'
  | 'pending'
  | 'submitted'
  | 'under_review'
  | 'approved'
  | 'releasing'
  | 'overdue'
  | 'disputed';

export interface BadgeProps {
  variant: BadgeVariant;
  text?: string;
  className?: string;
  id?: string;
}

// ─── Variant Config ──────────────────────────────────────────────────────────

interface VariantConfig {
  label: string;
  dot: string;           // dot bg color class
  bg: string;            // pill background
  textColor: string;     // pill text color
  border: string;        // pill border
  pulse?: boolean;       // animate dot with pulse
}

const variantConfig: Record<BadgeVariant, VariantConfig> = {
  active: {
    label: 'Active',
    dot: 'bg-[#2563EB]',
    bg: 'bg-[rgba(37,99,235,0.12)]',
    textColor: 'text-[#60A5FA]',
    border: 'border-[rgba(37,99,235,0.3)]',
  },
  completed: {
    label: 'Completed',
    dot: 'bg-[#0D9488]',
    bg: 'bg-[rgba(13,148,136,0.12)]',
    textColor: 'text-[#2DD4BF]',
    border: 'border-[rgba(13,148,136,0.3)]',
  },
  frozen: {
    label: 'Frozen',
    dot: 'bg-[#DC2626]',
    bg: 'bg-[rgba(220,38,38,0.10)]',
    textColor: 'text-[#F87171]',
    border: 'border-[rgba(220,38,38,0.25)]',
  },
  pending: {
    label: 'Pending',
    dot: 'bg-[#6B7280]',
    bg: 'bg-[rgba(107,114,128,0.10)]',
    textColor: 'text-[#9CA3AF]',
    border: 'border-[rgba(107,114,128,0.25)]',
  },
  submitted: {
    label: 'Submitted',
    dot: 'bg-[#7C3AED]',
    bg: 'bg-[rgba(124,58,237,0.10)]',
    textColor: 'text-[#A78BFA]',
    border: 'border-[rgba(124,58,237,0.25)]',
  },
  under_review: {
    label: 'Under Review',
    dot: 'bg-[#D97706]',
    bg: 'bg-[rgba(217,119,6,0.10)]',
    textColor: 'text-[#FCD34D]',
    border: 'border-[rgba(217,119,6,0.25)]',
  },
  approved: {
    label: 'Approved',
    dot: 'bg-[#0D9488]',
    bg: 'bg-[rgba(13,148,136,0.12)]',
    textColor: 'text-[#2DD4BF]',
    border: 'border-[rgba(13,148,136,0.3)]',
    pulse: true,
  },
  releasing: {
    label: 'Releasing',
    dot: 'bg-[#2563EB]',
    bg: 'bg-[rgba(37,99,235,0.10)]',
    textColor: 'text-[#93C5FD]',
    border: 'border-[rgba(37,99,235,0.25)]',
    pulse: true,
  },
  overdue: {
    label: 'Overdue',
    dot: 'bg-[#DC2626]',
    bg: 'bg-[rgba(220,38,38,0.10)]',
    textColor: 'text-[#F87171]',
    border: 'border-[rgba(220,38,38,0.25)]',
  },
  disputed: {
    label: 'Disputed',
    dot: 'bg-[#DC2626]',
    bg: 'bg-[rgba(220,38,38,0.10)]',
    textColor: 'text-[#FCA5A5]',
    border: 'border-[rgba(220,38,38,0.3)]',
    pulse: true,
  },
};

// ─── Badge ────────────────────────────────────────────────────────────────────

export function Badge({ variant, text, className = '', id }: BadgeProps) {
  const config = variantConfig[variant];
  const label = text ?? config.label;

  return (
    <span
      id={id}
      className={[
        'inline-flex items-center gap-1.5',
        'px-2.5 py-0.5',
        'rounded-full',
        'border',
        'text-xs font-medium',
        'whitespace-nowrap',
        'select-none',
        config.bg,
        config.textColor,
        config.border,
        className,
      ]
        .filter(Boolean)
        .join(' ')}
    >
      {/* Status dot */}
      <span
        className={[
          'w-1.5 h-1.5 rounded-full flex-shrink-0',
          config.dot,
          config.pulse ? 'animate-pulse' : '',
        ]
          .filter(Boolean)
          .join(' ')}
        aria-hidden="true"
      />
      {label}
    </span>
  );
}

export default Badge;
