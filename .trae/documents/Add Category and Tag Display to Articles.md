I will implement the display of article categories and tags in both the article list (Home) and article detail pages to ensure this information is visible at a glance, followinackend Changesg a modern design.

### B

1. **Update Article Service**: Modify `backend/internal/service/article.go` to use `ListWithAssociations` instead of `List`. This ensures that `Category` and `Tags` data are preloaded and returned in the API response for the article list.

### Frontend Changes

1. **Update Article Card (`frontend/src/components/ArticleCard.vue`)**:

   * Update the `article-header` section to iterate through `article.tags` and display them as badges/pills.

   * Style the tags to be consistent with the modern glassmorphism design (e.g., using `el-tag` with appropriate types and effects).

   * Ensure the Category badge is visible (it relies on the backend fix).

2. **Update Article Detail (`frontend/src/views/ArticleDetail.vue`)**:

   * Update the `meta-tags` section in the hero header to display the list of tags alongside the category and date.

   * Apply consistent styling for the tags.

### Verification

* Verify that the Home page article cards now show both Category and Tags.

* Verify that the Article Detail page header shows both Category and Tags.

* Ensure the design is responsive and aesthetically pleasing ("modern").

