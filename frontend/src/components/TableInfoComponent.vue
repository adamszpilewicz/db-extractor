<template>
  <div class="container mt-4">
    <div class="shadow p-4 bg-white rounded">
      <!-- Download button placed above the table -->
      <div class="mb-3 text-right">
        <button @click="downloadCSV" class="btn btn-primary">Download CSV</button>
      </div>

      <table class="table table-hover table-striped">
        <thead class="thead-dark">
        <tr>
          <th>Table Name</th>
          <th>Column Name</th>
          <th>Column Type</th>
          <th>Primary Key</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="column in columns" :key="column.column_name">
          <td>{{ tableName }}</td>
          <td>{{ column.column_name }}</td>
          <td>{{ column.column_type }}</td>
          <td>
            <span v-if="column.primary_key === 'YES'">
              YES
              <img src="@/assets/key.png" alt="Primary Key" style="height: 16px; width: 32px; vertical-align: middle;">
            </span>
            <span v-else>NO</span>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>


<script>
export default {
  props: ['tableName'],
  data() {
    return {
      columns: [] // for storing columns info
    };
  },
  created() {
    this.loadColumns();
  },
  methods: {
    loadColumns() {
      const creds = this.$store.state.credentials;

      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(creds)
      };

      fetch(`http://localhost:8081/schemas/${this.tableName}`, requestOptions)
          .then(response => {
            if (!response.ok) {
              return response.text().then(text => {
                throw new Error(`Server error: ${text}`);
              });
            }
            return response.json();
          })
          .then(data => {
            if (!data.error && data.columns) {
              this.columns = data.columns;
            } else {
              console.error(data.message);
            }
          })
          .catch(error => console.error('Error fetching column data:', error));
    },

    downloadCSV() {
      const creds = this.$store.state.credentials;
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(creds)
      };

      fetch(`http://localhost:8081/download/${this.tableName}`, requestOptions)
          .then(response => response.blob())
          .then(blob => {
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement('a');
            a.href = url;
            a.download = `${this.tableName}.csv`;
            a.click();
          })
          .catch(error => console.error('Error downloading CSV:', error));
    }

  }
}
</script>

