import type { Metadata } from "next";
import { GeistSans } from "geist/font/sans";
import { GeistMono } from "geist/font/mono";
import "./globals.css";

export const metadata: Metadata = {
  title: "pgit - Selective GitHub Downloads",
  description: "A fast, concurrent GitHub repository downloader that fetches only what you need - files, directories, or entire repos without cloning full git history.",
  keywords: [
    "pgit",
    "partial-git",
    "github downloader",
    "selective download",
    "git clone alternative",
    "repository downloader",
    "github files",
    "concurrent downloads",
    "cli tool",
    "developer tools"
  ],
  authors: [{ name: "rushikeshg25" }],
  creator: "pgit",
  publisher: "pgit",
  robots: {
    index: true,
    follow: true,
    googleBot: {
      index: true,
      follow: true,
      "max-video-preview": -1,
      "max-image-preview": "large",
      "max-snippet": -1,
    },
  },
  openGraph: {
    type: "website",
    locale: "en_US",
    url: "https://pgit.dev",
    title: "pgit - Selective GitHub Downloads",
    description: "A fast, concurrent GitHub repository downloader that fetches only what you need - files, directories, or entire repos without cloning full git history.",
    siteName: "pgit",
    images: [
      {
        url: "/og-image.png",
        width: 1200,
        height: 630,
        alt: "pgit - Selective GitHub Downloads",
      },
    ],
  },
  twitter: {
    card: "summary_large_image",
    title: "pgit - Selective GitHub Downloads",
    description: "A fast, concurrent GitHub repository downloader that fetches only what you need - files, directories, or entire repos without cloning full git history.",
    images: ["/og-image.png"],
  },
  metadataBase: new URL("https://pgit.dev"),
  alternates: {
    canonical: "/",
  },
  icons: {
    icon: [
      { url: "/favicon.ico", sizes: "any" },
      { url: "/favicon-16x16.png", sizes: "16x16", type: "image/png" },
      { url: "/favicon-32x32.png", sizes: "32x32", type: "image/png" },
    ],
    apple: [
      { url: "/apple-touch-icon.png", sizes: "180x180", type: "image/png" },
    ],
    other: [
      {
        rel: "android-chrome",
        url: "/android-chrome-192x192.png",
        sizes: "192x192",
        type: "image/png",
      },
      {
        rel: "android-chrome",
        url: "/android-chrome-512x512.png",
        sizes: "512x512",
        type: "image/png",
      },
    ],
  },
  manifest: "/site.webmanifest",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={`${GeistSans.variable} ${GeistMono.variable}`} suppressHydrationWarning>
      <head>
        <link rel="icon" href="/favicon.ico" sizes="any" />
        <link rel="icon" href="/favicon-16x16.png" sizes="16x16" type="image/png" />
        <link rel="icon" href="/favicon-32x32.png" sizes="32x32" type="image/png" />
        <link rel="apple-touch-icon" href="/apple-touch-icon.png" />
        <link rel="manifest" href="/site.webmanifest" />
      </head>
      <body className="font-sans antialiased" suppressHydrationWarning>{children}</body>
    </html>
  );
}
