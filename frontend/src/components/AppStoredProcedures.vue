<template>
  <div class="container mt-4">
    <div class="shadow p-4 bg-white rounded">
      <h1 class="mb-3">Database Stored Procedures</h1>
      <button @click="loadProcedures" class="btn btn-primary mb-3">Reload</button>
      <hr>

      <!-- Table to display the stored procedures -->
      <table class="table table-bordered table-striped table-hover">
        <thead class="thead-dark">
        <tr>
          <th>Schema Name</th>
          <th>Procedure Name</th>
          <th>Definition</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="procedure in procedures" :key="procedure.procedure_name">
          <td>{{ procedure.schema_name }}</td>
          <td>{{ procedure.procedure_name }}</td>
          <td class="definition-column">
            <pre>{{ procedure.definition }}</pre>
            <!-- Add Copy button here -->
            <button @click="copyToClipboard(procedure.definition)" class="btn btn-sm btn-secondary">Copy</button>
          </td>
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
      procedures: []
    };
  },
  beforeMount() {
    this.loadProcedures();
  },
  methods: {
    loadProcedures() {
      const creds = this.$store.state.credentials;
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(creds)
      };

      fetch('http://localhost:8081/stored-procedures', requestOptions)
          .then(response => response.json())
          .then(data => {
            this.procedures = data.data; // use the "data" property from the response
          })
          .catch(error => console.error('Error fetching stored procedure data:', error));
    },
    copyToClipboard(text) {
      navigator.clipboard.writeText(text).then(() => {
        notie.alert({type: 'success', text: "Stored procedure copied to clipboard.", time: 2});
      }).catch(err => {
        console.error('Failed to copy error message: ', err);
        notie.alert({type: 'error', text: 'Failed to copy.', time: 2});
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
