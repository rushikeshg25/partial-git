import Link from "next/link";
import {
  Github,
  Rocket,
  Terminal,
  FileDown,
  FolderTree,
  Timer,
  GitBranch,
  HardDrive,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Card } from "@/components/ui/card";
import { Separator } from "@/components/ui/separator";

import Hero from "@/components/hero";
import FeatureGrid from "@/components/feature-grid";
import CodeBlock from "@/components/code-block";
import StepsTimeline from "@/components/steps-timeline";
import VideoPlaceholder from "@/components/video-placeholder";

export default function Page() {
  const curlInstall = `curl -fsSL https://raw.githubusercontent.com/rushikeshg25/partial-git/main/scripts/install.sh | bash`;
  const wgetInstall = `wget -qO- https://raw.githubusercontent.com/rushikeshg25/partial-git/main/scripts/install.sh | bash`;

  const usageBasic = `# Download entire repository
pgit https://github.com/user/repo

# Download specific directory
pgit https://github.com/user/repo/tree/main/src

# Download specific file
pgit https://github.com/user/repo/blob/main/README.md

# Download from specific branch
pgit https://github.com/user/repo/tree/develop`;

  const examples = `# Download VS Code's common utilities
pgit https://github.com/microsoft/vscode/tree/main/src/vs/base/common

# Download React's source code
pgit https://github.com/facebook/react/tree/main/packages/react/src

# Download a specific config file
pgit https://github.com/facebook/react/blob/main/package.json`;

  return (
    <main className="relative text-neutral-900 dark:text-neutral-100">
      {/* Skip to main content link for keyboard navigation */}
      <a
        href="#main-content"
        className="sr-only focus:not-sr-only focus:absolute focus:top-4 focus:left-4 focus:z-50 focus:px-4 focus:py-2 focus:bg-emerald-600 focus:text-white focus:rounded-md focus:shadow-lg"
      >
        Skip to main content
      </a>

      {/* Background */}
      <div
        className="pointer-events-none absolute inset-0 -z-10"
        aria-hidden="true"
      >
        <div className="absolute inset-0 bg-gradient-to-b from-white via-white to-neutral-50 dark:from-neutral-950 dark:via-neutral-950 dark:to-black" />
        <div className="absolute left-1/2 top-[-10%] size-[300px] sm:size-[500px] md:size-[600px] lg:size-[700px] -translate-x-1/2 rounded-full bg-[radial-gradient(ellipse_at_center,rgba(16,185,129,0.15),rgba(255,255,255,0))] dark:bg-[radial-gradient(ellipse_at_center,rgba(16,185,129,0.12),rgba(0,0,0,0))] animate-pulse" />
      </div>

      {/* Navbar */}
      <header className="sticky top-0 z-30 border-b bg-white/70 backdrop-blur dark:border-neutral-800 dark:bg-neutral-950/60">
        <nav
          className="mx-auto flex max-w-6xl items-center justify-between px-4 py-3"
          aria-label="Main navigation"
        >
          <Link
            href="#top"
            className="flex items-center gap-2 font-semibold focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 rounded-md p-1 -m-1"
          >
            <div
              className="flex size-7 items-center justify-center rounded-md bg-emerald-600 text-white"
              aria-hidden="true"
            >
              <FileDown className="h-4 w-4" />
            </div>
            <span>pgit</span>
            <Badge variant="secondary" className="ml-1 hidden sm:inline-flex">
              partial-git
            </Badge>
          </Link>
          <div className="flex items-center gap-2">
            <Button asChild variant="outline" size="sm">
              <a
                href="https://github.com/rushikeshg25/partial-git"
                target="_blank"
                rel="noreferrer"
                aria-label="View pgit source code on GitHub (opens in new tab)"
              >
                <Github className="mr-2 h-4 w-4" aria-hidden="true" />
                GitHub
              </a>
            </Button>
            <Button
              asChild
              size="sm"
              className="bg-emerald-600 hover:bg-emerald-700"
            >
              <a href="#install" aria-label="Go to installation section">
                <Rocket className="mr-1.5 h-4 w-4" aria-hidden="true" />
                Install
              </a>
            </Button>
          </div>
        </nav>
      </header>

      {/* Hero */}
      <section
        id="main-content"
        className="relative z-10 mx-auto max-w-6xl px-4"
      >
        <Hero />
      </section>

      {/* Single video placeholder: half visible over the hero, full visible on scroll */}
      <section
        id="video"
        className="relative mx-auto -mt-20 max-w-4xl px-4 md:mt-8 mt-1"
      >
        <div className="relative">
          <VideoPlaceholder
            title="See pgit in action"
            subtitle="Quick overview"
            cta="Watch preview"
            videoSrc="/demo.mp4"
          />
        </div>
      </section>

      {/* Features - "Why pgit?" */}
      <section
        id="features"
        className="mx-auto max-w-6xl px-4 py-12 md:py-16"
        aria-labelledby="features-heading"
      >
        <div className="mx-auto max-w-3xl text-center">
          <h2
            id="features-heading"
            className="text-2xl font-semibold tracking-tight md:text-3xl"
          >
            Why pgit?
          </h2>
          <p className="mt-2 text-neutral-700 dark:text-neutral-200">
            A fast, concurrent GitHub repository downloader that fetches only
            what you need files, directories, or full repos without cloning
            history.
          </p>
        </div>
        <div className="mt-8">
          <FeatureGrid />
        </div>
      </section>

      <Separator className="my-6 dark:bg-neutral-800" />

      {/* Install (Supported Platforms removed) */}
      <section
        id="install"
        className="mx-auto max-w-6xl px-4 py-12 md:py-16"
        aria-labelledby="install-heading"
      >
        <div className="mx-auto max-w-3xl">
          <h2
            id="install-heading"
            className="text-2xl font-semibold tracking-tight md:text-3xl"
          >
            Install pgit
          </h2>
          <p className="mt-2 text-neutral-700 dark:text-neutral-200">
            Quick install is recommended. Copy one of the commands below.
          </p>
          <div className="mt-4 space-y-3">
            <Badge
              variant="outline"
              className="bg-emerald-50 text-emerald-700 dark:bg-transparent dark:text-emerald-300"
            >
              Using curl
            </Badge>
            <CodeBlock language="bash" code={curlInstall} />
            <Badge
              variant="outline"
              className="bg-emerald-50 text-emerald-700 dark:bg-transparent dark:text-emerald-300"
            >
              Using wget
            </Badge>
            <CodeBlock language="bash" code={wgetInstall} />
          </div>
        </div>
      </section>

      {/* Usage (GitHub Token Setup removed) */}
      <section
        id="usage"
        className="mx-auto max-w-6xl px-4 py-12 md:py-16"
        aria-labelledby="usage-heading"
      >
        <div className="mx-auto max-w-3xl">
          <h2
            id="usage-heading"
            className="text-2xl font-semibold tracking-tight md:text-3xl"
          >
            Basic Usage
          </h2>
          <p className="mt-2 text-neutral-700 dark:text-neutral-200">
            Download entire repos, directories, files, or from specific
            branches.
          </p>
          <div className="mt-4">
            <CodeBlock language="bash" code={usageBasic} />
          </div>
        </div>

        <div className="mx-auto mt-10 max-w-3xl">
          <h3 className="text-xl font-semibold">Examples</h3>
          <p className="mt-2 text-neutral-700 dark:text-neutral-200">
            Some quick links to get you started.
          </p>
          <div className="mt-4">
            <CodeBlock language="bash" code={examples} />
          </div>
        </div>
      </section>

      {/* How it works */}
      <section
        id="how-it-works"
        className="mx-auto max-w-6xl px-4 py-12 md:py-16"
        aria-labelledby="how-it-works-heading"
      >
        <div className="mx-auto max-w-3xl text-center">
          <h2
            id="how-it-works-heading"
            className="text-2xl font-semibold tracking-tight md:text-3xl"
          >
            How It Works
          </h2>
          <p className="mt-2 text-neutral-700 dark:text-neutral-200">
            Under the hood, pgit is optimized for speed and reliability.
          </p>
        </div>
        <div className="mt-10">
          <StepsTimeline
            steps={[
              {
                icon: <Terminal className="h-4 w-4" />,
                title: "GitHub API Integration",
                desc: "Uses GitHub Contents API to fetch repository metadata and resolve paths.",
              },
              {
                icon: <Rocket className="h-4 w-4" />,
                title: "Concurrent Downloads",
                desc: "Downloads multiple files simultaneously with Go goroutines for maximum throughput.",
              },
              {
                icon: <FolderTree className="h-4 w-4" />,
                title: "Smart Path Handling",
                desc: "Automatically detects files vs directories and handles nested structures gracefully.",
              },
              {
                icon: <Timer className="h-4 w-4" />,
                title: "Rate Limiting + Timeouts",
                desc: "Respects GitHub rate limits and uses a 60s timeout to avoid hanging downloads.",
              },
              {
                icon: <GitBranch className="h-4 w-4" />,
                title: "Branch/Ref Support",
                desc: "Download from specific branches or commits with precise ref resolution.",
              },
            ]}
          />
        </div>
      </section>

      {/* CTA Footer */}
      <section className="mx-auto max-w-6xl px-4 pb-16">
        <Card className="overflow-hidden border-emerald-200 dark:border-emerald-900/40">
          <div className="relative">
            <div className="absolute inset-0 bg-gradient-to-r from-emerald-50 to-white dark:from-transparent dark:to-transparent" />
            <div className="relative grid gap-6 p-6 md:grid-cols-[1fr_auto] md:items-center">
              <div>
                <h3 className="text-xl font-semibold">
                  Ready to fetch only what you need?
                </h3>
                <p className="mt-1 text-neutral-700 dark:text-neutral-300">
                  Install pgit and start downloading files and folders without
                  cloning full history.
                </p>
              </div>
              <div className="flex gap-2">
                <Button asChild className="bg-emerald-600 hover:bg-emerald-700">
                  <a href="#install">
                    <Rocket className="mr-2 h-4 w-4" />
                    Install Now
                  </a>
                </Button>
                <Button asChild variant="outline">
                  <a
                    href="https://github.com/rushikeshg25/partial-git"
                    target="_blank"
                    rel="noreferrer"
                  >
                    <Github className="mr-2 h-4 w-4" />
                    Star on GitHub
                  </a>
                </Button>
              </div>
            </div>
          </div>
        </Card>
      </section>

      <footer className="border-t" role="contentinfo">
        <div className="mx-auto flex max-w-6xl flex-col items-center gap-4 px-4 py-6 text-sm text-neutral-600 dark:text-neutral-400 sm:flex-row sm:justify-between sm:gap-0">
          <div className="flex items-center gap-2">
            <HardDrive className="h-4 w-4" aria-hidden="true" />
            <span>pgit — partial-git</span>
          </div>
          <nav
            className="flex flex-wrap items-center justify-center gap-3 sm:justify-end"
            aria-label="Footer navigation"
          >
            <a
              className="hover:text-neutral-900 dark:hover:text-white focus:text-neutral-900 dark:focus:text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 rounded px-1 py-0.5"
              href="#features"
            >
              Features
            </a>
            <a
              className="hover:text-neutral-900 dark:hover:text-white focus:text-neutral-900 dark:focus:text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 rounded px-1 py-0.5"
              href="#install"
            >
              Install
            </a>
            <a
              className="hover:text-neutral-900 dark:hover:text-white focus:text-neutral-900 dark:focus:text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 rounded px-1 py-0.5"
              href="#usage"
            >
              Usage
            </a>
            <a
              className="hover:text-neutral-900 dark:hover:text-white focus:text-neutral-900 dark:focus:text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 rounded px-1 py-0.5"
              href="#how-it-works"
            >
              How it works
            </a>
            <a
              className="hover:text-neutral-900 dark:hover:text-white focus:text-neutral-900 dark:focus:text-white focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 rounded px-1 py-0.5"
              href="https://github.com/rushikeshg25/partial-git"
              target="_blank"
              rel="noreferrer"
              aria-label="View pgit source code on GitHub (opens in new tab)"
            >
              <Github className="mr-1 inline h-3.5 w-3.5" aria-hidden="true" />
              GitHub
            </a>
          </nav>
        </div>
      </footer>
    </main>
  );
}
