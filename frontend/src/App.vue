<template>
  <div id="app">
    <AppHeader />
    <div class="content-wrapper">
      <router-view />
    </div>
    <AppFooter />
  </div>
</template>

<script>
import AppHeader from './components/AppHeader.vue'
import AppFooter from './components/AppFooter.vue'

export default {
  name: 'App',
  components: {
    AppHeader,
    AppFooter
  }
}
</script>

<style>
#app {
  display: flex;
  flex-direction: column;
  min-height: 100vh; /* ensure it takes up at least the full height of the viewport */
}

.content-wrapper {
  flex: 1; /* make the content-wrapper take up all available space */
}
</style>
