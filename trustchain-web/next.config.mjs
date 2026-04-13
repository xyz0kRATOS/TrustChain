/** @type {import('next').NextConfig} */
const nextConfig = {
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
