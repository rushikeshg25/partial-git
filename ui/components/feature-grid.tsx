"use client"

import * as React from "react"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Rocket, FolderTree, LockKeyhole, Timer, GitBranch, HardDrive } from "lucide-react"
import { cn } from "@/lib/utils"

const features = [
  {
    icon: Rocket,
    title: "Fast concurrent downloads",
    desc: "Parallelized with Go goroutines for maximum throughput.",
  },
  {
    icon: FolderTree,
    title: "Selective download",
    desc: "Fetch files, directories, or entire repositories.",
  },
  {
    icon: LockKeyhole,
    title: "GitHub token support",
    desc: "Access private repos and lift API limits securely.",
  },
  {
    icon: Timer,
    title: "Smart timeouts",
    desc: "60-second timeout and cancellation to prevent hangs.",
  },
  {
    icon: GitBranch,
    title: "Branch/ref support",
    desc: "Download from specific branches or commits.",
  },
  {
    icon: HardDrive,
    title: "Efficient by design",
    desc: "Only downloads what you need—no git history.",
  },
]

export default function FeatureGrid() {
  const ref = React.useRef<HTMLDivElement | null>(null)
  const [visible, setVisible] = React.useState<boolean[]>(Array(features.length).fill(false))

  React.useEffect(() => {
    if (!ref.current) return
    const cards = Array.from(ref.current.querySelectorAll("[data-card='true']"))
    const io = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          const idxAttr = entry.target.getAttribute("data-idx")
          const idx = idxAttr ? Number.parseInt(idxAttr, 10) : -1
          if (entry.isIntersecting && idx >= 0) {
            setTimeout(() => {
              setVisible((prev) => {
                if (prev[idx]) return prev
                const next = [...prev]
                next[idx] = true
                return next
              })
            }, idx * 90) // stagger
          }
        })
      },
      { threshold: 0.2, rootMargin: "0px 0px -10% 0px" },
    )
    cards.forEach((el) => io.observe(el))
    return () => io.disconnect()
  }, [])

  return (
    <div ref={ref} className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3" role="list" aria-label="pgit features">
      {features.map((f, i) => {
        const Icon = f.icon
        return (
          <Card
            key={f.title}
            data-card="true"
            data-idx={i}
            role="listitem"
            tabIndex={0}
            className={cn(
              "group relative overflow-hidden border-neutral-200 transition dark:border-neutral-800",
              "hover:shadow-[0_10px_30px_-15px_rgba(16,185,129,0.35)]",
              "focus:shadow-[0_10px_30px_-15px_rgba(16,185,129,0.35)] focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2",
              "bg-white dark:bg-neutral-900",
              visible[i] ? "translate-y-0 opacity-100" : "translate-y-6 opacity-0",
            )}
            style={{ transitionDuration: "700ms", transitionProperty: "transform,opacity,box-shadow" }}
            aria-label={`Feature: ${f.title}. ${f.desc}`}
          >
            <div className="pointer-events-none absolute inset-0 bg-[radial-gradient(600px_200px_at_120%_-10%,rgba(16,185,129,0.08),rgba(255,255,255,0))] opacity-0 transition group-hover:opacity-100 group-focus:opacity-100 dark:bg-[radial-gradient(600px_200px_at_120%_-10%,rgba(16,185,129,0.08),rgba(0,0,0,0))]" aria-hidden="true" />
            <CardHeader className="flex flex-row items-center gap-3">
              <div 
                className="flex h-10 w-10 items-center justify-center rounded-md bg-emerald-600/10 text-emerald-400 dark:text-emerald-300"
                aria-hidden="true"
              >
                <Icon className="h-5 w-5" />
              </div>
              <CardTitle className="text-base">{f.title}</CardTitle>
            </CardHeader>
            <CardContent className="pt-0 text-sm text-neutral-700 dark:text-neutral-200">{f.desc}</CardContent>
          </Card>
        )
      })}
    </div>
  )
}
