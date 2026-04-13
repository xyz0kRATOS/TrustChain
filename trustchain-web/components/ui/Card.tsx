import React from 'react';

// ─── Types ───────────────────────────────────────────────────────────────────

export interface CardProps {
  children: React.ReactNode;
  className?: string;
  hover?: boolean;
  onClick?: () => void;
  id?: string;
  as?: 'div' | 'article' | 'section' | 'li';
}

// ─── Card ─────────────────────────────────────────────────────────────────────

export function Card({
  children,
  className = '',
  hover = false,
  onClick,
  id,
  as: Tag = 'div',
}: CardProps) {
  const baseStyles = [
    'bg-[#111827]',
    'border border-[#1F2937]',
    'rounded-xl',
    'transition-all duration-300 ease-in-out',
  ].join(' ');

  const hoverStyles = hover
    ? [
        'cursor-pointer',
        'hover:border-[#2563EB]/40',
        'hover:shadow-[0_0_24px_rgba(37,99,235,0.12),0_0_8px_rgba(37,99,235,0.06)]',
        'hover:-translate-y-0.5',
      ].join(' ')
    : '';

  const clickableStyles = onClick && !hover ? 'cursor-pointer' : '';

  return (
    <Tag
      id={id}
      onClick={onClick}
      className={[baseStyles, hoverStyles, clickableStyles, className]
        .filter(Boolean)
        .join(' ')}
    >
      {children}
    </Tag>
  );
}

export default Card;
