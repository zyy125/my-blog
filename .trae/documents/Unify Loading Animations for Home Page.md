I will unify the loading animations for both the Article List and Sidebar components on the Home page, implementing a modern "scale-up & fade-in" effect as requested.

### Animation Specification
*   **Style**: "Scale-up & Fade-in" (从小到大逐渐放大 + 渐显).
*   **Keyframes**:
    *   Start: `opacity: 0`, `transform: scale(0.95) translateY(20px)`
    *   End: `opacity: 1`, `transform: scale(1) translateY(0)`
*   **Timing**: `0.8s cubic-bezier(0.34, 1.56, 0.64, 1)` (Soft spring-like effect) or `ease-out`.
*   **Staggering**: Items should appear one after another with a slight delay.

### Implementation Steps

1.  **Define Global Animation Class**:
    *   In `frontend/src/style.css`, define a reusable animation class `.anim-scale-in` and the `@keyframes scaleIn`.

2.  **Apply to Article List (`Home.vue`)**:
    *   Update the `ArticleCard` loop to use this new class.
    *   Implement staggered delays using `style="animation-delay: ${index * 0.1}s"`.
    *   Remove the old `IntersectionObserver` logic if it conflicts or is redundant (the user requested a unified *loading* animation, usually on mount). *Correction*: The user mentioned "loading animation", implying the initial render. I will replace the complex scroll-reveal with this unified mount animation for simplicity and consistency, or adapt the scroll reveal to use this new style. Given the request for "unified", I will prioritize the initial load effect first.

3.  **Apply to Sidebar (`Home.vue`)**:
    *   Update the sidebar cards to use the same `.anim-scale-in` class.
    *   Apply staggered delays for sidebar items (Profile -> Latest -> SiteInfo -> Categories -> Tags).

### File Changes
*   `frontend/src/style.css`: Add global keyframes and utility classes.
*   `frontend/src/views/Home.vue`:
    *   Remove old specific animation styles (`.scroll-reveal`, `.sidebar-card` animation).
    *   Apply the new global classes to `ArticleCard` and Sidebar components.
    *   Add inline styles for `animation-delay`.

### Verification
*   Check that both left (articles) and right (sidebar) areas animate in with the same "scale up" style.
*   Verify the timing is slow and soft as requested.
