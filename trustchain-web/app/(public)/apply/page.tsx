'use client';

/**
 * /apply — Campaign Application Page
 *
 * This form collects campaign details off-chain only.
 * On valid submission it would POST to: POST /api/campaigns/apply
 * Body: FormData with all fields + files
 * creator_wallet is sourced from the connected wagmi wallet address.
 *
 * Currently operates in mock/demo mode — no real backend call is made.
 */

import React, { useState } from 'react';
import Link from 'next/link';
import { useAccount } from 'wagmi';
import { Card } from '@/components/ui/Card';
import { Button } from '@/components/ui/Button';
import { WalletButton } from '@/components/blockchain/WalletButton';
import { MilestoneFormRow, MilestoneValue } from '@/components/campaign/MilestoneFormRow';
import { FileDropzone } from '@/components/ui/FileDropzone';
import { shortenAddress } from '@/lib/utils';

// ─── Design tokens (Tailwind classes) ────────────────────────────────────────

const inputClass = [
  'w-full bg-[#111827] border border-[#1F2937] rounded-lg px-4 py-2.5',
  'text-[#F9FAFB] placeholder-[#4B5563]',
  'focus:outline-none focus:border-[#2563EB] focus:ring-1 focus:ring-[#2563EB]',
  'transition-colors duration-200 text-sm',
].join(' ');

const labelClass = 'text-sm font-medium text-[#9CA3AF] mb-1.5 block';
const errorClass = 'text-xs text-[#F87171] mt-1';

// ─── Types ────────────────────────────────────────────────────────────────────

type Category =
  | ''
  | 'Humanitarian Aid'
  | 'Education'
  | 'Environment'
  | 'Medical'
  | 'Infrastructure'
  | 'Other';

interface FormValues {
  // Section 1
  campaignName: string;
  category: Category;
  description: string;
  targetGoalUSD: string;
  imageUrl: string;
  // Section 3
  fullName: string;
  organisation: string;
  email: string;
  country: string;
  bio: string;
}

type FormErrors = Partial<Record<keyof FormValues, string>> & {
  milestones?: string;
};

// ─── Countries list (abbreviated — key subset) ────────────────────────────────

const COUNTRIES = [
  'Afghanistan', 'Albania', 'Algeria', 'Argentina', 'Australia', 'Austria',
  'Bangladesh', 'Belgium', 'Bolivia', 'Brazil', 'Cambodia', 'Cameroon',
  'Canada', 'Chile', 'China', 'Colombia', 'Congo', 'Czech Republic',
  'Denmark', 'Ecuador', 'Egypt', 'Ethiopia', 'Finland', 'France',
  'Germany', 'Ghana', 'Greece', 'Guatemala', 'Haiti', 'Honduras',
  'Hungary', 'India', 'Indonesia', 'Iran', 'Iraq', 'Ireland', 'Israel',
  'Italy', 'Japan', 'Jordan', 'Kenya', 'Lebanon', 'Libya', 'Malaysia',
  'Mexico', 'Morocco', 'Mozambique', 'Myanmar', 'Nepal', 'Netherlands',
  'New Zealand', 'Nigeria', 'Norway', 'Pakistan', 'Palestine', 'Peru',
  'Philippines', 'Poland', 'Portugal', 'Romania', 'Russia', 'Rwanda',
  'Saudi Arabia', 'Senegal', 'Serbia', 'Sierra Leone', 'Somalia',
  'South Africa', 'South Korea', 'South Sudan', 'Spain', 'Sri Lanka',
  'Sudan', 'Sweden', 'Switzerland', 'Syria', 'Tanzania', 'Thailand',
  'Tunisia', 'Turkey', 'Uganda', 'Ukraine', 'United Arab Emirates',
  'United Kingdom', 'United States', 'Venezuela', 'Vietnam', 'Yemen',
  'Zambia', 'Zimbabwe',
];

// ─── Section header component ─────────────────────────────────────────────────

function SectionHeader({
  number,
  title,
  subtitle,
}: {
  number: number;
  title: string;
  subtitle?: string;
}) {
  return (
    <div className="flex items-start gap-4 mb-6">
      <span
        className="w-8 h-8 rounded-full flex items-center justify-center text-sm font-bold text-white flex-shrink-0 mt-0.5"
        style={{ backgroundColor: '#2563EB' }}
      >
        {number}
      </span>
      <div>
        <h2 className="text-lg font-semibold text-[#F9FAFB]">{title}</h2>
        {subtitle && (
          <p className="text-sm text-[#9CA3AF] mt-1 leading-relaxed">{subtitle}</p>
        )}
      </div>
    </div>
  );
}

// ─── Empty milestone ──────────────────────────────────────────────────────────

function emptyMilestone(): MilestoneValue {
  return {
    name: '',
    description: '',
    amountUSD: '',
    deadline: '',
    requiredEvidence: '',
  };
}

// ─── Animated checkmark SVG ───────────────────────────────────────────────────

function AnimatedCheckmark() {
  return (
    <div className="flex justify-center">
      <svg
        width="80"
        height="80"
        viewBox="0 0 80 80"
        fill="none"
        aria-hidden="true"
      >
        <circle
          cx="40"
          cy="40"
          r="36"
          stroke="#0D9488"
          strokeWidth="3"
          fill="rgba(13,148,136,0.08)"
        />
        <path
          d="M24 40l12 12 20-24"
          stroke="#0D9488"
          strokeWidth="3.5"
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeDasharray="60"
          strokeDashoffset="0"
          style={{
            animation: 'drawCheck 0.6s ease-out 0.2s both',
          }}
        />
        <style>{`
          @keyframes drawCheck {
            from { stroke-dashoffset: 60; }
            to   { stroke-dashoffset: 0;  }
          }
        `}</style>
      </svg>
    </div>
  );
}

// ─── Success State ────────────────────────────────────────────────────────────

function SuccessState({ email, appId }: { email: string; appId: string }) {
  return (
    <div className="max-w-3xl mx-auto px-6 pb-24 pt-16 text-center">
      <AnimatedCheckmark />

      <h1 className="text-3xl font-bold text-white mt-8">
        Application Submitted!
      </h1>
      <p className="text-[#9CA3AF] mt-4 text-base leading-relaxed">
        We&apos;ll review your application within 2–5 business days.
        <br />
        You&apos;ll receive an email at{' '}
        <span className="text-[#F9FAFB] font-medium">{email}</span>{' '}
        with our decision.
      </p>
      <p className="text-sm text-[#4B5563] mt-3">
        Application ID: <span className="font-mono">{appId}</span>
      </p>

      <div className="flex flex-col sm:flex-row items-center justify-center gap-4 mt-10">
        <Link href="/campaigns">
          <Button variant="primary" size="lg" id="success-browse-campaigns">
            Browse Campaigns
          </Button>
        </Link>
        <Link href="/">
          <Button variant="ghost" size="lg" id="success-back-home">
            Back to Home
          </Button>
        </Link>
      </div>
    </div>
  );
}

// ─── Page ─────────────────────────────────────────────────────────────────────

export default function ApplyPage() {
  const { address, isConnected, connector } = useAccount();

  // ── Form state ─────────────────────────────────────────────────────────────
  const [form, setForm] = useState<FormValues>({
    campaignName: '',
    category: '',
    description: '',
    targetGoalUSD: '',
    imageUrl: '',
    fullName: '',
    organisation: '',
    email: '',
    country: '',
    bio: '',
  });
  const [errors, setErrors] = useState<FormErrors>({});
  const [milestones, setMilestones] = useState<MilestoneValue[]>([emptyMilestone()]);
  const [docFiles, setDocFiles] = useState<File[]>([]);
  const [submitting, setSubmitting] = useState(false);
  const [submitted, setSubmitted] = useState(false);
  const [appId, setAppId] = useState('');

  // ── Field updater ──────────────────────────────────────────────────────────
  function updateField<K extends keyof FormValues>(key: K, value: FormValues[K]) {
    setForm((prev) => ({ ...prev, [key]: value }));
    if (errors[key]) {
      setErrors((prev) => { const next = { ...prev }; delete next[key]; return next; });
    }
  }

  // ── Milestone handlers ────────────────────────────────────────────────────
  function handleMilestoneChange(
    index: number,
    field: keyof MilestoneValue,
    value: string,
  ) {
    setMilestones((prev) =>
      prev.map((m, i) => (i === index ? { ...m, [field]: value } : m)),
    );
    if (errors.milestones) {
      setErrors((prev) => { const next = { ...prev }; delete next.milestones; return next; });
    }
  }

  function addMilestone() {
    if (milestones.length < 5) {
      setMilestones((prev) => [...prev, emptyMilestone()]);
    }
  }

  function removeMilestone(index: number) {
    setMilestones((prev) => prev.filter((_, i) => i !== index));
  }

  // ── Milestone totals ──────────────────────────────────────────────────────
  const milestoneTotal = milestones.reduce((sum, m) => {
    const n = parseFloat(m.amountUSD);
    return sum + (isNaN(n) ? 0 : n);
  }, 0);
  const goalNumber = parseFloat(form.targetGoalUSD);
  const goalValid = !isNaN(goalNumber) && goalNumber > 0;
  const totalsMatch = goalValid && Math.abs(milestoneTotal - goalNumber) < 0.01;

  // ── Validation ────────────────────────────────────────────────────────────
  function validate(): boolean {
    const newErrors: FormErrors = {};
    const today = new Date().toISOString().split('T')[0];

    if (!form.campaignName.trim()) newErrors.campaignName = 'Campaign name is required';
    if (!form.category) newErrors.category = 'Please select a category';
    if (!form.description.trim()) newErrors.description = 'Description is required';
    if (!form.targetGoalUSD || isNaN(parseFloat(form.targetGoalUSD)) || parseFloat(form.targetGoalUSD) <= 0)
      newErrors.targetGoalUSD = 'Enter a valid target goal';

    // Milestones
    let milestoneError = '';
    milestones.forEach((m, i) => {
      if (!m.name.trim()) milestoneError = `Milestone ${i + 1}: name is required`;
      else if (!m.description.trim()) milestoneError = `Milestone ${i + 1}: description is required`;
      else if (!m.amountUSD || parseFloat(m.amountUSD) <= 0) milestoneError = `Milestone ${i + 1}: enter a valid amount`;
      else if (!m.deadline || m.deadline < today) milestoneError = `Milestone ${i + 1}: deadline must be a future date`;
      else if (!m.requiredEvidence.trim()) milestoneError = `Milestone ${i + 1}: required evidence is required`;
    });
    if (!milestoneError && !totalsMatch && goalValid) {
      milestoneError = `Milestone amounts ($${milestoneTotal.toLocaleString()}) must equal your goal ($${goalNumber.toLocaleString()})`;
    }
    if (milestoneError) newErrors.milestones = milestoneError;

    if (!form.fullName.trim()) newErrors.fullName = 'Full name is required';
    if (!form.email.trim() || !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email))
      newErrors.email = 'Enter a valid email address';
    if (!form.country) newErrors.country = 'Please select your country';
    if (!form.bio.trim()) newErrors.bio = 'Please tell us about your campaign';

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  }

  // ── Submit ────────────────────────────────────────────────────────────────
  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    if (!validate()) {
      // Scroll to first error
      setTimeout(() => {
        const firstError = document.querySelector('[data-error="true"]');
        firstError?.scrollIntoView({ behavior: 'smooth', block: 'center' });
      }, 50);
      return;
    }

    setSubmitting(true);

    /**
     * TODO (real backend): POST /api/campaigns/apply
     * const formData = new FormData();
     * formData.append('campaignName', form.campaignName);
     * formData.append('category', form.category);
     * formData.append('description', form.description);
     * formData.append('targetGoalUSD', form.targetGoalUSD);
     * formData.append('imageUrl', form.imageUrl);
     * formData.append('milestones', JSON.stringify(milestones));
     * formData.append('fullName', form.fullName);
     * formData.append('organisation', form.organisation);
     * formData.append('email', form.email);
     * formData.append('country', form.country);
     * formData.append('bio', form.bio);
     * formData.append('creatorWallet', address ?? '');
     * docFiles.forEach((f) => formData.append('documents', f));
     * await fetch('/api/campaigns/apply', { method: 'POST', body: formData });
     */

    // Mock delay
    await new Promise((r) => setTimeout(r, 1200));

    const id = `TC-${Math.floor(100000 + Math.random() * 900000)}`;
    setAppId(id);
    setSubmitting(false);
    setSubmitted(true);
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }

  // ── Char counters ─────────────────────────────────────────────────────────
  const descLen = form.description.length;
  const bioLen = form.bio.length;

  function charCountColor(len: number, warn: number, max: number): string {
    if (len >= max) return '#F87171';
    if (len >= warn) return '#D97706';
    return '#4B5563';
  }

  // ── Render success ────────────────────────────────────────────────────────
  if (submitted) {
    return (
      <div className="min-h-screen" style={{ backgroundColor: '#0A0F1E' }}>
        <SuccessState email={form.email} appId={appId} />
      </div>
    );
  }

  // ── Render form ───────────────────────────────────────────────────────────
  return (
    <div className="min-h-screen" style={{ backgroundColor: '#0A0F1E' }}>
      {/* ── Page Header ──────────────────────────────────────────────── */}
      <div className="pt-24 pb-10 max-w-3xl mx-auto px-6">
        {/* Breadcrumb */}
        <p className="text-sm text-[#4B5563]">
          <Link href="/" className="hover:text-[#9CA3AF] transition-colors">
            Home
          </Link>
          <span className="mx-1.5 text-[#374151]">/</span>
          <Link href="/campaigns" className="hover:text-[#9CA3AF] transition-colors">
            Campaigns
          </Link>
          <span className="mx-1.5 text-[#374151]">/</span>
          <span className="text-[#6B7280]">Apply for Funding</span>
        </p>

        {/* Title */}
        <h1 className="text-4xl font-bold text-white mt-2 tracking-tight">
          Apply for Funding
        </h1>

        {/* Subtitle */}
        <p className="text-[#9CA3AF] mt-2 text-base leading-relaxed">
          Applications are reviewed by our team before going live.
          All milestones are locked on-chain at approval — plan carefully.
        </p>

        {/* Info banner */}
        <div
          className="mt-6 rounded-xl px-5 py-4 flex items-start gap-3"
          style={{
            backgroundColor: '#1E3A5F',
            border: '1px solid rgba(37,99,235,0.4)',
          }}
        >
          <span className="text-xl flex-shrink-0 mt-0.5" aria-hidden="true">ℹ️</span>
          <p className="text-sm text-[#93C5FD] leading-relaxed">
            Your application goes to our review team first. If approved, your
            campaign is deployed to the Base blockchain. This usually takes{' '}
            <strong className="text-white font-semibold">2–5 business days</strong>.
          </p>
        </div>
      </div>

      {/* ── Form ─────────────────────────────────────────────────────── */}
      <form
        onSubmit={handleSubmit}
        noValidate
        className="max-w-3xl mx-auto px-6 pb-24 space-y-6"
      >
        {/* ════════════════════════════════════════════════════════════
            CARD 1 — Campaign Basics
            ════════════════════════════════════════════════════════════ */}
        <Card className="p-6 sm:p-8">
          <SectionHeader
            number={1}
            title="Campaign Basics"
            subtitle="Tell us what your campaign is about and how much you need."
          />

          <div className="space-y-5">
            {/* Campaign Name */}
            <div>
              <label htmlFor="campaign-name" className={labelClass}>
                Campaign Name <span className="text-[#F87171]">*</span>
              </label>
              <input
                id="campaign-name"
                type="text"
                maxLength={100}
                placeholder="e.g. Clean Water Project — Nairobi"
                value={form.campaignName}
                onChange={(e) => updateField('campaignName', e.target.value)}
                className={inputClass}
                data-error={!!errors.campaignName}
              />
              {errors.campaignName && (
                <p className={errorClass} role="alert">{errors.campaignName}</p>
              )}
            </div>

            {/* Category */}
            <div>
              <label htmlFor="category" className={labelClass}>
                Category <span className="text-[#F87171]">*</span>
              </label>
              <select
                id="category"
                value={form.category}
                onChange={(e) => updateField('category', e.target.value as Category)}
                className={`${inputClass} cursor-pointer`}
                data-error={!!errors.category}
              >
                <option value="" disabled>Select a category…</option>
                <option>Humanitarian Aid</option>
                <option>Education</option>
                <option>Environment</option>
                <option>Medical</option>
                <option>Infrastructure</option>
                <option>Other</option>
              </select>
              {errors.category && (
                <p className={errorClass} role="alert">{errors.category}</p>
              )}
            </div>

            {/* Description */}
            <div>
              <div className="flex items-center justify-between mb-1.5">
                <label htmlFor="description" className="text-sm font-medium text-[#9CA3AF]">
                  Campaign Description <span className="text-[#F87171]">*</span>
                </label>
                <span
                  className="text-xs tabular-nums"
                  style={{ color: charCountColor(descLen, 1800, 2000) }}
                >
                  {descLen} / 2000
                </span>
              </div>
              <textarea
                id="description"
                rows={5}
                maxLength={2000}
                placeholder="Describe your campaign, the problem you're solving, who it helps, and how the funds will be used."
                value={form.description}
                onChange={(e) => updateField('description', e.target.value)}
                className={`${inputClass} resize-none`}
                data-error={!!errors.description}
              />
              {errors.description && (
                <p className={errorClass} role="alert">{errors.description}</p>
              )}
            </div>

            {/* Target Goal */}
            <div>
              <label htmlFor="target-goal" className={labelClass}>
                Target Goal (USD) <span className="text-[#F87171]">*</span>
              </label>
              <div className="relative">
                <span className="absolute left-3.5 top-1/2 -translate-y-1/2 text-[#4B5563] text-sm font-medium pointer-events-none select-none">
                  $
                </span>
                <input
                  id="target-goal"
                  type="number"
                  min="1"
                  step="any"
                  placeholder="10000"
                  value={form.targetGoalUSD}
                  onChange={(e) => updateField('targetGoalUSD', e.target.value)}
                  className={`${inputClass} pl-7`}
                  data-error={!!errors.targetGoalUSD}
                />
              </div>
              <p className="text-xs text-[#4B5563] mt-1.5">
                Final ETH amount is calculated at approval time based on the exchange rate.
              </p>
              {errors.targetGoalUSD && (
                <p className={errorClass} role="alert">{errors.targetGoalUSD}</p>
              )}
            </div>

            {/* Campaign Image URL */}
            <div>
              <label htmlFor="image-url" className={labelClass}>
                Campaign Image URL{' '}
                <span className="text-[#4B5563] font-normal">(optional)</span>
              </label>
              <input
                id="image-url"
                type="url"
                placeholder="https://example.com/campaign-image.jpg"
                value={form.imageUrl}
                onChange={(e) => updateField('imageUrl', e.target.value)}
                className={inputClass}
              />
              <p className="text-xs text-[#4B5563] mt-1.5">
                Optional. A public image URL for your campaign.
              </p>
            </div>
          </div>
        </Card>

        {/* ════════════════════════════════════════════════════════════
            CARD 2 — Milestones
            ════════════════════════════════════════════════════════════ */}
        <Card className="p-6 sm:p-8">
          <SectionHeader
            number={2}
            title="Milestones"
            subtitle="Define 1–5 milestones. Amounts must sum to your total goal. These are permanently locked when your campaign is approved."
          />

          {/* Amber warning callout */}
          <div
            className="rounded-lg px-4 py-3 text-sm mb-6"
            style={{
              backgroundColor: '#451A03',
              border: '1px solid rgba(217,119,6,0.4)',
            }}
          >
            <span className="text-[#FCD34D]">
              ⚠️{' '}
              <strong>Milestones cannot be changed after approval.</strong>{' '}
              Plan each one carefully.
            </span>
          </div>

          {/* milestone error */}
          {errors.milestones && (
            <div
              className="rounded-lg px-4 py-3 text-sm mb-4"
              style={{
                backgroundColor: 'rgba(220,38,38,0.08)',
                border: '1px solid rgba(220,38,38,0.3)',
              }}
              role="alert"
              data-error="true"
            >
              <span className="text-[#F87171]">⚠ {errors.milestones}</span>
            </div>
          )}

          {/* Milestone rows */}
          <div className="space-y-4">
            {milestones.map((m, i) => (
              <MilestoneFormRow
                key={i}
                index={i}
                value={m}
                onChange={handleMilestoneChange}
                onRemove={removeMilestone}
                isOnly={milestones.length === 1}
              />
            ))}
          </div>

          {/* Running total */}
          <div className="mt-5 flex items-center justify-between px-1">
            <span className="text-sm font-medium text-[#9CA3AF]">
              Milestone Total:
              <span className="text-[#F9FAFB] font-semibold ml-1.5">
                ${milestoneTotal.toLocaleString(undefined, { minimumFractionDigits: 0, maximumFractionDigits: 2 })}
              </span>
            </span>
            {goalValid ? (
              totalsMatch ? (
                <span className="text-sm text-[#0D9488] flex items-center gap-1.5">
                  <svg width="14" height="14" viewBox="0 0 14 14" fill="none" aria-hidden="true">
                    <path d="M2 7l3.5 3.5 6.5-7" stroke="#0D9488" strokeWidth="1.6" strokeLinecap="round" strokeLinejoin="round"/>
                  </svg>
                  Amounts match goal
                </span>
              ) : (
                <span className="text-sm text-[#D97706]">
                  ⚠️ Must equal ${goalNumber.toLocaleString()} goal
                </span>
              )
            ) : null}
          </div>

          {/* Add Milestone button */}
          <div className="mt-5">
            {milestones.length < 5 ? (
              <Button
                type="button"
                variant="ghost"
                size="sm"
                onClick={addMilestone}
                id="add-milestone"
              >
                <svg width="14" height="14" viewBox="0 0 14 14" fill="none" className="mr-1" aria-hidden="true">
                  <path d="M7 1v12M1 7h12" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round"/>
                </svg>
                Add Milestone
              </Button>
            ) : (
              <div className="group relative inline-block">
                <Button
                  type="button"
                  variant="ghost"
                  size="sm"
                  disabled
                  id="add-milestone-disabled"
                >
                  <svg width="14" height="14" viewBox="0 0 14 14" fill="none" className="mr-1" aria-hidden="true">
                    <path d="M7 1v12M1 7h12" stroke="currentColor" strokeWidth="1.8" strokeLinecap="round"/>
                  </svg>
                  Add Milestone
                </Button>
                <span
                  className="absolute -top-9 left-1/2 -translate-x-1/2 whitespace-nowrap text-xs px-2.5 py-1 rounded-lg pointer-events-none opacity-0 group-hover:opacity-100 transition-opacity duration-200 z-10"
                  style={{ backgroundColor: '#1F2937', color: '#9CA3AF', border: '1px solid #374151' }}
                  role="tooltip"
                >
                  Maximum 5 milestones
                </span>
              </div>
            )}
          </div>
        </Card>

        {/* ════════════════════════════════════════════════════════════
            CARD 3 — Creator Identity
            ════════════════════════════════════════════════════════════ */}
        <Card className="p-6 sm:p-8">
          <SectionHeader
            number={3}
            title="Creator Identity"
            subtitle="We need to know who is running this campaign."
          />

          <div className="space-y-5">
            {/* Row 1: Full Name + Organisation (2-col on desktop) */}
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-5">
              <div>
                <label htmlFor="full-name" className={labelClass}>
                  Full Name <span className="text-[#F87171]">*</span>
                </label>
                <input
                  id="full-name"
                  type="text"
                  placeholder="Your legal full name"
                  value={form.fullName}
                  onChange={(e) => updateField('fullName', e.target.value)}
                  className={inputClass}
                  data-error={!!errors.fullName}
                />
                {errors.fullName && (
                  <p className={errorClass} role="alert">{errors.fullName}</p>
                )}
              </div>

              <div>
                <label htmlFor="organisation" className={labelClass}>
                  Organisation{' '}
                  <span className="text-[#4B5563] font-normal">(optional)</span>
                </label>
                <input
                  id="organisation"
                  type="text"
                  placeholder="NGO, charity, or company name"
                  value={form.organisation}
                  onChange={(e) => updateField('organisation', e.target.value)}
                  className={inputClass}
                />
              </div>
            </div>

            {/* Row 2: Email + Country (2-col on desktop) */}
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-5">
              <div>
                <label htmlFor="email" className={labelClass}>
                  Email Address <span className="text-[#F87171]">*</span>
                </label>
                <input
                  id="email"
                  type="email"
                  placeholder="you@example.com"
                  value={form.email}
                  onChange={(e) => updateField('email', e.target.value)}
                  className={inputClass}
                  data-error={!!errors.email}
                />
                {errors.email && (
                  <p className={errorClass} role="alert">{errors.email}</p>
                )}
              </div>

              <div>
                <label htmlFor="country" className={labelClass}>
                  Country <span className="text-[#F87171]">*</span>
                </label>
                <select
                  id="country"
                  value={form.country}
                  onChange={(e) => updateField('country', e.target.value)}
                  className={`${inputClass} cursor-pointer`}
                  data-error={!!errors.country}
                >
                  <option value="" disabled>Select your country…</option>
                  {COUNTRIES.map((c) => (
                    <option key={c}>{c}</option>
                  ))}
                </select>
                {errors.country && (
                  <p className={errorClass} role="alert">{errors.country}</p>
                )}
              </div>
            </div>

            {/* Bio (full width) */}
            <div>
              <div className="flex items-center justify-between mb-1.5">
                <label htmlFor="bio" className="text-sm font-medium text-[#9CA3AF]">
                  Bio / Why this campaign <span className="text-[#F87171]">*</span>
                </label>
                <span
                  className="text-xs tabular-nums"
                  style={{ color: charCountColor(bioLen, 450, 500) }}
                >
                  {bioLen} / 500
                </span>
              </div>
              <textarea
                id="bio"
                rows={4}
                maxLength={500}
                placeholder="Tell us a bit about yourself and your connection to this cause."
                value={form.bio}
                onChange={(e) => updateField('bio', e.target.value)}
                className={`${inputClass} resize-none`}
                data-error={!!errors.bio}
              />
              {errors.bio && (
                <p className={errorClass} role="alert">{errors.bio}</p>
              )}
            </div>
          </div>
        </Card>

        {/* ════════════════════════════════════════════════════════════
            CARD 4 — Verification Documents
            ════════════════════════════════════════════════════════════ */}
        <Card className="p-6 sm:p-8">
          <SectionHeader
            number={4}
            title="Verification Documents"
            subtitle="Identity documents, registration certificates, evidence of the cause."
          />

          <FileDropzone
            files={docFiles}
            onChange={setDocFiles}
            maxFiles={5}
            maxSizeMB={10}
            accept={['PDF', 'JPG', 'PNG']}
          />
        </Card>

        {/* ════════════════════════════════════════════════════════════
            SUBMIT SECTION
            ════════════════════════════════════════════════════════════ */}
        <div className="mt-8">
          {isConnected && address ? (
            /* ── Connected wallet ─────────────────────────────────── */
            <div className="space-y-4">
              {/* Connected wallet display */}
              <div
                className="rounded-lg px-4 py-3 flex items-center gap-3"
                style={{
                  backgroundColor: '#111827',
                  border: '1px solid #1F2937',
                }}
              >
                <span className="relative flex h-2 w-2 flex-shrink-0">
                  <span className="absolute inline-flex h-full w-full rounded-full bg-[#0D9488] opacity-60 animate-ping" />
                  <span className="relative inline-flex h-2 w-2 rounded-full bg-[#0D9488]" />
                </span>
                <div>
                  <p className="text-sm text-[#9CA3AF]">
                    Submitting as:{' '}
                    <span className="text-[#F9FAFB] font-mono font-medium">
                      {shortenAddress(address)}
                    </span>
                  </p>
                  {connector?.name && (
                    <p className="text-xs text-[#4B5563] mt-0.5">
                      Connected via {connector.name}
                    </p>
                  )}
                </div>
              </div>

              <Button
                type="submit"
                variant="primary"
                size="lg"
                loading={submitting}
                className="w-full"
                id="submit-application"
              >
                {submitting ? 'Submitting…' : 'Submit Application'}
              </Button>

              <p className="text-xs text-[#4B5563] text-center leading-relaxed">
                By submitting, you agree that all provided information is accurate.
                False applications result in permanent wallet blacklisting.
              </p>
            </div>
          ) : (
            /* ── Wallet not connected ─────────────────────────────── */
            <div
              className="rounded-xl px-6 py-6 text-center"
              style={{
                backgroundColor: 'rgba(69,26,3,0.5)',
                border: '1px solid rgba(217,119,6,0.4)',
              }}
            >
              <p className="text-base font-semibold text-[#FCD34D]">
                Connect your wallet to submit
              </p>
              <p className="text-sm text-[#9CA3AF] mt-2 leading-relaxed">
                Your wallet address identifies you as the campaign creator on the blockchain.
              </p>
              <div className="mt-5 flex justify-center">
                <WalletButton />
              </div>
            </div>
          )}
        </div>
      </form>
    </div>
  );
}
