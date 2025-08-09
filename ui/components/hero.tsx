import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { Github, Rocket, Download } from "lucide-react"
import CodeBlock from "./code-block"

export default function Hero({
  subtitle = "A fast, concurrent GitHub repository downloader for files, directories, or entire repos without cloning full git history.",
}: {
  subtitle?: string
} = {}) {
  const quick = `curl -fsSL https://raw.githubusercontent.com/rushikeshg25/partial-git/main/scripts/install.sh | bash`
  return (
    <div className="relative isolate py-12 md:py-20">
      <div className="mx-auto max-w-3xl text-center">
        <div className="inline-flex items-center gap-2 rounded-full border bg-white px-3 py-1 text-xs text-neutral-700 shadow-sm dark:bg-neutral-900 dark:border-neutral-800 dark:text-neutral-300">
          <Badge
            variant="secondary"
            className="bg-emerald-100 text-emerald-800 dark:bg-emerald-900 dark:text-emerald-300"
          >
            New
          </Badge>
          <span>{"Selective File/Folder Download"}</span>
        </div>
        <h1 className="mt-4 text-4xl font-bold tracking-tight md:text-5xl">
          <span className="bg-gradient-to-r from-neutral-900 to-emerald-600 bg-clip-text text-transparent dark:from-white dark:to-emerald-400">
            pgit
          </span>{" "}
          <span className="text-neutral-900 dark:text-neutral-300">{"— partial-git"}</span>
        </h1>
        <p className="mx-auto mt-3 max-w-2xl text-neutral-600 dark:text-neutral-300">{subtitle}</p>
        <div className="mt-6 flex flex-col items-center justify-center gap-3 sm:flex-row">
          <Button asChild className="bg-emerald-600 hover:bg-emerald-700">
            <a href="#install" aria-label="Go to installation section">
              <Rocket className="mr-2 h-4 w-4" aria-hidden="true" />
              Quick Install
            </a>
          </Button>
          <Button asChild variant="outline">
            <a 
              href="https://github.com/rushikeshg25/partial-git" 
              target="_blank" 
              rel="noreferrer"
              aria-label="View pgit source code on GitHub (opens in new tab)"
            >
              <Github className="mr-2 h-4 w-4" aria-hidden="true" />
              View on GitHub
            </a>
          </Button>
        </div>
        <div className="mx-auto mt-6 max-w-2xl">
          <CodeBlock language="bash" code={quick} />
          <p className="mt-2 flex items-center justify-center gap-2 text-xs text-neutral-700 dark:text-neutral-200">
            <Download className="h-3.5 w-3.5 text-emerald-600 dark:text-emerald-400" aria-hidden="true" />
            {"Try a quick directory download after installation"}
          </p>
        </div>
      </div>
    </div>
  )
}
