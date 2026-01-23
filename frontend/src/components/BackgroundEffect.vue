<template>
  <div class="background-effect">
    <!-- Background Image Layer -->
    <div class="bg-image" :style="{ backgroundImage: `url(/images/${bgImage})` }"></div>
    
    <!-- Firefly Particles Layer -->
    <canvas ref="canvas" class="firefly-canvas"></canvas>
  </div>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount } from 'vue';

const canvas = ref(null);
const bgImage = ref('starry-bg.jpg'); // Default
let ctx = null;
let animationFrameId = null;
let particles = [];

// Configuration
const PARTICLE_COUNT = 80; // Increased count
const MIN_SPEED = 0.3;
const MAX_SPEED = 1.0;
const MIN_SIZE = 1.5; // Slightly larger
const MAX_SIZE = 4;

const updateBgImage = () => {
    const stored = localStorage.getItem('bg_image');
    if (stored) {
        bgImage.value = stored;
    }
};

class Particle {
    constructor(w, h) {
        this.reset(w, h, true);
    }

    reset(w, h, initial = false) {
        this.x = Math.random() * w;
        this.y = initial ? Math.random() * h : h + 10;
        this.radius = Math.random() * (MAX_SIZE - MIN_SIZE) + MIN_SIZE;
        this.speedY = Math.random() * (MAX_SPEED - MIN_SPEED) + MIN_SPEED;
        this.speedX = (Math.random() - 0.5) * 0.8; 
        this.opacity = 0;
        this.fadeState = 'in'; 
        this.fadeSpeed = Math.random() * 0.015 + 0.005;
        this.maxOpacity = Math.random() * 0.6 + 0.4; // Brighter
        // Color variation: mainly yellow/white, some slight blue/purple tint
        const colors = [
            `255, 255, 200`, // Warm yellow
            `255, 255, 255`, // White
            `200, 230, 255`  // Cool blueish
        ];
        this.color = colors[Math.floor(Math.random() * colors.length)];
    }

    update(w, h) {
        this.y -= this.speedY;
        this.x += this.speedX;

        // Opacity Animation
        if (this.fadeState === 'in') {
            this.opacity += this.fadeSpeed;
            if (this.opacity >= this.maxOpacity) {
                this.opacity = this.maxOpacity;
                this.fadeState = 'hold';
            }
        } else if (this.fadeState === 'hold') {
            if (Math.random() < 0.01) this.fadeState = 'out';
        } else if (this.fadeState === 'out') {
            this.opacity -= this.fadeSpeed;
            if (this.opacity <= 0) {
                this.opacity = 0;
                this.fadeState = 'in';
                if (this.y < -10) this.reset(w, h);
            }
        }

        if (this.y < -20) {
            this.reset(w, h);
        }
    }

    draw(ctx) {
        ctx.beginPath();
        ctx.arc(this.x, this.y, this.radius, 0, Math.PI * 2);
        
        // Stronger glow effect
        const gradient = ctx.createRadialGradient(this.x, this.y, 0, this.x, this.y, this.radius * 3);
        gradient.addColorStop(0, `rgba(${this.color}, ${this.opacity})`);
        gradient.addColorStop(0.4, `rgba(${this.color}, ${this.opacity * 0.5})`);
        gradient.addColorStop(1, `rgba(${this.color}, 0)`);
        
        ctx.fillStyle = gradient;
        ctx.fill();
    }
}

const initCanvas = () => {
    const el = canvas.value;
    if (!el) return;

    ctx = el.getContext('2d');
    resizeCanvas();
    
    particles = [];
    for (let i = 0; i < PARTICLE_COUNT; i++) {
        particles.push(new Particle(el.width, el.height));
    }

    animate();
};

const resizeCanvas = () => {
    if (!canvas.value) return;
    canvas.value.width = window.innerWidth;
    canvas.value.height = window.innerHeight;
};

const animate = () => {
    if (!ctx || !canvas.value) return;
    
    const w = canvas.value.width;
    const h = canvas.value.height;

    ctx.clearRect(0, 0, w, h);

    particles.forEach(p => {
        p.update(w, h);
        p.draw(ctx);
    });

    animationFrameId = requestAnimationFrame(animate);
};

onMounted(() => {
    updateBgImage();
    window.addEventListener('bg-image-changed', updateBgImage); // Listen for changes
    initCanvas();
    window.addEventListener('resize', resizeCanvas);
});

onBeforeUnmount(() => {
    window.removeEventListener('bg-image-changed', updateBgImage);
    window.removeEventListener('resize', resizeCanvas);
    if (animationFrameId) cancelAnimationFrame(animationFrameId);
});
</script>

<style scoped>
.background-effect {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    z-index: 0; /* Changed from -1 to 0 to ensure visibility if body bg is set */
    overflow: hidden;
    pointer-events: none;
}

.bg-image {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    /* Custom Starry Night Background Image (user configurable in /public/images/starry-bg.jpg) */
    background-color: #000; 
    /* Background Image is set via inline style now */
    background-position: center;
    background-size: cover;
    background-repeat: no-repeat;
    transition: background-image 0.5s ease;
}

.firefly-canvas {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 1;
}
</style>
