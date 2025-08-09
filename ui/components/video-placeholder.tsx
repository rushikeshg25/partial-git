"use client";

import * as React from "react";
import Image from "next/image";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { AspectRatio } from "@/components/ui/aspect-ratio";
import { Play } from "lucide-react";
import { cn } from "@/lib/utils";

export default function VideoPlaceholder({
  title = "See pgit in action",
  subtitle = "Quick overview",
  cta = "Watch preview",
  className = "",
  videoSrc = null,
  posterSrc = null,
}: {
  title?: string;
  subtitle?: string;
  cta?: string;
  className?: string;
  videoSrc?: string | null;
  posterSrc?: string | null;
}) {
  const [open, setOpen] = React.useState(false);

  return (
    <section className={cn("w-full", className)} aria-labelledby="video-title">
      <div className="overflow-hidden rounded-lg border-2 border-emerald-200 bg-white shadow-sm hover:border-emerald-300 hover:shadow-[0_10px_30px_-15px_rgba(16,185,129,0.35)] transition-all duration-300 dark:bg-neutral-900 dark:border-emerald-800 dark:hover:border-emerald-700">
        <div className="relative">
          <AspectRatio ratio={16 / 9}>
            <Image
              src={
                "/placeholder.svg?height=720&width=1280&query=pgit%20partial%20git%20video%20placeholder%20demo"
              }
              alt={`Video thumbnail showing pgit demonstration. ${title} - ${subtitle}`}
              fill
              className="object-cover"
              priority
            />
            <div
              className="absolute inset-0 bg-gradient-to-t from-black/40 via-emerald-900/5 to-black/10"
              aria-hidden="true"
            />
            <div className="absolute inset-0 flex items-center justify-center">
              <div className="relative">
                <span
                  className="absolute inset-0 -z-10 m-auto inline-flex h-12 w-12 sm:h-14 sm:w-14 animate-ping rounded-full bg-emerald-500/40"
                  aria-hidden="true"
                />
                <Dialog open={open} onOpenChange={setOpen}>
                  <DialogTrigger asChild>
                    <Button
                      size="lg"
                      className="gap-2 bg-emerald-600 hover:bg-emerald-700 focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2"
                      aria-label={`${cta} - Opens video dialog for ${title}`}
                    >
                      <Play className="h-4 w-4" aria-hidden="true" />
                      <span className="hidden sm:inline">{cta}</span>
                      <span className="sm:hidden">Play</span>
                    </Button>
                  </DialogTrigger>
                  <DialogContent
                    className="max-w-[95vw] sm:max-w-2xl md:max-w-3xl lg:max-w-4xl"
                    aria-describedby="video-description"
                  >
                    <DialogHeader>
                      <DialogTitle>{title}</DialogTitle>
                    </DialogHeader>
                    <div className="rounded-md border border-emerald-200 dark:border-emerald-800 overflow-hidden">
                      <AspectRatio ratio={16 / 9}>
                        {videoSrc ? (
                          <video
                            className="w-full h-full object-cover"
                            controls
                            preload="metadata"
                            poster={posterSrc || undefined}
                            aria-label={`Video: ${title}`}
                          >
                            <source src={videoSrc} type="video/mp4" />
                            <source
                              src={videoSrc.replace(".mp4", ".webm")}
                              type="video/webm"
                            />
                            {/* Fallback for browsers that don't support video */}
                            <p className="flex items-center justify-center h-full text-neutral-600 dark:text-neutral-300">
                              Your browser doesn't support video playback.
                            </p>
                          </video>
                        ) : (
                          <Image
                            src={
                              "/placeholder.svg?height=720&width=1280&query=pgit%20demo%20video%20coming%20soon"
                            }
                            alt="Placeholder image indicating that the demo video is coming soon"
                            fill
                            className="object-cover"
                          />
                        )}
                      </AspectRatio>
                    </div>
                    <p
                      id="video-description"
                      className="text-sm text-neutral-700 dark:text-neutral-200 flex items-center gap-2"
                    >
                      <span
                        className="inline-block w-2 h-2 bg-emerald-500 rounded-full"
                        aria-hidden="true"
                      ></span>
                      {
                        "Demo video is coming soon. This is a placeholder showing where the pgit demonstration video will be displayed."
                      }
                    </p>
                  </DialogContent>
                </Dialog>
              </div>
            </div>
          </AspectRatio>
        </div>
        <div className="border-t border-neutral-200 px-4 py-3 bg-gradient-to-r from-emerald-50/50 to-white dark:border-neutral-800 dark:from-emerald-950/20 dark:to-neutral-900">
          <h3
            id="video-title"
            className="font-medium text-neutral-900 dark:text-neutral-100"
          >
            {title}
          </h3>
          <p className="text-sm text-neutral-700 dark:text-neutral-200 flex items-center gap-1">
            <span
              className="inline-block w-1.5 h-1.5 bg-emerald-500 rounded-full animate-pulse"
              aria-hidden="true"
            ></span>
            {subtitle}
          </p>
        </div>
      </div>
    </section>
  );
}
