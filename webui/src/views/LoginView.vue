<script>


export default {
  data: function () {
    return {
      errormsg: null,
      loading: false,
      some_data: null,
      username: '',
    };
  },
  methods: {
    async login() {
      try {
        let response = await this.$axios.post('/session', {
          username: this.username
        });
        localStorage.setItem("token", response.data)
        this.$router.push({ path: '/home' });

      }
      catch (e) {
        alert("Error: " + e);
      }
    },
  },
}
</script>

<template>
  <div class="container d-flex justify-content-between mt-5">
    <!-- Left Box -->
    <div class="card p-4" style="width: 45%;">
      <h1 class="display-4">WASA Photo</h1>
      <p class="lead">Let your friends know you</p>
    </div>

    <!-- Right Box -->
    <div class="card p-4" style="width: 45%;">
      <h2 class="mb-2">Login</h2>
      <p class="lead">Insert your username:</p>
      <div class="mb-3">
        <input type="text" class="form-control" id="username-input" placeholder="Username" v-model="username">
      </div>

        <button type="button" class="btn btn-success" id="login-button" @click="this.login">Login</button>

    </div>
  </div>
</template>