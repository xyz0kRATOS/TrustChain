'use client';

import React, { useCallback, useRef, useState } from 'react';

// ─── Types ────────────────────────────────────────────────────────────────────

export interface FileDropzoneProps {
  files: File[];
  onChange: (files: File[]) => void;
  maxFiles?: number;
  maxSizeMB?: number;
  accept?: string[];
}

// ─── Helpers ──────────────────────────────────────────────────────────────────

function formatBytes(bytes: number): string {
  if (bytes < 1024) return `${bytes} B`;
  if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(1)} KB`;
  return `${(bytes / (1024 * 1024)).toFixed(1)} MB`;
}

function FileTypeIcon({ name }: { name: string }) {
  const ext = name.split('.').pop()?.toLowerCase() ?? '';
  const isPdf = ext === 'pdf';

  return (
    <span
      className="flex-shrink-0 w-8 h-8 rounded flex items-center justify-center text-[10px] font-bold uppercase"
      style={{
        backgroundColor: isPdf ? 'rgba(220,38,38,0.12)' : 'rgba(37,99,235,0.12)',
        color: isPdf ? '#F87171' : '#60A5FA',
        border: isPdf ? '1px solid rgba(220,38,38,0.25)' : '1px solid rgba(37,99,235,0.25)',
      }}
    >
      {ext || '?'}
    </span>
  );
}

// ─── Dropzone ─────────────────────────────────────────────────────────────────

export function FileDropzone({
  files,
  onChange,
  maxFiles = 5,
  maxSizeMB = 10,
  accept = ['PDF', 'JPG', 'PNG'],
}: FileDropzoneProps) {
  const inputRef = useRef<HTMLInputElement>(null);
  const [dragOver, setDragOver] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const acceptMimeTypes = accept
    .map((t) => {
      switch (t.toUpperCase()) {
        case 'PDF': return 'application/pdf';
        case 'JPG':
        case 'JPEG': return 'image/jpeg';
        case 'PNG': return 'image/png';
        default: return '';
      }
    })
    .filter(Boolean)
    .join(',');

  const addFiles = useCallback(
    (incoming: FileList | null) => {
      if (!incoming) return;
      setError(null);

      const maxBytes = maxSizeMB * 1024 * 1024;
      const toAdd: File[] = [];

      for (const file of Array.from(incoming)) {
        if (files.length + toAdd.length >= maxFiles) {
          setError(`Maximum ${maxFiles} files allowed`);
          break;
        }
        if (file.size > maxBytes) {
          setError(`"${file.name}" exceeds the ${maxSizeMB}MB limit`);
          continue;
        }
        toAdd.push(file);
      }

      if (toAdd.length > 0) {
        onChange([...files, ...toAdd]);
      }
    },
    [files, maxFiles, maxSizeMB, onChange],
  );

  function removeFile(index: number) {
    const next = [...files];
    next.splice(index, 1);
    onChange(next);
    setError(null);
    // Reset input so the same file can be re-added after removal
    if (inputRef.current) inputRef.current.value = '';
  }

  function handleDrop(e: React.DragEvent<HTMLDivElement>) {
    e.preventDefault();
    setDragOver(false);
    addFiles(e.dataTransfer.files);
  }

  function handleDragOver(e: React.DragEvent<HTMLDivElement>) {
    e.preventDefault();
    setDragOver(true);
  }

  function handleDragLeave() {
    setDragOver(false);
  }

  function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
    addFiles(e.target.files);
  }

  const canAddMore = files.length < maxFiles;

  return (
    <div className="space-y-3">
      {/* ── Drop area ───────────────────────────────────────────────── */}
      <div
        role="button"
        tabIndex={0}
        aria-label="File upload area — click or drag files here"
        onClick={() => canAddMore && inputRef.current?.click()}
        onKeyDown={(e) => {
          if ((e.key === 'Enter' || e.key === ' ') && canAddMore) {
            inputRef.current?.click();
          }
        }}
        onDrop={handleDrop}
        onDragOver={handleDragOver}
        onDragLeave={handleDragLeave}
        className="rounded-xl py-12 px-6 text-center transition-all duration-200 select-none"
        style={{
          backgroundColor: dragOver ? 'rgba(37,99,235,0.05)' : '#111827',
          border: `2px dashed ${dragOver ? '#2563EB' : '#374151'}`,
          cursor: canAddMore ? 'pointer' : 'not-allowed',
          opacity: canAddMore ? 1 : 0.6,
        }}
      >
        {/* Hidden native file input */}
        <input
          ref={inputRef}
          type="file"
          multiple
          accept={acceptMimeTypes}
          className="sr-only"
          onChange={handleInputChange}
          tabIndex={-1}
        />

        {/* Upload cloud icon */}
        <div className="flex justify-center">
          <svg
            width="40"
            height="40"
            viewBox="0 0 40 40"
            fill="none"
            aria-hidden="true"
          >
            <path
              d="M13.3 28.3A8.3 8.3 0 1 1 20 13.3a8 8 0 0 1 7.8 6"
              stroke="#4B5563"
              strokeWidth="1.8"
              strokeLinecap="round"
            />
            <path
              d="M28.3 28.3a5 5 0 0 0 0-10H26"
              stroke="#4B5563"
              strokeWidth="1.8"
              strokeLinecap="round"
            />
            <path
              d="M16 25l4-4 4 4"
              stroke="#4B5563"
              strokeWidth="1.8"
              strokeLinecap="round"
              strokeLinejoin="round"
            />
            <path
              d="M20 21v9"
              stroke="#4B5563"
              strokeWidth="1.8"
              strokeLinecap="round"
            />
          </svg>
        </div>

        <p className="text-[#9CA3AF] mt-3 text-sm font-medium">
          {canAddMore
            ? 'Drop files here or click to browse'
            : `Maximum ${maxFiles} files reached`}
        </p>
        <p className="text-[#4B5563] text-xs mt-1">
          {accept.join(', ')} — max {maxSizeMB}MB each
        </p>
      </div>

      {/* ── Inline error ─────────────────────────────────────────────── */}
      {error && (
        <p className="text-xs text-[#F87171] flex items-center gap-1.5" role="alert">
          <span aria-hidden="true">⚠</span>
          {error}
        </p>
      )}

      {/* ── Selected files list ──────────────────────────────────────── */}
      {files.length > 0 && (
        <ul className="space-y-2" aria-label="Selected files">
          {files.map((file, index) => (
            <li
              key={`${file.name}-${index}`}
              className="flex items-center gap-3 px-4 py-2.5 rounded-lg text-sm transition-colors duration-200"
              style={{
                backgroundColor: '#111827',
                border: '1px solid #1F2937',
              }}
              onMouseEnter={(e) => {
                (e.currentTarget as HTMLLIElement).style.borderColor = '#374151';
              }}
              onMouseLeave={(e) => {
                (e.currentTarget as HTMLLIElement).style.borderColor = '#1F2937';
              }}
            >
              <FileTypeIcon name={file.name} />

              <div className="flex-1 min-w-0">
                <p className="text-[#F9FAFB] text-sm truncate font-medium">
                  {file.name}
                </p>
                <p className="text-[#4B5563] text-xs mt-0.5">
                  {formatBytes(file.size)}
                </p>
              </div>

              <button
                type="button"
                onClick={() => removeFile(index)}
                aria-label={`Remove ${file.name}`}
                className="flex-shrink-0 w-6 h-6 rounded flex items-center justify-center text-[#4B5563] hover:text-[#F87171] hover:bg-[rgba(220,38,38,0.08)] transition-all duration-200 cursor-pointer"
              >
                <svg width="12" height="12" viewBox="0 0 12 12" fill="none" aria-hidden="true">
                  <path
                    d="M2 2l8 8M10 2l-8 8"
                    stroke="currentColor"
                    strokeWidth="1.6"
                    strokeLinecap="round"
                  />
                </svg>
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default FileDropzone;
