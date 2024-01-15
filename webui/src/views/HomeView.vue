<script>
import LoadingSpinner from "../components/LoadingSpinner.vue"
import { RouterLink } from 'vue-router';
import ErrorMsg from "../components/ErrorMsg.vue"

export default {
	components: LoadingSpinner, ErrorMsg,
	data: function () {
		return {
			errormsg: null,
			msg: null,
			loading: false,
			stream: {},
			username: "",
			myID: parseInt(localStorage.getItem("token")),
			textComments: [],
		}
	},
	methods: {
		async refresh() {
			if (!(localStorage.getItem("token") > 0)) {
				this.$router.push({ path: '/' });
				return
			}
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users/" + this.myID + "/photos",
					{ headers: { Authorization: this.myID } });
				this.stream = response.data;
				this.username = this.stream.user.username

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async deleteComment(ph, c) {
			try {
				await this.$axios.delete('/users/' + ph.photo.user.userID + "/photos/" + ph.photo.photoID + "/comments/" + c.commentID,
					{
						headers: {
							Authorization: this.myID
						}
					});
				ph.comments = ph.comments.filter(com => com.commentID !== c.commentID);
				ph.nComments--;
			}
			catch (e) {
				this.errormsg = "There are problems with the uncomment request...";
			}

		},
		async comment(ph) {
			let txt = this.textComments[ph.photo.photoID]
			if (txt == "") { return }
			try {
				let response = await this.$axios.post('/users/' + ph.photo.user.userID + "/photos/" + ph.photo.photoID + "/comments", {
					text: txt,
					user: { userID: this.myID },
				},
					{
						headers: {
							Authorization: this.myID
						}
					});
				this.textComments[ph.photo.photoID] = "";
				if (ph.comments == null) {
					ph.comments = [];
					ph.comments[0] = response.data
				}
				else {
					ph.comments.push(response.data)
				}
				ph.nComments++;
			}
			catch (e) {
				this.errormsg = "There are problems with the comment request...";

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
			let indexP = this.stream.photos.findIndex(p => p.photo.photoID == ph.photo.photoID);
			if (l != 0) {
				try {
					await this.$axios.delete('/users/' + ph.photo.user.userID + "/photos/" + ph.photo.photoID + "/likes/" + l,
						{
							headers: {
								Authorization: this.myID
							}
						});
					ph.likes = ph.likes.filter(li => li.likeID !== l);
					ph.nLikes--;
				}
				catch (e) {
					this.errormsg = "There are problems with the unlike request...";

				}
			}
			else {

				try {
					let response = await this.$axios.post('/users/' + ph.photo.user.userID + "/photos/" + ph.photo.photoID + "/likes", {
						userID: parseInt(this.myID),
						photoID: ph.photo.photoID,
					},
						{
							headers: {
								Authorization: this.myID
							}
						});
					if (ph.likes == null) {
						ph.likes = [];
						ph.likes[0] = response.data
					}
					else {
						ph.likes.push(response.data)
					}
					ph.nLikes++;
				}
				catch (e) {
					this.errormsg = "There are problems with the unlike request...";

				}
			}

		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<h2 class="my-3">Benvenuto/a {{ this.username }}!</h2>

	<LoadingSpinner :loading="this.loading"></LoadingSpinner>

	<div class="card mt-5 container" v-for="ph in this.stream.photos" :key="ph.photo.photoID">
		<div class="card-header">
			<RouterLink :to="'/profile/' + ph.photo.user.userID" class="nav-link" :us="ph.photo.user.userID">
				<h3>{{ ph.photo.user.username }}</h3>
			</RouterLink>
		</div>
		<div class="overflow-auto">
			<!-- Left Box -->
			<div class="row overflow-auto">
				<div class="col">
					<div class="custom-square">
						<div class="custom-content">
							<div class="image-wrapper">
								<img :src="'data:image/*;base64,' + ph.photo.binary" alt="Image" class="custom-image">
							</div>
						</div>
					</div>
				</div>

				<!-- Right Box -->
				<div class="col-md-8 overflow-auto" style="width:60%">
					<h2 class="mb-2">Comments</h2>
					<div style="height:500px; max-height: 70%; overflow-y:auto">
						<div class="card mb-3" v-for="c in ph.comments" :key="c.commentID">
							<div class="container">
								<div class="row p-3">
									<RouterLink :to="'/profile/' + c.user.userID" class="nav-link" :us="c.user.userID">
										<div class="col-md-8">
											<h4>{{ c.user.username }}</h4>
										</div>
									</RouterLink>
								</div>
								<div class="row">
									<div class="col">{{ c.text }}</div>
								</div>
								<div class="row p-2">
									<div class="col-9 col-md-9" v-if="c.user.userID == this.myID">
										<button type="button" class="btn btn-outline-danger delete-btn"
											@click="this.deleteComment(ph, c)">Delete
											comment</button>
									</div>
									<div class="col-3 text-muted"><small v-text=this.getReadableDate(c.timestamp)></small>
									</div>
								</div>
							</div>
						</div>
					</div>
					<div class="my-3 mx-2">
						Type your comment: <input type="text" v-model="textComments[ph.photo.photoID]" required
							placeholder="Your comment here...">
						<button class="btn btn-primary ms-2" @click="this.comment(ph)">Comment!</button>
					</div>
				</div>
			</div>
		</div>
		<div class="card-footer">

			<div class="align-middle mx-5">
				<label class="display-6 align-middle  p-2">{{ ph.nLikes }}</label>
				<svg style="width:50px; cursor:pointer" :fill="ph.likes == null || this.getLikeID(ph) == 0 ? 'none' : 'red'"
					stroke="black" @click="this.like(ph)">
					<use href="/feather-sprite-v4.29.0.svg#heart" />
				</svg>
				<label class="display-6 align-middle p-2">{{ ph.nComments }}</label>
				<svg style="width:50px;" fill="none" stroke="black">
					<use href="/feather-sprite-v4.29.0.svg#message-circle" />
				</svg>
			</div>
		</div>
	</div>


	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
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