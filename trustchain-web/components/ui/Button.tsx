import React from 'react';

// ─── Types ───────────────────────────────────────────────────────────────────

export type ButtonVariant = 'primary' | 'secondary' | 'ghost' | 'danger';
export type ButtonSize = 'sm' | 'md' | 'lg';

export interface ButtonProps {
  variant?: ButtonVariant;
  size?: ButtonSize;
  loading?: boolean;
  disabled?: boolean;
  onClick?: (event: React.MouseEvent<HTMLButtonElement>) => void;
  children: React.ReactNode;
  type?: 'button' | 'submit' | 'reset';
  className?: string;
  id?: string;
}

// ─── Style Maps ──────────────────────────────────────────────────────────────

const variantStyles: Record<ButtonVariant, string> = {
  primary: [
    'bg-[#2563EB] text-[#F9FAFB]',
    'hover:bg-[#1D4ED8]',
    'active:bg-[#1E40AF]',
    'shadow-[0_0_0_0_rgba(37,99,235,0)] hover:shadow-[0_0_16px_rgba(37,99,235,0.4)]',
    'border border-transparent',
  ].join(' '),

  secondary: [
    'bg-transparent text-[#2563EB]',
    'border border-[#2563EB]',
    'hover:bg-[rgba(37,99,235,0.08)]',
    'active:bg-[rgba(37,99,235,0.15)]',
  ].join(' '),

  ghost: [
    'bg-transparent text-[#9CA3AF]',
    'border border-transparent',
    'hover:bg-[#1F2937] hover:text-[#F9FAFB]',
    'active:bg-[#374151]',
  ].join(' '),

  danger: [
    'bg-[#DC2626] text-[#F9FAFB]',
    'hover:bg-[#B91C1C]',
    'active:bg-[#991B1B]',
    'shadow-[0_0_0_0_rgba(220,38,38,0)] hover:shadow-[0_0_16px_rgba(220,38,38,0.35)]',
    'border border-transparent',
  ].join(' '),
};

const sizeStyles: Record<ButtonSize, string> = {
  sm: 'h-8 px-3 text-xs gap-1.5',
  md: 'h-10 px-4 text-sm gap-2',
  lg: 'h-12 px-6 text-base gap-2.5',
};

const spinnerSizes: Record<ButtonSize, string> = {
  sm: 'w-3 h-3',
  md: 'w-4 h-4',
  lg: 'w-5 h-5',
};

// ─── Spinner ─────────────────────────────────────────────────────────────────

function Spinner({ size }: { size: ButtonSize }) {
  return (
    <svg
      className={`${spinnerSizes[size]} animate-spin flex-shrink-0`}
      style={{ animation: 'spin 0.7s linear infinite' }}
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      aria-hidden="true"
    >
      <circle
        className="opacity-25"
        cx="12"
        cy="12"
        r="10"
        stroke="currentColor"
        strokeWidth="4"
      />
      <path
        className="opacity-75"
        fill="currentColor"
        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"
      />
    </svg>
  );
}

// ─── Button ──────────────────────────────────────────────────────────────────

export function Button({
  variant = 'primary',
  size = 'md',
  loading = false,
  disabled = false,
  onClick,
  children,
  type = 'button',
  className = '',
  id,
}: ButtonProps) {
  const isDisabled = disabled || loading;

  const baseStyles = [
    'inline-flex items-center justify-center',
    'rounded-lg font-medium',
    'transition-all duration-200 ease-in-out',
    'focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-[#2563EB]',
    'select-none cursor-pointer',
    'relative overflow-hidden',
  ].join(' ');

  const disabledStyles = isDisabled
    ? 'opacity-50 cursor-not-allowed pointer-events-none'
    : '';

  return (
    <button
      id={id}
      type={type}
      disabled={isDisabled}
      onClick={isDisabled ? undefined : onClick}
      aria-disabled={isDisabled}
      aria-busy={loading}
      className={[
        baseStyles,
        variantStyles[variant],
        sizeStyles[size],
        disabledStyles,
        className,
      ]
        .filter(Boolean)
        .join(' ')}
    >
      {loading && <Spinner size={size} />}
      <span className={loading ? 'opacity-80' : ''}>{children}</span>
    </button>
  );
}

export default Button;
