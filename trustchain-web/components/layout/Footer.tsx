import React from 'react';
import Link from 'next/link';

// ─── Types ───────────────────────────────────────────────────────────────────

interface FooterLink {
  href: string;
  label: string;
  external?: boolean;
}

interface StatItem {
  label: string;
  value: string;
}

// ─── Data ─────────────────────────────────────────────────────────────────────

const COLUMN_LINKS: FooterLink[] = [
  { href: '/campaigns', label: 'Campaigns' },
  { href: '/how-it-works', label: 'How It Works' },
  { href: '/apply', label: 'Apply for Funding' },
  {
    href: 'https://github.com/trustchain/audit-report',
    label: 'Audit Report',
    external: true,
  },
];

// Placeholder stats — will be populated from The Graph subgraph later
const PLATFORM_STATS: StatItem[] = [
  { label: 'Total Raised', value: '0 ETH' },
  { label: 'Campaigns Completed', value: '0' },
  { label: 'Total Donors', value: '0' },
];

// ─── Base Badge ───────────────────────────────────────────────────────────────

function BaseBadge() {
  return (
    <span
      className={[
        'inline-flex items-center gap-1.5',
        'px-2.5 py-1 rounded-full',
        'border border-[#1F2937]',
        'bg-[#111827]',
        'text-xs font-medium text-[#9CA3AF]',
        'select-none',
      ].join(' ')}
    >
      {/* Base network blue circle logo mark */}
      <svg
        width="14"
        height="14"
        viewBox="0 0 14 14"
        fill="none"
        aria-label="Base network"
      >
        <circle cx="7" cy="7" r="7" fill="#0052FF" />
        <path
          d="M7 3.5C5.067 3.5 3.5 5.067 3.5 7C3.5 8.933 5.067 10.5 7 10.5C8.843 10.5 10.356 9.077 10.487 7.27H7.583V6.73H10.487C10.356 4.923 8.843 3.5 7 3.5Z"
          fill="white"
        />
      </svg>
      Built on Base
    </span>
  );
}

// ─── Footer ───────────────────────────────────────────────────────────────────

export function Footer() {
  const currentYear = new Date().getFullYear();

  return (
    <footer
      id="site-footer"
      className="bg-[#111827] border-t border-[#1F2937] mt-auto"
      aria-label="Site footer"
    >
      <div className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 pt-12 pb-8">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-10 lg:gap-16">

          {/* ── Col 1: Brand ──────────────────────────────────────────── */}
          <div className="flex flex-col gap-4">
            {/* Logo */}
            <Link
              href="/"
              className="inline-flex items-center gap-2 group w-fit"
              aria-label="TrustChain home"
            >
              <svg
                width="20"
                height="20"
                viewBox="0 0 24 24"
                fill="none"
                aria-hidden="true"
              >
                <path
                  d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"
                  stroke="#2563EB"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
                <path
                  d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"
                  stroke="#60A5FA"
                  strokeWidth="2"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
              <span className="text-base font-bold tracking-tight">
                <span className="text-[#F9FAFB]">Trust</span>
                <span className="text-[#2563EB]">Chain</span>
              </span>
            </Link>

            {/* Description */}
            <p className="text-sm text-[#9CA3AF] leading-relaxed max-w-xs">
              Transparent, milestone-driven blockchain donations on Base — where
              every fund release is verified and every donor is protected.
            </p>

            {/* Base badge */}
            <BaseBadge />
          </div>

          {/* ── Col 2: Links ──────────────────────────────────────────── */}
          <div className="flex flex-col gap-4">
            <h3 className="text-xs font-semibold uppercase tracking-widest text-[#4B5563]">
              Platform
            </h3>
            <nav aria-label="Footer links">
              <ul className="flex flex-col gap-2.5" role="list">
                {COLUMN_LINKS.map((link) => (
                  <li key={link.href} role="listitem">
                    {link.external ? (
                      <a
                        href={link.href}
                        target="_blank"
                        rel="noopener noreferrer"
                        id={`footer-link-${link.label.toLowerCase().replace(/\s+/g, '-')}`}
                        className={[
                          'inline-flex items-center gap-1.5',
                          'text-sm text-[#9CA3AF]',
                          'hover:text-[#F9FAFB] transition-colors duration-200',
                          'group',
                        ].join(' ')}
                      >
                        {link.label}
                        {/* External link arrow */}
                        <svg
                          className="w-3 h-3 opacity-40 group-hover:opacity-80 transition-opacity"
                          viewBox="0 0 12 12"
                          fill="none"
                          aria-hidden="true"
                        >
                          <path
                            d="M2.5 9.5L9.5 2.5M9.5 2.5H5.5M9.5 2.5V6.5"
                            stroke="currentColor"
                            strokeWidth="1.5"
                            strokeLinecap="round"
                            strokeLinejoin="round"
                          />
                        </svg>
                      </a>
                    ) : (
                      <Link
                        href={link.href}
                        id={`footer-link-${link.label.toLowerCase().replace(/\s+/g, '-')}`}
                        className="text-sm text-[#9CA3AF] hover:text-[#F9FAFB] transition-colors duration-200"
                      >
                        {link.label}
                      </Link>
                    )}
                  </li>
                ))}
              </ul>
            </nav>
          </div>

          {/* ── Col 3: Platform Stats ─────────────────────────────────── */}
          <div className="flex flex-col gap-4">
            <h3 className="text-xs font-semibold uppercase tracking-widest text-[#4B5563]">
              Platform Stats
            </h3>
            <ul className="flex flex-col gap-3" role="list">
              {PLATFORM_STATS.map((stat) => (
                <li
                  key={stat.label}
                  role="listitem"
                  className="flex items-center justify-between py-2.5 border-b border-[#1F2937] last:border-0"
                >
                  <span className="text-sm text-[#9CA3AF]">{stat.label}</span>
                  <span
                    className="text-sm font-semibold text-[#F9FAFB] tabular-nums"
                    id={`stat-${stat.label.toLowerCase().replace(/\s+/g, '-')}`}
                    aria-live="polite"
                  >
                    {stat.value}
                  </span>
                </li>
              ))}
            </ul>
            <p className="text-xs text-[#4B5563]">
              Live data from The Graph subgraph — coming soon.
            </p>
          </div>
        </div>

        {/* ── Bottom bar ────────────────────────────────────────────── */}
        <div className="mt-10 pt-6 border-t border-[#1F2937] flex flex-col sm:flex-row items-center justify-between gap-3">
          <p className="text-xs text-[#4B5563]">
            © {currentYear} TrustChain. Open source. MIT licensed.
          </p>
          <div className="flex items-center gap-4">
            <a
              href="https://github.com/trustchain"
              target="_blank"
              rel="noopener noreferrer"
              aria-label="GitHub"
              className="text-[#4B5563] hover:text-[#9CA3AF] transition-colors duration-200"
            >
              <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor" aria-hidden="true">
                <path d="M12 2C6.477 2 2 6.477 2 12c0 4.418 2.865 8.165 6.839 9.489.5.092.682-.217.682-.482 0-.237-.008-.866-.013-1.7-2.782.603-3.369-1.34-3.369-1.34-.454-1.156-1.11-1.463-1.11-1.463-.908-.62.069-.608.069-.608 1.003.07 1.531 1.03 1.531 1.03.892 1.529 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.11-4.555-4.943 0-1.091.39-1.984 1.029-2.683-.103-.253-.446-1.27.098-2.647 0 0 .84-.269 2.75 1.025A9.578 9.578 0 0 1 12 6.836c.85.004 1.705.115 2.504.337 1.909-1.294 2.747-1.025 2.747-1.025.546 1.377.202 2.394.1 2.647.64.699 1.028 1.592 1.028 2.683 0 3.842-2.339 4.687-4.566 4.935.359.309.678.919.678 1.852 0 1.336-.012 2.415-.012 2.743 0 .267.18.578.688.48C19.138 20.161 22 16.416 22 12c0-5.523-4.477-10-10-10z" />
              </svg>
            </a>
          </div>
        </div>
      </div>
    </footer>
  );
}

export default Footer;
