'use client';

import React, { useState } from 'react';
import Link from 'next/link';

// ─── Types ────────────────────────────────────────────────────────────────────

interface FAQItem {
  q: string;
  a: string;
}

// ─── Data ─────────────────────────────────────────────────────────────────────

const FAQ_ITEMS: FAQItem[] = [
  {
    q: 'What happens if a milestone is not completed?',
    a: 'If a creator misses a deadline or submits insufficient evidence, our admin team can freeze the campaign. When a campaign is frozen, the smart contract automatically calculates each donor\'s proportional share of the remaining balance and sends it back to their wallets. No action is needed from donors.',
  },
  {
    q: 'Can campaign milestones be changed after approval?',
    a: 'No. Milestone amounts, deadlines, and evidence requirements are written into the smart contract at creation and are immutable. We allow one formal amendment request per campaign in exceptional circumstances — but it cannot increase the milestone amount, remove evidence requirements, or extend a deadline by more than 30 days. Any amendment is recorded on-chain permanently.',
  },
  {
    q: "How do I know TrustChain won't just take the money?",
    a: "TrustChain never holds your funds. When you donate, ETH goes directly from your wallet to the campaign's smart contract address. We cannot withdraw it. Only the TimelockController contract can release funds — and only after our multi-sig approval plus a 48-hour wait. Our smart contracts are open source and verified on Basescan. You can read every line.",
  },
];

// ─── FAQ Accordion Item ───────────────────────────────────────────────────────

function FAQAccordionItem({ item, index }: { item: FAQItem; index: number }) {
  const [open, setOpen] = useState(false);

  return (
    <div
      className="rounded-xl border overflow-hidden transition-colors duration-200"
      style={{
        background: '#111827',
        borderColor: open ? '#374151' : '#1F2937',
      }}
    >
      <button
        id={`faq-toggle-${index}`}
        type="button"
        onClick={() => setOpen((v) => !v)}
        aria-expanded={open}
        aria-controls={`faq-body-${index}`}
        className="w-full flex items-center justify-between gap-4 px-6 py-5 text-left"
      >
        <span className="font-semibold text-[#F9FAFB] text-base">{item.q}</span>
        <span
          className="flex-shrink-0 transition-transform duration-300"
          style={{ transform: open ? 'rotate(180deg)' : 'rotate(0deg)' }}
          aria-hidden="true"
        >
          <svg width="18" height="18" viewBox="0 0 18 18" fill="none">
            <path
              d="M4 6.5l5 5 5-5"
              stroke="#9CA3AF"
              strokeWidth="1.8"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
          </svg>
        </span>
      </button>

      <div
        id={`faq-body-${index}`}
        role="region"
        aria-labelledby={`faq-toggle-${index}`}
        className="overflow-hidden transition-all duration-300 ease-in-out"
        style={{ maxHeight: open ? '400px' : '0px', opacity: open ? 1 : 0 }}
      >
        <p className="px-6 pb-6 text-[#9CA3AF] leading-relaxed text-sm">{item.a}</p>
      </div>
    </div>
  );
}

// ─── Chevron Right SVG ────────────────────────────────────────────────────────

function ChevronRight() {
  return (
    <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
      <path
        d="M5 2.5L9.5 7 5 11.5"
        stroke="currentColor"
        strokeWidth="1.6"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  );
}

// ─── Arrow Right SVG ─────────────────────────────────────────────────────────

function ArrowRight() {
  return (
    <svg width="16" height="16" viewBox="0 0 16 16" fill="none" aria-hidden="true">
      <path
        d="M3 8h10M9 4l4 4-4 4"
        stroke="currentColor"
        strokeWidth="1.5"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  );
}

// ─── Section 1: Hero ─────────────────────────────────────────────────────────

function HeroSection() {
  return (
    <section
      id="hiw-hero"
      aria-labelledby="hiw-hero-heading"
      className="pt-32 pb-16 text-center"
    >
      <div className="max-w-4xl mx-auto px-6">
        {/* Breadcrumb */}
        <nav className="flex items-center justify-center gap-2 text-sm text-[#4B5563] mb-4" aria-label="Breadcrumb">
          <Link href="/" className="hover:text-[#9CA3AF] transition-colors">
            Home
          </Link>
          <ChevronRight />
          <span className="text-[#9CA3AF]">How It Works</span>
        </nav>

        {/* Eyebrow */}
        <p className="text-xs font-semibold tracking-widest text-[#2563EB] uppercase mt-4">
          The Process
        </p>

        {/* Title */}
        <h1
          id="hiw-hero-heading"
          className="text-5xl font-bold text-[#F9FAFB] mt-3 leading-tight"
        >
          Accountability built into every step
        </h1>

        {/* Subtitle */}
        <p
          className="max-w-2xl mx-auto mt-6 text-xl text-[#9CA3AF] leading-relaxed"
        >
          TrustChain doesn&apos;t ask you to trust anyone. The smart contract enforces
          the rules. Here is exactly how it works.
        </p>

        {/* Stat pills */}
        <div className="mt-8 flex gap-4 justify-center flex-wrap">
          <span
            className="inline-flex items-center gap-2 rounded-full px-5 py-2 text-sm text-[#93C5FD] border"
            style={{
              background: 'rgba(30,58,95,0.6)',
              borderColor: 'rgba(37,99,235,0.4)',
            }}
          >
            🔒 $0 lost to fraud or misuse
          </span>
          <span
            className="inline-flex items-center gap-2 rounded-full px-5 py-2 text-sm text-[#2DD4BF] border"
            style={{
              background: 'rgba(19,78,74,0.6)',
              borderColor: 'rgba(13,148,136,0.4)',
            }}
          >
            ✅ 100% milestone-verified releases
          </span>
        </div>
      </div>
    </section>
  );
}

// ─── Section 2: The 5 Steps ───────────────────────────────────────────────────

function CheckRow({ text }: { text: string }) {
  return (
    <div className="flex items-start gap-3">
      <span className="flex-shrink-0 mt-0.5 text-[#2DD4BF] text-sm">✓</span>
      <span className="text-sm text-[#9CA3AF]">{text}</span>
    </div>
  );
}

function StepsSection() {
  const steps = [
    {
      n: 1,
      label: 'STEP 1',
      title: 'Every campaign is manually verified before going live',
      desc: 'TrustChain is a curated platform. We review every application — identity documents, organisational credentials, and the plausibility of each milestone. Nothing is automated.',
      detail: (
        <div className="flex flex-col gap-2.5">
          <p className="text-xs text-[#4B5563] font-semibold uppercase tracking-wider mb-1">
            What we check:
          </p>
          <CheckRow text="Identity documents and proof of registration" />
          <CheckRow text="Milestone clarity — is completion objectively verifiable?" />
          <CheckRow text="Amount plausibility — does the budget make sense?" />
          <CheckRow text="Evidence requirements — what proof must be submitted per milestone?" />
        </div>
      ),
    },
    {
      n: 2,
      label: 'STEP 2',
      title: 'Approved campaigns are deployed as smart contracts',
      desc: 'When we approve an application, a new smart contract is deployed to the Base blockchain. Milestones are permanently written into the contract at this point and cannot ever be changed.',
      detail: (
        <div className="flex flex-col gap-2">
          <p className="text-xs text-[#4B5563] font-semibold uppercase tracking-wider mb-1">
            What gets locked forever:
          </p>
          {[
            '• Milestone names, descriptions, and required evidence',
            '• ETH amount to release per milestone',
            '• Milestone deadlines (Unix timestamps)',
            '• Creator wallet address — the only address that receives funds',
            '• SHA-256 hash of all verified documents',
          ].map((line) => (
            <p key={line} className="text-sm text-[#9CA3AF]">
              {line}
            </p>
          ))}
        </div>
      ),
    },
    {
      n: 3,
      label: 'STEP 3',
      title: 'ETH goes directly to the smart contract — not to us',
      desc: "When you donate, your ETH goes from your wallet directly to the campaign's smart contract address. TrustChain never holds your funds. The contract holds them until milestones are verified.",
      detail: (
        <div className="flex flex-col gap-2.5">
          <CheckRow text="Every donation mints a soulbound NFT receipt to your wallet" />
          <CheckRow text="Your donation is recorded permanently on the Base blockchain" />
          <CheckRow text="You can verify the contract balance yourself on Basescan at any time" />
        </div>
      ),
    },
    {
      n: 4,
      label: 'STEP 4',
      title: 'Creators submit evidence. Our team verifies it.',
      desc: 'When a milestone deadline arrives, the campaign creator submits completion evidence — photos, invoices, signed letters, receipts. Our team reviews against the evidence requirements locked at campaign creation.',
      detail: (
        <div>
          <p className="text-sm text-[#9CA3AF] mb-4">
            The approval requires 2-of-3 admin signatures (Safe multi-sig).
          </p>
          {/* Mini timeline */}
          <div className="flex items-start gap-0 overflow-x-auto pb-1" role="list" aria-label="Approval timeline">
            {[
              { label: 'Creator submits evidence', color: '#D97706' },
              { label: 'Admin team reviews (2–5 days)', color: '#D97706' },
              { label: '2-of-3 signatures approve', color: '#2563EB' },
              { label: '48-hour safety window starts', color: '#2563EB' },
            ].map((item, i, arr) => (
              <div key={item.label} className="flex items-start" role="listitem">
                <div className="flex flex-col items-center">
                  <div
                    className="w-3 h-3 rounded-full flex-shrink-0 mt-1"
                    style={{ background: item.color }}
                    aria-hidden="true"
                  />
                  <p className="text-xs text-[#9CA3AF] text-center mt-2 max-w-[80px] leading-tight">
                    {item.label}
                  </p>
                </div>
                {i < arr.length - 1 && (
                  <div
                    className="h-px mt-[6px] flex-shrink-0"
                    style={{ width: '32px', background: '#374151' }}
                    aria-hidden="true"
                  />
                )}
              </div>
            ))}
          </div>
        </div>
      ),
    },
    {
      n: 5,
      label: 'STEP 5',
      title: 'After 48 hours, funds release — automatically and permanently',
      desc: "Once our multi-sig approves a milestone, a 48-hour countdown begins. This window exists so donors can see the pending release and raise an alarm if anything looks wrong. After 48 hours, the funds transfer automatically. No one can stop it — not us, not anyone.",
      detail: (
        <div className="flex flex-col gap-3">
          <p
            className="text-sm font-medium"
            style={{ color: '#2DD4BF' }}
          >
            ✓ During 48 hours: admin can cancel if fraud is detected
          </p>
          <p
            className="text-sm font-medium"
            style={{ color: '#2DD4BF' }}
          >
            ✓ After 48 hours: transfer is automatic and irreversible
          </p>
          <p className="text-sm text-[#4B5563] mt-1">
            If a campaign is frozen at any point, all remaining funds are automatically
            and proportionally refunded to every donor.
          </p>
        </div>
      ),
    },
  ];

  return (
    <section
      id="hiw-steps"
      aria-labelledby="hiw-steps-heading"
      className="py-20"
    >
      <div className="max-w-4xl mx-auto px-6">
        <h2
          id="hiw-steps-heading"
          className="text-3xl font-bold text-[#F9FAFB] mb-14 text-center"
        >
          The 5-step trust model
        </h2>

        {/* Vertical timeline */}
        <ol className="relative flex flex-col gap-12">
          {/* The vertical connecting line */}
          <div
            className="absolute left-5 top-6 bottom-6 w-px"
            style={{ background: '#1F2937' }}
            aria-hidden="true"
          />

          {steps.map((step) => {
            const isLast = step.n === 5;
            return (
              <li
                key={step.n}
                id={`step-${step.n}`}
                className="relative flex gap-8 items-start"
              >
                {/* Number circle */}
                <div
                  className="relative z-10 flex-shrink-0 w-12 h-12 rounded-full flex items-center justify-center font-bold text-lg border-2"
                  style={
                    isLast
                      ? {
                          borderColor: '#0D9488',
                          background: '#134E4A',
                          color: '#2DD4BF',
                        }
                      : {
                          borderColor: '#2563EB',
                          background: '#1E3A5F',
                          color: '#60A5FA',
                        }
                  }
                  aria-hidden="true"
                >
                  {step.n}
                </div>

                {/* Content card */}
                <div
                  className="flex-1 rounded-2xl p-6 border"
                  style={{
                    background: '#111827',
                    borderColor: '#1F2937',
                  }}
                >
                  <p className="text-xs font-semibold tracking-widest text-[#4B5563] uppercase">
                    {step.label}
                  </p>
                  <h3 className="text-xl font-bold text-[#F9FAFB] mt-1">
                    {step.title}
                  </h3>
                  <p className="text-[#9CA3AF] mt-3 leading-relaxed text-sm">
                    {step.desc}
                  </p>

                  {/* Detail box */}
                  <div
                    className="mt-4 rounded-xl p-4 border"
                    style={{
                      background: '#0D1117',
                      borderColor: '#1F2937',
                    }}
                  >
                    {step.detail}
                  </div>
                </div>
              </li>
            );
          })}
        </ol>
      </div>
    </section>
  );
}

// ─── Section 3: For Donors / For Creators ────────────────────────────────────

function AudienceSection() {
  const donorFeatures = [
    'Know exactly where every milestone leads',
    'Get an NFT receipt permanently in your wallet',
    '48-hour window to raise concerns before any release',
    'Automatic refund if campaign is frozen',
    'Public, permanent record on the blockchain',
  ];

  const creatorFeatures = [
    'Credibility from the first campaign',
    'Donors can verify your entire history on-chain',
    'Reputation score follows your wallet forever',
    'Clear milestone structure keeps you accountable',
    'Our team supports you through the evidence process',
  ];

  return (
    <section
      id="hiw-audience"
      aria-labelledby="hiw-audience-heading"
      className="py-20"
      style={{ background: '#0D1117' }}
    >
      <div className="max-w-5xl mx-auto px-6">
        <h2
          id="hiw-audience-heading"
          className="text-3xl font-bold text-[#F9FAFB] text-center"
        >
          Whether you are donating or fundraising
        </h2>

        <div className="mt-12 grid grid-cols-1 md:grid-cols-2 gap-8">
          {/* Donors card */}
          <div
            id="hiw-donors-card"
            className="rounded-2xl p-8 flex flex-col border-t-4"
            style={{
              background: '#111827',
              borderTopColor: '#2563EB',
              border: '1px solid #1F2937',
              borderTop: '4px solid #2563EB',
            }}
          >
            <span className="text-4xl" aria-hidden="true">💙</span>
            <h3 className="text-2xl font-bold text-[#F9FAFB] mt-4">
              For Donors
            </h3>
            <p className="text-[#9CA3AF] mt-1 text-sm">
              Donate with verifiable confidence
            </p>

            <ul className="mt-6 flex flex-col gap-4" role="list">
              {donorFeatures.map((feat) => (
                <li key={feat} className="flex items-start gap-3" role="listitem">
                  <span className="flex-shrink-0 text-[#2563EB] font-bold mt-0.5">✓</span>
                  <span className="text-[#9CA3AF] text-sm">{feat}</span>
                </li>
              ))}
            </ul>

            <div className="mt-8">
              <Link
                href="/campaigns"
                id="hiw-browse-campaigns"
                className="inline-flex items-center gap-2 h-11 px-6 rounded-xl bg-[#2563EB] text-white font-semibold text-sm hover:bg-[#1D4ED8] hover:shadow-[0_0_20px_rgba(37,99,235,0.4)] transition-all duration-200"
              >
                Browse Campaigns
                <ArrowRight />
              </Link>
            </div>
          </div>

          {/* Creators card */}
          <div
            id="hiw-creators-card"
            className="rounded-2xl p-8 flex flex-col"
            style={{
              background: '#111827',
              border: '1px solid #1F2937',
              borderTop: '4px solid #0D9488',
            }}
          >
            <span className="text-4xl" aria-hidden="true">📋</span>
            <h3 className="text-2xl font-bold text-[#F9FAFB] mt-4">
              For Campaign Creators
            </h3>
            <p className="text-[#9CA3AF] mt-1 text-sm">
              Build a verifiable track record
            </p>

            <ul className="mt-6 flex flex-col gap-4" role="list">
              {creatorFeatures.map((feat) => (
                <li key={feat} className="flex items-start gap-3" role="listitem">
                  <span className="flex-shrink-0 text-[#0D9488] font-bold mt-0.5">✓</span>
                  <span className="text-[#9CA3AF] text-sm">{feat}</span>
                </li>
              ))}
            </ul>

            <div className="mt-8">
              <Link
                href="/apply"
                id="hiw-apply-funding"
                className="inline-flex items-center gap-2 h-11 px-6 rounded-xl border border-[#2563EB] text-[#2563EB] font-semibold text-sm hover:bg-[rgba(37,99,235,0.08)] transition-all duration-200"
              >
                Apply for Funding
                <ArrowRight />
              </Link>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

// ─── Section 4: FAQ ───────────────────────────────────────────────────────────

function FAQSection() {
  return (
    <section
      id="hiw-faq"
      aria-labelledby="hiw-faq-heading"
      className="py-20"
    >
      <div className="max-w-3xl mx-auto px-6">
        <h2
          id="hiw-faq-heading"
          className="text-3xl font-bold text-[#F9FAFB] text-center mb-10"
        >
          Common questions
        </h2>

        <div className="flex flex-col gap-3">
          {FAQ_ITEMS.map((item, i) => (
            <FAQAccordionItem key={i} item={item} index={i} />
          ))}
        </div>
      </div>
    </section>
  );
}

// ─── Section 5: Final CTA ─────────────────────────────────────────────────────

function FinalCTA() {
  return (
    <section
      id="hiw-cta"
      aria-labelledby="hiw-cta-heading"
      className="py-20 text-center relative overflow-hidden"
    >
      {/* Background glow */}
      <div
        className="absolute inset-0 pointer-events-none"
        style={{
          background:
            'radial-gradient(ellipse 60% 60% at 50% 50%, rgba(37,99,235,0.08) 0%, transparent 70%)',
        }}
        aria-hidden="true"
      />
      {/* Top border glow */}
      <div
        className="absolute top-0 left-1/2 -translate-x-1/2 w-1/2 h-px"
        style={{
          background:
            'linear-gradient(90deg, transparent, #2563EB, transparent)',
          opacity: 0.4,
        }}
        aria-hidden="true"
      />

      <div className="relative z-10 max-w-2xl mx-auto px-6">
        <h2
          id="hiw-cta-heading"
          className="text-4xl font-bold text-[#F9FAFB]"
        >
          Ready to see it in action?
        </h2>
        <p className="text-[#9CA3AF] mt-4 text-lg">
          Browse live campaigns or apply to run one.
        </p>

        <div className="mt-8 flex gap-4 justify-center flex-wrap">
          <Link
            href="/campaigns"
            id="hiw-cta-browse"
            className="inline-flex items-center gap-2 h-13 px-8 rounded-xl bg-[#2563EB] text-white font-bold text-base hover:bg-[#1D4ED8] hover:shadow-[0_0_28px_rgba(37,99,235,0.45)] transition-all duration-200"
            style={{ height: '52px' }}
          >
            Browse Campaigns
            <ArrowRight />
          </Link>
          <Link
            href="/apply"
            id="hiw-cta-apply"
            className="inline-flex items-center gap-2 h-13 px-8 rounded-xl border border-[#1F2937] text-[#9CA3AF] font-bold text-base hover:border-[#374151] hover:text-[#F9FAFB] hover:bg-[#111827] transition-all duration-200"
            style={{ height: '52px' }}
          >
            Apply for Funding
            <ArrowRight />
          </Link>
        </div>
      </div>
    </section>
  );
}

// ─── Page ─────────────────────────────────────────────────────────────────────

export default function HowItWorksPage() {
  return (
    <>
      <HeroSection />
      <StepsSection />
      <AudienceSection />
      <FAQSection />
      <FinalCTA />
    </>
  );
}
