<template>
    
</template>

<script>
export default {
  components: { RepositoriesFilters, RepositoriesSortBy, RepositoriesList },
  props: {
    user: { type: String }
  },
  data () {
    return {
      repositories: [], // 1
      filters: { ... }, // 3
      searchQuery: '' // 2
    }
    /**
     * ......
     * very long code
     * ......
     */
  },

  computed: {
    filteredRepositories () { ... }, // 3
    repositoriesMatchingSearchQuery () { ... }, // 2
    /**
     * ......
     * very long code
     * ......
     */
  },

  watch: {
     /**
     * ......
     * very long code
     * ......
     */
    user: 'getUserRepositories' // 1
  },

  methods: {
    /**
     * ......
     * very long code
     * ......
     */
    getUserRepositories () {
      // 使用 `this.user` 获取用户仓库
    }, // 1
    updateFilters () { ... }, // 3
  },

  mounted () {
    /**
     * ......
     * very long code
     * ......
     */
    this.getUserRepositories() // 1
  }
}
</script>

