/**
 * lib/api.ts
 * ──────────
 * Single source of truth for all backend API calls.
 * Base URL is driven by NEXT_PUBLIC_API_URL so it works in every env
 * without code changes.
 */

const BASE_URL =
  process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

// ── Types ─────────────────────────────────────────────────────────────────────

export interface ApiResponse<T> {
  data: T | null
  error: string | null
}

export type CampaignStatus =
  | 'PENDING'
  | 'ACTIVE'
  | 'COMPLETED'
  | 'FROZEN'
  | 'REJECTED'

export type MilestoneStatus =
  | 'PENDING'
  | 'SUBMITTED'
  | 'UNDER_REVIEW'
  | 'APPROVED'
  | 'RELEASING'
  | 'COMPLETED'
  | 'OVERDUE'
  | 'DISPUTED'

export interface Campaign {
  id: string
  contractAddress?: string
  creatorWallet: string
  name: string
  description: string
  goalAmountWei: string
  status: CampaignStatus
  documentHash?: string
  imageUrl?: string
  ipfsEvidenceHash?: string
  createdAt: string
  updatedAt: string
  milestones?: Milestone[]
}

export interface Milestone {
  id: string
  campaignId: string
  sequenceIndex: number
  description: string
  amountWei: string
  deadline: string
  requiredEvidence: string
  status: MilestoneStatus
  submittedAt?: string
  approvedAt?: string
  releasedAt?: string
}

export interface HealthResponse {
  status: string
  service?: string
  database?: string
  timestamp?: string
}

export interface CampaignListParams {
  status?: CampaignStatus
  limit?: number
  offset?: number
}

// ── Core Fetch ────────────────────────────────────────────────────────────────

/**
 * Wraps every fetch call with:
 *  - correct Content-Type header (skipped for FormData)
 *  - unified error envelope matching BackendResponse[T]
 *  - network error catch so callers never see an unhandled rejection
 */
async function apiFetch<T>(
  path: string,
  options?: RequestInit,
): Promise<ApiResponse<T>> {
  try {
    const isFormData = options?.body instanceof FormData

    const res = await fetch(`${BASE_URL}${path}`, {
      ...options,
      headers: {
        // Let the browser set Content-Type for FormData (includes boundary)
        ...(isFormData ? {} : { 'Content-Type': 'application/json' }),
        ...options?.headers,
      },
    })

    const json = await res.json()

    if (!res.ok) {
      // Backend wraps errors as { error: "..." }
      return {
        data: null,
        error: json?.error ?? `Request failed with status ${res.status}`,
      }
    }

    // Backend wraps success as { data: T } or { data: T, meta: ... }
    return { data: json?.data ?? json, error: null }
  } catch (err) {
    // Network failure, CORS block, backend not running, etc.
    const message =
      err instanceof Error ? err.message : 'Network error — backend unreachable'
    console.error(`[api] ${path}:`, message)
    return { data: null, error: message }
  }
}

// ── Exports ───────────────────────────────────────────────────────────────────

export const api = {
  /** GET /api/health — fast liveness check, no DB involved */
  health: () => apiFetch<HealthResponse>('/api/health'),

  /** GET /api/health/db — verifies DB connectivity */
  healthDb: () => apiFetch<HealthResponse>('/api/health/db'),

  campaigns: {
    /**
     * GET /api/campaigns
     * Optional params: status, limit, offset
     */
    list: (params?: CampaignListParams) => {
      const query = params
        ? new URLSearchParams(
            Object.fromEntries(
              Object.entries(params)
                .filter(([, v]) => v !== undefined)
                .map(([k, v]) => [k, String(v)]),
            ),
          ).toString()
        : ''
      return apiFetch<Campaign[]>(
        `/api/campaigns${query ? '?' + query : ''}`,
      )
    },

    /** GET /api/campaigns/:id */
    get: (id: string) => apiFetch<Campaign>(`/api/campaigns/${id}`),

    /**
     * POST /api/campaigns/apply
     * Sends multipart/form-data — Content-Type header is omitted so the
     * browser sets it with the correct boundary automatically.
     */
    apply: (formData: FormData) =>
      apiFetch<{ applicationId: string }>('/api/campaigns/apply', {
        method: 'POST',
        body: formData,
      }),
  },
}
