'use client';

import React, { useEffect, useRef, useState } from 'react';
import Link from 'next/link';
import { CampaignCard, Campaign } from '@/components/campaign/CampaignCard';

// ─── Mock Data ────────────────────────────────────────────────────────────────

const FEATURED_CAMPAIGNS: Campaign[] = [
  {
    id: 'clean-water-nairobi',
    name: 'Clean Water Project — Nairobi',
    description:
      'Bringing clean, filtered water infrastructure to 3 underserved communities across Nairobi, reducing waterborne illness and travel burden.',
    goalAmount: 5,
    totalRaised: 3.2,
    status: 'active',
    donorCount: 47,
    milestoneCount: 3,
  },
  {
    id: 'school-rebuild-kenya',
    name: 'School Rebuild — Rural Kenya',
    description:
      'Rebuilding a primary school destroyed by floods, restoring education for over 400 children with safe, permanent classrooms.',
    goalAmount: 8,
    totalRaised: 8,
    status: 'completed',
    donorCount: 124,
    milestoneCount: 4,
  },
  {
    id: 'medical-aid-gaza',
    name: 'Medical Aid — Gaza',
    description:
      'Delivering critical medical supplies and field equipment to overwhelmed hospitals and clinics in the Gaza Strip.',
    goalAmount: 10,
    totalRaised: 6.1,
    status: 'active',
    donorCount: 89,
    milestoneCount: 3,
  },
];

const STATS = [
  { value: '12', label: 'Campaigns Live', suffix: '' },
  { value: '48,200', label: 'USD Raised', prefix: '$' },
  { value: '847', label: 'Donors', suffix: '' },
  { value: '3', label: 'Completed', suffix: '' },
];

const HOW_IT_WORKS_STEPS = [
  {
    n: 1,
    title: 'Campaign Verified',
    desc: 'Your team reviews applications and supporting documents before anything goes live on the platform.',
  },
  {
    n: 2,
    title: 'Donors Contribute',
    desc: 'ETH sent directly to the campaign smart contract. Every donation mints an NFT receipt — permanent proof on-chain.',
  },
  {
    n: 3,
    title: 'Evidence Submitted',
    desc: 'The campaign creator submits verified proof of milestone completion — photos, invoices, reports.',
  },
  {
    n: 4,
    title: 'Admin Approves',
    desc: 'A 2-of-3 multi-sig approval triggers a 48-hour countdown — a safety window for donors.',
  },
  {
    n: 5,
    title: 'Funds Released',
    desc: 'After 48 hours the smart contract automatically releases funds to the campaign creator. Transparent and automatic.',
  },
];

const TRUST_FEATURES = [
  {
    icon: '🔒',
    title: 'Milestone-locked Funds',
    desc: 'Money only moves when verified milestones complete. No exceptions, no workarounds.',
  },
  {
    icon: '⏱️',
    title: '48-hour Safety Window',
    desc: 'Every fund release has a 48-hour timelock. Enough time for donors to raise an alarm.',
  },
  {
    icon: '📄',
    title: 'NFT Donation Receipts',
    desc: 'Every donation mints a permanent, non-transferable proof of contribution on the blockchain.',
  },
  {
    icon: '⭐',
    title: 'On-chain Reputation',
    desc: 'Campaign creators build verifiable track records. Forever on the blockchain — no hiding a bad history.',
  },
];

// ─── Particle Canvas ──────────────────────────────────────────────────────────

function ParticleField() {
  const canvasRef = useRef<HTMLCanvasElement>(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const ctx = canvas.getContext('2d');
    if (!ctx) return;

    let animId: number;
    let W = 0, H = 0;

    interface Particle {
      x: number; y: number; r: number;
      vx: number; vy: number; alpha: number;
    }

    let particles: Particle[] = [];

    const resize = () => {
      W = canvas.width = canvas.offsetWidth;
      H = canvas.height = canvas.offsetHeight;
    };

    const init = () => {
      resize();
      particles = Array.from({ length: 55 }, () => ({
        x: Math.random() * W,
        y: Math.random() * H,
        r: Math.random() * 1.4 + 0.3,
        vx: (Math.random() - 0.5) * 0.18,
        vy: (Math.random() - 0.5) * 0.18,
        alpha: Math.random() * 0.4 + 0.08,
      }));
    };

    const draw = () => {
      ctx.clearRect(0, 0, W, H);
      for (const p of particles) {
        p.x += p.vx;
        p.y += p.vy;
        if (p.x < 0) p.x = W;
        if (p.x > W) p.x = 0;
        if (p.y < 0) p.y = H;
        if (p.y > H) p.y = 0;

        ctx.beginPath();
        ctx.arc(p.x, p.y, p.r, 0, Math.PI * 2);
        ctx.fillStyle = `rgba(96, 165, 250, ${p.alpha})`;
        ctx.fill();
      }
      animId = requestAnimationFrame(draw);
    };

    init();
    draw();
    window.addEventListener('resize', resize);
    return () => {
      cancelAnimationFrame(animId);
      window.removeEventListener('resize', resize);
    };
  }, []);

  return (
    <canvas
      ref={canvasRef}
      className="absolute inset-0 w-full h-full pointer-events-none"
      aria-hidden="true"
    />
  );
}

// ─── Scroll-triggered counter ─────────────────────────────────────────────────

function AnimatedNumber({ target, prefix = '', suffix = '' }: { target: string; prefix?: string; suffix?: string }) {
  const [display, setDisplay] = useState('0');
  const ref = useRef<HTMLSpanElement>(null);
  const animated = useRef(false);

  useEffect(() => {
    const el = ref.current;
    if (!el) return;
    const numeric = parseFloat(target.replace(/,/g, ''));
    if (isNaN(numeric)) { setDisplay(target); return; }

    const obs = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting && !animated.current) {
          animated.current = true;
          const duration = 1400;
          const start = performance.now();
          const tick = (now: number) => {
            const t = Math.min((now - start) / duration, 1);
            const eased = 1 - Math.pow(1 - t, 3);
            const cur = Math.round(eased * numeric);
            setDisplay(cur.toLocaleString());
            if (t < 1) requestAnimationFrame(tick);
          };
          requestAnimationFrame(tick);
          obs.disconnect();
        }
      },
      { threshold: 0.5 }
    );
    obs.observe(el);
    return () => obs.disconnect();
  }, [target]);

  return (
    <span ref={ref} aria-label={`${prefix}${target}${suffix}`}>
      {prefix}{display}{suffix}
    </span>
  );
}

// ─── Scroll Indicator ─────────────────────────────────────────────────────────

function ScrollIndicator() {
  return (
    <div className="flex flex-col items-center gap-1.5 text-[#4B5563]" aria-hidden="true">
      <span className="text-xs tracking-widest uppercase">Scroll</span>
      <svg
        width="20"
        height="20"
        viewBox="0 0 20 20"
        fill="none"
        style={{ animation: 'bounce-y 1.6s ease-in-out infinite' }}
      >
        <path d="M5 8l5 5 5-5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
      </svg>
      <style>{`
        @keyframes bounce-y {
          0%, 100% { transform: translateY(0); }
          50% { transform: translateY(6px); }
        }
      `}</style>
    </div>
  );
}

// ─── Section: Hero ────────────────────────────────────────────────────────────

function HeroSection() {
  return (
    <section
      id="hero"
      className="relative flex flex-col items-center justify-center min-h-screen px-4 text-center overflow-hidden"
      aria-labelledby="hero-heading"
    >
      {/* Radial glow background */}
      <div
        className="absolute inset-0 pointer-events-none"
        style={{
          background:
            'radial-gradient(ellipse 70% 55% at 50% 42%, rgba(37,99,235,0.13) 0%, rgba(10,15,30,0) 70%)',
        }}
        aria-hidden="true"
      />

      {/* Floating particles */}
      <ParticleField />

      {/* Content layer */}
      <div className="relative z-10 flex flex-col items-center gap-8 max-w-4xl mx-auto">

        {/* Eyebrow chip */}
        <div className="inline-flex items-center gap-2 px-4 py-1.5 rounded-full border border-[#2563EB]/25 bg-[rgba(37,99,235,0.07)] text-[#60A5FA] text-xs font-medium tracking-wide">
          <span className="w-1.5 h-1.5 rounded-full bg-[#2563EB] animate-pulse" aria-hidden="true" />
          Now live on Base Sepolia
        </div>

        {/* Main heading */}
        <h1
          id="hero-heading"
          className="text-5xl sm:text-6xl lg:text-7xl font-extrabold leading-tight tracking-tight"
        >
          <span className="text-[#F9FAFB]">Donations You Can</span>
          <br />
          <span className="text-[#2563EB]">Actually Verify</span>
        </h1>

        {/* Subheading */}
        <p className="text-lg sm:text-xl text-[#9CA3AF] max-w-2xl leading-relaxed">
          Every donation tracked on-chain. Every milestone verified before funds
          release.{' '}
          <span className="text-[#F9FAFB]/70">
            TrustChain makes charity transparent by default.
          </span>
        </p>

        {/* CTA Buttons */}
        <div className="flex flex-col sm:flex-row items-center gap-3 w-full sm:w-auto">
          <Link
            href="/campaigns"
            id="hero-cta-primary"
            className={[
              'inline-flex items-center justify-center gap-2',
              'h-12 px-8 rounded-xl',
              'bg-[#2563EB] text-white font-semibold text-base',
              'hover:bg-[#1D4ED8]',
              'hover:shadow-[0_0_24px_rgba(37,99,235,0.45)]',
              'transition-all duration-200',
              'w-full sm:w-auto',
            ].join(' ')}
          >
            Browse Campaigns
            <svg width="16" height="16" viewBox="0 0 16 16" fill="none" aria-hidden="true">
              <path d="M3 8h10M9 4l4 4-4 4" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
            </svg>
          </Link>

          <a
            href="#how-it-works"
            id="hero-cta-secondary"
            className={[
              'inline-flex items-center justify-center gap-2',
              'h-12 px-8 rounded-xl',
              'border border-[#1F2937] text-[#9CA3AF] font-semibold text-base',
              'hover:border-[#374151] hover:text-[#F9FAFB] hover:bg-[#111827]',
              'transition-all duration-200',
              'w-full sm:w-auto',
            ].join(' ')}
          >
            How It Works
          </a>
        </div>

        {/* Trust badges */}
        <div className="flex flex-wrap items-center justify-center gap-2 mt-2">
          {[
            { icon: '🔒', text: 'Funds locked until milestones' },
            { icon: '📋', text: 'Audited Smart Contracts' },
            { icon: '⛓️', text: 'Built on Base' },
          ].map(({ icon, text }) => (
            <span
              key={text}
              className={[
                'inline-flex items-center gap-1.5 px-3 py-1.5',
                'rounded-full border border-[#1F2937]',
                'bg-[#111827]/60 text-[#9CA3AF] text-xs font-medium',
                'select-none',
              ].join(' ')}
            >
              <span aria-hidden="true">{icon}</span>
              {text}
            </span>
          ))}
        </div>
      </div>

      {/* Scroll indicator */}
      <div className="absolute bottom-8 left-1/2 -translate-x-1/2 z-10">
        <ScrollIndicator />
      </div>
    </section>
  );
}

// ─── Section: Stats Bar ───────────────────────────────────────────────────────

function StatsSection() {
  return (
    <section
      id="stats"
      aria-label="Platform statistics"
      className="bg-[#111827] border-y border-[#1F2937]"
    >
      <div className="mx-auto max-w-5xl px-6 py-12">
        <dl className="grid grid-cols-2 md:grid-cols-4 gap-8 md:gap-4">
          {STATS.map((stat) => (
            <div
              key={stat.label}
              className="flex flex-col items-center text-center gap-1"
            >
              <dt className="text-sm text-[#9CA3AF] font-medium order-2 mt-1">
                {stat.label}
              </dt>
              <dd className="text-4xl sm:text-5xl font-extrabold text-[#F9FAFB] tabular-nums order-1">
                <AnimatedNumber
                  target={stat.value}
                  prefix={stat.prefix}
                  suffix={stat.suffix}
                />
              </dd>
            </div>
          ))}
        </dl>
      </div>
    </section>
  );
}

// ─── Section: Featured Campaigns ──────────────────────────────────────────────

function CampaignsSection() {
  return (
    <section
      id="featured-campaigns"
      className="py-20 px-4"
      aria-labelledby="campaigns-heading"
    >
      <div className="mx-auto max-w-6xl">
        {/* Header row */}
        <div className="flex items-end justify-between mb-10">
          <div>
            <div className="text-xs font-semibold uppercase tracking-widest text-[#2563EB] mb-2">
              Active Now
            </div>
            <h2
              id="campaigns-heading"
              className="text-3xl sm:text-4xl font-bold text-[#F9FAFB]"
            >
              Live Campaigns
            </h2>
          </div>
          <Link
            href="/campaigns"
            id="campaigns-view-all"
            className="inline-flex items-center gap-1.5 text-sm font-medium text-[#2563EB] hover:text-[#60A5FA] transition-colors"
          >
            View all
            <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
              <path d="M2.5 7h9M8 3.5L11.5 7 8 10.5" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round" />
            </svg>
          </Link>
        </div>

        {/* Cards grid */}
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {FEATURED_CAMPAIGNS.map((c) => (
            <CampaignCard key={c.id} campaign={c} />
          ))}
        </div>
      </div>
    </section>
  );
}

// ─── Section: How It Works ────────────────────────────────────────────────────

function HowItWorksSection() {
  return (
    <section
      id="how-it-works"
      className="py-20 px-4 bg-[#111827] border-y border-[#1F2937]"
      aria-labelledby="how-heading"
    >
      <div className="mx-auto max-w-6xl">
        {/* Header */}
        <div className="text-center mb-14">
          <div className="text-xs font-semibold uppercase tracking-widest text-[#2563EB] mb-2">
            The Process
          </div>
          <h2
            id="how-heading"
            className="text-3xl sm:text-4xl font-bold text-[#F9FAFB]"
          >
            How TrustChain Works
          </h2>
        </div>

        {/* Steps */}
        <div className="relative">
          {/* Connecting dashed line — desktop only */}
          <div
            className="hidden lg:block absolute top-6 left-[calc(10%+24px)] right-[calc(10%+24px)] h-px"
            style={{
              borderTop: '1.5px dashed #2563EB',
              opacity: 0.25,
            }}
            aria-hidden="true"
          />

          <ol className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-5 gap-8 lg:gap-4">
            {HOW_IT_WORKS_STEPS.map((step) => (
              <li
                key={step.n}
                id={`how-step-${step.n}`}
                className="flex flex-col items-center lg:items-center text-left lg:text-center gap-4"
              >
                {/* Numbered circle */}
                <div
                  className="relative z-10 flex-shrink-0 w-12 h-12 rounded-full flex items-center justify-center font-bold text-base text-white"
                  style={{
                    background: 'linear-gradient(135deg, #2563EB, #1D4ED8)',
                    boxShadow: '0 0 20px rgba(37,99,235,0.3)',
                  }}
                  aria-hidden="true"
                >
                  {step.n}
                </div>
                <div>
                  <h3 className="text-[#F9FAFB] font-semibold text-base mb-1.5">
                    {step.title}
                  </h3>
                  <p className="text-sm text-[#9CA3AF] leading-relaxed max-w-[200px] mx-auto lg:mx-auto">
                    {step.desc}
                  </p>
                </div>
              </li>
            ))}
          </ol>
        </div>
      </div>
    </section>
  );
}

// ─── Section: Trust ───────────────────────────────────────────────────────────

function TrustSection() {
  return (
    <section
      id="trust"
      className="py-20 px-4"
      aria-labelledby="trust-heading"
    >
      <div className="mx-auto max-w-6xl grid grid-cols-1 lg:grid-cols-2 gap-16 items-center">

        {/* Left: Feature list */}
        <div>
          <div className="text-xs font-semibold uppercase tracking-widest text-[#2563EB] mb-2">
            Trust & Safety
          </div>
          <h2
            id="trust-heading"
            className="text-3xl sm:text-4xl font-bold text-[#F9FAFB] mb-10 leading-tight"
          >
            Why donors trust TrustChain
          </h2>
          <ul className="flex flex-col gap-8" role="list">
            {TRUST_FEATURES.map(({ icon, title, desc }) => (
              <li key={title} className="flex items-start gap-4" role="listitem">
                <span
                  className="flex-shrink-0 w-10 h-10 rounded-xl flex items-center justify-center text-xl"
                  style={{ background: 'rgba(37,99,235,0.1)', border: '1px solid rgba(37,99,235,0.2)' }}
                  aria-hidden="true"
                >
                  {icon}
                </span>
                <div>
                  <h3 className="text-[#F9FAFB] font-semibold text-base mb-1">
                    {title}
                  </h3>
                  <p className="text-sm text-[#9CA3AF] leading-relaxed">{desc}</p>
                </div>
              </li>
            ))}
          </ul>
        </div>

        {/* Right: Security stats card */}
        <div
          id="security-stats-card"
          className="relative rounded-2xl p-8"
          style={{
            background: 'linear-gradient(145deg, #111827 0%, #0F1629 100%)',
            border: '1px solid rgba(37,99,235,0.3)',
            boxShadow: '0 0 40px rgba(37,99,235,0.12), 0 0 80px rgba(37,99,235,0.05)',
          }}
        >
          {/* Corner glow */}
          <div
            className="absolute top-0 right-0 w-32 h-32 rounded-full pointer-events-none"
            style={{
              background: 'radial-gradient(circle, rgba(37,99,235,0.15) 0%, transparent 70%)',
              transform: 'translate(30%, -30%)',
            }}
            aria-hidden="true"
          />

          <div className="relative z-10 flex flex-col gap-6">

            {/* Hero stat */}
            <div className="text-center pb-6 border-b border-[#1F2937]">
              <div
                className="text-5xl font-extrabold tabular-nums mb-1"
                style={{ color: '#2DD4BF' }}
              >
                $0
              </div>
              <div className="text-sm text-[#9CA3AF]">lost to fraud or misuse</div>
            </div>

            {/* Security specs */}
            <ul className="flex flex-col gap-4" role="list" aria-label="Security guarantees">
              {[
                { icon: '✅', stat: '100%', label: 'milestone-verified fund releases' },
                { icon: '🔐', stat: '2-of-3', label: 'multi-sig required for every approval' },
                { icon: '⏳', stat: '48h', label: 'timelock on every fund release' },
              ].map(({ icon, stat, label }) => (
                <li
                  key={label}
                  role="listitem"
                  className="flex items-center gap-3 py-3 border-b border-[#1F2937] last:border-0"
                >
                  <span className="text-lg flex-shrink-0" aria-hidden="true">{icon}</span>
                  <div className="flex items-baseline gap-2 min-w-0">
                    <span className="font-bold text-[#F9FAFB] whitespace-nowrap">{stat}</span>
                    <span className="text-sm text-[#9CA3AF] truncate">{label}</span>
                  </div>
                </li>
              ))}
            </ul>

            {/* Footer note */}
            <p className="text-xs text-[#4B5563] text-center pt-2">
              Smart contract audited · Open source · MIT licensed
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}

// ─── Section: CTA Banner ──────────────────────────────────────────────────────

function CTABanner() {
  return (
    <section
      id="cta-banner"
      className="py-20 px-4 relative overflow-hidden"
      aria-labelledby="cta-heading"
    >
      {/* Gradient background */}
      <div
        className="absolute inset-0"
        style={{
          background:
            'linear-gradient(135deg, #0D1B3E 0%, #0A0F1E 40%, #060D1F 100%)',
        }}
        aria-hidden="true"
      />
      {/* Blue glow */}
      <div
        className="absolute inset-0 pointer-events-none"
        style={{
          background:
            'radial-gradient(ellipse 60% 60% at 50% 50%, rgba(37,99,235,0.1) 0%, transparent 70%)',
        }}
        aria-hidden="true"
      />
      {/* Top border glow */}
      <div
        className="absolute top-0 left-1/2 -translate-x-1/2 w-1/2 h-px"
        style={{ background: 'linear-gradient(90deg, transparent, #2563EB, transparent)', opacity: 0.5 }}
        aria-hidden="true"
      />

      <div className="relative z-10 mx-auto max-w-3xl text-center flex flex-col items-center gap-6">
        <h2
          id="cta-heading"
          className="text-3xl sm:text-5xl font-extrabold text-[#F9FAFB] leading-tight"
        >
          Ready to donate{' '}
          <span className="text-[#2563EB]">with confidence?</span>
        </h2>
        <p className="text-lg text-[#9CA3AF] max-w-xl">
          Browse our verified campaigns. Every ETH you send is locked until
          milestones are proven — not just promised.
        </p>
        <Link
          href="/campaigns"
          id="cta-browse-campaigns"
          className={[
            'inline-flex items-center gap-2',
            'h-14 px-10 rounded-xl',
            'bg-[#2563EB] text-white font-bold text-lg',
            'hover:bg-[#1D4ED8]',
            'hover:shadow-[0_0_32px_rgba(37,99,235,0.5)]',
            'transition-all duration-200',
          ].join(' ')}
        >
          Browse Campaigns
          <svg width="18" height="18" viewBox="0 0 18 18" fill="none" aria-hidden="true">
            <path d="M3.5 9h11M10 5l4 4-4 4" stroke="currentColor" strokeWidth="1.6" strokeLinecap="round" strokeLinejoin="round" />
          </svg>
        </Link>
      </div>
    </section>
  );
}

// ─── Page ─────────────────────────────────────────────────────────────────────

export default function HomePage() {
  return (
    <>
      <HeroSection />
      <StatsSection />
      <CampaignsSection />
      <HowItWorksSection />
      <TrustSection />
      <CTABanner />
    </>
  );
}
