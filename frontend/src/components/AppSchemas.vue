<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5 mb-3">Schemas</h1>
        <button @click="loadSchemas" class="btn btn-primary mb-3">Reload</button>
        <hr>

        <!-- Loading Spinner -->
        <div v-if="loading" class="d-flex justify-content-center mb-3">
          <div class="spinner-border" role="status"></div>
        </div>

        <!-- Table to display the schemas -->
        <table v-if="!loading" class="table table-bordered table-striped table-hover">
          <thead class="thead-dark">
          <tr>
            <th>Schema Name</th>
            <th>Table</th>
          </tr>
          </thead>
          <tbody>
          <template v-for="schema in schemas.data">
            <tr v-for="(table, index) in schema.table_names" :key="schema.schema_name + index">
              <td v-if="index === 0" :rowspan="schema.table_names.length">{{ schema.schema_name }}</td>
              <td>
                <router-link :to="{ name: 'TableInfo', params: { tableName: table, schemaName: schema.schema_name } }">{{ table }}</router-link>
              </td>
            </tr>
          </template>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>


<script>
export default {
  data() {
    return {
      schemas: [],
      loading: false
    }
  },
  beforeMount() {
    this.loadSchemas(); // call the method on mount
  },
  methods: {
    loadSchemas() {
      this.schemas = []; // Empty the schemas array
      this.loading = true; // Set loading to true at the start
      const creds = this.$store.state.credentials;
      const requestOptions = {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(creds)
      };

      fetch('http://localhost:8081/schemas', requestOptions)
          .then(response => response.json())
          .then(response => {
            this.schemas = response;
            this.loading = false; // Set loading to false once data is fetched
          })
          .catch(error => {
            console.log(error);
            this.loading = false; // Also set loading to false in case of error
          })
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

.table tbody td {
  font-size: 14px;  /* Adjust the value as needed */
}
</style>
