<script setup lang="ts">
import type { NuxtError } from '#app'

defineProps({
  error: Object as () => NuxtError
})

const handleError = () => {
  clearError()
  navigateTo('/')
}
</script>

<template>
  <div class="min-h-screen relative overflow-hidden">
    <!-- Dark gradient background -->
    <div class="absolute inset-0 bg-gradient-to-b from-gray-900 via-gray-800 to-gray-900"></div>

    <!-- Vector strokes and patterns -->
    <div class="absolute inset-0">
      <!-- Grid pattern -->
      <div class="absolute inset-0 [background:radial-gradient(#2a3366_1px,transparent_1px)] [background-size:32px_32px] opacity-30">
      </div>

      <!-- Animated vector lines -->
      <div class="absolute inset-0">
        <svg class="w-full h-full" viewBox="0 0 1200 800" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path d="M-100 200L1300 -100" stroke="url(#line-gradient)" stroke-width="1.5" stroke-opacity="0.15" />
          <path d="M-100 400L1300 100" stroke="url(#line-gradient)" stroke-width="1.5" stroke-opacity="0.15" />
          <path d="M-100 600L1300 300" stroke="url(#line-gradient)" stroke-width="1.5" stroke-opacity="0.15" />
          <path d="M-100 800L1300 500" stroke="url(#line-gradient)" stroke-width="1.5" stroke-opacity="0.15" />
          
          <defs>
            <linearGradient id="line-gradient" x1="0%" y1="0%" x2="100%" y2="0%">
              <stop offset="0%" style="stop-color: rgb(129, 132, 255)" />
              <stop offset="100%" style="stop-color: rgb(188, 105, 267)" />
            </linearGradient>
          </defs>
        </svg>
      </div>

      <!-- Subtle glow effects -->
      <div class="absolute top-0 -left-4 w-72 h-72 bg-indigo-500/15 rounded-full blur-3xl"></div>
      <div class="absolute bottom-0 -right-4 w-72 h-72 bg-purple-500/15 rounded-full blur-3xl"></div>
    </div>

    <!-- Content overlay gradient -->
    <div class="absolute inset-0 bg-gradient-to-t from-gray-900 via-gray-900/80 to-gray-900/40"></div>

    <!-- Main content -->
    <slot />
  </div>
</template>