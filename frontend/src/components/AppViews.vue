<template>
  <div class="container mt-4">
    <div class="shadow p-4 bg-white rounded">
      <h1 class="mb-3">Database Views</h1>
      <button @click="loadViews" class="btn btn-primary mb-3">Reload</button>
      <hr>

      <!-- Table to display the views -->
      <table class="table table-bordered table-striped table-hover">
        <thead class="thead-dark">
        <tr>
          <th>View Name</th>
          <th>Owner</th>
          <th>Definition</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="view in views" :key="view.view_name">
          <td>{{ view.view_name }}</td>
          <td>{{ view.owner }}</td>
          <td class="definition-column">
            <pre>{{ view.definition }}</pre>
            <!-- Add Copy button here -->
            <button @click="copyToClipboard(view.definition)" class="btn btn-sm btn-secondary">Copy</button>
          </td>
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
      views: []
    };
  },
  beforeMount() {
    this.loadViews();
  },
  methods: {
    loadViews() {
      const creds = this.$store.state.credentials;
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(creds)
      };

      fetch('http://localhost:8081/views', requestOptions)
          .then(response => response.json())
          .then(data => {
            this.views = data.data; // use the "data" property from the response
          })
          .catch(error => console.error('Error fetching view data:', error));
    },
    copyToClipboard(text) {
      navigator.clipboard.writeText(text).then(() => {
        alert('Copied to clipboard!');
      }).catch(err => {
        console.error('Failed to copy!', err);
      });
    }
  }
}
</script>

<style scoped>
.definition-column pre {
  white-space: pre-wrap;
  max-height: 100px;
  overflow-y: auto;
}
</style>
