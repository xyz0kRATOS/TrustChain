'use client'

import { useEffect, useState } from 'react'
import { api } from '@/lib/api'

type Status = 'checking' | 'connected' | 'unreachable'

/**
 * BackendStatus
 * ─────────────
 * Dev-only floating pill (bottom-left) that pings /api/health every
 * 30 seconds and shows a green/red indicator.
 * Renders nothing in production.
 */
export function BackendStatus() {
  const [status, setStatus] = useState<Status>('checking')

  async function check() {
    const res = await api.health()
    setStatus(res.error ? 'unreachable' : 'connected')
  }

  useEffect(() => {
    check()
    const interval = setInterval(check, 30_000)
    return () => clearInterval(interval)
  }, [])

  // Never render in production
  if (process.env.NODE_ENV !== 'development') return null
  // Hide while first check is in flight
  if (status === 'checking') return null

  const connected = status === 'connected'

  return (
    <div
      role="status"
      aria-label={connected ? 'Backend connected' : 'Backend unreachable'}
      className={[
        'fixed bottom-4 left-4 z-50',
        'flex items-center gap-2',
        'px-3 py-1.5 rounded-full',
        'text-xs font-medium select-none',
        'border backdrop-blur-sm',
        connected
          ? 'bg-[#134E4A]/80 text-[#2DD4BF] border-[#0D9488]/60'
          : 'bg-[#450A0A]/80 text-[#FCA5A5] border-[#DC2626]/60',
      ].join(' ')}
    >
      {/* Pulsing dot */}
      <span
        className={[
          'relative flex h-2 w-2',
        ].join(' ')}
      >
        {connected && (
          <span className="animate-ping absolute inline-flex h-full w-full rounded-full bg-[#0D9488] opacity-60" />
        )}
        <span
          className={[
            'relative inline-flex rounded-full h-2 w-2',
            connected ? 'bg-[#0D9488]' : 'bg-[#DC2626]',
          ].join(' ')}
        />
      </span>

      {connected ? 'Backend connected' : 'Backend unreachable'}
    </div>
  )
}
