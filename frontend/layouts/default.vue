<template>
  <div class="min-h-screen bg-gray-900 flex flex-col">
    <!-- Navbar -->
    <nav class="bg-gray-800/80 backdrop-blur-sm border-b border-gray-700/50 sticky top-0 z-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
          <!-- Logo -->
          <div class="flex-shrink-0">
            <NuxtLink to="/" class="group">
              <h1 class="text-xl sm:text-2xl font-bold bg-gradient-to-r from-indigo-500 to-purple-500 bg-clip-text text-transparent group-hover:from-indigo-400 group-hover:to-purple-400 transition-all duration-300">
                <span class="flex items-center gap-2">
                  <Icon name="heroicons:lock-closed" class="w-6 h-6 text-indigo-500" />
                  Share Secret
                </span>
              </h1>
            </NuxtLink>
          </div>

          <!-- Mobile menu button -->
          <div class="flex sm:hidden">
            <button @click="isMobileMenuOpen = !isMobileMenuOpen" class="text-gray-400 hover:text-white p-2 transition-colors duration-200">
              <span class="sr-only">Open menu</span>
              <Icon :name="isMobileMenuOpen ? 'heroicons:x-mark' : 'heroicons:bars-3'" class="w-6 h-6" />
            </button>
          </div>

          <!-- Desktop menu -->
          <div class="hidden sm:flex sm:items-center sm:space-x-8">
            <NuxtLink to="/about" class="group text-gray-300 hover:text-white px-3 py-2 text-sm font-medium transition-colors duration-200">
              About
              <span class="block max-w-0 group-hover:max-w-full transition-all duration-500 h-0.5 bg-indigo-500"></span>
            </NuxtLink>
            <a href="https://github.com/yourusername/project" target="_blank" class="text-gray-300 hover:text-white transition-colors duration-200">
              <Icon name="mdi:github" class="w-6 h-6" />
            </a>
          </div>
        </div>
      </div>

      <!-- Mobile menu -->
      <div 
        v-show="isMobileMenuOpen" 
        ref="mobileMenuRef"
        class="sm:hidden absolute w-full bg-gray-800/95 backdrop-blur-sm border-b border-gray-700/50"
      >
        <div class="px-2 pt-2 pb-3 space-y-1">
          <NuxtLink to="/about" class="block text-gray-300 hover:text-white hover:bg-gray-700/50 px-3 py-2 rounded-md text-base font-medium transition-all duration-200">
            About
          </NuxtLink>
        </div>
      </div>
    </nav>

    <!-- Page Content -->
    <div class="flex-1">
      <slot />
    </div>

    <!-- Footer -->
    <footer class="mt-auto bg-gray-800/80 backdrop-blur-sm border-t border-gray-700/50">
      <div class="max-w-7xl mx-auto py-8 px-4 sm:px-6 lg:px-8">
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
          <!-- Logo and Description -->
          <div class="space-y-4">
            <h3 class="text-xl font-bold bg-gradient-to-r from-indigo-500 to-purple-500 bg-clip-text text-transparent">
                <span class="flex items-center gap-2">
                  <Icon name="heroicons:lock-closed" class="w-6 h-6 text-indigo-500" />
                  Share Secret
                </span>
            </h3>
            <p class="text-gray-400 text-sm">
              Share secrets securely with end-to-end encryption and automatic expiration.
            </p>
          </div>

          <!-- Quick Links -->
          <div class="space-y-4">
            <h4 class="text-white font-semibold">Quick Links</h4>
            <ul class="space-y-2">
              <li>
                <NuxtLink to="/" class="text-gray-400 hover:text-white transition-colors duration-200">
                  Home
                </NuxtLink>
              </li>
              <li>
                <NuxtLink to="/about" class="text-gray-400 hover:text-white transition-colors duration-200">
                  About
                </NuxtLink>
              </li>
            </ul>
          </div>

          <!-- Social Links -->
          <div class="space-y-4">
            <h4 class="text-white font-semibold">Connect</h4>
            <div class="flex space-x-4">
              <a
                href="https://github.com/yourusername/project"
                target="_blank"
                class="text-gray-400 hover:text-white transition-colors duration-200"
              >
                <Icon name="mdi:github" class="w-6 h-6" />
              </a>
            </div>
          </div>
        </div>

        <!-- Copyright -->
        <div class="mt-8 pt-8 border-t border-gray-700">
          <p class="text-center text-sm text-gray-400">
            © {{ new Date().getFullYear() }} Share Secret. All rights reserved.
          </p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const isMobileMenuOpen = ref(false)
const mobileMenuRef = ref(null)

useClickOutside(mobileMenuRef, () => {
  isMobileMenuOpen.value = false
})
</script>
