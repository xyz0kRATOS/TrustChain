'use client';

import React, { useMemo, useState } from 'react';
import { Button } from '@/components/ui/Button';
import { Badge } from '@/components/ui/Badge';
import { formatTimeAgo, shortenAddress } from '@/lib/utils';
import type { CampaignResponse } from '@/types/campaign';

interface ApplicationCardProps {
  campaign: CampaignResponse;
  onApprove: (campaignId: string, notes: string) => Promise<void>;
  onReject: (campaignId: string, reason: string) => Promise<void>;
}

export function ApplicationCard({ campaign, onApprove, onReject }: ApplicationCardProps) {
  const [expandedDesc, setExpandedDesc] = useState(false);
  const [expandedBio, setExpandedBio] = useState(false);
  const [notes, setNotes] = useState('');
  const [rejectReason, setRejectReason] = useState('');
  const [confirmApprove, setConfirmApprove] = useState(false);
  const [confirmReject, setConfirmReject] = useState(false);
  const [isBusy, setIsBusy] = useState(false);

  const submittedAgo = useMemo(() => {
    const epoch = new Date(campaign.createdAt).getTime();
    return formatTimeAgo(epoch);
  }, [campaign.createdAt]);

  async function handleApprove() {
    setIsBusy(true);
    try {
      await onApprove(campaign.id, notes);
    } finally {
      setIsBusy(false);
      setConfirmApprove(false);
    }
  }

  async function handleReject() {
    if (!rejectReason.trim()) return;
    setIsBusy(true);
    try {
      await onReject(campaign.id, rejectReason);
    } finally {
      setIsBusy(false);
      setConfirmReject(false);
    }
  }

  const descriptionNeedsToggle = campaign.description.length > 220;
  const bioNeedsToggle = (campaign.creatorBio || '').length > 180;

  return (
    <article className="bg-[#111827] border border-[#1F2937] rounded-2xl p-5 transition-all duration-300 hover:border-[#374151]">
      <div className="flex items-start justify-between gap-3">
        <div className="flex items-center gap-2">
          <h3 className="text-white font-semibold text-lg">{campaign.name}</h3>
          <Badge variant="pending" text={campaign.category || 'General'} />
        </div>
        <p className="text-sm text-[#4B5563]">Submitted {submittedAgo}</p>
      </div>

      <div className="mt-4 grid grid-cols-1 md:grid-cols-2 gap-3 text-sm">
        <p className="text-[#9CA3AF]">Creator wallet: <span className="text-white font-mono">{shortenAddress(campaign.creatorWallet)}</span></p>
        <p className="text-[#9CA3AF]">Email: <span className="text-white">{campaign.creatorEmail || '-'}</span></p>
        <p className="text-[#9CA3AF]">Country: <span className="text-white">{campaign.creatorCountry || '-'}</span></p>
        <p className="text-[#9CA3AF]">Goal: <span className="text-white">${Number(campaign.goalAmountUsd).toLocaleString()}</span></p>
        <p className="text-[#9CA3AF]">Milestones: <span className="text-white">{campaign.milestones.length}</span></p>
        <p className="text-[#9CA3AF]">Organisation: <span className="text-white">{campaign.creatorOrg || 'Individual'}</span></p>
      </div>

      <div className="mt-4">
        <p className="text-[#9CA3AF] leading-relaxed text-sm">
          {descriptionNeedsToggle && !expandedDesc ? `${campaign.description.slice(0, 220)}...` : campaign.description}
        </p>
        {descriptionNeedsToggle && (
          <button className="text-xs text-[#60A5FA] mt-2" onClick={() => setExpandedDesc((v) => !v)}>
            {expandedDesc ? 'Show less' : 'Show more'}
          </button>
        )}
      </div>

      <div className="mt-4 rounded-xl border border-[#1F2937] bg-[#0D1117] overflow-hidden">
        <table className="w-full text-sm">
          <thead className="text-[#6B7280] border-b border-[#1F2937]">
            <tr>
              <th className="text-left px-3 py-2">Name</th>
              <th className="text-left px-3 py-2">Amount</th>
              <th className="text-left px-3 py-2">Deadline</th>
              <th className="text-left px-3 py-2">Evidence Required</th>
            </tr>
          </thead>
          <tbody>
            {campaign.milestones.map((m) => (
              <tr key={m.id} className="border-b border-[#1F2937] last:border-b-0">
                <td className="px-3 py-2 text-[#E5E7EB]">{m.name}</td>
                <td className="px-3 py-2 text-[#9CA3AF]">${Number(m.amountUsd).toLocaleString()}</td>
                <td className="px-3 py-2 text-[#9CA3AF]">{new Date(m.deadline).toLocaleDateString()}</td>
                <td className="px-3 py-2 text-[#9CA3AF]">{m.requiredEvidence}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      <div className="mt-4">
        <p className="text-xs text-[#4B5563] uppercase tracking-wide">Creator Background</p>
        <p className="text-sm text-[#9CA3AF] mt-1">
          {bioNeedsToggle && !expandedBio
            ? `${(campaign.creatorBio || '').slice(0, 180)}...`
            : campaign.creatorBio || 'No bio provided.'}
        </p>
        {bioNeedsToggle && (
          <button className="text-xs text-[#60A5FA] mt-2" onClick={() => setExpandedBio((v) => !v)}>
            {expandedBio ? 'Show less' : 'Show more'}
          </button>
        )}
      </div>

      <div className="mt-4">
        <p className="text-xs text-[#4B5563] uppercase tracking-wide">Submitted Documents</p>
        <ul className="mt-2 text-sm text-[#9CA3AF] list-disc pl-5">
          {(campaign.documentFileNames || []).length === 0 ? (
            <li>No document names provided</li>
          ) : (
            campaign.documentFileNames?.map((f) => <li key={f}>{f}</li>)
          )}
        </ul>
      </div>

      <div className="mt-6 border-t border-[#1F2937] pt-4 flex flex-col md:flex-row gap-3 md:items-end md:justify-between">
        <textarea
          placeholder="Internal notes (optional - not shown to creator)"
          value={confirmReject ? rejectReason : notes}
          onChange={(e) => (confirmReject ? setRejectReason(e.target.value) : setNotes(e.target.value))}
          className="w-full md:max-w-lg h-24 bg-[#0D1117] border border-[#1F2937] rounded-lg px-3 py-2 text-sm text-[#E5E7EB]"
        />

        <div className="flex items-center gap-2">
          {!confirmApprove && !confirmReject && (
            <>
              <Button variant="danger" size="sm" onClick={() => setConfirmReject(true)}>Reject</Button>
              <Button variant="primary" size="sm" onClick={() => setConfirmApprove(true)}>Approve Campaign</Button>
            </>
          )}

          {confirmApprove && (
            <div className="flex items-center gap-2">
              <Button variant="ghost" size="sm" onClick={() => setConfirmApprove(false)}>Cancel</Button>
              <Button variant="primary" size="sm" loading={isBusy} onClick={handleApprove}>Confirm Approve</Button>
            </div>
          )}

          {confirmReject && (
            <div className="flex items-center gap-2">
              <Button variant="ghost" size="sm" onClick={() => setConfirmReject(false)}>Cancel</Button>
              <Button variant="danger" size="sm" loading={isBusy} onClick={handleReject}>Confirm Reject</Button>
            </div>
          )}
        </div>
      </div>
    </article>
  );
}

export default ApplicationCard;
