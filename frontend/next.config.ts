import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  webpack: (config) => {
    // wagmi bundles all connectors; only injected (MetaMask/Rabby) is used.
    // Stub out unused connector dependencies to avoid "Module not found" errors.
    config.resolve.alias = {
      ...config.resolve.alias,
      // wagmi Tempo dynamic import — must be stubbed
      accounts: false,
      // Unused wagmi connector deps
      "@safe-global/safe-apps-provider": false,
      "@safe-global/safe-apps-sdk": false,
      "@base-org/account": false,
      "@coinbase/wallet-sdk": false,
      "@metamask/connect-evm": false,
      "@walletconnect/ethereum-provider": false,
    };
    return config;
  },
};

export default nextConfig;
