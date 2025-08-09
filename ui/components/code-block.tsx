"use client"

import * as React from "react"
import { Button } from "@/components/ui/button"
import { Check, Copy } from "lucide-react"
import { cn } from "@/lib/utils"

export default function CodeBlock({
  code = "",
  language = "bash",
  className = "",
}: {
  code?: string
  language?: string
  className?: string
}) {
  const [copied, setCopied] = React.useState(false)
  const preRef = React.useRef<HTMLPreElement | null>(null)

  const onCopy = async () => {
    try {
      await navigator.clipboard.writeText(code)
      setCopied(true)
      setTimeout(() => setCopied(false), 1500)
    } catch (e) {
      console.error("Copy failed", e)
      // Could add user notification here in the future
    }
  }

  return (
    <div
      className={cn(
        "relative rounded-md border bg-neutral-50 dark:bg-neutral-900/60 dark:border-neutral-800",
        className,
      )}
      role="region"
      aria-label={`${language} code block`}
    >
      <div className="flex items-center justify-between border-b px-3 py-2 dark:border-neutral-800">
        <span 
          className="text-xs font-medium uppercase tracking-wide text-neutral-500 dark:text-neutral-400"
          aria-label={`Programming language: ${language}`}
        >
          {language}
        </span>
        <Button 
          onClick={onCopy} 
          variant="outline" 
          size="sm" 
          className="h-7 gap-1 bg-transparent focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2"
          aria-label={copied ? "Code copied to clipboard" : "Copy code to clipboard"}
          aria-live="polite"
        >
          {copied ? <Check className="h-3.5 w-3.5" aria-hidden="true" /> : <Copy className="h-3.5 w-3.5" aria-hidden="true" />}
          <span className="text-xs">{copied ? "Copied" : "Copy"}</span>
        </Button>
      </div>
      <pre
        ref={preRef}
        className="max-h-[420px] overflow-auto p-3 text-sm leading-relaxed text-neutral-800 dark:text-neutral-200 font-mono"
        aria-label={`${language} code: ${code}`}
        tabIndex={0}
      >
        <code>{code}</code>
      </pre>
    </div>
  )
}
