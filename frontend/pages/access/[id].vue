<template>
    <div>
        <h1>Access</h1>
        <p>Accessing secret with id: {{ id }}</p>
    </div>
</template>

<script setup lang="ts">
const { id } = useRoute().params
</script>