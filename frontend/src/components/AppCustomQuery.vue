<template>
  <div class="container mt-5">
    <h2>Custom Query</h2>
    <div class="form-group mt-3">
      <label for="sqlQuery">Enter your SQL Query:</label>
      <textarea class="form-control" id="sqlQuery" v-model="query" rows="4"></textarea>
    </div>

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
      <button class="btn btn-primary" @click="sendQuery">Execute</button>
      <button class="btn btn-secondary" @click="copyTableToClipboard">Copy Table</button>
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
        <tr v-for="row in result" :key="row.id">
          <td v-for="column in Object.keys(row)" :key="column">{{ row[column] }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      query: '',
      result: [],
      loading: false,
      errorMessage: ''
    };
  },
  methods: {
    async sendQuery() {
      this.errorMessage = '';  // Clear the error message
      this.result = [];  // Clear previous results
      this.loading = true;  // Start loading spinner

      try {
        const creds = this.$store.state.credentials;
        const dataToSend = {
          credentials: creds,
          query: this.query
        };
        const requestOptions = {
          method: 'POST',
          headers: {'Content-Type': 'application/json'},
          body: JSON.stringify(dataToSend)
        };

        const response = await fetch('http://localhost:8081/custom-query', requestOptions);
        const result = await response.json();

        // Handle error from the server
        if (result.error) {
          this.errorMessage = result.message || 'Unknown error occurred.';
          return;
        }

        this.result = result.data;
      } catch (error) {
        console.error("Error executing query:", error);
        this.errorMessage = "Failed to communicate with the server.";
      } finally {
        this.loading = false;  // Stop loading spinner
      }
    },
    async copyErrorToClipboard() {
      try {
        await navigator.clipboard.writeText(this.errorMessage);
        console.log("Error message copied to clipboard.");
      } catch (err) {
        console.error('Failed to copy error message: ', err);
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
      } catch (err) {
        console.error('Failed to copy table data: ', err);
      }
    }

  }
}
</script>

<style scoped>
/* Custom styles for the transparent blue spinner */
.spinner-grow.text-primary {
  background-color: rgba(0, 123, 255, 0.7);  /* Blue with 70% opacity */
}

.table-container {
  max-width: 100%; /* Or you can set this to a fixed pixel value if desired */
  overflow-x: auto; /* Enable horizontal scrolling */
  font-size: 0.8rem;
}
</style>
