I will modify the `ArticleDetail.vue` page to modernize the Table of Contents (TOC) display and ensure it appears on the right side as requested.

### Changes
1.  **Restructure Sidebar**:
    *   Remove the global `sticky` positioning from the entire sidebar container to allow individual components to control their stickiness.
    *   Make the **TOC (Directory)** card sticky so it stays visible on the right while reading the article.

2.  **Modernize TOC Styling**:
    *   Transform the TOC into a clean, modern "Navigation Rail" style.
    *   Remove the heavy card background (or make it more subtle/glass-like).
    *   Use a vertical line design with active state highlighting (color change + border indicator) for better readability.
    *   Improve the interaction (hover effects, smooth scrolling).

3.  **Prevent "Left" Display**:
    *   Add CSS to explicitly hide any default TOC that might be generated within the markdown content area (e.g., if `[TOC]` is used in the article), ensuring the only TOC is the modern one on the right.

### File to Modify
*   `frontend/src/views/ArticleDetail.vue`

### Verification
*   Verify the TOC is positioned on the right column.
*   Verify the TOC sticks to the top of the viewport when scrolling.
*   Verify the new modern styling is applied.
