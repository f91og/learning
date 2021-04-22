<template>
    
</template>

<script>
import { fetchUserRepositories } from '@/api/repositories'
import { ref, onMounted, watch, computed } from 'vue'

export default {
  components: { RepositoriesFilters, RepositoriesSortBy, RepositoriesList },
  props: {
    user: { type: String }
  },
  setup(props) {
    // 使用 `toRefs` 创建对prop的 `user` property 的响应式引用
    const { user } = toRefs(props)

    const repositories = ref([])
    const getUserRepositories = async () => {
      repositories = await fetchUserRepositories(props.user)
    }

    onMounted(getUserRepositories)

    // 在用户 prop 的响应式引用上设置一个侦听器
    watch(user, getUserRepositories)

    const searchQuery = ref('')
    const repositoriesMatchingSearchQuery = computed(() => {
      return repositories.value.filter(
        repository => repository.name.includes(searchQuery.value)
      )
    })

    return {
      repositories,
      getUserRepositories,
      searchQuery,
      repositoriesMatchingSearchQuery
    } // 这里返回的任何内容都可以用于组件的其余部分
  },
  data () {
    return {
      filters: { ... }, // 3
    }
  },
  computed: {
    filteredRepositories () { ... }, // 3
  },
  methods: {
    updateFilters () { ... }, // 3
  },
}
</script>

