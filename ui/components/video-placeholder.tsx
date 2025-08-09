"use client";

import * as React from "react";
import Image from "next/image";
import { AspectRatio } from "@/components/ui/aspect-ratio";
import { Loader2 } from "lucide-react";
import { cn } from "@/lib/utils";

export default function VideoPlaceholder({
  title = "See pgit in action",
  subtitle = "Quick overview",
  className = "",
  videoSrc = null,
  posterSrc = null,
}: {
  title?: string;
  subtitle?: string;
  className?: string;
  videoSrc?: string | null;
  posterSrc?: string | null;
}) {
  const videoRef = React.useRef<HTMLVideoElement>(null);
  const containerRef = React.useRef<HTMLDivElement>(null);
  const [isLoading, setIsLoading] = React.useState(!!videoSrc);
  const [hasError, setHasError] = React.useState(false);
  const [canPlay, setCanPlay] = React.useState(false);

  React.useEffect(() => {
    const video = videoRef.current;
    const container = containerRef.current;

    if (!video || !container || !videoSrc) {
      setIsLoading(false);
      return;
    }

    // Reset states when videoSrc changes
    setIsLoading(true);
    setHasError(false);
    setCanPlay(false);

    // Video event handlers
    const handleLoadStart = () => {
      setIsLoading(true);
      setHasError(false);
    };

    const handleCanPlay = () => {
      setCanPlay(true);
      setIsLoading(false);
    };

    const handleLoadedData = () => {
      setIsLoading(false);
    };

    const handleError = (e: Event) => {
      console.error("Video loading error:", e);
      setIsLoading(false);
      setHasError(true);
    };

    const handleLoadedMetadata = () => {
      setIsLoading(false);
      setCanPlay(true);
    };

    // Add event listeners
    video.addEventListener('loadstart', handleLoadStart);
    video.addEventListener('canplay', handleCanPlay);
    video.addEventListener('loadeddata', handleLoadedData);
    video.addEventListener('loadedmetadata', handleLoadedMetadata);
    video.addEventListener('error', handleError);

    // Force load if video is already ready
    if (video.readyState >= 3) { // HAVE_FUTURE_DATA or higher
      setIsLoading(false);
      setCanPlay(true);
    }

    // Fallback timeout to prevent infinite loading
    const timeoutId = setTimeout(() => {
      if (isLoading) {
        console.warn("Video loading timeout, showing error state");
        setIsLoading(false);
        setHasError(true);
      }
    }, 10000); // 10 second timeout

    const observer = new IntersectionObserver(
      (entries) => {
        entries.forEach((entry) => {
          if (entry.isIntersecting && canPlay) {
            // Video is in viewport and ready, play it
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
      clearTimeout(timeoutId);
      video.removeEventListener('loadstart', handleLoadStart);
      video.removeEventListener('canplay', handleCanPlay);
      video.removeEventListener('loadeddata', handleLoadedData);
      video.removeEventListener('loadedmetadata', handleLoadedMetadata);
      video.removeEventListener('error', handleError);
    };
  }, [videoSrc, canPlay]);

  return (
    <section className={cn("w-full", className)} aria-labelledby="video-title">
      <div 
        ref={containerRef}
        className="overflow-hidden rounded-lg border-2 border-emerald-200 bg-white shadow-sm hover:border-emerald-300 hover:shadow-[0_10px_30px_-15px_rgba(16,185,129,0.35)] transition-all duration-300 dark:bg-neutral-900 dark:border-emerald-800 dark:hover:border-emerald-700"
      >
        <div className="relative">
          <AspectRatio ratio={16 / 9}>
            {videoSrc ? (
              <>
                <video
                  ref={videoRef}
                  className={cn(
                    "w-full h-full object-cover transition-opacity duration-300",
                    isLoading || hasError ? "opacity-0" : "opacity-100"
                  )}
                  muted
                  loop
                  playsInline
                  preload="metadata"
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

                {/* Loading overlay */}
                {isLoading && (
                  <div className="absolute inset-0 flex items-center justify-center bg-neutral-100 dark:bg-neutral-800">
                    <div className="flex flex-col items-center gap-3">
                      <Loader2 className="h-8 w-8 animate-spin text-emerald-600" />
                      <p className="text-sm text-neutral-600 dark:text-neutral-300">
                        Loading video...
                      </p>
                    </div>
                  </div>
                )}

                {/* Error overlay */}
                {hasError && (
                  <div className="absolute inset-0 flex items-center justify-center bg-neutral-100 dark:bg-neutral-800">
                    <div className="flex flex-col items-center gap-3 text-center px-4">
                      <div className="w-12 h-12 rounded-full bg-red-100 dark:bg-red-900/20 flex items-center justify-center">
                        <span className="text-red-600 dark:text-red-400 text-xl">⚠</span>
                      </div>
                      <div>
                        <p className="text-sm font-medium text-neutral-900 dark:text-neutral-100">
                          Failed to load video
                        </p>
                        <p className="text-xs text-neutral-600 dark:text-neutral-400 mt-1">
                          Please check your connection and try again
                        </p>
                      </div>
                    </div>
                  </div>
                )}

                {/* Poster/placeholder when no video is loading */}
                {!isLoading && !hasError && posterSrc && (
                  <Image
                    src={posterSrc}
                    alt={`Video poster for ${title}`}
                    fill
                    className={cn(
                      "object-cover transition-opacity duration-300",
                      canPlay ? "opacity-0" : "opacity-100"
                    )}
                  />
                )}
              </>
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
