'use client';

import React from 'react';
import { Card } from '@/components/ui/Card';

// ─── Types ────────────────────────────────────────────────────────────────────

export interface MilestoneValue {
  name: string;
  description: string;
  amountUSD: string;
  deadline: string;
  requiredEvidence: string;
}

export interface MilestoneFormRowProps {
  index: number;
  value: MilestoneValue;
  onChange: (index: number, field: keyof MilestoneValue, value: string) => void;
  onRemove: (index: number) => void;
  isOnly: boolean;
  errors?: Partial<Record<keyof MilestoneValue, string>>;
}

// ─── Shared field styles ──────────────────────────────────────────────────────

const inputClass = [
  'w-full bg-[#111827] border border-[#1F2937] rounded-lg px-4 py-2.5',
  'text-[#F9FAFB] placeholder-[#4B5563]',
  'focus:outline-none focus:border-[#2563EB] focus:ring-1 focus:ring-[#2563EB]',
  'transition-colors duration-200 text-sm',
].join(' ');

const labelClass = 'text-sm font-medium text-[#9CA3AF] mb-1.5 block';
const errorClass = 'text-xs text-[#F87171] mt-1';

// ─── Today's date string (for min= on date input) ─────────────────────────────

function todayString(): string {
  return new Date().toISOString().split('T')[0];
}

// ─── MilestoneFormRow ─────────────────────────────────────────────────────────

export function MilestoneFormRow({
  index,
  value,
  onChange,
  onRemove,
  isOnly,
  errors = {},
}: MilestoneFormRowProps) {
  const n = index + 1; // 1-indexed label

  function handle(field: keyof MilestoneValue) {
    return (
      e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
    ) => onChange(index, field, e.target.value);
  }

  return (
    <Card hover className="p-5" id={`milestone-row-${index}`}>
      {/* ── Header row ─────────────────────────────────────────────── */}
      <div className="flex items-center justify-between mb-5">
        <div className="flex items-center gap-3">
          {/* Blue circle number */}
          <span
            className="w-7 h-7 rounded-full flex items-center justify-center text-sm font-bold text-white flex-shrink-0"
            style={{ backgroundColor: '#2563EB' }}
          >
            {n}
          </span>
          <span className="text-[#F9FAFB] font-semibold text-sm">
            Milestone {n}
          </span>
        </div>

        {/* Remove button — hidden when only milestone */}
        {!isOnly && (
          <button
            type="button"
            onClick={() => onRemove(index)}
            id={`milestone-remove-${index}`}
            aria-label={`Remove milestone ${n}`}
            className="text-xs font-medium text-[#F87171] hover:text-[#FCA5A5] px-3 py-1.5 rounded-lg hover:bg-[rgba(220,38,38,0.08)] border border-transparent hover:border-[rgba(220,38,38,0.2)] transition-all duration-200 cursor-pointer"
          >
            Remove
          </button>
        )}
      </div>

      {/* ── Fields grid ────────────────────────────────────────────── */}
      <div className="space-y-4">
        {/* Row 1: Name + Amount */}
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          {/* Milestone Name */}
          <div>
            <label
              htmlFor={`milestone-${index}-name`}
              className={labelClass}
            >
              Milestone Name <span className="text-[#F87171]">*</span>
            </label>
            <input
              id={`milestone-${index}-name`}
              type="text"
              maxLength={100}
              placeholder="e.g. Groundwater survey complete"
              value={value.name}
              onChange={handle('name')}
              className={inputClass}
            />
            {errors.name && (
              <p className={errorClass}>{errors.name}</p>
            )}
          </div>

          {/* Amount in USD */}
          <div>
            <label
              htmlFor={`milestone-${index}-amount`}
              className={labelClass}
            >
              Amount (USD) <span className="text-[#F87171]">*</span>
            </label>
            <div className="relative">
              <span className="absolute left-3.5 top-1/2 -translate-y-1/2 text-[#4B5563] text-sm font-medium pointer-events-none select-none">
                $
              </span>
              <input
                id={`milestone-${index}-amount`}
                type="number"
                min="1"
                step="any"
                placeholder="0"
                value={value.amountUSD}
                onChange={handle('amountUSD')}
                className={`${inputClass} pl-7`}
              />
            </div>
            {errors.amountUSD && (
              <p className={errorClass}>{errors.amountUSD}</p>
            )}
          </div>
        </div>

        {/* Row 2: Description (full width) */}
        <div>
          <label
            htmlFor={`milestone-${index}-description`}
            className={labelClass}
          >
            What will be achieved <span className="text-[#F87171]">*</span>
          </label>
          <textarea
            id={`milestone-${index}-description`}
            rows={2}
            placeholder="Describe the outcome of this milestone in concrete, verifiable terms."
            value={value.description}
            onChange={handle('description')}
            className={`${inputClass} resize-none`}
          />
          {errors.description && (
            <p className={errorClass}>{errors.description}</p>
          )}
        </div>

        {/* Row 3: Deadline + Evidence */}
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
          {/* Deadline */}
          <div>
            <label
              htmlFor={`milestone-${index}-deadline`}
              className={labelClass}
            >
              Deadline <span className="text-[#F87171]">*</span>
            </label>
            <input
              id={`milestone-${index}-deadline`}
              type="date"
              min={todayString()}
              value={value.deadline}
              onChange={handle('deadline')}
              className={`${inputClass} [color-scheme:dark]`}
            />
            {errors.deadline && (
              <p className={errorClass}>{errors.deadline}</p>
            )}
          </div>

          {/* Required Evidence */}
          <div>
            <label
              htmlFor={`milestone-${index}-evidence`}
              className={labelClass}
            >
              Required Evidence <span className="text-[#F87171]">*</span>
            </label>
            <input
              id={`milestone-${index}-evidence`}
              type="text"
              placeholder="e.g. Invoice + delivery photo, Bank receipt, Signed letter"
              value={value.requiredEvidence}
              onChange={handle('requiredEvidence')}
              className={inputClass}
            />
            {errors.requiredEvidence && (
              <p className={errorClass}>{errors.requiredEvidence}</p>
            )}
          </div>
        </div>
      </div>
    </Card>
  );
}

export default MilestoneFormRow;
