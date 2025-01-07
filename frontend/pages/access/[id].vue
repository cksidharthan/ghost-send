<template>
    <main class="h-full">
        <div class="relative min-h-[calc(100vh-theme(spacing.16))]">
            <!-- Dark gradient background -->
            <div class="absolute inset-0 bg-gradient-to-b from-gray-900 via-gray-800 to-gray-900"></div>

            <!-- Vector patterns -->
            <div class="absolute inset-0">
                <!-- Grid pattern -->
                <div class="absolute inset-0 [background:radial-gradient(#2a3366_1px,transparent_1px)] [background-size:32px_32px] opacity-30">
                </div>

                <!-- Subtle glow effects -->
                <div class="absolute top-0 -left-4 w-72 h-72 bg-indigo-500/15 rounded-full blur-3xl"></div>
                <div class="absolute bottom-0 -right-4 w-72 h-72 bg-purple-500/15 rounded-full blur-3xl"></div>
            </div>

            <!-- Content overlay gradient -->
            <div class="absolute inset-0 bg-gradient-to-t from-gray-900 via-gray-900/80 to-gray-900/40"></div>

            <!-- Content Container -->
            <div class="relative z-10 mx-auto max-w-lg px-6 py-12 sm:py-16 lg:py-20">
                <div class="bg-gray-800/80 backdrop-blur-sm p-6 sm:p-8 rounded-xl shadow-2xl border border-gray-700/50">
                    <!-- Loading State -->
                    <div v-if="loading" class="text-center py-8">
                        <div class="animate-spin w-10 h-10 border-4 border-indigo-500 border-t-transparent rounded-full mx-auto"></div>
                        <p class="mt-4 text-gray-300">Loading secret...</p>
                    </div>

                    <!-- Error State -->
                    <div v-else-if="error" class="p-6 text-center">
                        <div class="w-12 h-12 rounded-full bg-red-500/10 flex items-center justify-center mx-auto mb-4">
                            <Icon name="heroicons:exclamation-triangle" class="w-6 h-6 text-red-500" />
                        </div>
                        <h2 class="text-xl font-semibold text-red-500 mb-2">Error</h2>
                        <p class="text-gray-300">{{ error }}</p>
                    </div>

                    <!-- Password Input State -->
                    <div v-else-if="!secretText" class="space-y-6">
                        <div class="text-center">
                            <div class="w-12 h-12 rounded-full bg-indigo-500/10 flex items-center justify-center mx-auto mb-4">
                                <Icon name="heroicons:lock-closed" class="w-6 h-6 text-indigo-500" />
                            </div>
                            <h2 class="text-xl font-semibold text-white mb-2">Access Secret</h2>
                            <p class="text-gray-400">Enter the password to view this secret</p>
                        </div>

                        <div class="space-y-4">
                            <div class="relative">
                                <input 
                                    v-model="password"
                                    :type="showPassword ? 'text' : 'password'"
                                    class="block w-full px-3 py-2 border border-gray-600 rounded-lg text-white bg-gray-700/50 backdrop-blur-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 pr-10"
                                    placeholder="Enter password"
                                    @keyup.enter="accessSecret"
                                />
                                <button 
                                    type="button" 
                                    @click="showPassword = !showPassword"
                                    class="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-400 hover:text-white transition-colors duration-200"
                                >
                                    <Icon :name="showPassword ? 'heroicons:eye-slash' : 'heroicons:eye'" class="w-5 h-5" />
                                </button>
                            </div>
                            
                            <button 
                                @click="accessSecret"
                                class="w-full flex justify-center items-center gap-2 py-2 px-4 bg-indigo-600 hover:bg-indigo-500 text-white rounded-lg transition-colors duration-200"
                            >
                                <Icon name="heroicons:key" class="w-5 h-5" />
                                Access Secret
                            </button>
                        </div>
                    </div>

                    <!-- Secret Display State -->
                    <div v-else class="space-y-4">
                        <div class="text-center">
                            <div class="w-12 h-12 rounded-full bg-green-500/10 flex items-center justify-center mx-auto mb-4">
                                <Icon name="heroicons:document-text" class="w-6 h-6 text-green-500" />
                            </div>
                            <h2 class="text-xl font-semibold text-white mb-2">Secret Retrieved</h2>
                            <p class="text-gray-400">Here's your secret message:</p>
                        </div>
                        
                        <div class="bg-gray-700/50 backdrop-blur-sm p-4 rounded-lg border border-gray-600">
                            <p class="text-white break-words">{{ secretText }}</p>
                        </div>
                        
                        <p class="text-sm text-gray-400 text-center">
                            This secret will be destroyed after viewing. Make sure to save it if needed.
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </main>
</template>

<script setup lang="ts">
const route = useRoute()
const { id } = route.params

const loading = ref(true)
const error = ref<string | null>(null)
const password = ref('')
const secretText = ref<string | null>(null)
const showPassword = ref(false)

// Check if secret exists
async function checkSecretStatus() {
    try {
        const response = await fetch(`http://localhost:8080/secrets/${id}/status`)
        console.log(response)
        
        if (response.status === 404) {
            error.value = 'Secret not found or has expired'
        } else if (response.status === 401) {
            error.value = 'Unauthorized access'
        } else if (response.status === 500) {
            error.value = 'Server error occurred'
        } else if (response.status !== 200) {
            error.value = 'An unexpected error occurred'
        }
        
    } catch (e) {
        error.value = 'Failed to check secret status'
    } finally {
        loading.value = false
    }
}

// Access secret with password
async function accessSecret() {
    if (!password.value) return
    
    loading.value = true
    error.value = null
    
    try {
        const response = await fetch(`http://localhost:8080/secrets/${id}?password=${encodeURIComponent(password.value)}`)
        
        if (response.status === 200) {
            const data = await response.json()
            secretText.value = data.secret_text
        } else if (response.status === 401) {
            error.value = 'Incorrect password'
        } else if (response.status === 404) {
            error.value = 'Secret not found or has expired'
        } else {
            error.value = 'Failed to access secret'
        }
    } catch (e) {
        error.value = 'Failed to access secret'
    } finally {
        loading.value = false
    }
}

// Check secret status on page load
onMounted(() => {
    checkSecretStatus()
})
</script>