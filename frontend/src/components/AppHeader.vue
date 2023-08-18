<template>
  <nav class="navbar fixed-top navbar-expand-lg bg-body-tertiary navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Navbar</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav">
          <li class="nav-item">
            <router-link class="nav-link active" aria-current="page" to="/" >Home</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/login">Login</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/schemas">Schemas</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/custom-query">Custom Query</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/views">Views</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/stored-procedures">Stored Procedures</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/info">Info</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/db-stats">DB Stats</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/db-slots">DB Slots</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/pg-settings">Pg Settings</router-link>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>
<script setup>
</script>