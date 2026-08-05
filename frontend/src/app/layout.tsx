import type { Metadata, Viewport } from "next";
import "./globals.css";
import { Providers } from "./providers";
import { Header } from "@/components/layout/Header";
import { Footer } from "@/components/layout/Footer";
import { COPY } from "@/lib/copy";

export const metadata: Metadata = {
  title: {
    default: `${COPY.appName} · 有担保的 Swap`,
    template: `%s · ${COPY.appName}`,
  },
  description: COPY.tagline,
};

export const viewport: Viewport = {
  themeColor: "#0b0612",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="zh-CN" className="h-full antialiased">
      <body className="bg-aurora bg-grid min-h-full">
        <Providers>
          <div className="flex min-h-screen flex-col">
            <Header />
            <main className="mx-auto w-full max-w-3xl flex-1 px-4 py-8">
              {children}
            </main>
            <Footer />
          </div>
        </Providers>
      </body>
    </html>
  );
}