<script>
import { RouterLink } from 'vue-router';
import  Images from '../components/Images.vue';

export default {
    props:['userID'],
    component:Images,
    data: function () {
        return {
            errormsg: null,
            loading: false,
            profile: {},
            myProfile: {},
            us: this.userID,
            myID: localStorage.getItem("token"),
            mine: this.myID==this.us,
            modalComments: {},
            textComment:"",
        };
    },
    methods: {
        async refresh() {
			    this.loading = true;
			    this.errormsg = null;
			    try {
		   	  	let response = await this.$axios.get("/profile");
			    	this.some_data = response.data;
			    } catch (e) {
		  	  	this.errormsg = e.toString();
			    } 
			    this.loading = false;
		    },
        async load() {
            try {
                let response =await this.$axios.get('/users/'+this.myID,
                {headers: {Authorization: this.myID}});
                this.myProfile = response.data;

                if( this.us == 0 || this.us == this.myID){
                  this.mine = true
                  this.us = this.myID
                  this.profile = this.myProfile
                }
                else{
                  let response =await this.$axios.get('/users/'+this.us,
                  {headers: {Authorization: this.myID}});
                  this.profile = response.data;
                }
                console.log(this.profile)

            }
            catch (e) {
                alert("Error: " + e);
            }
            
        },
        async uploadPhoto(){
          this.image = this.$refs.file.files[0];
        },
        async submitPhoto(){
          if (this.images === null) {
				    alert("No photo selected");
			    } else {
				    try {
					    let response = await this.$axios.post("/users/" + this.myID + "/photos", this.image, {
						    headers: {
							    Authorization: this.myID
						    }
					    })
              this.binary = response.data.binary
            }catch(e){
              alert(e)
            }
          }
          this.load();
        },
        isFollowed(id) {
          let yes = false;
          let f;
          for (f in this.myProfile.followings){
            if (f.followed == id){
              yes= true
            }
          }
          return yes;
        },
        isBanned(id) {  //implementare ban e unban
          let yes = false;
          let b;
          for (b in this.myProfile.bans){
            if (b.banned == id){
              yes= true
            }
          }
          return yes;
        },
        async unfollow(id){
          try {
              let response =await this.$axios.get('/users/'+this.myID+"/followID/"+id,
              {
              headers: {
							    Authorization: this.myID
						  }
              });
              let followID = response.data.followID
              
              let res =await this.$axios.delete('/users/'+this.myID+"/follow/"+followID,
              {
              headers: {
							    Authorization: this.myID
						  }
              });
              
              this.load();
          }
          catch (e) {
            alert("Error: " + e);
          }
        },
        async follow(id){
          try {
              await this.$axios.post('/users/'+this.myID+"/follow",{
                followed:id,
              },
              {
              headers: {
							    Authorization: this.myID
						  }
              });
          }
          catch(e){
            alert(e);
          }
          this.load()
        },
        async ban(id){
          try {
              let response = await this.$axios.post('/users/'+this.myID+"/bans",{
                banned:id,
              },
              {
              headers: {
							    Authorization: this.myID
						  }
              });
          }
          catch(e){
            alert(e);
          }
          this.load()
        },
        async unban(id){
          try {
              let response =await this.$axios.get('/users/'+this.myID+"/banID/"+id,
              {
              headers: {
							    Authorization: this.myID
						  }
              });
              let banID = response.data.banID
              let res =await this.$axios.delete('/users/'+this.myID+"/bans/"+banID,
              {
              headers: {
							    Authorization: this.myID
						  }
              });
              
              this.load();
          }
          catch (e) {
            alert("Error: " + e);
          }
        },
        getLikeID(ph) {
          let like = 0;
          if (ph.likes == null){return 0}
          for (let i=0; i<ph.likes.length; i++){
            if (ph.likes[i].userID == this.myID){
              like = ph.likes[i].likeID
            } 
          }
          return like;
        },
        async like(ph){
          let l = this.getLikeID(ph)
          if(l!=0){
            try {
              await this.$axios.delete('/users/'+ph.photo.userID+"/photos/"+ph.photo.photoID+"/likes/"+l,
              {
              headers: {
							    Authorization: this.myID
						  }
              });
            }
            catch(e){
              alert(e);
            }
          }
          else{

          try {
              await this.$axios.post('/users/'+ph.photo.userID+"/photos/"+ph.photo.photoID+"/likes",{ 
                userID : parseInt(this.myID),
                photoID: ph.photo.photoID,
              },
              {
              headers: {
							    Authorization: this.myID
						  }
              });
            }
          catch(e){
            alert(e);
          }
        }
          this.load()
        },
        async deletePhoto(ph){
          try {
              await this.$axios.delete('/users/'+ph.photo.userID+"/photos/"+ph.photo.photoID,
              {
              headers: {
							    Authorization: this.myID
						  }
              });
              this.load()
            }
            catch(e){
              alert(e);
            }
            
        },
        getReadableDate(d){
          const dateObject = new Date(d);

          const year = dateObject.getFullYear();
          const month = dateObject.getMonth() + 1; // I mesi in JavaScript vanno da 0 a 11
          const day = dateObject.getDate();
          const hours = dateObject.getHours();
          const minutes = dateObject.getMinutes();
          const seconds = dateObject.getSeconds();

          const formattedDate = day+"-"+month+"-"+year+" | "+ hours+":"+minutes;
          return formattedDate
        },
        async comment(){
          let txt = this.textComment;
          let ph = this.modalComments;
          try{
            let response=await this.$axios.post('/users/'+ph.photo.userID+"/photos/"+ph.photo.photoID+"/comments",{ 
                text: txt,
                user:{userID : parseInt(this.myID)},
              },
              {
              headers: {
							    Authorization: this.myID
						  }
              });
              this.load();
              this.textComment="";
              console.log(response.data);
              this.modalComments.comments.push(response.data);
          }
          catch(e){
            alert(e);
          }
          
        },
        async deleteComment(c){
          let ph = this.modalComments;
          try {
              await this.$axios.delete('/users/'+ph.photo.userID+"/photos/"+ph.photo.photoID+"/comments/"+c.commentID,
              {
              headers: {
							    Authorization: this.myID
						  }
              });
              this.load();
              this.modalComments.comments = this.modalComments.comments.filter( com => com.commentID !== c.commentID );
            }
            catch(e){
              alert(e);
            }
        },
      },
    mounted() {
        this.load();
    }
}
</script>

<template>

<!--
<img :src="'data:image/*;base64,' + this.binary">  Modo di caricare le foto 
-->
{{ this.us }}
<div class="container mt-5 text-center">
  <div class="row">
    <div class="col fs-3">
      <p>{{ this.profile && this.profile.user && this.profile.user.username }}</p>
    </div>
    <div class="col fs-5">
      <p>Photos</p>
      <p>{{ this.profile.photos ? this.profile.nPhotos : 0  }}</p>
    </div>
    <div class="col fs-5" data-toggle="modal" data-target="#followersModal" style="cursor: pointer;">
      <p>Followers</p>
      <p>{{ this.profile.followers ? this.profile.followers.length : 0  }}</p>
    </div>
    <div class="col fs-5" data-toggle="modal" data-target="#followingsModal" style="cursor: pointer;">
      <p>Followings</p>
      <p>{{ this.profile.followings ? this.profile.followings.length : 0  }}</p>
    </div>
  </div>

  <div class="position-absolute bottom-0 start-50 translate-middle-x" v-if="this.mine"> <!-- Nuova foto -->
      <button class="btn btn-success fs-3"  data-toggle="modal" data-target="#addPhotoModal">+</button>
  </div>
</div>

<!--Sezione Foto -->
<div class="border-top mt-5 mb-5"></div>


<div class="container">
  <div class="row">
        <!-- Photo -->

        <!-- List of photos -->
    <div class="col-md-4 mb-3" v-for="ph in profile.photos">
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
          <svg style="width:30px; cursor:pointer" :fill="ph.likes == null || this.getLikeID(ph)==0 ? 'none' : 'red'" stroke="white" @click="like(ph)"><use href="/feather-sprite-v4.29.0.svg#heart" /></svg>
          {{ ph.nComments }}
          <svg style="width:30px; cursor:pointer" fill="none" stroke="white" @click="this.modalComments = ph" data-toggle="modal" data-target="#commentsModal"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
        </div>
      <div class="d-flex flex-row-reverse" style="margin-left:55%" >
        <svg style=" width:30px; cursor:pointer" fill="none" stroke="white" @click="deletePhoto(ph)" v-if="mine"><use href="/feather-sprite-v4.29.0.svg#trash-2" /></svg>
        </div>
      </div>
    </div>
  </div>
</div>


<!-- Upload photo modal -->
<div class="modal fade" id="addPhotoModal" tabindex="-1" role="dialog" aria-labelledby="addPhotoModalLabel" aria-hidden="true">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="addPhotoModalLabel">Add a new photo</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">
          <label for="photoInput">Load photo:</label>
          <input type="file" accept="image/*" id="photoInput" class="form-control" @change="uploadPhoto" ref="file">
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-dismiss="modal">Cancel</button>
        <button type="button" class="btn btn-primary" data-dismiss="modal" @click="submitPhoto">Post Photo</button>
      </div>
    </div>
  </div>
</div>


<!-- List of followers
curl -X POST -H "Content-Type: application/json" -H "Authorization:4" -d '{"followed":2}' http://localhost:3000/users/4/follow -i
-->

<div class="modal fade" id="followersModal" tabindex="-1" role="dialog" aria-labelledby="followersModalLabel" aria-hidden="true">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="followersModalLabel">Followers</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">
        <div class="container mt-5">
          <div class="row">
            <div class="col">


      <div class="card mb-3" v-for="user in this.profile.followers" v-if="this.profile.followers!=null">
        <div class="card-body d-flex align-items-center justify-content-between">
          <div class="me-5 p-2 bd-highlight">
            <h5 class="card-title">{{ user.username }}</h5>
          </div>
          <div class="p-2 bd-highlight">
            <button type="button" class="btn btn-outline-success me-3" @click="follow(user.userID)" v-if="isFollowed(user.id)==false">Follow</button>
            <button v-else type="button" class="btn btn-primary me-3" @click="unfollow(user.userID)"> Followed </button>
            <button type="button" class="btn btn-outline-danger" v-if="isBanned(user.id)==false" @click="ban(user.userID)">Ban</button>
            <button type="button" class="btn btn-outline-danger" v-else @click="unban(user.userID)">Banned</button>
          </div>
        </div>
      </div>
      <div v-else class="d-flex align-items-center"><p>No followers yet</p></div>
    </div>
  </div>
</div>
</div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
      </div>
    </div>
  </div>
</div>

<!-- List of followings modal -->
<div class="modal fade" id="followingsModal" tabindex="-1" role="dialog" aria-labelledby="followingsModalLabel" aria-hidden="true">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="followingsModalLabel">Followings</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">
        <div class="container mt-5">
          <div class="row">
            <div class="col">


      <div class="card mb-3" v-for="user in this.profile.followings" v-if="this.profile.followings!=null">
        <div class="card-body d-flex align-items-center justify-content-between">
          <div class="me-5 p-2 bd-highlight">
            <h5 class="card-title">{{ user.username }}</h5>
          </div>
          <div class="p-2 bd-highlight">
            <button type="button" class="btn btn-primary me-3" @click="unfollow(user.userID)"> Unfollow </button>
            <button type="button" class="btn btn-outline-danger" v-if="isBanned(user.id)==false" @click="ban(user.userID)">Ban</button>
            <button type="button" class="btn btn-outline-danger" v-else @click="unban(user.userID)">Banned</button>
          </div>
        </div>
      </div>
      <div v-else class="d-flex align-items-center"><p>No followings yet</p></div>
    </div>
  </div>
</div>
</div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
      </div>
    </div>
  </div>
</div>


<!-- Comments modal 

<div class="modal fade" id="commentsModal" tabindex="-1" role="dialog" aria-labelledby="commentsModalLabel" aria-hidden="true">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="followingsModalLabel">Comments</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">
        <div class="container mt-5">
          <div class="row">
            <div class="col">

              <div class="comment-container">
  <div class="comment">
    <div class="username">Username</div>
    <div class="text">Testo del commento lungo le righe. Testo del commento lungo le righe. Testo del commento lungo le righe.</div>
    <button type="button" class="btn btn-outline-danger delete-btn">Elimina commento</button>
  </div>

</div>

      <div class="card mb-3" v-for="c in modalComments.comments" v-if="modalComments.comments!=null">

        <div class="card-body d-flex align-items-center justify-content-between">
          <div class="me-5 p-2 bd-highlight">
            <p class="card-title">{{ c.username }}</p>
          </div>
          <div>Testo del</div>
          <div class="p-2 bd-highlight">
            Delete comment
          </div>
        </div>
      </div>
      <div v-else class="d-flex align-items-center"><p>No comments yet</p></div>
    
    </div>
  </div>
</div>
</div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
      </div>
    </div>
  </div>
</div>
-->


<!-- Finestra modale per la lista di commenti -->
<div class="modal fade" id="commentsModal" tabindex="-1" role="dialog" aria-labelledby="commentsModalLabel" aria-hidden="true">
  <div class="modal-dialog modal-lg" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="commentsModalLabel">Comments</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body overflow-auto">

        <div class="card mb-3" v-for="c in modalComments.comments">
          <div class="container">
  <!-- Stack the columns on mobile by making one full-width and the other half-width -->
  <div class="row p-3">
    <RouterLink :to="'/profile/'+c.user.userID" class="nav-link" :us="c.user.userID"  data-dismiss="modal" @click="this.us=c.user.userID; this.load()">
      <div class="col-md-8"><h5>{{ c.user.username }}</h5></div>
    </RouterLink>
    
  </div>

  <!-- Columns start at 50% wide on mobile and bump up to 33.3% wide on desktop -->
  <div class="row">
    <div class="col">{{ c.text }}</div>
  </div>

  <!-- Columns are always 50% wide, on mobile and desktop -->
  <div class="row p-2" >
    <div class="col-9 col-md-9"><button type="button" class="btn btn-outline-danger delete-btn" v-if="c.user.userID==this.myID" @click="this.deleteComment(c)">Delete comment</button></div>
    <div class="col-3 text-muted"><small v-text=this.getReadableDate(c.timestamp)></small></div>
  </div>
</div>
        </div>
      </div>
      <div class="border-top mt-5 mb-5"></div>
        <div class="my-3 mx-2">
          Type your comment: <input type="text" v-model="textComment"> 
          <button class="btn btn-primary ms-2" @click="this.comment()">Comment!</button>
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