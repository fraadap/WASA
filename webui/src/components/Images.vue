<script>
export default {
props:["mine","profile"],
    data: function () {
        return {
            
        };
    },
    methods: {
    }
}
</script>

<template>




<div class="container">
  <div class="row">
        <!-- Photo -->

        <!-- Photo -->
    <div class="col-md-4 mb-3" v-for="ph in profile.photos" :key="ph.photo.photoID">
      <div class="custom-square">
        <div class="custom-content">
          <div class="image-wrapper">
            <img :src="'data:image/*;base64,' + ph.photo.binary" alt="Image" class="custom-image">
          </div>
        </div>
      </div>
      <div style="background-color: #4871c3; height:40px;" class="d-flex align-items-center p-3">
        <div style="color:white;">
          {{ ph.nLikes }}
          <svg style="width:15%;" fill="none" stroke="white"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg>
          {{ ph.nComments }}
          <svg style=" width:15%;" fill="none" stroke="white"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
        </div>
        <div></div>
        <div class="d-flex flex-row-reverse" >
          <svg style=" width:20%;" fill="none" stroke="white"><use href="/feather-sprite-v4.29.0.svg#trash-2" v-if="mine"/></svg>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<style>
  .custom-square {
    padding-top: calc(100%);
    position: relative;
  }

  .custom-content {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    padding: 0;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .image-wrapper {
    width: 100%;
    height: 100%;
    overflow: hidden;
  }

  .custom-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
</style>