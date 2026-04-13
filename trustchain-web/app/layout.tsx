import type { Metadata } from 'next';
import localFont from 'next/font/local';
import './globals.css';
import { Providers } from '@/components/layout/Providers';
import { Navbar } from '@/components/layout/Navbar';
import { Footer } from '@/components/layout/Footer';
import { LiveActivityFeed } from '@/components/layout/LiveActivityFeed';

const geistSans = localFont({
  src: './fonts/GeistVF.woff',
  variable: '--font-geist-sans',
  weight: '100 900',
});
const geistMono = localFont({
  src: './fonts/GeistMonoVF.woff',
  variable: '--font-geist-mono',
  weight: '100 900',
});

export const metadata: Metadata = {
  title: 'TrustChain — Transparent Blockchain Donations',
  description:
    'Every donation tracked on-chain. Every milestone verified before funds release. TrustChain makes charity transparent by default.',
  keywords: ['blockchain', 'donations', 'charity', 'Base', 'crypto', 'transparent', 'smart contracts'],
  openGraph: {
    title: 'TrustChain — Transparent Blockchain Donations',
    description: 'Milestone-locked, on-chain charity donations built on Base.',
    type: 'website',
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased bg-[#0A0F1E] text-[#F9FAFB]`}
      >
        <Providers>
          {/* Fixed top navbar */}
          <Navbar />

          {/* Page content + right activity panel */}
          <div className="flex min-h-screen">
            {/* Main scrollable content — padded right on xl to avoid overlap with feed */}
            <main className="flex-1 flex flex-col min-w-0 xl:pr-80">
              {children}
              <Footer />
            </main>

            {/* Fixed right panel — only visible xl+ */}
            <LiveActivityFeed />
          </div>
        </Providers>
      </body>
    </html>
  );
}
