"use client";

import * as React from "react";
import Image from "next/image";
import { AspectRatio } from "@/components/ui/aspect-ratio";
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
  const videoRef = React.useRef<HTMLVideoElement>(null);
  const containerRef = React.useRef<HTMLDivElement>(null);

  React.useEffect(() => {
    const video = videoRef.current;
    const container = containerRef.current;

    if (!video || !container || !videoSrc) return;

    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting) {
            // Video is in viewport, play it
            video.play().catch((error) => {
              console.log("Auto-play failed:", error);
            });
          } else {
            // Video is out of viewport, pause it
            video.pause();
          }
        });
      },
      {
        threshold: 0.5, // Play when 50% of video is visible
      }
    );

    observer.observe(container);

    return () => {
      observer.disconnect();
    };
  }, [videoSrc]);

  return (
    <section className={cn("w-full", className)} aria-labelledby="video-title">
      <div 
        ref={containerRef}
        className="overflow-hidden rounded-lg border-2 border-emerald-200 bg-white shadow-sm hover:border-emerald-300 hover:shadow-[0_10px_30px_-15px_rgba(16,185,129,0.35)] transition-all duration-300 dark:bg-neutral-900 dark:border-emerald-800 dark:hover:border-emerald-700"
      >
        <div className="relative">
          <AspectRatio ratio={16 / 9}>
            {videoSrc ? (
              <video
                ref={videoRef}
                className="w-full h-full object-cover"
                muted
                loop
                playsInline
                poster={posterSrc || undefined}
                aria-label={`Video: ${title}`}
              >
                <source src={videoSrc} type="video/mp4" />
                <source
                  src={videoSrc.replace(".mp4", ".webm")}
                  type="video/webm"
                />
                <p className="flex items-center justify-center h-full text-neutral-600 dark:text-neutral-300">
                  Your browser doesn't support video playback.
                </p>
              </video>
            ) : (
              <Image
                src={
                  "/placeholder.svg?height=720&width=1280&query=pgit%20partial%20git%20video%20placeholder%20demo"
                }
                alt={`Video thumbnail showing pgit demonstration. ${title} - ${subtitle}`}
                fill
                className="object-cover"
                priority
              />
            )}
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
