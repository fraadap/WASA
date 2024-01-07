  
  <script>
  export default {
    name: "CustomAlert",
    props: {
      msg: {
        type: String,
        required: true,
      },
    },
  };
  </script>

  <template>
    <div class="alert alert-success alert-dismissible fade show" role="alert">
      {{ msg }}
      <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
    </div>
  </template>

  <style scoped>
  .alert {
    position: fixed;
    left: 0;
    right: 0;
    border-radius: 0;
    text-align: center;
  }
  </style>
  
  