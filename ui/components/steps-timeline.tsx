"use client"

import * as React from "react"
import { cn } from "@/lib/utils"

type Step = {
  icon: React.ReactNode
  title: string
  desc: string
}

export default function StepsTimeline({
  steps = [],
  className = "",
}: {
  steps?: Step[]
  className?: string
}) {
  const containerRef = React.useRef<HTMLDivElement | null>(null)
  const [visibleIndexes, setVisibleIndexes] = React.useState<number[]>([])

  React.useEffect(() => {
    if (!containerRef.current) return
    const els = Array.from(containerRef.current.querySelectorAll('[data-step="true"]'))
    const io = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          const idxAttr = entry.target.getAttribute("data-idx")
          const idx = idxAttr ? Number.parseInt(idxAttr, 10) : -1
          if (entry.isIntersecting && idx >= 0) {
            setTimeout(() => {
              setVisibleIndexes((prev) => (prev.includes(idx) ? prev : [...prev, idx]))
            }, idx * 120) // stagger
          }
        })
      },
      { rootMargin: "0px 0px -10% 0px", threshold: 0.15 },
    )
    els.forEach((el) => io.observe(el))
    return () => io.disconnect()
  }, [])

  return (
    <div ref={containerRef} className={cn("relative", className)}>
      <div className="relative mx-auto max-w-3xl">
        {/* Vertical line */}
        <div className="absolute left-4 top-0 bottom-0 hidden w-px bg-gradient-to-b from-emerald-300 via-neutral-200 to-neutral-200 dark:from-emerald-800 dark:via-neutral-800 dark:to-neutral-900 md:block" />
        <ul className="space-y-6">
          {steps.map((step, idx) => {
            const isVisible = visibleIndexes.includes(idx)
            return (
              <li
                key={idx}
                data-step="true"
                data-idx={idx}
                className={cn(
                  "relative grid gap-3 rounded-lg border bg-white p-4 md:grid-cols-[auto_1fr] md:gap-5",
                  "dark:bg-neutral-900 dark:border-neutral-800",
                  "transition-all duration-700 ease-out",
                  isVisible
                    ? "translate-y-0 opacity-100 shadow-[0_12px_40px_-20px_rgba(16,185,129,0.35)]"
                    : "translate-y-6 opacity-0",
                )}
              >
                {/* Bullet */}
                <div className="relative hidden md:block">
                  <div className="absolute -left-[27px] top-2 h-5 w-5 rounded-full border-2 border-emerald-500 bg-white dark:bg-neutral-900 dark:border-emerald-500">
                    <div
                      className={cn(
                        "absolute left-1/2 top-1/2 h-2 w-2 -translate-x-1/2 -translate-y-1/2 rounded-full bg-emerald-500",
                        isVisible ? "animate-ping" : "",
                      )}
                    />
                    <div className="absolute left-1/2 top-1/2 h-2 w-2 -translate-x-1/2 -translate-y-1/2 rounded-full bg-emerald-500" />
                  </div>
                </div>

                {/* Content */}
                <div className="flex items-start gap-3 md:col-span-2 md:ml-0">
                  <div className="mt-0.5 flex h-9 w-9 items-center justify-center rounded-md bg-emerald-600/10 text-emerald-700 dark:bg-emerald-500/10 dark:text-emerald-400">
                    {step.icon}
                  </div>
                  <div>
                    <div className="text-base font-semibold dark:text-neutral-100">{step.title}</div>
                    <p className="mt-1 text-sm text-neutral-600 dark:text-neutral-400">{step.desc}</p>
                  </div>
                </div>
              </li>
            )
          })}
        </ul>
      </div>
    </div>
  )
}
