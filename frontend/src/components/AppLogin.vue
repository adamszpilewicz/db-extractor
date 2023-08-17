<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr>
        <form-tag v-on:myevent="submitHandler" name="myform" event="myevent">

          <text-input
              v-model="email"
              label="Email"
              type="email"
              name="email"
              required="true">
          </text-input>

          <text-input
              v-model="host"
              label="Host"
              type="text"
              name="host"
              required="true">
          </text-input>

          <text-input
              v-model="database"
              label="Database name"
              type="text"
              name="database"
              required="true">
          </text-input>

          <text-input
              v-model="user"
              label="User name"
              type="text"
              name="user"
              required="true">
          </text-input>

          <text-input
              v-model="password"
              label="Password"
              type="password"
              name="password"
              required="true">
          </text-input>

          <hr>
          <input type="submit" class="btn btn-primary" value="Login">
        </form-tag>
      </div>
    </div>
  </div>
</template>

<script>
import TextInput from './forms/TextInput.vue'
import FormTag from './forms/FormTag.vue'
import notie from 'notie'

export default {
  name: 'UserLogin',
  components: {
    TextInput,
    FormTag
  },
  data() {
    return {
      email: "",
      host: "",
      database: "",
      user: "",
      password: "",
    }
  },
  methods: {
    submitHandler() {
      console.log('submitHandler called - success!');

      const payload = {
        email: this.email,
        host: this.host,
        database: this.database,
        user: this.user,
        password: this.password,
      }
      console.log("Payload: ", JSON.stringify(payload));

      const requestOptions = {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(payload)
      };

      fetch('http://localhost:8081/users/login', requestOptions)
          .then(response => response.json())
          .then(data => {
            if (data.error) {
              console.log("Error: ", data.message);
              notie.alert({type: 'error', text: data.message, time: 2});
            } else {
              console.log("Success: ", data.message);
              notie.alert({type: 'success', text: data.message, time: 2});
              this.$store.commit('setCredentials', payload);
              // this.$router.push({name: 'Home'});
            }
          })
    }
  }
}
</script>