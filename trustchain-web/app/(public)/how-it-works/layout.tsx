import type { Metadata } from 'next';

export const metadata: Metadata = {
  title: 'How It Works — TrustChain',
  description:
    'See exactly how TrustChain ensures every donation reaches its intended milestone. Five transparent steps, enforced by smart contract.',
};

export default function HowItWorksLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <>{children}</>;
}
