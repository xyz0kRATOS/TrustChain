'use client';

import React, { useCallback, useEffect, useMemo, useState } from 'react';
import Link from 'next/link';
import { useAccount } from 'wagmi';
import { WalletButton } from '@/components/blockchain/WalletButton';
import { Badge } from '@/components/ui/Badge';
import { Button } from '@/components/ui/Button';
import { ApplicationCard } from '@/components/admin/ApplicationCard';
import { shortenAddress } from '@/lib/utils';
import type { ApiEnvelope, CampaignResponse } from '@/types/campaign';

type Tab = 'pending' | 'all' | 'live';

interface AdminStats {
  pendingCount: number;
  liveCount: number;
  completedCount: number;
  frozenCount: number;
  totalDonations: number;
  totalRaisedUsd: string;
}

const emptyStats: AdminStats = {
  pendingCount: 0,
  liveCount: 0,
  completedCount: 0,
  frozenCount: 0,
  totalDonations: 0,
  totalRaisedUsd: '0',
};

export default function AdminPage() {
  const { address, isConnected } = useAccount();
  const adminWallet = process.env.NEXT_PUBLIC_ADMIN_WALLET;
  const connectedWallet = address || '';
  const isAdmin =
    isConnected &&
    !!address &&
    !!adminWallet &&
    address.toLowerCase() === adminWallet.toLowerCase();

  const [tab, setTab] = useState<Tab>('pending');
  const [stats, setStats] = useState<AdminStats>(emptyStats);
  const [pending, setPending] = useState<CampaignResponse[]>([]);
  const [allCampaigns, setAllCampaigns] = useState<CampaignResponse[]>([]);
  const [loading, setLoading] = useState(false);
  const [toast, setToast] = useState('');
  const [fetchError, setFetchError] = useState<string | null>(null);

  const adminHeaders = useMemo(
    () => ({
      'Content-Type': 'application/json',
      'X-Admin-Wallet': connectedWallet,
    }),
    [connectedWallet],
  );

  const adminFetch = useCallback(
    (url: string, options?: RequestInit) => {
      if (!address) {
        throw new Error('No wallet connected')
      }

      return fetch(url, {
        ...options,
        headers: {
          'Content-Type': 'application/json',
          'X-Admin-Wallet': address,
          ...options?.headers,
        },
      })
    },
    [address],
  )

  const loadDashboard = useCallback(async () => {
    if (!isAdmin) return;
    setLoading(true);
    setFetchError(null);
    try {
      const [statsRes, pendingRes, allRes] = await Promise.all([
        adminFetch(`${process.env.NEXT_PUBLIC_API_URL}/api/admin/stats`),
        adminFetch(`${process.env.NEXT_PUBLIC_API_URL}/api/admin/applications`),
        adminFetch(`${process.env.NEXT_PUBLIC_API_URL}/api/admin/campaigns`),
      ]);

      const statsJson = (await statsRes.json()) as ApiEnvelope<AdminStats>;
      const pendingJson = (await pendingRes.json()) as ApiEnvelope<CampaignResponse[]>;
      const allJson = (await allRes.json()) as ApiEnvelope<CampaignResponse[]>;

      if (!statsRes.ok) {
        setFetchError(`API error ${statsRes.status}: ${statsJson.error || 'Unknown error'}`)
        return
      }
      if (!pendingRes.ok) {
        setFetchError(`API error ${pendingRes.status}: ${pendingJson.error || 'Unknown error'}`)
        return
      }
      if (!allRes.ok) {
        setFetchError(`API error ${allRes.status}: ${allJson.error || 'Unknown error'}`)
        return
      }

      setStats(statsJson.data || emptyStats);
      setPending(Array.isArray(pendingJson.data) ? pendingJson.data : []);
      setAllCampaigns(Array.isArray(allJson.data) ? allJson.data : []);
    } catch (err) {
      setFetchError(
        err instanceof Error ? err.message : 'Network error — is the backend running?',
      )
    } finally {
      setLoading(false);
    }
  }, [adminFetch, isAdmin]);

  useEffect(() => {
    loadDashboard();
  }, [loadDashboard]);

  useEffect(() => {
    if (!toast) return;
    const t = setTimeout(() => setToast(''), 3000);
    return () => clearTimeout(t);
  }, [toast]);

  async function approveCampaign(id: string, notes: string) {
    const res = await adminFetch(`${process.env.NEXT_PUBLIC_API_URL}/api/admin/campaigns/${id}/approve`, {
      method: 'POST',
      body: JSON.stringify({ adminWallet: connectedWallet.toLowerCase(), notes: notes || null }),
    });
    const json = await res.json();
    if (!res.ok) throw new Error(json.error || `API error ${res.status}`);
    setToast(`✓ ${json.data.name} approved and now live`);
    await loadDashboard();
  }

  async function rejectCampaign(id: string, reason: string) {
    const res = await adminFetch(`${process.env.NEXT_PUBLIC_API_URL}/api/admin/campaigns/${id}/reject`, {
      method: 'POST',
      body: JSON.stringify({ adminWallet: connectedWallet.toLowerCase(), reason }),
    });
    const json = await res.json();
    if (!res.ok) throw new Error(json.error || `API error ${res.status}`);
    const target = allCampaigns.find((c) => c.id === id);
    setToast(`${target?.name || 'Campaign'} rejected`);
    await loadDashboard();
  }

  if (!isConnected) {
    return (
      <div className="min-h-screen bg-[#0A0F1E] flex items-center justify-center px-6">
        <div className="text-center">
          <svg width="64" height="64" viewBox="0 0 24 24" fill="none" className="mx-auto text-[#374151]">
            <path d="M7 11V8a5 5 0 0 1 10 0v3" stroke="currentColor" strokeWidth="1.6" />
            <rect x="5" y="11" width="14" height="10" rx="2" stroke="currentColor" strokeWidth="1.6" />
          </svg>
          <h1 className="text-2xl text-white mt-5 font-semibold">Admin Access Required</h1>
          <p className="text-[#9CA3AF] mt-2">Connect the TrustChain admin wallet to continue.</p>
          <div className="mt-6 flex justify-center"><WalletButton /></div>
        </div>
      </div>
    );
  }

  if (!isAdmin) {
    return (
      <div className="min-h-screen bg-[#0A0F1E] flex items-center justify-center px-6">
        <div className="max-w-xl rounded-2xl border border-red-900/40 bg-red-950/20 p-6 text-center">
          <h1 className="text-xl text-[#FCA5A5] font-semibold">Access Denied</h1>
          <p className="text-[#9CA3AF] mt-2">This page is restricted to the TrustChain admin wallet.</p>
          <p className="text-[#9CA3AF] mt-2">Connected: {shortenAddress(connectedWallet.toLowerCase())}</p>
          <div className="mt-6 flex justify-center"><WalletButton /></div>
        </div>
      </div>
    );
  }

  const liveCampaigns = allCampaigns.filter((c) => c.status.toLowerCase() === 'live');

  return (
    <div className="min-h-screen bg-[#0A0F1E] px-6 py-24">
      <div className="max-w-7xl mx-auto">
        {toast && (
          <div className="mb-4 rounded-lg border border-teal-800/40 bg-teal-900/20 text-teal-200 px-4 py-3 text-sm">
            {toast}
          </div>
        )}

        <div className="flex items-start justify-between gap-3">
          <div>
            <h1 className="text-3xl text-white font-bold">Admin Dashboard</h1>
            <p className="text-[#9CA3AF] mt-1">TrustChain Developer Console</p>
          </div>
          <div className="px-3 py-1.5 rounded-full border border-teal-900/50 bg-teal-950/30 text-teal-300 text-sm font-mono">
            {shortenAddress(connectedWallet.toLowerCase())}
          </div>
        </div>

        <div className="mt-6 grid grid-cols-2 md:grid-cols-4 gap-4">
          <div className="rounded-xl border border-[#1F2937] bg-[#111827] p-4"><p className="text-[#9CA3AF] text-sm">Pending Review</p><p className="text-white text-2xl font-semibold mt-1">{stats.pendingCount}</p></div>
          <div className="rounded-xl border border-[#1F2937] bg-[#111827] p-4"><p className="text-[#9CA3AF] text-sm">Live</p><p className="text-white text-2xl font-semibold mt-1">{stats.liveCount}</p></div>
          <div className="rounded-xl border border-[#1F2937] bg-[#111827] p-4"><p className="text-[#9CA3AF] text-sm">Completed</p><p className="text-white text-2xl font-semibold mt-1">{stats.completedCount}</p></div>
          <div className="rounded-xl border border-[#1F2937] bg-[#111827] p-4"><p className="text-[#9CA3AF] text-sm">Total Raised</p><p className="text-white text-2xl font-semibold mt-1">${Number(stats.totalRaisedUsd).toLocaleString()}</p></div>
        </div>

        <div className="mt-8 border-b border-[#1F2937] flex items-center gap-6">
          <button className={`pb-3 text-sm ${tab === 'pending' ? 'text-white border-b-2 border-[#2563EB]' : 'text-[#9CA3AF]'}`} onClick={() => setTab('pending')}>Pending Review ({pending.length})</button>
          <button className={`pb-3 text-sm ${tab === 'all' ? 'text-white border-b-2 border-[#2563EB]' : 'text-[#9CA3AF]'}`} onClick={() => setTab('all')}>All Campaigns</button>
          <button className={`pb-3 text-sm ${tab === 'live' ? 'text-white border-b-2 border-[#2563EB]' : 'text-[#9CA3AF]'}`} onClick={() => setTab('live')}>Live Campaigns</button>
        </div>

        {loading ? (
          <div className="mt-6 text-[#9CA3AF]">Loading dashboard...</div>
        ) : null}

        {!loading && tab === 'pending' && (
          <div className="mt-6 space-y-4">
            {fetchError && (
              <div className="rounded-xl border border-[#DC2626] bg-[#450A0A] px-5 py-4 text-sm text-[#FCA5A5]">
                <span className="font-semibold">Error loading applications: </span>
                {fetchError}
              </div>
            )}
            {!fetchError && pending.length === 0 && !loading ? (
              <div className="py-16 text-center text-[#4B5563]">
                <p className="text-lg">No pending applications</p>
                <p className="mt-1 text-sm">Submitted campaigns will appear here</p>
              </div>
            ) : null}
            {pending.map((campaign) => (
              <ApplicationCard key={campaign.id} campaign={campaign} onApprove={approveCampaign} onReject={rejectCampaign} />
            ))}
          </div>
        )}

        {!loading && tab === 'all' && (
          <div className="mt-6 overflow-x-auto rounded-2xl border border-[#1F2937] bg-[#111827]">
            <table className="w-full text-sm">
              <thead className="text-[#9CA3AF] border-b border-[#1F2937]">
                <tr>
                  <th className="text-left px-4 py-3">Name</th>
                  <th className="text-left px-4 py-3">Status</th>
                  <th className="text-left px-4 py-3">Creator</th>
                  <th className="text-left px-4 py-3">Goal</th>
                  <th className="text-left px-4 py-3">Submitted</th>
                  <th className="text-left px-4 py-3">Actions</th>
                </tr>
              </thead>
              <tbody>
                {allCampaigns.map((c) => (
                  <tr key={c.id} className="border-b border-[#1F2937] last:border-b-0">
                    <td className="px-4 py-3 text-white">{c.name}</td>
                    <td className="px-4 py-3"><Badge variant={c.status.toLowerCase() === 'live' ? 'active' : c.status.toLowerCase() === 'completed' ? 'completed' : c.status.toLowerCase() === 'frozen' ? 'frozen' : 'pending'} text={c.status} /></td>
                    <td className="px-4 py-3 text-[#9CA3AF] font-mono">{shortenAddress(c.creatorWallet)}</td>
                    <td className="px-4 py-3 text-[#9CA3AF]">${Number(c.goalAmountUsd).toLocaleString()}</td>
                    <td className="px-4 py-3 text-[#9CA3AF]">{new Date(c.createdAt).toLocaleDateString()}</td>
                    <td className="px-4 py-3 flex items-center gap-2">
                      <Link href={`/campaigns/${c.id}`}><Button variant="ghost" size="sm">View</Button></Link>
                      {c.status.toLowerCase() === 'pending' ? (
                        <Button variant="primary" size="sm" onClick={() => approveCampaign(c.id, '')}>Approve</Button>
                      ) : null}
                      {c.status.toLowerCase() === 'live' ? (
                        <Button variant="danger" size="sm" disabled>Freeze</Button>
                      ) : null}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}

        {!loading && tab === 'live' && (
          <div className="mt-6 space-y-3">
            {liveCampaigns.map((c) => (
              <div key={c.id} className="rounded-xl border border-[#1F2937] bg-[#111827] p-4 flex flex-col md:flex-row md:items-center md:justify-between gap-3">
                <div>
                  <h3 className="text-white font-medium">{c.name}</h3>
                  <p className="text-sm text-[#9CA3AF] mt-1">Raised: 0 ETH · Donors: {c.donorCount} · Milestones: {c.milestones.length}</p>
                </div>
                <Button variant="secondary" size="sm" disabled>Deploy to Blockchain</Button>
              </div>
            ))}
            {liveCampaigns.length === 0 ? <p className="text-[#6B7280]">No live campaigns yet.</p> : null}
          </div>
        )}
      </div>
    </div>
  );
}
