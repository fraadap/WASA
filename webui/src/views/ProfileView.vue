<script>
import { RouterLink } from 'vue-router';
import ErrorMsg from '../components/ErrorMsg.vue';
import Msg from '../components/Msg.vue';


export default {
  props: ['userID'],
  components: ErrorMsg, Msg,
  data: function () {
    return {
      errormsg: null,
      msg: null,
      profile: {},
      myProfile: {},
      us: parseInt(this.userID),
      myID: localStorage.getItem("token"),
      modalComments: {},
      textComment: "",
      imBanned: false,
    };
  },
  methods: {
    async amIBanned() {
      this.imBanned = false
      if (this.profile.bans == null) { return false }
      for (let i = 0; i < this.profile.bans.length; i++) {

        if (this.profile.bans[i].userID == this.myID) {
          this.imBanned = true
        }
      }
    },
    async load() {
      try {
        let response = await this.$axios.get('/users/' + this.myID,
          { headers: { Authorization: this.myID } });
        this.myProfile = response.data;

        if (this.us == 0 || this.us == this.myID) {
          this.us = this.myID
          this.profile = this.myProfile
        }
        else {
          let response = await this.$axios.get('/users/' + this.us,
            { headers: { Authorization: this.myID } });
          this.profile = response.data;
        }
        console.log(this.profile)
        this.amIBanned()
      }
      catch (e) {
        this.errormsg = "There are problems with the profile's loading...";

      }

    },
    async uploadPhoto() {
      this.image = this.$refs.file.files[0];
    },
    async submitPhoto() {
      if (this.images === null) {
        this.errormsg = "Please select a photo";
      } else {
        try {
          await this.$axios.post("/users/" + this.myID + "/photos", this.image,
            {
              headers: {
                Authorization: this.myID
              }
            })
          this.load();
          this.msg = "Photo uploaded succesfully";
        } catch (e) {
          this.errormsg = "There are problems with the photo's uploading...";

        }
      }

    },
    isFollowed(id) {
      let yes = false;
      let f;
      if (this.myProfile.followings == null) { return false }
      for (let i = 0; i < this.myProfile.followings.length; i++) {

        if (this.myProfile.followings[i].userID == id) {
          yes = true;
        }
      }
      return yes;
    },
    isBanned(id) {  //implementare ban e unban
      let yes = false;
      if (this.myProfile.bans == null) { return false }
      for (let i = 0; i < this.myProfile.bans.length; i++) {

        if (this.myProfile.bans[i].userID == id) {
          yes = true;
        }
      }
      return yes;
    },
    async unfollow(id) {
      try {
        let response = await this.$axios.get('/users/' + this.myID + "/followID/" + id,
          {
            headers: {
              Authorization: this.myID
            }
          });
        let followID = response.data.followID

        await this.$axios.delete('/users/' + this.myID + "/follow/" + followID,
          {
            headers: {
              Authorization: this.myID
            }
          });

        this.load();
        this.msg = "User unfollowed succesfully";
      }
      catch (e) {
        this.errormsg = "There are problems with the unfollow request...";

      }
    },
    async follow(id) {
      try {
        await this.$axios.post('/users/' + this.myID + "/follow", {
          followed: id,
        },
          {
            headers: {
              Authorization: this.myID
            }
          });
        this.load();
      }
      catch (e) {
        alert("You can't follow the user because he banned you :(");
      }
    },
    async ban(id) {
      try {
        await this.$axios.post('/users/' + this.myID + "/bans", {
          banned: id,
        },
          {
            headers: {
              Authorization: this.myID
            }
          });
        this.load()
      }
      catch (e) {
        this.errormsg = "There are problems with the ban request...";

      }

    },
    async unban(id) {
      try {
        let response = await this.$axios.get('/users/' + this.myID + "/banID/" + id,
          {
            headers: {
              Authorization: this.myID
            }
          });
        let banID = response.data.banID
        await this.$axios.delete('/users/' + this.myID + "/bans/" + banID,
          {
            headers: {
              Authorization: this.myID
            }
          });

        this.load();
      }
      catch (e) {
        this.errormsg = "There are problems with the unban request...";

      }
    },
    getLikeID(ph) {
      let like = 0;
      if (ph.likes == null) { return 0 }
      for (let i = 0; i < ph.likes.length; i++) {
        if (ph.likes[i].userID == this.myID) {
          like = ph.likes[i].likeID
        }
      }
      return like;
    },
    async like(ph) {
      let l = this.getLikeID(ph)
      if (l != 0) {
        try {
          await this.$axios.delete('/users/' + ph.photo.user.userID + "/photos/" + ph.photo.photoID + "/likes/" + l,
            {
              headers: {
                Authorization: this.myID
              }
            });
        }
        catch (e) {
          this.errormsg = "There are problems with the unlike request...";

        }
      }
      else {

        try {
          await this.$axios.post('/users/' + ph.photo.user.userID + "/photos/" + ph.photo.photoID + "/likes", {
            userID: parseInt(this.myID),
            photoID: ph.photo.photoID,
          },
            {
              headers: {
                Authorization: this.myID
              }
            });
        }
        catch (e) {
          this.errormsg = "There are problems with the like request...";

        }
      }
      this.load()
    },
    async deletePhoto(ph) {
      try {
        await this.$axios.delete('/users/' + ph.photo.user.userID + "/photos/" + ph.photo.photoID,
          {
            headers: {
              Authorization: this.myID
            }
          });
        this.load()
        this.msg = "Photo deleted succesfully";
      }
      catch (e) {
        this.errormsg = "There are problems with deleting the photo..";
      }

    },
    getReadableDate(d) {
      const dateObject = new Date(d);

      const year = dateObject.getFullYear();
      const month = dateObject.getMonth() + 1; // I mesi in JavaScript vanno da 0 a 11
      const day = dateObject.getDate();
      const hours = dateObject.getHours();
      const minutes = dateObject.getMinutes();
      const seconds = dateObject.getSeconds();

      const formattedDate = day + "-" + month + "-" + year + " | " + hours + ":" + minutes;
      return formattedDate
    },
    async comment() {
      let txt = this.textComment;
      if (txt == "") { return }
      let ph = this.modalComments;
      try {
        let response = await this.$axios.post('/users/' + ph.photo.user.userID + "/photos/" + ph.photo.photoID + "/comments", {
          text: txt,
          user: { userID: parseInt(this.myID) },
        },
          {
            headers: {
              Authorization: this.myID
            }
          });
        this.load();
        this.textComment = "";

        if (this.modalComments.comments == null) {
          this.modalComments.comments = [];
          this.modalComments.comments[0] = response.data
        } else {
          this.modalComments.comments.push(response.data);
        }
        this.msg = "Commented succesfully";
      }
      catch (e) {
        this.errormsg = "There are problems with the comment's uploading...";

      }

    },
    async deleteComment(c) {
      let ph = this.modalComments;
      try {
        await this.$axios.delete('/users/' + ph.photo.user.userID + "/photos/" + ph.photo.photoID + "/comments/" + c.commentID,
          {
            headers: {
              Authorization: this.myID
            }
          });
        this.load();
        this.modalComments.comments = this.modalComments.comments.filter(com => com.commentID !== c.commentID);
        this.msg = "Comment deleted succesfully";
      }
      catch (e) {
        this.errormsg = "There are problems with the comment's deleting...";

      }
    },
  },
  mounted() {
    if (!(localStorage.getItem("token") > 0)) {
      this.$router.push({ path: '/#/' });
      return
    }
    this.load();
  },
  watch: {
    userID: {
      handler() {
        // Questa funzione verrà eseguita ogni volta che userID cambierà
        this.us = parseInt(this.userID);
        // Esegui la funzione load() ogni volta che userID cambia
        this.load();
      },
    },
  },
}
</script>

<template>
  <Msg :msg="this.msg" v-if="this.msg"></Msg>

  <div v-if="!this.imBanned">
    <div class="container mt-5 text-center">
      <div class="row">
        <div class="col fs-3">
          <p>{{ this.profile && this.profile.user && this.profile.user.username }}</p>
        </div>
        <div class="col fs-5">
          <p>Photos</p>
          <p>{{ this.profile.photos ? this.profile.nPhotos : 0 }}</p>
        </div>
        <div class="col fs-5" data-toggle="modal" data-target="#followersModal" style="cursor: pointer;">
          <p>Followers</p>
          <p>{{ this.profile.followers ? this.profile.followers.length : 0 }}</p>
        </div>
        <div class="col fs-5" data-toggle="modal" data-target="#followingsModal" style="cursor: pointer;">
          <p>Followings</p>
          <p>{{ this.profile.followings ? this.profile.followings.length : 0 }}</p>
        </div>
        <div class="col fs-5">
          <div class="p-2 bd-highlight" v-if="this.us != this.myID && this.us != 0">
            <button type="button" class="btn btn-outline-success me-3" @click="follow(this.us)"
              v-if="isFollowed(this.us) == false">Follow</button>
            <button v-else type="button" class="btn btn-primary me-3" @click="unfollow(this.us)"> Followed </button>
            <button type="button" class="btn btn-outline-danger" v-if="isBanned(this.us) == false"
              @click="ban(this.us)">Ban</button>
            <button type="button" class="btn btn-outline-danger" v-else @click="unban(this.us)">Banned</button>
          </div>
          <div class="p-2 bd-highlight" v-else> <!-- Nuova foto -->
            <button class="btn btn-success fs-3" data-toggle="modal" data-target="#addPhotoModal">+</button>
          </div>
        </div>
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
              <svg style="width:30px; cursor:pointer" :fill="ph.likes == null || this.getLikeID(ph) == 0 ? 'none' : 'red'"
                stroke="white" @click="like(ph)">
                <use href="/feather-sprite-v4.29.0.svg#heart" />
              </svg>
              {{ ph.nComments }}
              <svg style="width:30px; cursor:pointer" fill="none" stroke="white" @click="this.modalComments = ph"
                data-toggle="modal" data-target="#commentsModal">
                <use href="/feather-sprite-v4.29.0.svg#message-circle" />
              </svg>
            </div>
            <div class="d-flex flex-row-reverse" style="margin-left:55%">
              <svg style=" width:30px; cursor:pointer" fill="none" stroke="white" @click="deletePhoto(ph)"
                v-if="this.us == this.myID || this.us == 0">
                <use href="/feather-sprite-v4.29.0.svg#trash-2" />
              </svg>
            </div>
          </div>
        </div>
      </div>

    </div>


    <!-- Upload photo modal -->
    <div class="modal fade" id="addPhotoModal" tabindex="-1" role="dialog" aria-labelledby="addPhotoModalLabel"
      aria-hidden="true">
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

    <div class="modal fade" id="followersModal" tabindex="-1" role="dialog" aria-labelledby="followersModalLabel"
      aria-hidden="true">
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

                  <div class="card mb-3" v-for="user in this.profile.followers" v-if="this.profile.followers != null">
                    <div class="card-body d-flex align-items-center justify-content-between">
                      <div class="me-5 p-2 bd-highlight">
                        <RouterLink :to="'/profile/' + user.userID" class="nav-link" :us="user.userID"
                          data-dismiss="modal" @click="this.us = user.userID; this.load()">
                          <h5 class="card-title">{{ user.username }}</h5>
                        </RouterLink>
                      </div>
                      <div class="p-2 bd-highlight" v-if="user.userID != this.myID && user.userID != 0">
                        <button type="button" class="btn btn-outline-success me-3" @click="follow(user.userID)"
                          v-if="isFollowed(user.userID) == false">Follow</button>
                        <button v-else type="button" class="btn btn-primary me-3" @click="unfollow(user.userID)"> Followed
                        </button>
                        <button type="button" class="btn btn-outline-danger" v-if="isBanned(user.userID) == false"
                          @click="ban(user.userID)">Ban</button>
                        <button type="button" class="btn btn-outline-danger" v-else
                          @click="unban(user.userID)">Banned</button>
                      </div>
                    </div>
                  </div>
                  <div v-else class="d-flex align-items-center">
                    <p>No followers yet</p>
                  </div>
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
    <div class="modal fade" id="followingsModal" tabindex="-1" role="dialog" aria-labelledby="followingsModalLabel"
      aria-hidden="true">
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


                  <div class="card mb-3" v-for="user in this.profile.followings" v-if="this.profile.followings != null">
                    <div class="card-body d-flex align-items-center justify-content-between">
                      <div class="me-5 p-2 bd-highlight">
                        <RouterLink :to="'/profile/' + user.userID" class="nav-link" :us="user.userID"
                          data-dismiss="modal" @click="this.us = user.userID; this.load()">
                          <h5 class="card-title">{{ user.username }}</h5>
                        </RouterLink>
                      </div>
                      <div class="p-2 bd-highlight" v-if="user.userID != this.myID">
                        <button type="button" class="btn btn-outline-success me-3" @click="follow(user.userID)"
                          v-if="isFollowed(user.userID) == false">Follow</button>
                        <button v-else type="button" class="btn btn-primary me-3" @click="unfollow(user.userID)"> Followed
                        </button>
                        <button type="button" class="btn btn-outline-danger" v-if="isBanned(user.userID) == false"
                          @click="ban(user.userID)">Ban</button>
                        <button type="button" class="btn btn-outline-danger" v-else
                          @click="unban(user.userID)">Banned</button>
                      </div>
                    </div>
                  </div>
                  <div v-else class="d-flex align-items-center">
                    <p>No followings yet</p>
                  </div>
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

    <!-- Finestra modale per la lista di commenti -->
    <div class="modal fade" id="commentsModal" tabindex="-1" role="dialog" aria-labelledby="commentsModalLabel"
      aria-hidden="true">
      <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="commentsModalLabel">Comments</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body overflow-auto">

            <div class="mb-3 card" v-for="c in modalComments.comments">
              <div class="container">
                <div class="row p-1">

                  <div class="col mt-1">

                    <RouterLink :to="'/profile/' + c.user.userID" class="nav-link" :us="c.user.userID"
                      data-dismiss="modal" @click="this.us = c.user.userID; this.load()">
                      <h5 style=""> {{ c.user.username }} </h5>
                    </RouterLink>
                    <div class="col">
                      <button type="button" class="btn btn-outline-danger delete-btn" v-if="c.user.userID == this.myID"
                        @click="this.deleteComment(c)">Delete comment</button>
                    </div>
                  </div>
                </div>
                <div class="row">
                  <div class="col">{{ c.text }}</div>
                </div>
                <div class="row p-2">
                  <div class="col-9 col-md-9"></div>
                  <div class="col-3 text-muted"><small v-text=this.getReadableDate(c.timestamp)></small></div>
                </div>
              </div>
            </div>
          </div>
          <div class="border-top mt-5 mb-5"></div>
          <div class="my-3 mx-2">
            Type your comment: <input type="text" v-model="textComment" required placeholder="Your comment here...">
            <button class="btn btn-primary ms-2" @click="this.comment()">Comment!</button>
          </div>
        </div>
      </div>
    </div>
  </div>
  <ErrorMsg v-else :msg='"Ops! Its seem like you are banned from the user " + profile.user.username + " :("'></ErrorMsg>
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