<template>
  <div class="container mt-4">
    <div class="shadow p-4 bg-white rounded">
      <h1 class="mb-3">Database Statistics</h1>
      <div class="mb-3 d-flex justify-content-between align-items-center">
        <button @click="loadDBStats" class="btn btn-primary">Reload</button>
        <button @click="copyToClipboard" class="btn btn-secondary">Copy</button>
      </div>
      <hr>

      <!-- Table to display the database statistics -->
      <table class="table table-bordered table-hover table-striped">
        <thead>
        <tr>
          <th>Statistic Name</th>
          <th>Value</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(value, key) in dbStats" :key="key">
          <td>{{ formatKey(key) }}</td>
          <td>{{ getDisplayValue(value) }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import notie from 'notie'

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
    },
    async copyToClipboard() {
      let tableString = 'Statistic Name\tValue\n';  // Headers with a tab in between
      for (let key in this.dbStats) {
        tableString += `${this.formatKey(key)}\t${this.getDisplayValue(this.dbStats[key])}\n`;
      }

      try {
        await navigator.clipboard.writeText(tableString);
        notie.alert({ type: 'success', text: 'Table data copied to clipboard!', time: 3 });
      } catch (err) {
        console.error('Failed to copy table data: ', err);
        notie.alert({ type: 'error', text: 'Failed to copy table data!', time: 3 });
      }
    }
  }
}
</script>

<style scoped>
/* Add padding to cells for better readability */
.table > tbody > tr > td,
.table > thead > tr > th {
  padding: 15px;
}

/* Highlight row background on hover */
.table-hover > tbody > tr:hover {
  background-color: #f5f5f5;
}

/* Optional: Give a border-bottom to each row for clear separation */
.table > tbody > tr {
  border-bottom: 1px solid #dee2e6;
}
</style>