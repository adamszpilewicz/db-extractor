<template>
  <div class="container mt-4">
    <div class="shadow p-4 bg-white rounded">
      <h1 class="mb-3">Database Statistics</h1>
      <button @click="loadDBStats" class="btn btn-primary mb-3">Reload</button>
      <hr>

      <!-- Card to display the database statistics -->
      <div class="card mb-3" v-for="(value, key) in dbStats" :key="key">
        <div class="card-header custom-header">
          <strong>{{ formatKey(key) }}</strong>
        </div>
        <ul class="list-group list-group-flush">
          <li class="list-group-item">{{ getDisplayValue(value) }}</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      dbStats: {}
    };
  },
  beforeMount() {
    this.loadDBStats();
  },
  methods: {
    loadDBStats() {
      const creds = this.$store.state.credentials;
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(creds)
      };

      fetch('http://localhost:8081/db-stats', requestOptions)
          .then(response => response.json())
          .then(data => {
            this.dbStats = data.data;
          })
          .catch(error => console.error('Error fetching database statistics:', error));
    },
    getDisplayValue(value) {
      if (value && value.Valid) {
        return value.String;
      }
      return value;
    },
    formatKey(key) {
      return key.replace(/([A-Z])/g, ' $1') // insert a space before all capital letters
          .replace(/^./, str => str.toUpperCase()); // capitalize the first letter
    }
  }
}
</script>

<style scoped>
.card {
  margin-bottom: 10px;
}
</style>
