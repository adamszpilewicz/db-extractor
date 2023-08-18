<template>
  <div class="container mt-5">
    <h2>Fetch PostgreSQL Settings</h2>

    <!-- Error Message -->
    <div v-if="errorMessage" class="mt-3 alert alert-danger d-flex justify-content-between align-items-center">
      <span>{{ errorMessage }}</span>
      <button class="btn btn-sm btn-secondary" @click="copyErrorToClipboard">Copy</button>
    </div>

    <!-- Loading Spinner -->
    <div v-if="loading" class="mt-5 d-flex justify-content-center">
      <div class="spinner-grow text-primary" role="status" style="opacity: 0.7;"></div>
    </div>

    <div class="d-flex justify-content-between mt-3">
      <button class="btn btn-primary mt-3" @click="fetchStats">Fetch Settings</button>
      <button class="btn btn-secondary mt-3" @click="copyTableToClipboard">Copy Table</button>
    </div>

    <!-- Result Table -->
    <div v-if="result && result.length" class="mt-5 table-container">
      <table class="table table-striped table-bordered">
        <thead class="thead-dark">
        <tr>
          <th v-for="column in Object.keys(result[0])" :key="column">{{ column }}</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="row in result" :key="row.name">  <!-- Changed the key to row.name -->
          <td v-for="column in Object.keys(row)" :key="column">{{ row[column] }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import notie from 'notie';

export default {
  data() {
    return {
      result: [],
      loading: false,
      errorMessage: ''
    };
  },
  mounted() {
    this.fetchStats();
  },
  methods: {
    async fetchStats() {
      this.errorMessage = '';
      this.result = [];
      this.loading = true;

      try {
        const creds = this.$store.state.credentials;
        const requestOptions = {
          method: 'POST',
          headers: {'Content-Type': 'application/json'},
          body: JSON.stringify(creds)
        };

        const response = await fetch('http://localhost:8081/pg-settings', requestOptions); // Modified the URL to 'pg-settings'
        const result = await response.json();

        if (result.error) {
          this.errorMessage = result.message || 'Unknown error occurred.';
          return;
        }

        this.result = result.data;
      } catch (error) {
        console.error("Error fetching stats:", error);
        this.errorMessage = "Failed to communicate with the server.";
      } finally {
        this.loading = false;
      }
    },
    async copyErrorToClipboard() {
      try {
        await navigator.clipboard.writeText(this.errorMessage);
        console.log("Error message copied to clipboard.");
        notie.alert({position:'bottom',type: 'success', text: "Error message copied to clipboard.", time: 2});
      } catch (err) {
        console.error('Failed to copy error message: ', err);
        notie.alert({position:'bottom', type: 'error', text: 'Failed to copy error message.', time: 2});
      }
    },
    async copyTableToClipboard() {
      try {
        // Convert table data to CSV
        const header = Object.keys(this.result[0]).join(",");
        const rows = this.result.map(row => Object.values(row).join(",")).join("\n");
        const csvData = `${header}\n${rows}`;

        // Copy to clipboard
        await navigator.clipboard.writeText(csvData);
        console.log("Table data copied to clipboard.");
        notie.alert({position:'bottom', type: 'success', text: "Table data copied to clipboard.", time: 2});
      } catch (err) {
        console.error('Failed to copy table data: ', err);
        notie.alert({position:'bottom', type: 'error', text: 'Failed to copy table data.', time: 2});
      }
    }
  }
}
</script>

<style scoped>
/* Custom styles for the transparent blue spinner */
.spinner-grow.text-primary {
  background-color: rgba(0, 123, 255, 0.7); /* Blue with 70% opacity */
}

.table-container {
  max-width: 100%; /* Or you can set this to a fixed pixel value if desired */
  overflow-x: auto; /* Enable horizontal scrolling */
  font-size: 0.8rem;
}
</style>
