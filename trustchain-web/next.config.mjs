/** @type {import('next').NextConfig} */
const nextConfig = {
  /**
   * Proxy /api/* → backend during development.
   * This is a safety net for any direct fetch('/api/...') calls;
   * lib/api.ts uses the full BASE_URL already.
   */
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: `${process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'}/api/:path*`,
      },
    ];
  },

  webpack: (config) => {
    // MetaMask SDK pulls in @react-native-async-storage via wagmi connectors.
    // Provide an empty stub so the browser build doesn't fail.
    config.resolve.alias = {
      ...config.resolve.alias,
      '@react-native-async-storage/async-storage': false,
    };
    return config;
  },
};

export default nextConfig;
