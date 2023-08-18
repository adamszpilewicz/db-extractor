<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5 mb-3">Schemas</h1> <!-- Adjust margin for better spacing -->
        <button @click="loadSchemas" class="btn btn-primary mb-3">Reload</button> <!-- The reload button -->
        <hr>

        <!-- Table to display the schemas -->
        <table class="table table-bordered table-striped table-hover"> <!-- Added table-striped and table-hover -->
          <thead class="thead-dark"> <!-- Made the header dark for better contrast -->
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
      schemas: []
    }
  },
  beforeMount() {
    this.loadSchemas(); // call the method on mount
  },
  methods: {
    loadSchemas() {
      const creds = this.$store.state.credentials;
      const requestOptions = {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(creds)
      };

      console.log("Request options: ", requestOptions);

      fetch('http://localhost:8081/schemas', requestOptions)
          .then(response => response.json())
          .then(response => {
            this.schemas = response
          })
          .catch(error => console.log(error))
    }
  }
}

</script>
