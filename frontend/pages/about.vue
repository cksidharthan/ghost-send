<template>
  <div class="min-h-screen bg-gray-900 flex flex-col">
    <!-- Main Content -->
    <main class="flex-1 flex flex-col">
      <!-- Hero Section with animated background -->
      <div class="relative isolate overflow-hidden">
        <!-- Dark gradient background -->
        <div class="absolute inset-0 bg-gradient-to-b from-gray-900 via-gray-800 to-gray-900"></div>
        
        <!-- Vector strokes and patterns -->
        <div class="absolute inset-0">
          <!-- Grid pattern -->
          <div class="absolute inset-0 [background:radial-gradient(#1c274f_1px,transparent_1px)] [background-size:32px_32px] opacity-20"></div>
          
          <!-- Animated vector lines -->
          <div class="absolute inset-0">
            <svg class="w-full h-full" viewBox="0 0 1200 800" fill="none" xmlns="http://www.w3.org/2000/svg">
              <!-- Diagonal lines -->
              <path 
                d="M-100 200L1300 -100" 
                stroke="url(#line-gradient)" 
                stroke-width="1" 
                stroke-opacity="0.1"
              />
              <path 
                d="M-100 400L1300 100" 
                stroke="url(#line-gradient)" 
                stroke-width="1" 
                stroke-opacity="0.1"
              />
              <path 
                d="M-100 600L1300 300" 
                stroke="url(#line-gradient)" 
                stroke-width="1" 
                stroke-opacity="0.1"
              />
              
              <!-- Curved accent lines -->
              <path
                d="M-100 700Q300 600 600 700T1300 700"
                stroke="url(#line-gradient)"
                stroke-width="1"
                stroke-opacity="0.2"
                class="animate-pulse"
              />
              
              <!-- Gradient definitions -->
              <defs>
                <linearGradient id="line-gradient" x1="0%" y1="0%" x2="100%" y2="0%">
                  <stop offset="0%" style="stop-color: rgb(99, 102, 241)" />
                  <stop offset="100%" style="stop-color: rgb(168, 85, 247)" />
                </linearGradient>
              </defs>
            </svg>
          </div>
          
          <!-- Subtle glow effects -->
          <div class="absolute top-0 -left-4 w-72 h-72 bg-indigo-500/10 rounded-full blur-3xl"></div>
          <div class="absolute bottom-0 -right-4 w-72 h-72 bg-purple-500/10 rounded-full blur-3xl"></div>
        </div>

        <!-- Content overlay gradient -->
        <div class="absolute inset-0 bg-gradient-to-t from-gray-900 via-gray-900/80 to-gray-900/40"></div>

        <div class="mx-auto max-w-7xl px-6 lg:px-8 relative z-10">
          <div class="mx-auto max-w-3xl py-12 sm:py-16 lg:py-20">
            <div class="text-center">
              <h1 class="text-4xl font-bold tracking-tight text-white sm:text-6xl animate-fade-in-up">
                About Share Text
              </h1>
              <p class="mt-6 text-lg leading-8 text-gray-300 animate-fade-in-up delay-150">
                A secure platform for sharing sensitive information with end-to-end encryption
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Features Section -->
      <div class="py-12 sm:py-16 relative">
        <div class="mx-auto max-w-7xl px-6 lg:px-8">
          <div class="mx-auto max-w-2xl lg:text-center animate-fade-in-up delay-300">
            <h2 class="text-base font-semibold leading-7 text-indigo-400">Secure by Design</h2>
            <p class="mt-2 text-3xl font-bold tracking-tight text-white sm:text-4xl">
              Everything you need to share securely
            </p>
          </div>

          <div class="mx-auto mt-16 max-w-2xl sm:mt-20 lg:mt-24 lg:max-w-4xl">
            <dl class="grid max-w-xl grid-cols-1 gap-x-8 gap-y-10 lg:max-w-none lg:grid-cols-2 lg:gap-y-16">
              <!-- Feature 1 -->
              <div class="relative pl-16 animate-fade-in-up delay-[400ms]">
                <dt class="text-base font-semibold leading-7 text-white">
                  <div class="absolute left-0 top-0 flex h-10 w-10 items-center justify-center rounded-lg bg-indigo-600">
                    <svg class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M16.5 10.5V6.75a4.5 4.5 0 10-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 002.25-2.25v-6.75a2.25 2.25 0 00-2.25-2.25H6.75a2.25 2.25 0 00-2.25 2.25v6.75a2.25 2.25 0 002.25 2.25z" />
                    </svg>
                  </div>
                  End-to-End Encryption
                </dt>
                <dd class="mt-2 text-base leading-7 text-gray-300">
                  Your data is encrypted before it leaves your browser and can only be decrypted by the intended recipient.
                </dd>
              </div>

              <!-- Feature 2 -->
              <div class="relative pl-16 animate-fade-in-up delay-[500ms]">
                <dt class="text-base font-semibold leading-7 text-white">
                  <div class="absolute left-0 top-0 flex h-10 w-10 items-center justify-center rounded-lg bg-indigo-600">
                    <svg class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  Auto-Expiration
                </dt>
                <dd class="mt-2 text-base leading-7 text-gray-300">
                  Set expiration times for your shared content. Once expired, the data is permanently deleted.
                </dd>
              </div>

              <!-- Feature 3 -->
              <div class="relative pl-16 animate-fade-in-up delay-[600ms]">
                <dt class="text-base font-semibold leading-7 text-white">
                  <div class="absolute left-0 top-0 flex h-10 w-10 items-center justify-center rounded-lg bg-indigo-600">
                    <svg class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M2.036 12.322a1.012 1.012 0 010-.639C3.423 7.51 7.36 4.5 12 4.5c4.638 0 8.573 3.007 9.963 7.178.07.207.07.431 0 .639C20.577 16.49 16.64 19.5 12 19.5c-4.638 0-8.573-3.007-9.963-7.178z" />
                      <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                  </div>
                  View Limits
                </dt>
                <dd class="mt-2 text-base leading-7 text-gray-300">
                  Control how many times your shared content can be viewed before it's automatically deleted.
                </dd>
              </div>

              <!-- Feature 4 -->
              <div class="relative pl-16 animate-fade-in-up delay-[700ms]">
                <dt class="text-base font-semibold leading-7 text-white">
                  <div class="absolute left-0 top-0 flex h-10 w-10 items-center justify-center rounded-lg bg-indigo-600">
                    <svg class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75L11.25 15 15 9.75m-3-7.036A11.959 11.959 0 013.598 6 11.99 11.99 0 003 9.749c0 5.592 3.824 10.29 9 11.623 5.176-1.332 9-6.03 9-11.622 0-1.31-.21-2.571-.598-3.751h-.152c-3.196 0-6.1-1.248-8.25-3.285z" />
                    </svg>
                  </div>
                  Password Protection
                </dt>
                <dd class="mt-2 text-base leading-7 text-gray-300">
                  Add an extra layer of security with optional password protection for your shared content.
                </dd>
              </div>
            </dl>
          </div>
        </div>
      </div>

      <!-- CTA Section -->
      <div class="relative isolate mt-12 px-6 py-12 sm:py-16 lg:px-8">
        <div class="mx-auto max-w-2xl text-center animate-fade-in-up delay-[800ms]">
          <h2 class="text-3xl font-bold tracking-tight text-white sm:text-4xl">
            Ready to share securely?
          </h2>
          <p class="mx-auto mt-6 max-w-xl text-lg leading-8 text-gray-300">
            Start sharing sensitive information with confidence today.
          </p>
          <div class="mt-10 flex items-center justify-center gap-x-6">
            <NuxtLink
              to="/"
              class="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 transition-all duration-300 hover:scale-105"
            >
              Get started
            </NuxtLink>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>
