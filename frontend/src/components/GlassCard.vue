<template>
  <div class="glass-card" :class="{ 'hover-effect': hoverEffect }">
    <slot></slot>
  </div>
</template>

<script setup>
defineProps({
  hoverEffect: {
    type: Boolean,
    default: false
  }
})
</script>

<style scoped>
.glass-card {
  background: var(--card-bg);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 20px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.05);
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
}

:global(.dark) .glass-card {
  background: var(--card-bg);
  border-color: rgba(255, 255, 255, 0.1);
}

.hover-effect:hover {
  transform: translateY(-5px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.12);
  border-color: var(--primary-color);
}
</style>
