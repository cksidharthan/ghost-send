<template>
    <main class="h-full">
        <!-- Hero Section with Form -->
        <div class="relative min-h-[calc(100vh-theme(spacing.16)-theme(spacing.32))]">
            <!-- Dark gradient background -->
            <div class="absolute inset-0 bg-gradient-to-b from-gray-900 via-gray-800 to-gray-900"></div>

            <!-- Vector strokes and patterns -->
            <div class="absolute inset-0">
                <!-- Grid pattern -->
                <div
                    class="absolute inset-0 [background:radial-gradient(#2a3366_1px,transparent_1px)] [background-size:32px_32px] opacity-30">
                </div>

                <!-- Animated vector lines -->
                <div class="absolute inset-0">
                    <svg class="w-full h-full" viewBox="0 0 1200 800" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <!-- Diagonal lines -->
                        <path d="M-100 200L1300 -100" stroke="url(#line-gradient)" stroke-width="1.5"
                            stroke-opacity="0.15" />
                        <path d="M-100 400L1300 100" stroke="url(#line-gradient)" stroke-width="1.5"
                            stroke-opacity="0.15" />
                        <path d="M-100 600L1300 300" stroke="url(#line-gradient)" stroke-width="1.5"
                            stroke-opacity="0.15" />
                        <path d="M-100 800L1300 500" stroke="url(#line-gradient)" stroke-width="1.5"
                            stroke-opacity="0.15" />

                        <!-- Curved accent lines -->
                        <path d="M-100 700Q300 600 600 700T1300 700" stroke="url(#line-gradient)" stroke-width="1.5"
                            stroke-opacity="0.25" class="animate-pulse" />
                        <path d="M-100 750Q300 650 600 750T1300 750" stroke="url(#line-gradient)" stroke-width="1"
                            stroke-opacity="0.1" />

                        <!-- Gradient definitions -->
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

            <!-- Content Container -->
            <div class="mx-auto max-w-7xl px-6 py-12 sm:py-16 lg:py-20 relative z-10">
                <div class="flex flex-col lg:flex-row lg:items-center lg:gap-x-16">
                    <!-- Left Side - Hero Content -->
                    <div class="max-w-2xl lg:max-w-xl lg:pt-8 mb-12 lg:mb-0 animate-fade-in">
                        <h1 class="text-4xl font-bold tracking-tight text-white sm:text-5xl lg:text-6xl">
                            Share secrets
                            <span
                                class="block mt-2 bg-gradient-to-r from-indigo-500 to-purple-500 bg-clip-text text-transparent">
                                securely with confidence
                            </span>
                        </h1>
                        <p class="mt-6 text-lg sm:text-xl leading-8 text-gray-300 animate-fade-in-delay">
                            Share sensitive information securely with end-to-end encryption.
                            Set expiration times and view limits to ensure your data stays protected.
                        </p>

                        <!-- Feature List -->
                        <div class="mt-10 space-y-6">
                            <div class="flex items-center gap-x-3 animate-fade-in-up delay-200">
                                <div
                                    class="flex-shrink-0 w-8 h-8 flex items-center justify-center rounded-lg bg-indigo-500/10">
                                    <Icon name="heroicons:shield-check" class="w-5 h-5 text-indigo-500" />
                                </div>
                                <p class="text-base text-gray-300">End-to-end encryption for maximum security</p>
                            </div>
                            <div class="flex items-center gap-x-3 animate-fade-in-up delay-300">
                                <div
                                    class="flex-shrink-0 w-8 h-8 flex items-center justify-center rounded-lg bg-indigo-500/10">
                                    <Icon name="heroicons:clock" class="w-5 h-5 text-indigo-500" />
                                </div>
                                <p class="text-base text-gray-300">Automatic expiration after specified time</p>
                            </div>
                            <div class="flex items-center gap-x-3 animate-fade-in-up delay-400">
                                <div
                                    class="flex-shrink-0 w-8 h-8 flex items-center justify-center rounded-lg bg-indigo-500/10">
                                    <Icon name="heroicons:eye" class="w-5 h-5 text-indigo-500" />
                                </div>
                                <p class="text-base text-gray-300">Limited views for enhanced privacy</p>
                            </div>
                        </div>
                    </div>

                    <!-- Right Side - Form -->
                    <div class="lg:flex-1 lg:min-w-[500px]">
                        <div
                            class="bg-gray-800/80 backdrop-blur-sm p-6 sm:p-8 rounded-xl shadow-2xl transition-all duration-300 hover:shadow-indigo-500/10 border border-gray-700/50 animate-fade-in-up delay-200">
                            <div>
                                <h2 class="text-center text-2xl font-extrabold text-white">
                                    <span class="flex items-center justify-center gap-2">
                                        <Icon name="heroicons:shield-check" class="w-6 h-6 text-indigo-500" />
                                        Share a Secret
                                    </span>
                                </h2>
                                <p class="mt-2 text-center text-sm text-gray-400">
                                    Secure, encrypted, and temporary
                                </p>
                            </div>

                            <form @submit.prevent="submitSecret" class="mt-8 space-y-6">
                                <!-- Secret Input -->
                                <div class="space-y-2">
                                    <label for="secret" class="block text-sm font-medium text-gray-300">
                                        <span class="flex items-center gap-2">
                                            <Icon name="heroicons:document-text" class="w-5 h-5" />
                                            Your Secret
                                        </span>
                                    </label>
                                    <textarea v-model="formData.secret_text" id="secret" required
                                        class="appearance-none relative block w-full px-3 py-2 border border-gray-600 placeholder-gray-500 text-white rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 text-sm bg-gray-700/50 backdrop-blur-sm transition-all duration-300"
                                        placeholder="Enter the secret text or password you want to share"
                                        rows="3"></textarea>
                                </div>

                                <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
                                    <!-- Expiration Select -->
                                    <div class="space-y-2">
                                        <label for="expiration" class="block text-sm font-medium text-gray-300">
                                            <span class="flex items-center gap-2">
                                                <Icon name="heroicons:clock" class="w-5 h-5" />
                                                Expires After
                                            </span>
                                        </label>
                                        <select v-model="formData.expires_at" id="expiration" required
                                            class="appearance-none block w-full px-3 py-2 pr-10 text-sm border border-gray-600 rounded-lg text-white bg-gray-700/50 backdrop-blur-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all duration-300">
                                            <option value="5m">5 minutes</option>
                                            <option value="1h">1 hour</option>
                                            <option value="1d">1 day</option>
                                            <option value="7d">7 days</option>
                                        </select>
                                    </div>

                                    <!-- Views Input -->
                                    <div class="space-y-2">
                                        <label for="views" class="block text-sm font-medium text-gray-300">
                                            <span class="flex items-center gap-2">
                                                <Icon name="heroicons:eye" class="w-5 h-5" />
                                                Maximum Views
                                            </span>
                                        </label>
                                        <input v-model.number="formData.views" type="number" id="views" min="1"
                                            max="100" required
                                            class="block w-full px-3 py-2 border border-gray-600 rounded-lg text-white bg-gray-700/50 backdrop-blur-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 transition-all duration-300" />
                                    </div>
                                </div>

                                <!-- Password Input -->
                                <div class="space-y-2">
                                    <label for="password" class="block text-sm font-medium text-gray-300">
                                        <span class="flex items-center gap-2">
                                            <Icon name="heroicons:key" class="w-5 h-5" />
                                            Access Password
                                        </span>
                                    </label>
                                    <div class="relative">
                                        <input v-model="formData.password" :type="showPassword ? 'text' : 'password'"
                                            id="password"
                                            class="block w-full px-3 py-2 border border-gray-600 rounded-lg text-white bg-gray-700/50 backdrop-blur-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 pr-10 transition-all duration-300"
                                            placeholder="Set a password to protect this secret" required/>
                                        <button type="button" @click="showPassword = !showPassword"
                                            class="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-400 hover:text-white transition-colors duration-200">
                                            <Icon :name="showPassword ? 'heroicons:eye-slash' : 'heroicons:eye'"
                                                class="w-5 h-5" />
                                        </button>
                                    </div>
                                </div>

                                <!-- Submit Button -->
                                <button type="submit" :disabled="loading"
                                    class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-medium rounded-lg text-white bg-indigo-600 hover:bg-indigo-500 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed transition-all duration-300">
                                    <span class="absolute left-0 inset-y-0 flex items-center pl-3">
                                        <Icon name="heroicons:shield-check"
                                            class="w-5 h-5 text-indigo-500 group-hover:text-indigo-400 transition-colors duration-300" />
                                    </span>
                                    {{ loading ? 'Creating...' : 'Create Secure Link' }}
                                </button>
                            </form>

                            <!-- Result Section -->
                            <div v-if="result" class="mt-4 p-4 bg-gray-700/50 backdrop-blur-sm rounded-lg border border-gray-600">
                                <div class="flex items-center justify-between">
                                    <div class="flex-1 break-all">
                                        <p class="text-sm text-gray-400 mb-1">Share this secure link:</p>
                                        <p class="text-white font-mono">{{ accessUrl }}</p>
                                    </div>
                                    <button @click="copyToClipboard(accessUrl)"
                                        class="ml-4 p-2 text-gray-400 hover:text-white transition-colors duration-200 flex-shrink-0">
                                        <Icon name="heroicons:clipboard" class="w-5 h-5" />
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </main>
</template>

<script setup>
const isMobileMenuOpen = ref(false)
const showPassword = ref(false)
const loading = ref(false)
const result = ref('')

const formData = ref({
    secret_text: '',
    password: '',
    expires_at: '5m',
    views: 1
})

const config = useRuntimeConfig()
const accessUrl = computed(() => {
    if (!result.value) return ''
    const baseUrl = window.location.origin
    return `${baseUrl}/access/${result.value}`
})

const submitSecret = async () => {
    loading.value = true
    try {
        const response = await fetch(config.public.shareSecretApiUrl + '/secrets', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(formData.value),
        })
        const data = await response.json()
        result.value = data.secret_id // Store just the ID
    } catch (error) {
        result.value = 'Error creating secret. Please try again.'
    } finally {
        loading.value = false
    }
}

const toast = useToast()

const copyToClipboard = async (text) => {
    try {
        await navigator.clipboard.writeText(text)
        toast.add({
            title: 'Success',
            description: 'Link copied to clipboard',
            timeout: 2000
        })
    } catch (error) {
        console.error('Failed to copy text:', error)
        toast.add({
            title: 'Error',
            description: 'Failed to copy link to clipboard',
            timeout: 2000
        })
    }
}
</script>

<style>
.bg-grid-pattern {
    background-image: url('/grid.svg');
    background-repeat: repeat;
    mask-image: radial-gradient(circle at center, black, transparent 80%);
}
</style>
