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
        <div class="relative z-10 min-h-screen flex flex-col items-center justify-center px-4">
            <div class="bg-gray-800/80 backdrop-blur-sm p-8 rounded-xl shadow-2xl border border-gray-700/50 max-w-lg w-full animate-fade-in">
                <!-- 404 Text -->
                <div class="relative mb-8 text-center">
                    <h1 class="text-8xl font-black text-transparent bg-clip-text bg-gradient-to-r from-indigo-500 to-purple-500 animate-pulse">
                        404
                    </h1>
                    <div class="absolute top-0 left-0 w-full h-full flex justify-center items-center">
                        <span class="text-8xl font-black text-indigo-500/10 animate-bounce delay-75">404</span>
                        <span class="text-8xl font-black text-purple-500/10 animate-bounce delay-150">404</span>
                    </div>
                </div>

                <h2 class="text-2xl font-bold text-white text-center mb-4">Page Not Found</h2>
                <p class="text-gray-400 text-center mb-8">
                    The page you're looking for seems to have vanished into the digital void.
                </p>

                <NuxtLink 
                    to="/" 
                    class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-medium rounded-lg text-white bg-indigo-600 hover:bg-indigo-500 transition-all duration-300"
                >
                    <span class="absolute left-0 inset-y-0 flex items-center pl-3 transition-transform group-hover:-translate-x-1">
                        <Icon name="heroicons:arrow-left" class="w-5 h-5 text-indigo-500 group-hover:text-indigo-400" />
                    </span>
                    <span class="pl-6">Return Home</span>
                </NuxtLink>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
definePageMeta({
    layout: 'error'
})
</script>