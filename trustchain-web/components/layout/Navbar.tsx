'use client';

import React, { useEffect, useRef, useState } from 'react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { useAccount } from 'wagmi';
import { WalletButton } from '@/components/blockchain/WalletButton';

// ─── Types ───────────────────────────────────────────────────────────────────

interface NavLink {
  href: string;
  label: string;
}

const NAV_LINKS: NavLink[] = [
  { href: '/', label: 'Home' },
  { href: '/campaigns', label: 'Campaigns' },
  { href: '/how-it-works', label: 'How It Works' },
  { href: '/apply', label: 'Apply' },
];

// ─── Chain Link Icon ──────────────────────────────────────────────────────────

function ChainLinkIcon() {
  return (
    <svg
      width="22"
      height="22"
      viewBox="0 0 24 24"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      aria-hidden="true"
      className="flex-shrink-0"
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
  );
}

// ─── Hamburger Icon ───────────────────────────────────────────────────────────

function HamburgerIcon({ open }: { open: boolean }) {
  return (
    <svg
      width="22"
      height="22"
      viewBox="0 0 22 22"
      fill="none"
      aria-hidden="true"
      className="transition-transform duration-200"
    >
      {open ? (
        <>
          <line x1="4" y1="4" x2="18" y2="18" stroke="#F9FAFB" strokeWidth="2" strokeLinecap="round" />
          <line x1="18" y1="4" x2="4" y2="18" stroke="#F9FAFB" strokeWidth="2" strokeLinecap="round" />
        </>
      ) : (
        <>
          <line x1="3" y1="6" x2="19" y2="6" stroke="#F9FAFB" strokeWidth="2" strokeLinecap="round" />
          <line x1="3" y1="11" x2="19" y2="11" stroke="#F9FAFB" strokeWidth="2" strokeLinecap="round" />
          <line x1="3" y1="16" x2="19" y2="16" stroke="#F9FAFB" strokeWidth="2" strokeLinecap="round" />
        </>
      )}
    </svg>
  );
}

// ─── Navbar ───────────────────────────────────────────────────────────────────

export function Navbar() {
  const pathname = usePathname();
  const { address } = useAccount();
  const [scrolled, setScrolled] = useState(false);
  const [mobileOpen, setMobileOpen] = useState(false);
  const mobileMenuRef = useRef<HTMLDivElement>(null);
  const adminWallet = process.env.NEXT_PUBLIC_ADMIN_WALLET;
  const isAdmin =
    address && adminWallet && address.toLowerCase() === adminWallet.toLowerCase();

  // Glassmorphism on scroll
  useEffect(() => {
    const handler = () => setScrolled(window.scrollY > 12);
    window.addEventListener('scroll', handler, { passive: true });
    handler(); // init
    return () => window.removeEventListener('scroll', handler);
  }, []);

  // Close mobile menu on route change
  useEffect(() => {
    setMobileOpen(false);
  }, [pathname]);

  // Close mobile menu on outside click
  useEffect(() => {
    if (!mobileOpen) return;
    const handler = (e: MouseEvent) => {
      if (mobileMenuRef.current && !mobileMenuRef.current.contains(e.target as Node)) {
        setMobileOpen(false);
      }
    };
    document.addEventListener('mousedown', handler);
    return () => document.removeEventListener('mousedown', handler);
  }, [mobileOpen]);

  return (
    <>
      <header
        id="navbar"
        className={[
          'fixed top-0 left-0 right-0 z-50',
          'transition-all duration-300 ease-in-out',
          scrolled
            ? 'bg-[rgba(10,15,30,0.75)] backdrop-blur-xl border-b border-[#1F2937]/80 shadow-[0_4px_24px_rgba(0,0,0,0.3)]'
            : 'bg-[#0A0F1E] border-b border-[#1F2937]',
        ].join(' ')}
      >
        <nav
          className="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8 h-16 flex items-center justify-between gap-4"
          aria-label="Main navigation"
        >
          {/* ── Left: Logo ─────────────────────────────────────────────── */}
          <Link
            href="/"
            className="flex items-center gap-2 flex-shrink-0 group"
            id="nav-logo"
            aria-label="TrustChain home"
          >
            <ChainLinkIcon />
            <span className="text-lg font-bold tracking-tight select-none">
              <span className="text-[#F9FAFB] group-hover:text-white transition-colors">
                Trust
              </span>
              <span className="text-[#2563EB] group-hover:text-[#60A5FA] transition-colors">
                Chain
              </span>
            </span>
          </Link>

          {/* ── Center: Nav Links (desktop) ─────────────────────────────── */}
          <div className="hidden md:flex items-center gap-1" role="list">
            {NAV_LINKS.map((link) => {
              const isActive =
                link.href === '/'
                  ? pathname === '/'
                  : pathname === link.href || pathname.startsWith(link.href + '/');

              return (
                <Link
                  key={link.href}
                  href={link.href}
                  role="listitem"
                  id={`nav-link-${link.label.toLowerCase().replace(/\s+/g, '-')}`}
                  className={[
                    'relative px-4 py-2 rounded-lg text-sm font-medium',
                    'transition-all duration-200',
                    isActive
                      ? 'text-[#F9FAFB]'
                      : 'text-[#9CA3AF] hover:text-[#F9FAFB] hover:bg-[#1F2937]/60',
                  ].join(' ')}
                >
                  {link.label}
                  {/* Active underline indicator */}
                  {isActive && (
                    <span
                      className="absolute bottom-0 left-1/2 -translate-x-1/2 h-0.5 w-5 rounded-full bg-[#2563EB]"
                      aria-hidden="true"
                    />
                  )}
                </Link>
              );
            })}
          </div>

          {/* ── Right: Wallet + hamburger ────────────────────────────────── */}
          <div className="flex items-center gap-3">
            {/* Wallet button — hidden on mobile */}
            <div className="hidden sm:block">
              <WalletButton />
            </div>
            {isAdmin && (
              <Link
                href="/admin"
                id="nav-link-admin"
                className={[
                  'hidden sm:flex items-center gap-1.5 px-3 py-1.5 rounded-lg',
                  'text-xs font-medium transition-all duration-200',
                  'border border-[#374151] text-[#9CA3AF]',
                  'hover:border-[#2563EB] hover:text-[#60A5FA]',
                  'hover:bg-[rgba(37,99,235,0.08)]',
                ].join(' ')}
              >
                <svg
                  className="w-3 h-3"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                  />
                </svg>
                Admin
              </Link>
            )}

            {/* Hamburger — mobile only */}
            <button
              id="nav-hamburger"
              type="button"
              onClick={() => setMobileOpen((v) => !v)}
              aria-expanded={mobileOpen}
              aria-controls="mobile-menu"
              aria-label={mobileOpen ? 'Close menu' : 'Open menu'}
              className={[
                'md:hidden p-2 rounded-lg',
                'border border-[#1F2937]',
                'hover:bg-[#1F2937] transition-colors duration-200',
              ].join(' ')}
            >
              <HamburgerIcon open={mobileOpen} />
            </button>
          </div>
        </nav>

        {/* ── Mobile Menu ──────────────────────────────────────────────── */}
        <div
          id="mobile-menu"
          ref={mobileMenuRef}
          aria-hidden={!mobileOpen}
          className={[
            'md:hidden overflow-hidden transition-all duration-300 ease-in-out',
            'border-t border-[#1F2937]',
            'bg-[rgba(10,15,30,0.95)] backdrop-blur-xl',
            mobileOpen ? 'max-h-80 opacity-100' : 'max-h-0 opacity-0',
          ].join(' ')}
        >
          <div className="px-4 py-4 flex flex-col gap-1">
            {NAV_LINKS.map((link) => {
              const isActive =
                link.href === '/'
                  ? pathname === '/'
                  : pathname.startsWith(link.href);
              return (
                <Link
                  key={link.href}
                  href={link.href}
                  id={`mobile-nav-${link.label.toLowerCase().replace(/\s+/g, '-')}`}
                  className={[
                    'flex items-center gap-3 px-4 py-3 rounded-lg text-sm font-medium',
                    'transition-colors duration-200',
                    isActive
                      ? 'bg-[rgba(37,99,235,0.08)] text-[#60A5FA] border border-[#2563EB]/20'
                      : 'text-[#9CA3AF] hover:bg-[#1F2937] hover:text-[#F9FAFB]',
                  ].join(' ')}
                >
                  {isActive && (
                    <span className="w-1 h-1 rounded-full bg-[#2563EB]" aria-hidden="true" />
                  )}
                  {link.label}
                </Link>
              );
            })}

            {isAdmin && (
              <Link
                href="/admin"
                id="mobile-nav-admin"
                className={[
                  'flex items-center gap-1.5 px-4 py-3 rounded-lg',
                  'text-xs font-medium transition-all duration-200',
                  'border border-[#374151] text-[#9CA3AF]',
                  'hover:border-[#2563EB] hover:text-[#60A5FA]',
                  'hover:bg-[rgba(37,99,235,0.08)]',
                ].join(' ')}
              >
                <svg
                  className="w-3 h-3"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                  />
                </svg>
                Admin
              </Link>
            )}

            {/* Wallet on mobile */}
            <div className="pt-3 border-t border-[#1F2937] mt-2">
              <WalletButton />
            </div>
          </div>
        </div>
      </header>

      {/* Spacer so page content doesn't hide under fixed navbar */}
      <div className="h-16" aria-hidden="true" />
    </>
  );
}

export default Navbar;
