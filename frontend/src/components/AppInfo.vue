<template>
  <div class="container mt-4">
    <div class="shadow p-4 bg-white rounded">
      <h1 class="mb-3">Database Information</h1>
      <button @click="loadDBInfo" class="btn btn-primary mb-3">Reload</button>
      <hr>

      <!-- Card to display the database info -->
      <div class="card mb-3" v-for="(value, key) in dbInfo" :key="key">
        <div class="card-header custom-header">
          <strong>{{ key }}</strong>
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
      dbInfo: {}
    };
  },
  beforeMount() {
    this.loadDBInfo();
  },
  methods: {
    loadDBInfo() {
      const creds = this.$store.state.credentials;
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(creds)
      };

      fetch('http://localhost:8081/dbinfo', requestOptions)
          .then(response => response.json())
          .then(data => {
            this.dbInfo = data.data; // use the "data" property from the response
          })
          .catch(error => console.error('Error fetching database info:', error));
    },
    getDisplayValue(value) {
      if (value && value.Valid) {
        return value.String;
      }
      return value;
    }
  }
}
</script>

<style scoped>
.card {
  margin-bottom: 10px;
}

.card-header.custom-header {
  font-size: 0.9em;
}

.list-group-item {
  font-size: 0.8em;
}
</style>
